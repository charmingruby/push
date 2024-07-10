DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_USER ?= postgres
DATABASE_PASSWORD ?= postgres
DATABASE_SSL ?= disable
DATABASE_DATABASE = push-db
DATABASE_DSN := "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DATABASE}?sslmode=${DATABASE_SSL}"
MIGRATIONS_PATH="db/migrations"

DOCKER_IMAGE_NAME := push

###################
# Database        #
###################
.PHONY: mig-up
mig-up: ## Runs the migrations up
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} up

.PHONY: mig-down
mig-down: ## Runs the migrations down
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} down

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(NAME)

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 ${DOCKER_IMAGE_NAME}

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