GOCMD		?=go
CUR_BRANCH := $(shell git branch --show-current 2>/dev/null || git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")
# 当前目录
CUR_DIR		=$(shell pwd)
OUT_DIR?	=$(CUR_DIR)/bin

# example
.PHONY: example
# run example
example:
	cd example && $(GOCMD) run ./main.go

# patch
.PHONY: patch
# create a patch file
patch:
	cd example && $(GOCMD) build -buildmode=plugin -o=patch_v1.so ./patch_v1/main.go

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help