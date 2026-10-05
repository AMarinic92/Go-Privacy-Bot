# channel/manual

## Responsibility

A `Channel` that queues a task for the operator instead of contacting the broker. It covers ID uploads, phone calls and CAPTCHAs. The queue goes to the operator with the digest. The operator does the step and marks it done, which moves the record from `ManualQueue` to `Sent`.

## Imports

Adapter. Imports `engine` and `lifecycle`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Queue struct{}`
- `func (q *Queue) Submit(ctx context.Context, r lifecycle.Request) (engine.Receipt, error)`
- `var _ engine.Channel = (*Queue)(nil)`

## To implement

- `Queue` fields and `Submit`.

## Open questions

- Where queued tasks are kept. `engine.Store` has no method for them.
- How the operator marks a task done. The `queue` subcommand is planned but not specified.
- What `Receipt` means for a queued task.
