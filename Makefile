ifeq ($(origin VERSION), undefined)
  VERSION != git rev-parse --short HEAD
endif

ifeq ($(SYSTEMD), 1)
	BUILD_TAGS = -tags "systemd"
endif

ifdef VERBOSE
	V="-v -x"
endif

HOST_GOOS:=$(shell go env GOOS)
HOST_GOARCH:=$(shell go env GOARCH)
REPOPATH = github.com/zahfox/gourd
INSTALL_DIR = /usr/local/bin/

WHAT := gourd gourdd gsh

info:
			@echo "$(VERSION) $(HOST_GOOS) $(HOST_GOARCH) $(REPOPATH) $(BUILD_TAGS)"

build: clean-build
			for target in $(WHAT); do \
			    echo "building $$target"; \
				$(BUILD_ENV_FLAGS) go build $(BUILD_TAGS) $(V) -o bin/$$target -ldflags "-X $(REPOPATH).Version=$(VERSION)" ./cmd/$$target; \
			done

install: clean-install build
			sudo useradd -Urs /bin/false gourd 1>/dev/null 2>&1; \
			sudo mkdir -p /etc/gourd /opt/gourd; \
			sudo chown -R gourd:gourd /etc/gourd /opt/gourd; \
			sudo chmod 750 /etc/gourd; \
			sudo chmod -R 750 /etc/gourd; \
			for target in $(WHAT); do \
				sudo cp bin/$$target $(INSTALL_DIR)$$target; \
				sudo chown gourd:gourd $(INSTALL_DIR)$$target; \
				sudo chmod 550 $(INSTALL_DIR)$$target; \
			done; \
			if [ ! -z "$$SYSTEMD" ] && [ "$$SYSTEMD" -eq "1" ]; then \
				sudo cp init/gourdd.service /etc/systemd/system/gourdd.service; \
				sudo cp init/gourdd.socket /etc/systemd/system/gourdd.socket; \
				sudo systemctl daemon-reload; \
				sudo systemctl enable gourdd.socket; \
				sudo systemctl start gourdd.socket; \
				sudo systemctl enable gourdd.service; \
				sudo systemctl start gourdd.service; \
			fi

clean-build:
			rm -rf ./bin

clean-install:
			if [ ! -z "$$SYSTEMD" ] && [ "$$SYSTEMD" -eq "1" ]; 	then \
				sudo systemctl stop gourdd.socket > /dev/null; \
				sudo systemctl disable gourdd.socket > /dev/null; \
				sudo systemctl stop gourdd.service > /dev/null; \
				sudo systemctl disable gourdd.service > /dev/null; \
				sudo rm -f /etc/systemd/system/gourdd.service; \
				sudo rm -f /etc/systemd/system/gourdd.socket; \
				sudo systemctl daemon-reload; \
			fi
			for target in $(WHAT); do \
				sudo rm -f $(INSTALL_DIR)$$target; \
			done

