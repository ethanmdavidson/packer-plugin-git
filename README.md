# Git Packer Plugin

[![tests](https://github.com/ethanmdavidson/packer-plugin-git/actions/workflows/run-tests.yml/badge.svg)](https://github.com/ethanmdavidson/packer-plugin-git/actions/workflows/run-tests.yml)

A plugin for packer which provides access to git.

Right now, the only feature is a datasource that provides the current commit hash, because
this was the only feature I personally needed. If there is another feature
you want, feel free to open an issue or submit a PR.

## Usage

Add the plugin to your packer config:
```hcl
packer {
  required_plugins {
    git = {
      version = ">=v0.0.2"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}
```

Then configure the datasource (directory should be the root of the git repo, relative
to the location of the packer template):
```hcl
data "git-local" "test" {
  directory = "../"
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

## Registering Documentation on Packer.io

Documentation for a plugin is maintained within the `docs` directory and served on GitHub.
To include plugin docs on Packer.io a global pre-hook has been added to the main scaffolding .goreleaser.yml file, that if uncommented will generate and include a docs.zip file as part of the plugin release.

The `docs.zip` file will contain all of the `.mdx` files under the plugins root `docs/` directory that can be consumed remotely by Packer.io.

Once the first `docs.zip` file has been included into a release you will need to open a one time pull-request against [hashicorp/packer](https://github.com/hashicorp/packer) to register the plugin docs.
This is done by adding the block below for the respective plugin to the file [website/data/docs-remote-navigation.js](https://github.com/hashicorp/packer/blob/master/website/data/docs-remote-plugins.json).

```json
{
   "title": "Scaffolding",
   "path": "scaffolding",
   "repo": "hashicorp/packer-plugin-scaffolding",
   "version": "latest",
   "sourceBranch": "main"
 }
```

If a plugin maintainer wishes to only include a specific version of released docs then the `"version"` key in the above configuration should be set to a released version of the plugin. Otherwise it should be set to `"latest"`.

The `"sourceBranch"` key in the above configuration ensures potential contributors can link back to source files in the plugin repository from the Packer docs site. If a `"sourceBranch"` value is not present, it will default to `"main"`. 

The documentation structure needed for Packer.io can be generated manually, by creating a simple zip file called `docs.zip` of the docs directory and included in the plugin release.

```/bin/bash
[[ -d docs/ ]] && zip -r docs.zip docs/
```

Once the first `docs.zip` file has been included into a release you will need to open a one time pull-request against [hashicorp/packer](https://github.com/hashicorp/packer) to register the plugin docs.

# Requirements

-	[packer-plugin-sdk](https://github.com/hashicorp/packer-plugin-sdk) >= v0.1.0
-	[Go](https://golang.org/doc/install) >= 1.16

## Packer Compatibility
This plugin is compatible with Packer >= v1.7.0
