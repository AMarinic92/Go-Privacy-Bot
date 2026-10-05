# letter

## Responsibility

Renders the letters sent to brokers. Templates cite PIPEDA today. They must be swappable, because Bill C-36 would add a deletion right if it becomes law. Letters never cite CCPA.

- `Access` asks what personal information a broker holds. A broker with no public search gets one first, from an alias, with minimal identifiers.
- `Removal` asks for access, withdraws consent, and demands destruction of what is no longer needed.
- `Reminder` follows up once when 30 days pass without a reply.

## Imports

Core package. Imports `profile` and `broker`. Must not import `engine` or any adapter package. Signatures take no file paths, connections or clients. No network or disk access.

## API

- `type Kind string` with constants `Access`, `Removal` and `Reminder`.
- `type Letter struct{}`
- `func Render(k Kind, p profile.Profile, b broker.Broker) (Letter, error)`

## To implement

- `Letter` fields.
- The templates and `Render`.

## Open questions

- How templates are stored and swapped. This package cannot take a file path.
- Whether `Render` needs the `lifecycle.Request`. A reminder refers to the original request and its date.
- Which identifiers each kind discloses before a broker is confirmed to hold the operator's data.
- The stored form of `Kind`. Each constant's value is its own name for now.
