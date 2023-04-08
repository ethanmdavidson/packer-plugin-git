NAME=git
BINARY=packer-plugin-${NAME}
HASHICORP_PACKER_PLUGIN_SDK_VERSION?=$(shell go list -m github.com/hashicorp/packer-plugin-sdk | cut -d " " -f2)

COUNT?=1
TEST?=$(shell go list ./...)

.PHONY: dev

build:
	@go build -o ${BINARY}

dev: build
	@mkdir -p ~/.packer.d/plugins/
	@mv ${BINARY} ~/.packer.d/plugins/${BINARY}

run-example: dev
	@packer build ./example

test:
	@go test -race -count $(COUNT) $(TEST) -timeout=3m

# the acceptance tests have a weird habit of messing up the tty (e.g. turning off echo mode, so
# terminal stops showing what you type). If this happens to you, run `reset` or `stty sane` to fix.
testacc: dev
	@PACKER_ACC=1 go test -count $(COUNT) -v $(TEST) -timeout=120m

install-packer-sdc: ## Install packer software development command
	@go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@${HASHICORP_PACKER_PLUGIN_SDK_VERSION}

ci-release-docs: install-packer-sdc
	@packer-sdc renderdocs -src docs -partials docs-partials/ -dst docs/
	@/bin/sh -c "[ -d docs ] && zip -r docs.zip docs/"

plugin-check: install-packer-sdc build
	@packer-sdc plugin-check ${BINARY}

generate: install-packer-sdc
	@go generate -v ./...
	packer-sdc renderdocs -src ./docs -dst ./.docs -partials ./docs-partials
	# see the .docs folder for a preview of the docs
