# profile

## Responsibility

The operator's identifiers and per-broker email aliases. The identity profile is kept in the encrypted secrets file and decrypted at start. This package defines its shape, not how it is stored.

## Imports

Core package. Imports nothing. Must not import `engine` or any adapter package. Signatures take no file paths, connections or clients. No network or disk access.

## API

- `type Profile struct{}`

## To implement

- `Profile` fields.

## Open questions

- Which identifiers the profile holds.
- How a `Profile` is built from the decrypted secrets file. Secrets handling is not designed.
- How each broker's alias is created and recorded.
