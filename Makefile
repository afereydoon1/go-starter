# Project settings
APP_NAME := starter-app
DB_SERVICE := db
APP_SERVICE := app
COMPOSE_FILE := docker-compose.yml
ENV_FILE := .env

#---------------------------------------------------------------
# 🐳 Build production image
#---------------------------------------------------------------
docker-build:
	docker build -f Dockerfile.prod -t $(APP_NAME):latest .

#---------------------------------------------------------------
# 🚀 Start all services (DB + App + Admin)
#---------------------------------------------------------------

up:
	docker-compose -f $(COMPOSE_FILE) up -d
#---------------------------------------------------------------
# 🔄 Start only the application (useful for dev or testing)
#---------------------------------------------------------------
up-app:
	docker-compose -f $(COMPOSE_FILE) up -d $(APP_SERVICE)

#---------------------------------------------------------------
# 🔽 Stop all services
#---------------------------------------------------------------
down:
	docker-compose -f $(COMPOSE_FILE) down

#---------------------------------------------------------------
# 🧹 Stop services and remove volumes (clean DB and data)
#---------------------------------------------------------------
down-clean:
	docker-compose -f $(COMPOSE_FILE) down -v --remove-orphans

#---------------------------------------------------------------
# 📥 Pull the latest images
#---------------------------------------------------------------
docker-pull:
	docker-compose -f $(COMPOSE_FILE) pull

#---------------------------------------------------------------
# 📤 Push the image to Docker Hub or registry
#---------------------------------------------------------------
docker-push:
	docker push $(APP_NAME):latest

#---------------------------------------------------------------
# 🐳 View logs in real-time
#---------------------------------------------------------------
logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

#---------------------------------------------------------------
# 🔄 Restart a service quickly
#---------------------------------------------------------------
restart-app:
	docker-compose -f $(COMPOSE_FILE) restart $(APP_SERVICE)

#---------------------------------------------------------------
# 🧪 Execute commands inside the app container
#---------------------------------------------------------------
exec-app:
	docker-compose -f $(COMPOSE_FILE) exec $(APP_SERVICE) sh

.PHONY: docker-build up up-app down down-clean docker-pull docker-push logs restart-app exec-app
