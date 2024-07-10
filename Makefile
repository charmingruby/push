###################
# App             #
###################
.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: docs
docs:
	swag init -g swagger.go -d ./internal/infra -o ./docs --parseDependency --parseInternal

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/push ./cmd/api/main.go

###################
# Utils           #
###################
.PHONY: clear-notes
clear-notes:
	find . -type f -name "*_notes.md" -exec rm -f {} \;