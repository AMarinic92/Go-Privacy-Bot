# notify

## Responsibility

A `Notifier` that sends the operator the digest by push or email. The manual queue goes out with the digest.

## Imports

Adapter. Imports `engine`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Notifier struct{}`
- `func (n *Notifier) Send(ctx context.Context, d engine.Digest) error`
- `var _ engine.Notifier = (*Notifier)(nil)`

## To implement

- `Notifier` fields and `Send`.

## Open questions

- Push, email, or both. No push service is chosen.
- What the digest contains, and whether a run with no changes sends one.
