# scrub

A self-hosted bot that finds the operator's listings on data broker and people-search sites, sends removal requests, and tracks each request until the listing is gone. It runs once a day on a home server. `scrub` is a working name for the binary.

## Status

Scaffold. Nothing is implemented. Every function and method body is `panic("not implemented")`, structs have no fields, and `go.mod` has no dependencies. The repository holds the package layout, the declarations and the documentation. `go build ./...` and `go vet ./...` pass. Simple Github actions are implemented.

## How a daily run works

A systemd timer starts `scrub` once a day. The secrets file, which holds the identity profile and mail credentials, is decrypted at start, and the broker definitions are loaded from `brokers/`. The engine reads and writes request state in SQLite. It polls the catch-all inbox for broker replies and confirmation links, opens confirmation links, and checks deadlines. When a broker is due for a check, a Source looks for a listing. When a listing is found, a Channel submits a request by email or web form, or the task goes to the manual queue. A run moves each request record at most one step, and every request, reply and timestamp is stored. The operator receives a digest and the manual queue. Most runs only read the inbox and check deadlines.

## Layout

| Package                      | Responsibility                                                                                                                                                        |
| ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `cmd/scrub`                  | The binary. Wires adapters into the engine and calls `Tick`. Planned subcommands: `run`, `scan`, `status`, `queue`, `import`. The only package that imports adapters. |
| `internal/engine`            | Daily tick. Picks the next action per broker. Declares the five ports in `ports.go`.                                                                                  |
| `internal/lifecycle`         | Request states and deadlines.                                                                                                                                         |
| `internal/broker`            | Broker definition schema and validation.                                                                                                                              |
| `internal/profile`           | The operator's identifiers and aliases.                                                                                                                               |
| `internal/letter`            | Letter templates. PIPEDA today.                                                                                                                                       |
| `internal/registry`          | Turns public broker lists into draft definitions.                                                                                                                     |
| `internal/discover/probe`    | Source. Searches broker sites.                                                                                                                                        |
| `internal/discover/mailscan` | Source. Scans the operator's mailboxes for companies that hold the operator's data.                                                                                   |
| `internal/channel/email`     | Channel. Sends letters over SMTP.                                                                                                                                     |
| `internal/channel/webform`   | Channel. Runs a definition's form steps in headless Chromium.                                                                                                         |
| `internal/channel/manual`    | Channel. Queues a task for the operator.                                                                                                                              |
| `internal/inbox/imap`        | Inbox. Finds confirmation links and classifies replies.                                                                                                               |
| `internal/store/sqlite`      | Store. Requests, events and evidence.                                                                                                                                 |
| `internal/notify`            | Notifier. Digest by push or email.                                                                                                                                    |
| `brokers/`                   | Broker definitions as YAML: `example.yaml`, and `brokers.yaml` with 633 unreviewed drafts. The format is a draft.                                                     |
| `docs/`                      | Architecture diagrams.                                                                                                                                                |

Core packages are `engine`, `lifecycle`, `broker`, `profile` and `letter`. They import no adapter package, and their signatures take no file paths, connections or clients. Each package directory has a README with its API and open questions.

## Operating rules

- The bot runs daily. It sends each request once per listing. Most runs only read the inbox and check deadlines. A request is re-sent only when a listing reappears.
- Verify first. Confirm a broker holds the operator's data before sending it more identifiers. A broker with no public search gets an access request from an alias with minimal identifiers.
- Each broker gets its own email alias on a domain the operator controls. Replies route back to the right request, and an alias that starts receiving spam shows who resold it.
- The server sits on a residential connection because people-search sites often block datacenter addresses.
- Broker definitions are data in the repository. Opt-out pages change often, and a fix should not need a rebuild.
- Every request, reply and timestamp is stored. That log is the evidence file for a complaint.
- ID uploads, phone calls and CAPTCHAs go to a manual queue for the operator. The bot never tries to solve or bypass them.

## Legal notes

These are the operator's working notes. They are not legal advice.

