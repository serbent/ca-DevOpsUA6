# Define the variable for docker compose
DOCKER_COMPOSE = docker compose

# Default target: show available options (using just)
.PHONY: default
default:
	@just -l

# Stop services with arguments
.PHONY: stop
stop:
	@$(DOCKER_COMPOSE) stop $(args)

# Stop all services
.PHONY: stopall
stopall:
	@$(DOCKER_COMPOSE) stop

# Run services with arguments
.PHONY: run
run:
	@$(DOCKER_COMPOSE) up $(args) -d

# Run all services
.PHONY: runall
runall:
	@$(DOCKER_COMPOSE) up -d

# Pull the latest images
.PHONY: pull
pull:
	@$(DOCKER_COMPOSE) pull

# Restart services with arguments
.PHONY: restart
restart:
	@$(DOCKER_COMPOSE) restart $(args)

# Restart all services
.PHONY: restartall
restartall:
	@$(DOCKER_COMPOSE) restart

# Show logs with arguments
.PHONY: logs
logs:
	@$(DOCKER_COMPOSE) logs $(args)

# Show all logs
.PHONY: logsall
logsall:
	@$(DOCKER_COMPOSE) logs

# Build service(s) and run with arguments
.PHONY: build
build:
	@$(DOCKER_COMPOSE) up $(args) --build -d

# Build all services and run
.PHONY: buildall
buildall:
	@$(DOCKER_COMPOSE) up --build -d

# Show running containers
.PHONY: ps
ps:
	@$(DOCKER_COMPOSE) ps

