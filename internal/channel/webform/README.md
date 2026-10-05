# channel/webform

## Responsibility

A `Channel` that submits a request by running a definition's form steps (`open`, `fill`, `click`, `wait_for`) in headless Chromium. ID uploads, phone calls and CAPTCHAs go to the manual queue. The runner never tries to solve or bypass them.

## Imports

Adapter. Imports `engine` and `lifecycle`. Must not import another adapter. Only `cmd/scrub` imports this package.

## API

- `type Runner struct{}`
- `func (w *Runner) Submit(ctx context.Context, r lifecycle.Request) (engine.Receipt, error)`
- `var _ engine.Channel = (*Runner)(nil)`

## To implement

- `Runner` fields and `Submit`.

## Open questions

- Which Chromium library: `chromedp` or `rod`.
- How the runner recognizes an ID upload, phone call or CAPTCHA and reports it, so the request goes to the manual queue.
- How `Submit` gets the broker's steps and the values for placeholders. It takes only a `lifecycle.Request`.
- What evidence of a form submission is stored.
