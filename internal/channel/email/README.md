# channel/email

## Responsibility

A `Channel` that sends letters to a broker's privacy contact over SMTP, from that broker's alias. Replies arrive in the catch-all inbox.

## Imports

Adapter. Imports `engine` and `lifecycle`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Sender struct{}`
- `func (s *Sender) Submit(ctx context.Context, r lifecycle.Request) (engine.Receipt, error)`
- `var _ engine.Channel = (*Sender)(nil)`

## To implement

- `Sender` fields and `Submit`.

## Open questions

- The SMTP server and authentication. Mail credentials are in the secrets file, which is not designed.
- How `Submit` gets the rendered letter and the broker's alias. It takes only a `lifecycle.Request`.
- What `Receipt` records for a sent message.
