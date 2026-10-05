# engine

## Responsibility

Runs the daily tick and picks the next action for each broker. The engine holds every decision. It reaches the outside world only through the five ports in `ports.go`.

## Imports

Core package. Imports `lifecycle`, `broker` and `profile` today, and may import `letter`. Must not import any adapter package, `registry` or `cmd/scrub`. Signatures take no file paths, connections or clients. No network or disk access.

## API

- `ports.go`: interfaces `Source`, `Channel`, `Inbox`, `Store` and `Notifier`; empty structs `Listing`, `Receipt`, `Message` and `Digest`.
- `engine.go`: `type Engine struct{}`, which will hold the five ports, and `func (e *Engine) Tick(ctx context.Context, now time.Time) error`. There is no constructor.

## To implement

- `Engine` fields and `Tick`.
- Fields of `Listing`, `Receipt`, `Message` and `Digest`.

## Open questions

- How the engine picks a Channel for a broker, and whether it holds one Source or several.
- How `cmd/scrub` supplies the ports, given there is no constructor.
- Which port opens a confirmation link (`AwaitingConfirm` to `AwaitingResponse`). None of the five does.
- Where the `since` time for `Inbox.Fetch` is kept. `Store` has no method for it.
- How a rendered `letter.Letter` reaches a Channel. `Submit` takes only a `lifecycle.Request`.
- How a Channel reports that a request needs the operator (`Listed` to `ManualQueue`).
- Whether reply classification belongs here or in `inbox/imap`. Diagram 2 puts it in the adapter, but the engine is meant to hold every decision.
