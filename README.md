# syd-perf

Living under runway 34L sparked my obsession with Sydney's buzzing skies. This tool
started as a weekend project to decode takeoff performance and evolved into a
full-stack playground.

## Quick Start

```bash
make dev
```



## Architecture

```
        +-------+       +---------+       +---------+
        | Front | <---> |  API    | <---> |  DB     |
        +-------+       +---------+       +---------+
```

## FAQ

**Why does OAT sometimes jump by 7°C?** Sea-breeze bias applies on summer
afternoons. See `internal/perf/adjust.go`.

**Which runways close at night?** Curfew prevents 34L/R and 16L/R between
23:00–06:00 local unless exempted.

