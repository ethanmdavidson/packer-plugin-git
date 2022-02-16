data "git-repository" "test" {}

locals {
  refval = data.git-repository.test.value
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
      "echo 'value: ${local.refval}'",
    ]
  }
}
