# log — documentation

togo logging: levels, text/json, file output

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

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `LOG_FILE"` |
| `LOG_FORMAT"` |
| `LOG_LEVEL"` |
| `LOG_SERVICE"` |

## Usage

```go
// Structured logs/errors forward to the configured sink automatically
// once this driver is installed and its env is set.
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/log
- Full README: ../README.md
