data "git-datasource-local" "test" {
  directory = "../.."
}

locals {
  hash = data.git-datasource-local.test.commit_sha
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
      "echo hash: ${local.hash}",
    ]
  }
}
