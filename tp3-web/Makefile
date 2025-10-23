run:
	@echo "Ejecutando servidor Go, consultar logs en /logs/server.log"
	@mkdir logs
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
	until sudo docker exec tp3_db pg_isready -U admin -d gorun; do \
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
	$(MAKE) start-docker-db
	$(MAKE) test
	$(MAKE) run
	$(MAKE) stop

stop:
	@echo "Deteniendo contenedores..."
	$(MAKE) down-docker-db

hurl-test:
	@echo "Corriendo pruebas HURL"
	hurl --test hurl/requests.hurl
