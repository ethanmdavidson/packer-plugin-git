packer {}

data "git-commit" "test" {}

locals {
  hash = data.git-commit.test.hash
  # if message contains a single quote, it will mess up the echo command
  message = replace(data.git-commit.test.message, "'", "")
  branchesString = join(",", sort(data.git-commit.test.branches))
}

source "null" "git-plugin-test-commit" {
  communicator = "none"
}

build {
  sources = [
    "source.null.git-plugin-test-commit",
  ]
  provisioner "shell-local" {
    inline = [
      "echo 'hash: ${local.hash}'",
      "echo 'branches: ${local.branchesString}'",
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
