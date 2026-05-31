# netconf

[![CI](https://github.com/Mrflatt/netconf/actions/workflows/ci.yaml/badge.svg)](https://github.com/Mrflatt/netconf/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/Mrflatt/netconf/graph/badge.svg?token=TY1CJ0MYXY)](https://codecov.io/gh/Mrflatt/netconf)
[![Go Reference](https://pkg.go.dev/badge/github.com/Mrflatt/netconf.svg)](https://pkg.go.dev/github.com/Mrflatt/netconf)

Go implementation of the NETCONF protocol ([RFC 6241](https://www.rfc-editor.org/rfc/rfc6241.html)).

## Features

- Full NETCONF base operations: `get`, `get-config`, `edit-config`, `copy-config`, `delete-config`, `lock`, `unlock`, `commit`, `validate`, `discard-changes`, `kill-session`, `close-session`
- Confirmed commit with timeout, persist, and cancel support ([RFC 6241 Section 8.4](https://www.rfc-editor.org/rfc/rfc6241.html#section-8.4))
- Event notifications via `create-subscription` ([RFC 5277](https://www.rfc-editor.org/rfc/rfc5277.html))
- Custom RPC dispatch for vendor-specific operations
- Capability negotiation with automatic NETCONF 1.1 chunked framing upgrade
- SSH transport ([RFC 6242](https://www.rfc-editor.org/rfc/rfc6242.html))
- TLS transport ([RFC 7589](https://www.rfc-editor.org/rfc/rfc7589.html))

## Install

```
go get github.com/Mrflatt/netconf
```

## Usage

### SSH

```go
import (
    "context"

    "github.com/Mrflatt/netconf"
    ncssh "github.com/Mrflatt/netconf/transport/ssh"
    "golang.org/x/crypto/ssh"
)

config := &ssh.ClientConfig{
    User:            "admin",
    Auth:            []ssh.AuthMethod{ssh.Password("admin")},
    HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

ctx := context.Background()

tr, err := ncssh.Dial(ctx, "tcp", "router:830", config)
if err != nil {
    // handle error
}

sess, err := netconf.NewSession(ctx, netconf.WithTransport(tr))
if err != nil {
    // handle error
}
defer sess.Close(ctx)

reply, err := sess.GetConfig(ctx, netconf.Running)
if err != nil {
    // handle error
}
fmt.Println(reply)
```

### TLS

```go
import (
    "context"
    "crypto/tls"

    "github.com/Mrflatt/netconf"
    nctls "github.com/Mrflatt/netconf/transport/tls"
)

config := &tls.Config{
    // configure certificates
}

ctx := context.Background()

tr, err := nctls.Dial(ctx, "tcp", "router:6513", config)
if err != nil {
    // handle error
}

sess, err := netconf.NewSession(ctx, netconf.WithTransport(tr))
if err != nil {
    // handle error
}
defer sess.Close(ctx)
```

### Edit config

```go
config := `
<config>
    <interface>
        <name>eth0</name>
        <enabled>true</enabled>
    </interface>
</config>`

err := sess.EditConfig(ctx, netconf.Candidate, config,
    netconf.WithDefaultMergeStrategy(netconf.MergeConfig),
    netconf.WithTestStrategy(netconf.TestThenSet),
)
if err != nil {
    // handle error
}

err = sess.Commit(ctx)
```

### Subtree filtering

```go
filter := `<interfaces/>`
reply, err := sess.GetConfig(ctx, netconf.Running,
    netconf.WithSubtreeFilter(filter),
)
```

### Custom RPC

```go
type GetSchema struct {
    XMLName    xml.Name `xml:"urn:ietf:params:xml:ns:yang:ietf-netconf-monitoring get-schema"`
    Identifier string   `xml:"identifier"`
    Format     string   `xml:"format"`
}

reply, err := sess.Dispatch(ctx, &GetSchema{
    Identifier: "ietf-interfaces",
    Format:     "yang",
})
```

## Session options

| Option | Description |
|---|---|
| `WithTransport` | Set the transport (required) |
| `WithCapabilities` | Override default client capabilities |
| `WithNotificationHandler` | Set handler for async notifications |
| `WithLogger` | Set a logger |
| `WithErrorSeverity` | Configure which error severities are returned |
| `WithRPCAttr` | Add extra XML attributes to outgoing `<rpc>` elements |
| `WithHelloTimeout` | Set hello exchange timeout (default 30s) |

## License

[MIT](LICENSE)
