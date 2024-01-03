data "git-commit" "test" {
  path = ".."
}

locals {
  hash = data.git-commit.test.hash
  # if message contains a single quote, the test will fail
  message = replace(data.git-commit.test.message, "'", "")
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
      "echo 'timestamp: ${data.git-commit.test.timestamp}'",
      "echo 'pgp_signature: ${data.git-commit.test.pgp_signature}'",
      "echo 'message: ${local.message}'",
      "echo 'tree_hash: ${data.git-commit.test.tree_hash}'",
      "echo 'first_parent: ${data.git-commit.test.parent_hashes[0]}'",
    ]
  }
}
