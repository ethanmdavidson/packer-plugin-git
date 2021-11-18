packer {
  required_plugins {
    git = {
      version = ">=v0.2.0"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}

data "git-commit" "test" {}

locals {
  hash = data.git-commit.test.hash
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
      "echo 'hash: ${local.hash}'",
      "echo 'author: ${data.git-commit.test.author}'",
      "echo 'committer: ${data.git-commit.test.committer}'",
      "echo 'pgp_signature: ${data.git-commit.test.pgp_signature}'",
      "echo 'message: ${data.git-commit.test.message}'",
      "echo 'tree_hash: ${data.git-commit.test.tree_hash}'",
      "echo 'first_parent: ${data.git-commit.test.parent_hashes[0]}'",
    ]
  }
}
