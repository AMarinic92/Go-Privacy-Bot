# discover/mailscan

## Responsibility

A `Source` that reads the operator's mailboxes to list companies that already hold the operator's data. A later mass-unsubscribe feature would reuse it. Diagram 2 marks it "later".

## Imports

Adapter. Imports `engine`, `profile` and `broker`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Scanner struct{}`
- `func (s *Scanner) Find(ctx context.Context, p profile.Profile, b broker.Broker) ([]engine.Listing, error)`
- `var _ engine.Source = (*Scanner)(nil)`

## To implement

- `Scanner` fields and `Find`.

## Open questions

- Which mailboxes it reads, and where their credentials come from.
- Whether `Source.Find`, which runs per broker, fits a scan that covers all mail at once.
- How a company found in mail is matched to a broker definition.
- Whether this source or `discover/probe` is built first.
