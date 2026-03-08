DEV_COMPOSE_FILE=docker-compose-dev.yml

.PHONY: compose-build-dev
compose-build-dev:
	docker compose -f $(DEV_COMPOSE_FILE) build

.PHONY: compose-up-dev
compose-up-dev:
	docker compose -f $(DEV_COMPOSE_FILE) up

.PHONY: compose-up-build-dev
compose-up-build-dev:
	docker compose -f $(DEV_COMPOSE_FILE) up --build

.PHONY: compose-down-dev
compose-down-dev:
	docker compose -f $(DEV_COMPOSE_FILE) down

PROD_COMPOSE_FILE=docker-compose-prod.yml

.PHONY: compose-build-prod
compose-build-prod:
	docker compose -f $(PROD_COMPOSE_FILE) build

.PHONY: compose-up-prod
compose-up-prod:
	docker compose -f $(PROD_COMPOSE_FILE) up

.PHONY: compose-up-build-prod
compose-up-build-prod:
	docker compose -f $(PROD_COMPOSE_FILE) up --build

.PHONY: compose-down-prod
compose-down-prod:
	docker compose -f $(PROD_COMPOSE_FILE) down

