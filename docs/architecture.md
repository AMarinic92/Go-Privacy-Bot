# Architecture

Three diagrams: the system, the code, and the request lifecycle. They describe the planned design. Nothing in them is implemented yet.

## 1. System

```mermaid
flowchart TB
    you(["You"])

    subgraph home["Home server on a residential IP"]
        timer["systemd timer<br/>fires once a day"]
        bot["scrub<br/>single Go binary"]
        db[("SQLite<br/>requests, events, evidence")]
        defs[/"Broker definitions<br/>YAML in a git repo"/]
        vault["Encrypted secrets<br/>identity profile, mail credentials"]
        chrome["Headless Chromium"]
    end

    subgraph mail["Mail provider on your own domain"]
        smtp["SMTP<br/>one alias per broker"]
        imap["IMAP<br/>catch-all inbox"]
    end

    subgraph brokers["Brokers"]
        search["People-search sites<br/>search pages and opt-out forms"]
        priv["Privacy contacts<br/>registry brokers with no public search"]
    end

    feeds["Broker lists<br/>California registry CSV, curated opt-out list"]
    opc["Privacy commissioner"]

    timer -->|"starts the daily run"| bot
    vault -->|"decrypted at start"| bot
    defs -->|"loaded each run"| bot
    feeds -.->|"imported, then reviewed by you"| defs
    bot <-->|"reads and writes state"| db
    bot -->|"drives"| chrome
    chrome -->|"searches for your listing, submits forms"| search
    bot -->|"sends access and deletion letters"| smtp
    smtp --> priv
    priv -->|"replies"| imap
    search -->|"confirmation links"| imap
    imap -->|"polled for links and replies"| bot
    bot <-->|"digest and manual queue out, finished ID, phone and CAPTCHA steps back"| you
    you -.->|"complaint with the evidence log"| opc
```

One binary on a home server talks to brokers through a browser and a mailbox, and reads their answers back from the same mailbox. Dotted arrows are steps I do by hand.

Notes:

- The binary is built from `cmd/scrub`. The systemd timer belongs to deployment, which is not designed. The repository has no unit files.
- The server uses a residential connection because people-search sites often block datacenter addresses.
- Each broker gets its own alias on a domain the operator controls. Replies to an alias route back to the right request. An alias that starts receiving spam shows who resold it.
- The secrets file holds the identity profile and mail credentials. Its format and key handling are not designed. `filippo.io/age` is the candidate library.
- Broker definitions are data in the repository and are loaded each run, so a fix to a definition needs no rebuild.
- Broker lists are imported by `internal/registry` and become definitions only after review by hand.
- SQLite holds every request, reply and timestamp. That log is the evidence file for a complaint.
- Filing a complaint with the Office of the Privacy Commissioner of Canada is the operator's decision.

## 2. Code

```mermaid
flowchart LR
    cmd["cmd/scrub<br/>run, scan, status, queue, import"]
    registry["registry<br/>turns broker lists into YAML drafts"]

    subgraph core["Core packages, no network or disk access"]
        engine["engine<br/>daily tick, picks the next action per broker"]
        lifecycle["lifecycle<br/>request states and deadlines"]
        broker["broker<br/>definition schema and validation"]
        profile["profile<br/>your identifiers and aliases"]
        letter["letter<br/>legal templates, PIPEDA today"]
    end

    subgraph ports["Interfaces the engine depends on"]
        source{{"Source"}}
        channel{{"Channel"}}
        inbox{{"Inbox"}}
        store{{"Store"}}
        notifier{{"Notifier"}}
    end

    subgraph adapters["Adapters, one package per outside system"]
        probe["discover/probe<br/>searches broker sites"]
        mailscan["discover/mailscan<br/>scans your mailboxes, later"]
        email["channel/email<br/>SMTP"]
        webform["channel/webform<br/>runs YAML steps in Chromium"]
        manual["channel/manual<br/>queues a task for you"]
        imap["inbox/imap<br/>finds links, classifies replies"]
        sqlite["store/sqlite"]
        push["notify<br/>digest by push or email"]
    end

    cmd -->|"wires adapters, calls Tick"| engine
    cmd -->|"import command"| registry
    registry -->|"writes drafts for"| broker
    engine --> lifecycle
    engine --> broker
    engine --> profile
    engine --> letter
    engine -->|"find listings"| source
    engine -->|"submit request"| channel
    engine -->|"fetch replies"| inbox
    engine -->|"load and save"| store
    engine -->|"report"| notifier
    source -.-> probe
    source -.-> mailscan
    channel -.-> email
    channel -.-> webform
    channel -.-> manual
    inbox -.-> imap
    store -.-> sqlite
    notifier -.-> push
```

