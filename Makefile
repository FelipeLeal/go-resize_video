# Set 'help' as the default command when running 'make'
.DEFAULT_GOAL := help

# Use docker-compose, but can be easily switched to 'docker compose' (the new CLI plugin)
COMPOSE = docker-compose

# Define the primary service name for commands.
APP_SERVICE_NAME ?= resize

.PHONY: help build up down start stop restart logs ps shell prune

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the Docker image
	$(COMPOSE) build

up: ## Start the application in detached mode
	$(COMPOSE) up

down: ## Stop and remove containers
	$(COMPOSE) down

start: ## Start the containers
	$(COMPOSE) start

stop: ## Stop the containers
	$(COMPOSE) stop

restart: ## Restart the application
	$(COMPOSE) restart

logs: ## Show logs from the application
	$(COMPOSE) logs -f $(APP_SERVICE_NAME)

ps: ## Show running containers
	$(COMPOSE) ps

shell: ## Access the container shell
	$(COMPOSE) exec $(APP_SERVICE_NAME) /bin/sh

prune: ## Remove unused Docker resources
	docker system prune -f
	docker volume prune -f

# Additional development commands
dev-build: ## Build and start in development mode
	$(COMPOSE) up --build

dev-logs: ## Follow logs in development
	$(COMPOSE) logs -f

clean: ## Stop containers and remove images
	$(COMPOSE) down --rmi all --volumes --remove-orphans

rebuild: ## Rebuild from scratch
	$(COMPOSE) down --rmi all --volumes --remove-orphans
	$(COMPOSE) up --build
