# load .env file
include .env
export $(shell sed 's/=.*//' .env)

#
# SO
#
GOOSE_BIN=
UNAME_S=$(shell uname -s)
ifeq ($(UNAME_S),Linux)
	GOOSE_BIN=goose
endif
ifeq ($(UNAME_S),Darwin)
	GOOSE_BIN=goose-darwin
endif

#
# ARCH
#
CURRENT_ARCH=
UNAME_M=$(shell uname -m)
ifeq ($(UNAME_M),x86_64)
	CURRENT_ARCH=amd64
endif
ifneq ($(filter arm%,$(UNAME_M)),)
	CURRENT_ARCH=arm
endif


# API VALUES
SVC=ruvix-api
REGISTRY_URL=
VERSION=0.0.30
API_CONTAINER_NAME=$(SVC):$(VERSION)
BIN_PATH=$(PWD)/bin
BIN=$(BIN_PATH)/$(SVC)
GOPATH=$(HOME)/go

#
# GENERAL
#
clean:
	@echo "[clean] Cleaning bin folder..."
	@rm -rf bin/

run: 
	@echo "[running] Running service..."
	@go run cmd/$(SVC)/main.go

build:
	@echo "[build] Building service..."
	@GOOS=linux GOARCH=amd64 go build -o $(BIN)-v$(VERSION)-linux cmd/$(SVC)/main.go
	
#
# SERVICE API
#
dev:
	@echo "[run-dev] Running docker compose..."
	@docker-compose -f docker-compose.yml up -d --build
	@echo "[run-dev] running service in dev mode..."
	@go run cmd/$(SVC)/main.go

stop: 
	@echo "[stop] Stopping docker compose..."
	@docker-compose -f docker-compose.yml down || true

docker: 
	@echo "[docker] Login to docker..."
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	@echo "[docker] Building docker $(API_CONTAINER_NAME)..."
	@docker build -t $(API_CONTAINER_NAME) -f Dockerfile .
	@echo "[docker] Pushing $(REGISTRY_URL)/$(API_CONTAINER_NAME)"
	@docker tag $(API_CONTAINER_NAME) $(REGISTRY_URL)/$(API_CONTAINER_NAME)
	@docker push $(REGISTRY_URL)/$(API_CONTAINER_NAME)

#
# TEMPLATE
#
template tmpl:
	@echo "[template] Generating..."
	@qtc template

#
# DEPLOYMENT
#
deploy: linux
	@echo "[deploy] Deploying to v$(VERSION) version..."
	@git push dokku master

#
# TESTING
#
test-users tu:
	@echo ""
	@echo ""
	@echo "=========================="
	@echo "[test] Testing Users..."
	@echo "=========================="
	@echo ""
	@go test -count=1 -v ./pkg/users/*_test.go

test-authentication ta:
	@echo ""
	@echo ""
	@echo "==================================="
	@echo "[test] Testing Authentication..."
	@echo "==================================="
	@echo ""
	@HOST=$(AUTH_HOST) \
	 PORT=$(AUTH_PORT) \
	 go test -count=1 -v $(GOPATH)/src/github.com/cagodoy/ruvix-api/pkg/auth/client/auth_test.go

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
	 go test -count=1 -v $(GOPATH)/src/github.com/cagodoy/email-api/client/email_test.go

test t: test-users test-authentication test-profile test-savings test-goals

.PHONY: clean c run r build b linux l docker d docker-login dl push p compose co stop s clean-proto cp proto pro test t template tmpl