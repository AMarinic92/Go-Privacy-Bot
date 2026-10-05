# broker

## Responsibility

Broker definition schema and validation. A definition describes one data broker or people-search site: how to check for a listing, which channel to use, and how to submit. Definitions are YAML files under `brokers/`. The format is a draft described in `brokers/README.md`.

## Imports

Core package. Imports nothing today. Must not import `engine` or any adapter package. Signatures take no file paths, connections or clients, so `Parse` takes bytes. No network or disk access.

## API

- `type Broker struct{}`
- `func Parse(data []byte) (Broker, error)`
- `func (b Broker) Validate() error`

## To implement

- `Broker` fields, once the definition format is settled.
- `Parse`. Decoding YAML needs a library. `gopkg.in/yaml.v3` is the candidate. No dependency is added yet.
- `Validate`.

## Open questions

- The definition format. Every field in `brokers/README.md` is a proposal.
- Which package reads definition files from disk. Core packages cannot take file paths, and only `cmd/scrub` may import adapters.
- Whether `Parse` also validates, or callers must call `Validate`.
- How a multi-document stream is read. `brokers/brokers.yaml` holds 633 documents, and `Parse` returns one `Broker`.
- What a stale `last_verified` date does, if anything.
