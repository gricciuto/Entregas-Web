run:
	@echo "Ejecutando servidor Go, consultar logs en /logs/server.log"
	@mkdir -p  logs
	@go run main.go > logs/server.log 2>&1
	@echo "Se cerro el servidor Go"

test:
	@echo "Ejecutando tests..."
	go test -v ./...

build:
	@echo "Compilando binario..."
	go build -o bin/app main.go

start-docker-db:
	@echo "Levantando base de datos en Docker..."
	docker compose up -d
	@echo "Esperando que la base de datos esté lista..."
	until docker exec tp3_db pg_isready -U admin -d gorun; do \
    		echo "Esperando a Postgres..."; \
    		sleep 2; \
    	done
	docker compose ps

down-docker-db:
	@echo "Apagando base de datos..."
	docker compose down

start:
	@echo "Iniciando proyecto completo..."
	sqlc generate
	make start-docker-db
	make templ
	make run
	make stop

stop:
	@echo "Deteniendo contenedores..."
	make down-docker-db

hurl-test:
	@echo "Corriendo pruebas HURL"
	hurl --test hurl/requests.hurl

templ:
	templ generate