NAME=git
BINARY=packer-plugin-${NAME}
HASHICORP_PACKER_PLUGIN_SDK_VERSION?=$(shell go list -m github.com/hashicorp/packer-plugin-sdk | cut -d " " -f2)

COUNT?=1
TEST?=$(shell go list ./...)

prep: phony
	go install honnef.co/go/tools/cmd/staticcheck@latest

build: phony
	@go build -o ${BINARY}

dev: phony build
	@mkdir -p ~/.packer.d/plugins/
	@mv ${BINARY} ~/.packer.d/plugins/${BINARY}

run-example: phony dev
	@packer build ./example

test: phony
	go mod tidy
	go fmt ./...
	go vet ./...
	staticcheck -checks="all" -tests ./...
	go test -race -count $(COUNT) $(TEST) -timeout=3m

# the acceptance tests have a weird habit of messing up the tty (e.g. turning off echo mode, so
# terminal stops showing what you type). If this happens to you, run `reset` or `stty sane` to fix.
testacc: phony dev
	@PACKER_ACC=1 go test -count $(COUNT) -v $(TEST) -timeout=120m

install-packer-sdc: phony ## Install packer software development command
	@go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@${HASHICORP_PACKER_PLUGIN_SDK_VERSION}

ci-release-docs: phony install-packer-sdc
	@packer-sdc renderdocs -src docs -partials docs-partials/ -dst docs/
	@/bin/sh -c "[ -d docs ] && zip -r docs.zip docs/"

plugin-check: phony install-packer-sdc build
	@packer-sdc plugin-check ${BINARY}

generate: phony install-packer-sdc
	@go generate -v ./...
	packer-sdc renderdocs -src ./docs -dst ./.docs -partials ./docs-partials
	# see the .docs folder for a preview of the docs

# instead of listing every target in .PHONY, we create one
# 'phony' target which all the other targets depend on.
# This saves me from having to remember to add each new target
# to the .PHONY list, and is a little cleaner than putting
# `.PHONY: target` before each target
.PHONY: phony
phony:
