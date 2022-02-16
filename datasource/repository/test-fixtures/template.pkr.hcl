data "git-repository" "test" {}

source "null" "basic-example" {
  communicator = "none"
}

locals {
  branches = join(",", data.git-repository.test.branches)
  tags = join(",", data.git-repository.test.tags)
}

build {
  sources = [
    "source.null.basic-example"
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
