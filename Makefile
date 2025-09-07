# ===============================
# Project settings
# ===============================
APP_NAME := starter-app
DB_SERVICE := db
APP_SERVICE := app
COMPOSE_PROD := docker-compose.yml
COMPOSE_DEV := docker-compose.override.yml
ENV_FILE_PROD := .env
ENV_FILE_DEV := .env.dev

# ===============================
# 🐳 Build
# ===============================

docker-build:
	docker build -f Dockerfile.prod -t $(APP_NAME):latest .

docker-build-dev:
	docker build -f Dockerfile.dev -t $(APP_NAME):dev .

# ===============================
# 🚀 Start services
# ===============================

up:
	docker-compose -f $(COMPOSE_PROD) up -d

up-dev:
	docker-compose -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) up -d

up-app:
	docker-compose -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) up -d $(APP_SERVICE)

# ===============================
# 🔽 Stop services
# ===============================

down:
	docker-compose -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) down

down-clean:
	docker-compose -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) down -v --remove-orphans

# ===============================
# 📥 Pull / 📤 Push
# ===============================

docker-pull:
	docker-compose -f $(COMPOSE_PROD) pull

docker-push:
	docker push $(APP_NAME):latest

# ===============================
# 🐳 Logs / Restart / Exec
# ===============================

logs:
	docker-compose -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) logs -f

restart-app:
	docker-compose -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) restart $(APP_SERVICE)

exec-app:
	docker-compose -f $(COMPOSE_PROD) -f $(COMPOSE_DEV) exec $(APP_SERVICE) sh

# ===============================
# Phony targets
# ===============================
.PHONY: docker-build docker-build-dev up up-dev up-app down down-clean docker-pull docker-push logs restart-app exec-app
