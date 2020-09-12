#
# INTERNAL VARIABLES
#	note: should export env values form .bashrc/.zshrc
#				- DOCKER_USER
#				- DOCKER_PASS
#
VERSION=0.0.19
NAME=uluru
SVC=$(NAME)-api
BIN_PATH=$(PWD)/bin
BIN=$(BIN_PATH)/$(SVC)
REGISTRY_URL=$(DOCKER_USER)
GOPATH=$(HOME)/go

#
# ULURU SERVICE
#
HOST=localhost
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
REDIS_URL=redis://localhost:6379/0
PROVIDERS=sendgrid
PROVIDER_SENDGRID_API_KEY=SG.5-x4_3s-QKyPYQjPNP1l_Q.0lgbeslquQkSQSQNTd84vx2SjG-r90XYU35FKnK0Qqg===

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
	REDIS_URL=$(REDIS_URL) \
	PROVIDERS=$(PROVIDERS)	\
	PROVIDER_SENDGRID_API_KEY=$(PROVIDER_SENDGRID_API_KEY) \
	go run cmd/$(NAME)/main.go

build b: proto
	@echo "[build] Building service..."
	@cd cmd/$(NAME) && go build -o $(BIN)

linux l:
	@echo "[build-linux] Building service..."
	@cd cmd/$(NAME) && GOOS=linux GOARCH=amd64 go build -o $(BIN)

docker d: linux
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

compose-development code:
	@echo "[docker-compose] Running docker-compose in development mode..."
	@docker-compose -f docker-compose.development.yml build
	@docker-compose -f docker-compose.development.yml up

compose-local colo:
	@echo "[docker-compose] Running docker-compose in local mode..."
	@docker-compose -f docker-compose.local.yml build
	@docker-compose -f docker-compose.local.yml up

stop s: 
	@echo "[docker-compose] Stopping docker-compose..."
	@docker-compose down

clean-proto cp:
	@echo "[clean-proto] Cleaning proto files..."
	@rm -rf proto/*.pb.go || true

proto pro: clean-proto
	@echo "[proto] Generating proto file..."
	@protoc -I proto -I $(GOPATH)/src --go_out=plugins=grpc:./proto ./proto/*.proto 

test-users tu:
	@echo ""
	@echo ""
	@echo "=========================="
	@echo "[test] Testing Users..."
	@echo "=========================="
	@echo ""
	@HOST=$(USERS_HOST) \
	 PORT=$(USERS_PORT) \
	 go test -count=1 -v $(GOPATH)/src/github.com/microapis/users-api/client/users_test.go

test-authentication ta:
	@echo ""
	@echo ""
	@echo "==================================="
	@echo "[test] Testing Authentication..."
	@echo "==================================="
	@echo ""
	@HOST=$(AUTH_HOST) \
	 PORT=$(AUTH_PORT) \
	 go test -count=1 -v $(GOPATH)/src/github.com/microapis/authentication-api/client/auth_test.go

test-profile tp:
	@echo ""
	@echo ""
	@echo "============================"
	@echo "[test] Testing Profile..."
	@echo "============================"
	@echo ""
	@HOST=$(HOST) \
	 PORT=$(PORT) \
	 AUTH_HOST=$(AUTH_HOST) \
	 AUTH_PORT=$(AUTH_PORT) \
	 go test -count=1 -v ./pkg/profile/client_test.go

test-savings ts:
	@echo ""
	@echo ""
	@echo "=========================="
	@echo "[test] Testing Savings..."
	@echo "=========================="
	@echo ""
	@HOST=$(HOST) \
	 PORT=$(PORT) \
	 AUTH_HOST=$(AUTH_HOST) \
	 AUTH_PORT=$(AUTH_PORT) \
	 go test -count=1 -v ./pkg/savings/client_test.go

test-goals tg:
	@echo ""
	@echo ""
	@echo "=========================="
	@echo "[test] Testing Goals..."
	@echo "=========================="
	@echo ""
	@HOST=$(HOST) \
	 PORT=$(PORT) \
	 AUTH_HOST=$(AUTH_HOST) \
	 AUTH_PORT=$(AUTH_PORT) \
	 go test -count=1 -v ./pkg/goals/client_test.go

test-subscriptions tsu:
	@echo ""
	@echo ""
	@echo "=================================="
	@echo "[test] Testing Subscriptions..."
	@echo "=================================="
	@echo ""
	@HOST=$(HOST) \
	 PORT=$(PORT) \
	 AUTH_HOST=$(AUTH_HOST) \
	 AUTH_PORT=$(AUTH_PORT) \
	 go test -count=1 -v ./pkg/subscriptions/client_test.go
	
test-email te:
	@echo ""
	@echo ""
	@echo "=========================="
	@echo "[test] Testing Email..."
	@echo "=========================="
	@echo ""
	@HOST=$(EMAIL_HOST) \
	 PORT=$(EMAIL_PORT) \
	 go test -count=1 -v $(GOPATH)/src/github.com/microapis/email-api/client/email_test.go

test t: test-users test-authentication test-profile test-savings test-goals

template tmpl:
	@echo "[template] Generating..."
	@qtc template

deploy de: docker
	@echo "[deploy] Deploying to $(VERSION) version..."
	@git push dokku master

.PHONY: clean c run r build b linux l docker d docker-login dl push p compose co stop s clean-proto cp proto pro test t template tmpl