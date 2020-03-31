#
# INTERNAL VARIABLES
#	note: should export env values form .bashrc/.zshrc
#				- REMOTE_IP
# 			- REMOTE_USER
#				- DOCKER_USER
#				- DOCKER_PASS
#
VERSION=0.0.3
LAST_VERSION=0.0.2
NAME=uluru
SVC=$(NAME)-api
BIN_PATH=$(PWD)/bin
BIN=$(BIN_PATH)/$(SVC)
REGISTRY_URL=$(DOCKER_USER)

#
# ULURU SERVICE
#
HOST=0.0.0.0
PORT=5000
DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

#
# AUTH SERVICE
#
JWT_SECRET=mega-secret
AUTH_HOST=localhost
AUTH_PORT=5010

#
# USERS SERVICE
#
USERS_HOST=localhost
USERS_PORT=5020

#
# EMAIL SERVICE
#
EMAIL_HOST=localhost
EMAIL_PORT=5030
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DATABASE=1
PROVIDERS=sendgrid
PROVIDER_SENDGRID_API_KEY=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

clean c:
	@echo "[clean] Cleaning bin folder..."
	@rm -rf bin/

run r:
	@echo "[running] Running service..."
	@HOST=$(HOST) \
	PORT=$(PORT) \
	DATABASE_URL=$(DATABASE_URL) \
	JWT_SECRET=$(JWT_SECRET) \
	USERS_HOST=$(USERS_HOST) \
	USERS_PORT=$(USERS_PORT) \
	AUTH_HOST=$(AUTH_HOST) \
	AUTH_PORT=$(AUTH_PORT) \
	EMAIL_HOST=$(EMAIL_HOST) \
	EMAIL_PORT=$(EMAIL_PORT) \
	REDIS_HOST=$(REDIS_HOST) \
	REDIS_PORT=$(REDIS_PORT) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	REDIS_DATABASE=$(REDIS_DATABASE) \
	PROVIDERS=$(PROVIDERS)	\
	PROVIDER_SENDGRID_API_KEY=$(PROVIDER_SENDGRID_API_KEY) \
	go run cmd/$(NAME)/main.go

build b: proto
	@echo "[build] Building service..."
	@cd cmd/$(NAME) && go build -o $(BIN)

linux l:
	@echo "[build-linux] Building service..."
	@cd cmd/$(NAME) && GOOS=linux GOARCH=amd64 go build -o $(BIN)

add-migration am: 
	@echo "[add-migration] Adding migration"
	@goose -dir "./database/migrations" create $(name) sql

migrations m:
	@echo "[migrations] Runing migrations..."
	@cd database/migrations && goose postgres $(DSN) up

docker d:
	@echo "[docker] Building image..."
	@docker build -t $(SVC):$(VERSION) .
	
docker-login dl:
	@echo "[docker] Login to docker..."
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)

push p: linux docker docker-login
	@echo "[docker] pushing $(REGISTRY_URL)/$(SVC):$(VERSION)"
	@docker tag $(SVC):$(VERSION) $(REGISTRY_URL)/$(SVC):$(VERSION)
	@docker push $(REGISTRY_URL)/$(SVC):$(VERSION)

compose co:
	@echo "[docker-compose] Running docker-compose..."
	@docker-compose build
	@docker-compose up

stop s: 
	@echo "[docker-compose] Stopping docker-compose..."
	@docker-compose down

clean-proto cp:
	@echo "[clean-proto] Cleaning proto files..."
	@rm -rf proto/*.pb.go || true

proto pro: clean-proto
	@echo "[proto] Generating proto file..."
	@protoc -I proto -I $(GOPATH)/src --go_out=plugins=grpc:./proto ./proto/*.proto 

test t:
	@echo "[test] Testing $(NAME)..."
	@cd $(GOPATH)/src/github.com/microapis/users-api && make t
	@cd $(GOPATH)/src/github.com/microapis/auth-api && make t
	@cd $(GOPATH)/src/github.com/microapis/email-api && make t

template tmpl:
	@echo "[template] Generating..."
	@qtc template

update up: push
	@echo "[deploy] Update version on remote machine to $(VERSION) version..."
	@ssh $(REMOTE_USER)@$(REMOTE_IP) -T "cat > /remotefile.txt"

deploy de: 
	@echo "[deploy] Deploying to $(VERSION) version..."
	@make stop

.PHONY: clean c run r build b linux l add-migration am migrations m docker d docker-login dl push p compose co stop s clean-proto cp proto pro test t template tmpl