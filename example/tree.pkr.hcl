packer {
  required_plugins {
    git = {
      version = ">=v0.2.0"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}

data git-tree example {}

locals {
  numFiles     = length(data.git-tree.example.files)
  fileString   = join(",", sort(data.git-tree.example.files))
  fileChecksum = md5(local.fileString)
}

source "null" "git-plugin-test" {
  communicator = "none"
}

build {
  sources = ["source.null.git-plugin-test"]

  provisioner "shell-local" {
    inline = [
      "echo 'numFiles: ${local.numFiles}'",
      "echo 'files: ${local.fileString}'",
      "echo 'checksum: ${local.fileChecksum}'",
    ]
  }
}

