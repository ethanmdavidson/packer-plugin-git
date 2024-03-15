NAME=git
BINARY=packer-plugin-${NAME}
PLUGIN_FQN="$(shell grep -E '^module' <go.mod | sed -E 's/module *//')"
HASHICORP_PACKER_PLUGIN_SDK_VERSION?=$(shell go list -m github.com/hashicorp/packer-plugin-sdk | cut -d " " -f2)

COUNT?=1
TEST?=$(shell go list ./...)

prep: phony
	go install honnef.co/go/tools/cmd/staticcheck@latest

build: phony
	@go build -o ${BINARY}

dev: phony
	@go build -ldflags="-X '${PLUGIN_FQN}/version.VersionPrerelease=dev'" -o '${BINARY}'
	packer plugins install --path ${BINARY} "$(shell echo "${PLUGIN_FQN}" | sed 's/packer-plugin-//')"

run-example: phony dev
	@packer build ./example

test: phony
	go mod tidy
	go fmt ./...
	go vet ./...
	staticcheck -checks="all" -tests ./...
	go test -race -count $(COUNT) $(TEST) -timeout=3m

test-releaser: export API_VERSION = x5.0
test-releaser: phony
	go install github.com/goreleaser/goreleaser@latest
	goreleaser check
	goreleaser release --snapshot --clean --skip=sign

# the acceptance tests have a weird habit of messing up the tty (e.g. turning off echo mode, so
# terminal stops showing what you type). If this happens to you, run `reset` or `stty sane` to fix.
testacc: phony dev
	@PACKER_ACC=1 go test -count $(COUNT) -v $(TEST) -timeout=120m

install-packer-sdc: phony ## Install packer software development command
	@go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@${HASHICORP_PACKER_PLUGIN_SDK_VERSION}

plugin-check: phony install-packer-sdc build
	@packer-sdc plugin-check ${BINARY}

generate: phony install-packer-sdc
	@go generate -v ./...
	@if [ -d ".docs" ]; then rm -r ".docs"; fi
	@packer-sdc renderdocs -src "docs" -partials docs-partials/ -dst ".docs/"
	@./.web-docs/scripts/compile-to-webdocs.sh "." ".docs" ".web-docs" "ethanmdavidson"
	@rm -r ".docs"
	# see the .docs folder for a preview of the docs

# instead of listing every target in .PHONY, we create one
# 'phony' target which all the other targets depend on.
# This saves me from having to remember to add each new target
# to the .PHONY list, and is a little cleaner than putting
# `.PHONY: target` before each target
.PHONY: phony
phony:
