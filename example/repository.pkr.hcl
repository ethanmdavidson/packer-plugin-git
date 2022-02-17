packer {}

data "git-repository" "test" {}

locals {
  branches = join(",", data.git-repository.test.branches)
  tags = join(",", data.git-repository.test.tags)
}

source "null" "git-plugin-test-repository" {
  communicator = "none"
}

build {
  sources = [
    "source.null.git-plugin-test-repository",
  ]
  provisioner "shell-local" {
    inline = [
      "echo 'head: ${data.git-repository.test.head}'",
      "echo 'is_clean: ${data.git-repository.test.is_clean}'",
      "echo 'branches: ${local.branches}'",
      "echo 'tags: ${local.tags}'",
    ]
  }
}
