data "git-tree" "test" {}

source "null" "basic-example" {
  communicator = "none"
}

locals {
  numFiles = length(data.git-tree.test.files)
  files = join(",", data.git-tree.test.files)
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo 'hash: ${data.git-tree.test.hash}'",
      "echo 'fileCount: ${local.numFiles}'",
      "echo 'files: ${local.files}'",
    ]
  }
}
