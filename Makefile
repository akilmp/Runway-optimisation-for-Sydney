.PHONY: dev ci

dev:
docker compose up --build

ci:
go fmt ./...
npm --prefix frontend install
npm --prefix frontend run type-check
go test ./...