- The operator is a Canadian resident, and the governing statute is PIPEDA. It gives rights to access, to correct, and to withdraw consent, with a 30-day response clock.
- PIPEDA has no unqualified right to delete information an organization collected from someone other than the individual. Letters therefore ask for access, withdraw consent, and demand destruction of what is no longer needed.
- Escalation is a complaint to the Office of the Privacy Commissioner of Canada. Filing one is the operator's decision. The bot stops at `Escalated`.
- Bill C-36 was introduced in June 2026 and would add a deletion right. It is not law. Letter templates must be swappable for that reason.
- The operator is not a California resident. The bot never cites CCPA and never uses California's DROP platform. The California data broker registry is used only as a public list of brokers.

## Broker lists

Two public lists feed draft definitions. Both are reviewed by hand before anything becomes a definition.

- The California data broker registry CSV.
- The Big Ass Data Broker Opt-Out List on GitHub.

`brokers/brokers.yaml` holds 633 drafts generated from both lists on 2026-10-05. None has been reviewed. `brokers/README.md` describes its sources, conventions and license.

## What stays manual

- ID uploads, phone calls and CAPTCHAs. The bot queues them and sends the queue with the digest. The operator completes each one and marks it done.
- Reviewing imported broker lists before anything becomes a definition.
- Checking each definition against the live site. The date goes in the definition's `last_verified` field.
- Filing a complaint with the Office of the Privacy Commissioner of Canada after a request reaches `Escalated`.

## Candidate libraries

None of these is a dependency yet. `go.mod` has no `require` block.

| Need               | Candidate            |
| ------------------ | -------------------- |
| Driving Chromium   | `chromedp` or `rod`  |
| IMAP inbox         | `emersion/go-imap`   |
| SQLite without cgo | `modernc.org/sqlite` |
| Secrets file       | `filippo.io/age`     |
| Broker definitions | `gopkg.in/yaml.v3`   |

## Not designed yet

- Deployment: the systemd timer and the service it starts, and where the binary and database live. The repository has no unit files.
- Broker definition schema details. `brokers/README.md` is a draft.
- How the engine chooses a channel for a broker.
- Which discovery source is built first. `discover/probe` searches broker sites for the operator's identifiers. `discover/mailscan` reads the operator's mailboxes to list companies that already hold the operator's data, and a later mass-unsubscribe feature would reuse it.
- The mail provider, and how per-broker aliases are created and recorded.
- Gaps in the declared interfaces. These are listed in `internal/engine/README.md` and `internal/lifecycle/README.md`.

## Github actions

### Pull Request and Pushes

When a pull request or push is done to main the following github actions are run

```go
 go build -v ./...
 go vet -v ./...
 go test -v ./...

```

The user can run these same commands locally from the project repo to simulate the PR/Push requirements.

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

## Secrets

The identity profile and the mail credentials, encrypted with `age` and decrypted at start. Nothing in it is ever written to disk in the clear.

This is the file that would hurt. The bot exists to scrub your personal data off the internet, and to do that it has to keep a complete copy of it on a box in your house. Treat it accordingly.

### Editing

You do not open this file. No editor, no temp file, no decrypt-and-vim. The CLI owns it.

```sh
scrub secrets set smtp.password        # prompts, echo off
cat key | scrub secrets set imap.password
```

The value never goes in an argument. Arguments land in your shell history and in `/proc`, where anyone on the box can read them. Prompt or stdin, nothing else.

Writes are atomic. Decrypt in memory, change the field, re-encrypt, write beside the original, rename over it. A crash halfway through should cost nothing.

### Checking

You cannot read the file back, so the CLI tells you whether it works instead.

```sh
scrub secrets check
```

SMTP and IMAP get a real login. The alias domain gets an MX lookup. Those pass or they don't.

Your name and address get a format check and nothing more. Nobody can tell you whether you typed your own street right, which is why:

```sh
scrub secrets show
```

Credentials come back masked. Identity fields come back in full. A typo in your address means every broker rejects every request and none of them tell you why, so read it twice.

### The key

`secrets.age` is encrypted with `age`. The age identity that opens it is handed to the service by systemd as an encrypted credential.

```ini
[Service]
User=scrub
LoadCredentialEncrypted=age-identity:/etc/scrub/age-identity.cred
```

systemd encrypts that credential with the TPM and a host key, then drops the plaintext into a tmpfs only this service can read. Steal the disk and you get two files that open nothing.

Keep an offline copy of the age identity somewhere else. If the board dies the credential dies with it, and `secrets.age` is the only half you can restore on new hardware.

Swap goes encrypted or off. Decrypted secrets sitting in a swap file defeat the whole thing.
