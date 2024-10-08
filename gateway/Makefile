ORGANIZATION = "organization"
APP = "gateway"

VERSION := $(shell git describe --tags --always --match='v*' 2>/dev/null)
ifeq ($(VERSION),)
    VERSION = "latest"
endif

COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null)
ifeq ($(COMMIT),)
    COMMIT = ""
endif

DATE = `date "+%Y-%m-%d %H:%M:%S"`
ARCH = `go env GOARCH`

.PHONY: build
build:
	@go build -ldflags "-X 'gateway/cmd.Date=$(DATE)' -X 'gateway/cmd.Version=$(VERSION)' -X 'gateway/cmd.Commit=$(COMMIT)'" -o $(APP) main.go

.PHONY: docker
docker:
	@docker buildx build --platform linux/$(ARCH) --progress=plain -t $(ORGANIZATION)/$(APP):$(VERSION) . --load

.PHONY: push
push:
	@docker buildx create --use --name=mybuilder --driver docker-container --driver-opt image=dockerpracticesig/buildkit:master 2> /dev/null || true
	@docker buildx build --platform linux/amd64,linux/arm64 --progress=plain -t $(ORGANIZATION)/$(APP):$(VERSION) . --push