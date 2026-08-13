package netconf

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

type testServer struct {
	t   *testing.T
	in  chan []byte
	out chan []byte
}

func newTestServer(t *testing.T) *testServer {
	return &testServer{
		t:   t,
		in:  make(chan []byte),
		out: make(chan []byte),
	}
}

func (s *testServer) handle(r io.ReadCloser, w io.WriteCloser) {
	in, err := io.ReadAll(r)
	if err != nil {
		panic(fmt.Sprintf("testerver: failed to read incomming message: %v", err))
	}
	s.t.Logf("testserver recv: %s", in)
	go func() { s.in <- in }()

	out, ok := <-s.out
	if !ok {
		panic("testserver: no message to send")
	}
	s.t.Logf("testserver send: %s", out)

	_, err = w.Write(out)
	if err != nil {
		panic(fmt.Sprintf("testserver: failed to write message: %v", err))
	}

	if err := w.Close(); err != nil {
		panic("testserver: failed to close outbound message")
	}
}

func (s *testServer) queueResp(p []byte)         { go func() { s.out <- p }() }
func (s *testServer) queueRespString(str string) { s.queueResp([]byte(str)) }
func (s *testServer) popReq() ([]byte, error) {
	msg, ok := <-s.in
	if !ok {
		return nil, fmt.Errorf("testserver: no message to read")
	}
	return msg, nil
}

func (s *testServer) transport() *testTransport { return newTestTransport(s.handle) }

type testTransport struct {
	handler func(r io.ReadCloser, w io.WriteCloser)
	out     chan io.ReadCloser
	// msgReceived, msgSent int
}

func newTestTransport(handler func(r io.ReadCloser, w io.WriteCloser)) *testTransport {
	return &testTransport{
		handler: handler,
		out:     make(chan io.ReadCloser),
	}
}

func (s *testTransport) MsgReader() (io.ReadCloser, error) {
	return <-s.out, nil
}

func (s *testTransport) MsgWriter() (io.WriteCloser, error) {
	inr, inw := io.Pipe()
	outr, outw := io.Pipe()

	go func() { s.out <- outr }()
	go s.handler(inr, outw)

	return inw, nil
}

func (s *testTransport) Close() error {
	if len(s.out) > 0 {
		return fmt.Errorf("testtransport: remaining outboard messages not sent at close")
	}
	return nil
}

func TestSanitizeXML(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want string
	}{
		{name: "clean", in: "<a>ok</a>", want: "<a>ok</a>"},
		{name: "stx", in: "<a>aa-\nbb-cc\x02</a>", want: "<a>aa-\nbb-cc</a>"},
		{name: "nul", in: "<a>\x00x</a>", want: "<a>x</a>"},
		{name: "tab newline cr kept", in: "<a>\t\n\r</a>", want: "<a>\t\n\r</a>"},
		{name: "utf8 nonchar", in: "<a>\xEF\xBF\xBE</a>", want: "<a></a>"},
		{name: "empty", in: "", want: ""},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, []byte(tc.want), sanitizeXML([]byte(tc.in)))
		})
	}
}

func TestMessageIDFromRaw(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want uint64
	}{
		{name: "double quote", in: `<rpc-reply message-id="42">`, want: 42},
		{name: "single quote", in: `<rpc-reply message-id='7'>`, want: 7},
		{name: "missing", in: `<rpc-reply>`, want: 0},
		{name: "unquoted", in: `<rpc-reply message-id=3>`, want: 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, messageIDFromRaw([]byte(tc.in)))
		})
	}
}

func TestNonUTF8HelloMessage(t *testing.T) {
	ts := newTestServer(t)
	sess, err := newSession(WithTransport(ts.transport()))
	require.NoError(t, err)

	ts.queueRespString(`<?xml version="1.0" encoding="ISO-8859-1"?>` +
		`<hello xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">` +
		`<session-id>42</session-id>` +
		`<capabilities>` +
		`<capability>urn:ietf:params:netconf:base:1.0</capability>` +
		`</capabilities>` +
		`</hello>`)

	err = sess.handshake(t.Context())
	require.NoError(t, err)
	require.Equal(t, uint64(42), sess.SessionID())
}
