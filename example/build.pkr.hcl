packer {
  required_plugins {
    git = {
      version = ">=v0.0.1"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}

data "git-datasource-local" "test" {
  directory = "../"
}

locals {
  hash = data.git-datasource-local.test.commit_sha
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
