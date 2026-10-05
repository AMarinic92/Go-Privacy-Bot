# registry

## Responsibility

Turns public broker lists into draft broker definitions. Every draft is reviewed by hand before it becomes a file under `brokers/`. The two lists are the California data broker registry CSV and the Big Ass Data Broker Opt-Out List on GitHub. The California registry is used only as a public list of brokers.

## Imports

Not a core package and not an adapter. Imports `io` and `broker`. Must not import `engine` or any adapter package. Takes an `io.Reader`, not a file path.

## API

- `func ImportCSV(r io.Reader) ([]broker.Broker, error)`

## To implement

- `ImportCSV`.

## Source notes

The 2026 registry CSV (`https://cppa.ca.gov/data_broker_registry/registry.csv`) has 603 rows and 77 columns. `brokers/brokers.yaml` used the name, DBA, website, contact email, city, state, country and privacy page columns, plus the yes/no answers on data collected and sold. Multi-value cells are separated by `;`. Many websites have no scheme. The DBA column holds placeholders such as `NA`, `None` and `Yes`. Nine contact addresses are DROP mailboxes.

## Open questions

- Which columns `ImportCSV` maps, and how it handles the cases in the source notes.
- Who writes the YAML drafts. `ImportCSV` returns `broker.Broker` values, and `broker` declares no encoder.
- Whether the curated opt-out list gets an importer. `ImportCSV` covers CSV input only.
- How a draft is marked as not yet reviewed, so it is never used before review.
