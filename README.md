# syd-perf

Living under runway 34L sparked my obsession with Sydney's buzzing skies. This tool started as a weekend project to decode takeoff performance and evolved into a full-stack playground.

## What it does

`syd-perf` ingests live YSSY METAR, TAF and NOTAMs every minute and stores them in Postgres. The frontend lets you arrange passengers in a 737‑800 seat map and request a take‑off calculation. The backend responds with runway recommendations, V‑speeds, fuel burn deltas and an engine‑out SID line. Everything is strictly for enthusiasts and **NOT FOR OPERATIONAL DISPATCH**.

## How it works

- **Ingest service** – cron job in Go that pulls BoM Aviation JSON and Sydney NOTAM RSS. Parsed results are saved via simple SQL helpers.
- **Performance engine** – reads digitised CSV tables to compute balanced field length and speed settings. Adjusts OAT for the sea‑breeze bias described in `internal/perf/adjust.go`.
- **API** – Gin router exposing `/v1` endpoints. JWT auth uses a single demo secret. Handlers live under `backend/api`.
- **Frontend** – React + TypeScript application served by Vite. Redux Toolkit stores the seat map, while React Query hooks call the API.
- **Infrastructure** – Docker Compose for local dev, Kubernetes manifests and Terraform for AWS deployment, plus a Grafana dashboard.

## Repository layout

```
.
├── backend/    # Go API, ingest and perf logic
├── frontend/   # React client
├── infra/      # Terraform and k8s manifests
├── grafana/    # Dashboard JSON
└── docs/       # Architecture diagrams
```

Each directory is self‑contained with its own Dockerfile or build config. `Makefile` shortcuts let you run `make dev` to start the full stack locally.

## Design decisions

- **Gin + pgx** offer lightweight routing and efficient Postgres access without heavy frameworks.
- **Redux Toolkit** keeps the cabin manifest predictable and easily testable.
- **Multi‑stage Dockerfiles** minimise final image size and mirror production.
- **Terraform + k8s** provide an opinionated yet simple path to run on AWS EKS with RDS Postgres.

## Quick Start

```bash
make dev
```

Visit <http://localhost:5173> to use the app.

## Architecture

```
        +-------+       +---------+       +---------+
        | Front | <---> |  API    | <---> |  DB     |
        +-------+       +---------+       +---------+
```

High‑res diagram: [docs/architecture.png](docs/architecture.png)

Postman collection: <https://example.com/syd-perf.postman.json>

## FAQ

**Why does OAT sometimes jump by 7°C?** Sea‑breeze bias applies on summer afternoons. See `internal/perf/adjust.go`.

**Which runways close at night?** Curfew prevents 34L/R and 16L/R between 23:00–06:00 local unless exempted.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)
