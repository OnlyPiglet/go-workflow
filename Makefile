.PHONY: build image release chart deploy

# Build variables
REGISTRY_URI := registry.cn-hangzhou.aliyuncs.com
APP_NAME=zice-go-workflow

MAIN_VERSION=1.0
PATCH_VERSION=0
IMAGE_VERSION?=v$(MAIN_VERSION).$(PATCH_VERSION)
CHART_VERSION=$(MAIN_VERSION).$(PATCH_VERSION)
GIT_COMMIT_ID=$(shell git rev-parse --short HEAD)

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o bin/linux/amd64/mmp  zice.com/MeasurementManagerPlatform/cmd

.PHONY: devquickx86
devquickx86:
	@echo "Build & push linux/amd64 image"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o bin/linux/amd64/go-workflow main.go
	docker build --platform=linux/amd64 -t $(REGISTRY_URI)/fordisk/$(APP_NAME):$(IMAGE_VERSION)-amd64-dev -f Dockerfile-local-quick .
	docker push  $(REGISTRY_URI)/fordisk/$(APP_NAME):$(IMAGE_VERSION)-amd64-dev


release:
	@echo "Build & push linux/amd64 image"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o bin/linux/amd64/go-workflow main.go
	docker build --platform=linux/amd64 -t $(REGISTRY_URI)/fordisk/$(APP_NAME):$(IMAGE_VERSION)-amd64 -f Dockerfile-local-quick .
	docker push  $(REGISTRY_URI)/fordisk/$(APP_NAME):$(IMAGE_VERSION)-amd64
