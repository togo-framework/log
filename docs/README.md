# log — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package log is togo's configurable logging plugin: levels, text/JSON format,
and file output. It overrides the kernel's baseline logger. Error trackers
(Sentry, GlitchTip, …) ship as separate plugins that subscribe to the kernel
"error" hook — this package only configures the slog sink.

Install: `togo install togo-framework/log` (blank-import registers it).

## Install

```bash
togo install togo-framework/log
```

Set `LOG_DRIVER=<provider>` and install a driver (log-sentry, …).

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `LOG_FILE` | _see provider docs_ |
| `LOG_FORMAT` | _see provider docs_ |
| `LOG_LEVEL` | _see provider docs_ |
| `LOG_SERVICE` | _see provider docs_ |

## Usage

```go
// Structured logs/errors are forwarded to the configured sink automatically
// once this driver is installed and its env is set.
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/log
- README: ../README.md