The engine holds every decision and imports no adapter. Solid arrows are calls. Dotted arrows point from an interface to the packages that implement it. In the repo these packages live under `internal/`.

Notes:

- Core packages are `engine`, `lifecycle`, `broker`, `profile` and `letter`. They import no adapter package, and their signatures take no file paths, connections or clients.
- The five interfaces are declared in `internal/engine/ports.go`. Each adapter declares one struct and a compile-time assertion that the struct satisfies its interface.
- Only `cmd/scrub` imports adapters, so adapters do not import one another.
- `registry` sits outside the core. It imports `broker` and reads from an `io.Reader`.
- `discover/mailscan` is marked "later". Which discovery source is built first is not decided.
- How the engine picks among the three channels is not decided.
- The diagram places reply classification in `inbox/imap`. Whether that decision belongs in the engine is an open question in `internal/engine/README.md`.

## 3. Request lifecycle

```mermaid
stateDiagram-v2
    state "Not found" as NotFound
    state "Manual queue" as ManualQueue
    state "Awaiting confirmation" as AwaitingConfirm
    state "Awaiting response" as AwaitingResponse
    state "Verify removal" as Verify

    [*] --> Candidate: broker definition loaded
    Candidate --> Listed: Source finds a match
    Candidate --> NotFound: no match
    NotFound --> Candidate: re-check due
    Listed --> Sent: Channel submits by email or form
    Listed --> ManualQueue: needs ID, phone or CAPTCHA
    ManualQueue --> Sent: you mark it done
    Sent --> AwaitingConfirm: broker emails a link
    AwaitingConfirm --> AwaitingResponse: bot opens the link
    Sent --> AwaitingResponse: no confirmation step
    AwaitingResponse --> Verify: reply says removed
    AwaitingResponse --> Overdue: 30 days, no reply
    Overdue --> AwaitingResponse: one reminder sent
    Overdue --> Escalated: reminder also ignored
    Verify --> Removed: listing gone
    Verify --> Overdue: still listed
    Removed --> Listed: listing reappears at re-check
    Escalated --> [*]: complaint filed by you
```

Each broker has one request record. A daily run moves a record at most one step. The 30 days is PIPEDA's response clock. The re-check interval is set per broker.

Notes:

- Each state maps to a `lifecycle.State` constant:

  | Diagram label | Constant |
  |---|---|
  | Candidate | `Candidate` |
  | Not found | `NotFound` |
  | Listed | `Listed` |
  | Manual queue | `ManualQueue` |
  | Sent | `Sent` |
  | Awaiting confirmation | `AwaitingConfirm` |
  | Awaiting response | `AwaitingResponse` |
  | Verify removal | `Verify` |
  | Overdue | `Overdue` |
  | Removed | `Removed` |
  | Escalated | `Escalated` |

- A request is sent once per listing. It is re-sent only when a listing reappears, which is the `Removed` to `Listed` transition.
- `Overdue` allows one reminder, rendered as `letter.Reminder`. If the reminder is also ignored, the record moves to `Escalated`.
- `Escalated` is final for the bot. The complaint is filed by the operator.
- The diagram does not show the access request sent to a broker with no public search. This is an open question in `internal/lifecycle/README.md`.
