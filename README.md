# Git Packer Plugin

[![tests](https://github.com/ethanmdavidson/packer-plugin-git/actions/workflows/run-tests.yml/badge.svg)](https://github.com/ethanmdavidson/packer-plugin-git/actions/workflows/run-tests.yml)

A plugin for [Packer](https://www.packer.io/) which provides access to git. Compatible with Packer >= 1.7.0

Under the hood, it uses [go-git](https://github.com/go-git/go-git).

## Usage

Add the plugin to your packer config:

```hcl
packer {
  required_plugins {
    git = {
      version = ">= 0.6.2"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}
```

Add the data source:

```hcl
data "git-commit" "example" { }
```

Now you should have access to info about the commit:

```hcl
locals {
  hash = data.git-commit.example.hash
}
```

### Examples

See [the examples directory](example) for some example code.

### Components

See [the docs](docs/README.md) for a reference of all the available
components and their attributes.

## Development

The GNUmakefile has all the commands you need to work with this repo.
The typical development flow looks something like this:

1) Make code changes, and add test cases for these changes.
2) Run `make generate` to recreate generated code.
2) Run `make dev` to build the plugin and install it locally.
3) Run `make testacc` to run the acceptance tests. If there are failures, go back to step 1.
4) Update examples in `./example` directory if necessary.
5) Run `make run-example` to test examples.
6) Once the above steps are complete: commit, push, and open a PR!

For local development, you will need to install:
- [Packer](https://learn.hashicorp.com/tutorials/packer/get-started-install-cli) >= 1.7
- [Go](https://golang.org/doc/install) >= 1.21
- [GNU Make](https://www.gnu.org/software/make/)

Check out the [Packer docs on Developing Plugins](https://developer.hashicorp.com/packer/docs/plugins/creation)
for more detailed info on plugins.

