# Broker definitions

**Draft.** This format is a proposal. Nothing parses it yet, and any field name, type or rule below may change.

Definitions are data, not code: opt-out pages change often, and a fix to a definition should not need a rebuild. `internal/broker` will parse and validate them. Every definition is reviewed by hand before use.

## Files

| File | Contents |
|---|---|
| `example.yaml` | One fictional broker on `example.com`, showing every field. |
| `brokers.yaml` | 633 real brokers as unreviewed drafts, in one multi-document YAML stream. Each document is one definition with the fields below. |

## brokers.yaml

Generated on 2026-10-05 from the two broker lists. Nothing in it has been checked by hand.

- Section 1 has 47 people-search definitions from the Big Ass Data Broker Opt-Out List (BADBOOL), updated 2026-09-27, in BADBOOL's priority order. 26 of them are matched to their company's registry entry, which supplies jurisdiction and contact.
- Section 2 has the other 586 brokers in the 2026 California data broker registry, last updated 2026-07-29. The 17 registry rows behind section 1 are merged there.

Comments in each entry hold what the sources say: the opt-out procedure, other routes, registered websites, location and privacy pages, and the broker's own registry answers about what it collects and who it sells to. The file header lists the conventions:

- `last_verified: null` marks an unreviewed draft. Every entry has it.
- `null` means no source gives the value. `recheck_days` is null on every entry.
- `search_url: none` in section 2 means neither source names a public search. It was not checked against the site.
- A `search_url` marked `TODO` is a search page or site root, not a query pattern.
- `webform` entries hold only the `open` step. The other steps must be written by hand against the live page.
- `jurisdiction` is the ISO 3166-1 alpha-2 code of the registered address's country. `CA` is Canada.
- Nine registry contacts are California DROP mailboxes. They are left null, because the bot never uses DROP.

BADBOOL is licensed CC BY-NC-SA. Section 1 adapts it and is shared under the same license. The registry is used only as a public list of brokers. Not used: BADBOOL's search engine, special circumstances and identity theft sections, and the registry's earlier editions.

## Fields

| Field | Required | Meaning |
|---|---|---|
| `id` | yes | Stable identifier. Lowercase letters, digits and hyphens. Unique across definitions. |
| `name` | yes | The broker's display name. |
| `jurisdiction` | yes | Where the broker operates. |
| `search_url` | yes | URL pattern for the broker's public search, with `{placeholders}` for profile fields. The literal `none` means the broker has no public search. Such a broker gets an access request first, from an alias, with minimal identifiers. |
| `channel` | yes | How a request is sent: `email`, `webform` or `manual`. |
| `privacy_contact` | when `channel` is `email` or `search_url` is `none` | Email address of the broker's privacy contact. |
| `steps` | when `channel` is `webform` | Ordered list of form steps. See below. |
| `recheck_days` | yes | Days between re-checks of the broker. |
| `last_verified` | yes | Date, as `YYYY-MM-DD`, on which the definition was last checked by hand against the live site. |

## Form steps

Each step is a map with one key. Steps run in order.

| Step | Value | Action |
|---|---|---|
| `open` | URL | Load the page. |
| `fill` | `selector` and `value` | Type `value` into the element that matches the CSS selector. `value` may contain placeholders. |
| `click` | CSS selector | Click the element. |
| `wait_for` | CSS selector | Wait until the element is present. |

## Placeholders

`example.yaml` uses `{full_name}` and `{region}` in `search_url`, and `{alias}` and `{listing_url}` in form steps. `brokers.yaml` uses `{last_name}` and `{first_name}` once, for PeopleByName. These names are illustrations, not a fixed list.

## Open questions

- Placeholder syntax, and which profile fields a definition may reference.
- What the engine uses `jurisdiction` for. `brokers.yaml` uses the proposed country code.
- One file for all brokers or one file per broker. `brokers.yaml` uses one stream. If files stay separate, whether a file's name must equal its `id`.
- Whether one company with several sites is one broker or several. `brokers.yaml` has one entry per site, so seven sites share Mississippi Tornado Alley's contact, and BeenVerified's people and property opt-outs are two entries.
- Whether form steps need more actions, such as selecting an option or a timeout on `wait_for`.
- Whether a definition declares that a step needs the operator (ID upload, phone call, CAPTCHA), or the runner only detects it at run time.
- Whether a definition declares that the broker sends a confirmation link. ZoomInfo sends a code instead.
- The unit of the re-check interval. `recheck_days` is a proposal.
- Whether a null `last_verified` is the marker for an unreviewed draft, as proposed.
- How `brokers.yaml` is regenerated. The one-off script that wrote it is not in the repository. `registry.ImportCSV` is the planned replacement for the registry part.
