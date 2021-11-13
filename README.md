# Git Packer Plugin

[![tests](https://github.com/ethanmdavidson/packer-plugin-git/actions/workflows/run-tests.yml/badge.svg)](https://github.com/ethanmdavidson/packer-plugin-git/actions/workflows/run-tests.yml)

A plugin for packer which provides access to git. Compatible with Packer >= 1.7.0

Right now, the only feature is a datasource that provides the current commit hash, because
this was the only feature I personally needed. If there is another feature
you want, feel free to open an issue or submit a PR.

Under the hood, it uses [go-git](https://github.com/go-git/go-git).

## Usage

Add the plugin to your packer config:
```hcl
packer {
  required_plugins {
    git = {
      version = ">=v0.1.0"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}
```

Now you should have access to the commit hash:
```hcl
locals {
  hash = data.git-local.test.commit_sha
}
```

See docs for more detailed information.

## Development

The GNUmakefile has all the commands you need to work with this repo. 
The typical development flow looks something like this:

1) Make code changes, and add test cases for these changes.
2) Run `make generate` to recreate generated code.
2) Run `make dev` to build the plugin and install it locally.
3) Run `make testacc` to run the acceptance tests. If there are failures, go back to step 1.
4) If the acceptance tests pass, commit and push!

For local development, you will need to install:
- Packer >= 1.7
- Go >= 1.16
- Make
