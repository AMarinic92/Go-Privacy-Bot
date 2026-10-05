# inbox/imap

## Responsibility

An `Inbox` that polls the catch-all IMAP mailbox. It finds confirmation links and classifies broker replies. Each broker has its own alias, so replies route back to the right request.

## Imports

Adapter. Imports `engine`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Client struct{}`
- `func (c *Client) Fetch(ctx context.Context, since time.Time) ([]engine.Message, error)`
- `var _ engine.Inbox = (*Client)(nil)`

## To implement

- `Client` fields and `Fetch`. `emersion/go-imap` is the candidate library.

## Open questions

- Whether reply classification lives here or in `engine`. See `internal/engine/README.md`.
- What `Message` carries, and whether raw messages are stored as evidence.
- Where the `since` time comes from.
