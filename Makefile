LOCAL_BIN:=$(CURDIR)/bin

DOCKER_IMAGE:=notification-service
DOCKER_IMAGE_TAG:=v0.0.1
DOCKER_CONTAINER_NAME:=notification-service-container

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG) .

docker-run:
	docker run -d -p 4008:8080 --name $(DOCKER_CONTAINER_NAME) -t $(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG)

docker-rm:
	docker stop $(DOCKER_CONTAINER_NAME)
	docker rm $(DOCKER_CONTAINER_NAME)

test:
	go test ./... -v