package netconf

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	tt := []struct {
		name    string
		source  Datastore
		options []GetOption
		matches []*regexp.Regexp
	}{
		{
			name:   "get-config startup with-defaults",
			source: Startup,
			options: []GetOption{
				WithDefaultMode(DefaultsModeTrim),
			},
			matches: []*regexp.Regexp{
				regexp.MustCompile(`<source>\S*<startup/>\S*</source>`),
				regexp.MustCompile(`<with-defaults xmlns="urn:ietf:params:xml:ns:yang:ietf-netconf-with-defaults">trim</with-defaults>`),
			},
		},
		{
			name:   "get-config running filter",
			source: Running,
			options: []GetOption{
				WithSubtreeFilter(`<interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces"/>`),
			},
			matches: []*regexp.Regexp{
				regexp.MustCompile(`<source>\S*<running/>\S*</source>`),
				regexp.MustCompile(`<filter type="subtree"><interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces"/></filter>`),
			},
		},
		{
			name:   "get-config running no options",
			source: Running,
			matches: []*regexp.Regexp{
				regexp.MustCompile(`<get-config xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><source><running/></source></get-config>`),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ts := newTestServer(t)
			sess, _ := newSession(WithTransport(ts.transport()))
			sess.serverCaps = newCapabilitySet(append(DefaultCapabilities, "urn:ietf:params:netconf:capability:with-defaults:1.0?basic-mode=report-all&also-supported=explicit,trim")...)
			go sess.recv()

			ts.queueRespString(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1"><ok/></rpc-reply>`)

			reply, err := sess.GetConfig(t.Context(), tc.source, tc.options...)
			require.NoError(t, err)
			require.NotNil(t, reply)

			sentMsg, err := ts.popReq()
			require.NoError(t, err)

			for _, match := range tc.matches {
				require.Regexp(t, match, string(sentMsg))
			}
		})
	}
}

func TestGet(t *testing.T) {
	tt := []struct {
		name    string
		options []GetOption
		matches []*regexp.Regexp
	}{
		{
			name: "get ifm",
			options: []GetOption{
				WithSubtreeFilter(`<ifm xmlns="urn:huawei:yang:huawei-ifm"/>`),
			},
			matches: []*regexp.Regexp{
				regexp.MustCompile(`<get xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">\S*<filter type="subtree">\S*<ifm xmlns="urn:huawei:yang:huawei-ifm"/>\S*</filter>\S*</get>`),
			},
		},
		{
			name: "get devm",
			options: []GetOption{
				WithDefaultMode("report-all"),
				WithSubtreeFilter(`<devm xmlns="urn:huawei:yang:huawei-devm"/>`),
			},
			matches: []*regexp.Regexp{
				regexp.MustCompile(`<get xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">\S*<filter type="subtree">\S*<devm xmlns="urn:huawei:yang:huawei-devm"/>\S*</filter>\S*<with-defaults xmlns="urn:ietf:params:xml:ns:yang:ietf-netconf-with-defaults">report-all</with-defaults>\S*</get>`),
			},
		},
		{
			name: "get configuration with attribute",
			options: []GetOption{
				WithDefaultMode("report-all"),
				WithSubtreeFilter(`<configuration/>`),
				WithAttribute("format", "set"),
			},
			matches: []*regexp.Regexp{
				regexp.MustCompile(`<get xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" format="set">\S*<filter type="subtree">\S*<configuration/>\S*</filter>\S*<with-defaults xmlns="urn:ietf:params:xml:ns:yang:ietf-netconf-with-defaults">report-all</with-defaults>\S*</get>`),
			},
		},
		{
			name: "get with unsupported capability fails",
			options: []GetOption{
				WithDefaultMode("trim"),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ts := newTestServer(t)
			sess, _ := newSession(WithTransport(ts.transport()))
			sess.serverCaps = newCapabilitySet(append(DefaultCapabilities, "urn:ietf:params:netconf:capability:with-defaults:1.0?basic-mode=report-all&also-supported=explicit")...)
			go sess.recv()

			ts.queueRespString(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1"><data>daa</data></rpc-reply>`)

			reply, err := sess.Get(t.Context(), tc.options...)
			if len(tc.matches) == 0 {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, reply)

				sentMsg, err := ts.popReq()
				require.NoError(t, err)

				for _, match := range tc.matches {
					require.Regexp(t, match, string(sentMsg))
				}
			}
		})
	}
}

