packer {
  required_plugins {
    git = {
      version = ">=v0.1.0"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}

data "git-local" "test" { }

locals {
  hash = data.git-local.test.commit_sha
}

source "null" "git-plugin-test" {
  communicator = "none"
}

build {
  sources = [
    "source.null.git-plugin-test",
  ]
  provisioner "shell-local" {
    inline = [
      "echo hash: ${local.hash}",
    ]
  }
}
