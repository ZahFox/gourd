ifeq ($(origin VERSION), undefined)
  VERSION != git rev-parse --short HEAD
endif

HOST_GOOS=$(shell go env GOOS)
HOST_GOARCH=$(shell go env GOARCH)
REPOPATH = github.com/zahfox/gourd

WHAT := gourd

info:
			@echo "$(VERSION) $(HOST_GOOS) $(HOST_GOARCH) $(REPOPATH)"

build: clean
			for target in $(WHAT); do \
						$(BUILD_ENV_FLAGS) go build -v -x -o bin/$$target -ldflags "-X $(REPOPATH).Version=$(VERSION)" ./cmd/$$target.go; \
			done

install: build
			for target in $(WHAT); do \
						sudo cp bin/$$target /usr/local/bin/$$target; \
						sudo chown $(USER) /usr/local/bin/$$target; \
			done

clean:
			for target in $(WHAT); do \
						sudo rm -f /usr/local/bin/$$target; \
			done
			rm -rf ./bin

