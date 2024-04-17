DIST_DIR := .dist
DOCKER_IMAGE_NAME := user-management-api
DOCKER_TAG := latest

default: all

all: clean test build run

.PHONY: $(DIST_DIR)/server
$(DIST_DIR)/server: ./cmd/server/main.go
	CGO_ENABLED=0 go build -mod vendor -ldflags="-s -w" -o ./$(DIST_DIR)/server ./cmd/server/

.PHONY: build
build: $(DIST_DIR)/server

.PHONY: clean
clean:
	rm -rf $(DIST_DIR)
	mkdir $(DIST_DIR)
	@go mod vendor
	@go mod tidy

.PHONY: run
run: build
	$(DIST_DIR)/server


.PHONY: test
test:
	go test -mod vendor -race -cover -coverprofile=./$(DIST_DIR)/coverage.txt -covermode=atomic ./...

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run: docker-build
	docker run -p 8081:8081 $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)
