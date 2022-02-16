data "git-repository" "test" {}

source "null" "basic-example" {
  communicator = "none"
}

locals {
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
      "echo 'num_branches: ${length(data.git-repository.test.branches)}'",
      "echo 'tags: ${local.tags}'",
    ]
  }
}