func TestGetWithIllegalControlCharacter(t *testing.T) {
	ts := newTestServer(t)
	sess, err := newSession(WithTransport(ts.transport()))
	require.NoError(t, err)
	sess.serverCaps = newCapabilitySet(DefaultCapabilities...)
	go sess.recv()

	replyXML := "<rpc-reply xmlns=\"urn:ietf:params:xml:ns:netconf:base:1.0\" message-id=\"1\">" +
		"<data><If-list><id>eth1/47</id><adj-items><AdjEp-list>" +
		"<id>1</id><capability>bridge,router</capability>" +
		"<chassisIdT>7</chassisIdT><chassisIdV>aa-\nbb-cc\x02</chassisIdV>" +
		"</AdjEp-list></adj-items></If-list></data></rpc-reply>"
	ts.queueResp([]byte(replyXML))

	ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)
	defer cancel()

	reply, err := sess.Get(ctx, WithSubtreeFilter(`<state/>`))
	require.NoError(t, err)
	require.NotNil(t, reply)
	require.Equal(t, uint64(1), reply.MessageID)
	require.Contains(t, string(reply.Raw()), "<chassisIdV>aa-\nbb-cc</chassisIdV>")
	require.NotContains(t, string(reply.Raw()), "\x02")
}

func TestGetDeliversReplyWhenDecodeFails(t *testing.T) {
	ts := newTestServer(t)
	sess, err := newSession(WithTransport(ts.transport()))
	require.NoError(t, err)
	sess.serverCaps = newCapabilitySet(DefaultCapabilities...)
	go sess.recv()

	// Unescaped ]]> is illegal XML 1.0 and is not stripped by sanitizeXML.
	ts.queueRespString(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1"><data>foo]]>bar</data></rpc-reply>`)

	ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)
	defer cancel()

	reply, err := sess.Get(ctx, WithSubtreeFilter(`<state/>`))
	require.NoError(t, err)
	require.NotNil(t, reply)
	require.Equal(t, uint64(1), reply.MessageID)
	require.Contains(t, string(reply.Raw()), "foo]]>bar")
}

// errorCollectingLogger records Errorf calls so tests can assert the decode
// path did not have to fall back to an undecoded raw reply.
type errorCollectingLogger struct {
	noOpLogger
	errors []string
}

func (l *errorCollectingLogger) Errorf(format string, args ...any) {
	l.errors = append(l.errors, fmt.Sprintf(format, args...))
}

func TestGetWithInvalidUTF8Byte(t *testing.T) {
	ts := newTestServer(t)
	logger := &errorCollectingLogger{}
	sess, err := newSession(WithTransport(ts.transport()), WithLogger(logger))
	require.NoError(t, err)
	sess.serverCaps = newCapabilitySet(DefaultCapabilities...)
	go sess.recv()

	// A lone 0xCB byte (invalid UTF-8 continuation byte), as observed in
	// device-supplied free-text fields such as LLDP port descriptions.
	replyXML := "<rpc-reply xmlns=\"urn:ietf:params:xml:ns:netconf:base:1.0\" message-id=\"1\">" +
		"<data><port-description>uplink - foo\xCBbar</port-description></data></rpc-reply>"
	ts.queueResp([]byte(replyXML))

	ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)
	defer cancel()

	reply, err := sess.Get(ctx, WithSubtreeFilter(`<state/>`))
	require.NoError(t, err)
	require.NotNil(t, reply)
	require.Equal(t, uint64(1), reply.MessageID)
	require.Contains(t, string(reply.Raw()), "uplink - foo\uFFFDbar")
	require.NotContains(t, string(reply.Raw()), "\xCB")
	// The corrupted byte was repaired before decoding, so the full
	// rpc-reply parsed successfully instead of falling back to a raw,
	// mostly-undecoded reply.
	require.Empty(t, logger.errors)
}
