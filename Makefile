.PHONY: up down build logs restart clean \
	up-auth up-inventory up-purchasing up-gateway \
	build-auth build-inventory build-purchasing build-gateway \
	restart-auth restart-inventory restart-purchasing restart-gateway \
	stop-auth stop-inventory stop-purchasing stop-gateway \
	logs-auth logs-inventory logs-purchasing logs-gateway \
	fresh swag

# =========================
# GLOBAL
# =========================

up:
	docker compose up -d --build

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

restart:
	docker compose restart

clean:
	docker compose down -v --remove-orphans

fresh:
	make down && make up

# =========================
# AUTH SERVICE
# =========================

up-auth:
	docker compose up -d auth-service

build-auth:
	docker compose build auth-service

restart-auth:
	docker compose restart auth-service

stop-auth:
	docker compose stop auth-service

logs-auth:
	docker compose logs -f auth-service

# =========================
# INVENTORY SERVICE
# =========================

up-inventory:
	docker compose up -d inventory-service

build-inventory:
	docker compose build inventory-service

restart-inventory:
	docker compose restart inventory-service

stop-inventory:
	docker compose stop inventory-service

logs-inventory:
	docker compose logs -f inventory-service

# =========================
# PURCHASING SERVICE
# =========================

up-purchasing:
	docker compose up -d purchasing-service

build-purchasing:
	docker compose build purchasing-service

restart-purchasing:
	docker compose restart purchasing-service

stop-purchasing:
	docker compose stop purchasing-service

logs-purchasing:
	docker compose logs -f purchasing-service

# =========================
# API GATEWAY
# =========================

up-gateway:
	docker compose up -d api-gateway

build-gateway:
	docker compose build api-gateway

restart-gateway:
	docker compose restart api-gateway

stop-gateway:
	docker compose stop api-gateway

logs-gateway:
	docker compose logs -f api-gateway

# =========================
# SWAGGER (AUTH SERVICE)
# =========================

swag:
	docker exec -it auth-service swag init -g cmd/api/main.go
