# discover/probe

## Responsibility

A `Source` that searches broker sites for the operator's listings, using each definition's search URL pattern. In the system diagram, searches run in headless Chromium on the home server's residential connection.

## Imports

Adapter. Imports `engine`, `profile` and `broker`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Prober struct{}`
- `func (pr *Prober) Find(ctx context.Context, p profile.Profile, b broker.Broker) ([]engine.Listing, error)`
- `var _ engine.Source = (*Prober)(nil)`

## To implement

- `Prober` fields and `Find`.

## Open questions

- How a search result counts as a match, and how a near match is handled.
- Which Chromium library: `chromedp` or `rod`.
- What `Find` returns for a broker whose `search_url` is `none`.
- Whether this source or `discover/mailscan` is built first.
