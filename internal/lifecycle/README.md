# lifecycle

## Responsibility

Request states and deadlines. Defines the states of a request record and the rule that moves a record to its next state. Each broker has one request record, and a daily run moves it at most one step. The state diagram is in `docs/architecture.md`.

## Imports

Core package. Imports `time` only. Must not import `engine` (which imports this package) or any adapter package. Signatures take no file paths, connections or clients. No network or disk access.

## API

- `type State string` with constants `Candidate`, `NotFound`, `Listed`, `ManualQueue`, `Sent`, `AwaitingConfirm`, `AwaitingResponse`, `Verify`, `Overdue`, `Removed` and `Escalated`.
- `type Request struct{}`: the request record for one broker.
- `type Event struct{}`: something that happened to a record. Every event is stored as evidence.
- `func Next(r Request, e Event, now time.Time) (State, error)`.

## To implement

- `Request` and `Event` fields.
- `Next`, covering every transition in the state diagram. `Escalated` is final.

## Open questions

- The stored form of `State`. Each constant's value is its own name for now.
- How a deadline that passes with no reply reaches `Next`, which requires an `Event`.
- How `Next` knows a broker's re-check interval. The interval is set per broker, but `Next` takes no `broker.Broker`.
- Which timestamp starts the 30-day clock: the submission, or the opened confirmation link.
- Where the access request to a broker with no public search fits. The state diagram does not show it.
- Whether a broker with several listings needs several records. The diagram has one record per broker; the operating rules send one request per listing.
