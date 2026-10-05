# store/sqlite

## Responsibility

A `Store` backed by a local SQLite database. It holds requests, events and evidence: every request, reply and timestamp. That log is the evidence file for a complaint. Database files are git-ignored.

## Imports

Adapter. Imports `engine` and `lifecycle`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Store struct{}`
- `func (s *Store) Due(ctx context.Context, now time.Time) ([]lifecycle.Request, error)`
- `func (s *Store) Save(ctx context.Context, r lifecycle.Request, e lifecycle.Event) error`
- `var _ engine.Store = (*Store)(nil)`

## To implement

- `Store` fields, `Due` and `Save`. `modernc.org/sqlite` is the candidate driver because it needs no cgo.

## Open questions

- Schema and migrations.
- Where the database file lives. Config loading is not designed.
- What makes a record due at `now`: a deadline, a re-check, or new mail.
- `engine.Store` declares only `Due` and `Save`. The `status` and `queue` subcommands will need reads it does not declare.
