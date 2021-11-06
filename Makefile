ARTIFACTS := _out
GO := go
TARGET = reverse-shell
SRC = cmd/main.go
LDFLAGS = -ldflags="-s -w"

all: compile

compile-linux: ## Compile for linux 64bit
	GOOS=linux GOARCH=amd64 $(GO) build  -o $(ARTIFACTS)/$(TARGET)-linux $(LDFLAGS) $(SRC)

compile-darwin: ## Compile for darwin 64bit
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(ARTIFACTS)/$(TARGET)-darwin $(LDFLAGS) $(SRC)

compile-freebsd: ## Compile for freebsd 64bit
	GOOS=freebsd GOARCH=amd64 $(GO) build -o $(ARTIFACTS)/$(TARGET)-darwin $(LDFLAGS) $(SRC)

compile-windows: ## Compile for windows 64bit
	GOOS=windows GOARCH=amd64 $(GO) build -o $(ARTIFACTS)/$(TARGET)-windows.exe $(LDFLAGS) $(SRC)

compile: ## Compile for all targets
	@$(MAKE) compile-linux compile-darwin compile-windows

.PHONY: clean
clean: ## Cleans up all artifacts
	@-rm -rf $(ARTIFACTS)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
