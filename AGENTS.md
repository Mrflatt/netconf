# AGENTS.md

## Project overview

Go library implementing the NETCONF protocol (RFC 6241). Provides a session-based API for managing network devices over SSH and TLS transports.

## Package structure

- Root package (`netconf`): Session management, NETCONF operations (get, get-config, edit-config, lock, unlock, etc.), capability negotiation, notification subscriptions, XML message encoding/decoding.
- `transport/`: Transport interface definition — message-oriented reader/writer abstraction.
- `transport/ssh/`: SSH transport implementation.
- `transport/tls/`: TLS transport implementation.
- `internal/`: NETCONF message framing (chunked and EOM).

## Development

- Go 1.26+
- Linting: `golangci-lint` (vendored as Go tool dependency)
- Tests: `go test -race ./...`
- CI runs lint, build, and tests on every push/PR

## Conventions

- No external test frameworks beyond `testify`
- Transport implementations live in sub-packages under `transport/`
- Internal framing details stay in `internal/`
