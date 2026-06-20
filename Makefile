BIN := bin/concord-plugin-hello
VERSION ?= v0.1.0
INSTALL_DIR := $(HOME)/.concord/plugins/hello/$(VERSION)

.PHONY: build install clean

build:
	go build -o $(BIN) ./cmd/concord-plugin-hello

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BIN) $(INSTALL_DIR)/concord-plugin-hello

clean:
	rm -rf bin
