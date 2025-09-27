# ===============================
# Project settings
# ===============================
APP_NAME := starter-app
DB_SERVICE := db
APP_SERVICE := app

COMPOSE_PROD := deployment/docker-compose.yml
COMPOSE_DEV := deployment/docker-compose.override.yml
ENV_FILE := .env

DOCKERFILE_DEV := deployment/docker/app/Dockerfile.dev
DOCKERFILE_PROD := deployment/docker/app/Dockerfile.prod

# ===============================
# 🐳 Build
# ===============================

docker-build:
	docker build -f $(DOCKERFILE_PROD) -t $(APP_NAME):latest .

docker-build-dev:
	docker build -f $(DOCKERFILE_DEV) -t $(APP_NAME):dev .

# ===============================
# 🚀 Start services
# ===============================

up:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) up -d

up-dev:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) up -d

up-app:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) up -d $(APP_SERVICE)

# ===============================
# 🔽 Stop services
# ===============================

down:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) down

down-clean:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) down -v --remove-orphans
# ===============================
# 🐳 Logs / Restart / Exec
# ===============================

logs:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) logs -f

restart-app:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) restart $(APP_SERVICE)

exec-app:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) exec $(APP_SERVICE) sh

# ===============================
# ✅ Test
# ===============================
test:
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) exec $(APP_SERVICE) go test -v ./...
# ===============================
# Phony targets
# ===============================
.PHONY: docker-build docker-build-dev up up-dev up-app down down-clean logs restart-app exec-app
