# cmd/scrub

## Responsibility

The `scrub` binary. `scrub` is a working name. This is the only package that imports adapters. It wires them into the engine and calls `Tick`.

## Imports

May import any package in this module. Nothing imports it.

## API

- `func main()`. A comment lists the planned subcommands: `run`, `scan`, `status`, `queue` and `import`.

## To implement

- Subcommand dispatch.
- Wiring each adapter into the `engine.Engine` ports.
- The `import` command, which uses `registry.ImportCSV`.

## Open questions

- What `run`, `scan`, `status` and `queue` do. Diagram 2 links only `import` (to `registry`) and the call to `Tick`, and does not say which subcommand calls `Tick`.
- Where `import` reads its CSV and where it writes drafts.
- Config loading, secrets handling and deployment. See the root README.

## Config

### Example

The scrubber requires a configuration to be set with appropriate values. The config is a YAML file with the following example which also exists in the repo.

```yaml
database: /var/lib/scrub/scrub.db
brokers: /etc/scrub/brokers
secrets: /etc/scrub/secrets.age

smtp: { host: smtp.example.net, port: 587 }
imap: { host: imap.example.net, port: 993 }

alias_domain: aliases.example.net
digest_to: you@example.net
```

### Location

The config should be in the `/etc/scrub/` folder, but it can be overrideen with a flag.
