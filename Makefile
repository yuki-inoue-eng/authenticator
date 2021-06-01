SHELL=/bin/bash
.DEFAULT_GOAL := help

.PHONY: statik
statik: ## Embed config in statik.
	statik -src=configs -include=*.yml -dest=configs

.PHONY: generate-key-pair
generate-key-pair: ## Generate key pair used for digital signature.
	openssl genrsa 4096 > private.key
	openssl rsa -pubout < private.key > public.key

.PHONY: local-run
local-run: statik generate-key-pair ## Run in local environment.
	export PRIVATE_KEY=`cat private.key`; go run ./cmd/main/.

# See "Self-Documented Makefile" article
# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'