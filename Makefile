TAG=1.0.0
DOCKER_REGISTRY=docker.io
IMG_NAME=$(DOCKER_REGISTRY)/frenchben/label-exporter

build:
	docker buildx build --tag $(IMG_NAME):$(TAG) .

push: build
	docker image push $(IMG_NAME):$(TAG)

local-build:
	go build -o label-exporter cmd/label-exporter/main.go

coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

test:
	go test -v ./... -cover

clean:
	docker image rm $(IMG_NAME):$(TAG) || true
	_rm ./label-exporter
