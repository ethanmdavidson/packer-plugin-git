data "git-commit" "test" {
  path = ".."
}

locals {
  hash = data.git-commit.test.hash
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo 'hash: ${local.hash}'",
      "echo 'num_branches: ${length(data.git-commit.test.branches)}'",
      "echo 'author: ${data.git-commit.test.author}'",
      "echo 'committer: ${data.git-commit.test.committer}'",
      "echo 'pgp_signature: ${data.git-commit.test.pgp_signature}'",
      "echo 'message: ${data.git-commit.test.message}'",
      "echo 'tree_hash: ${data.git-commit.test.tree_hash}'",
      "echo 'first_parent: ${data.git-commit.test.parent_hashes[0]}'",
    ]
  }
}
