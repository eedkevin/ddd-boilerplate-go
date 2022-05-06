.PHONY: install
install:
	go install golang.org/x/lint/golint@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	brew install pre-commit
	pre-commit install

.PHONY: test
test:
	go test -v ./...

.PHONY: up
up:
	go run cmd/app/main.go

.PHONY: dev
dev:
	docker compose up watch-app

.PHONY: infra-up
infra-up:
	docker compose up redis-cluster redis-commander -d

.PHONY: infra-down
infra-down:
	docker compose down redis-cluster redis-commander
