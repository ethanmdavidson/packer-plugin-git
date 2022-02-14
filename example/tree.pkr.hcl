packer {}

data git-tree example {}

locals {
  numFiles     = length(data.git-tree.example.files)
  fileString   = join(",", sort(data.git-tree.example.files))
  fileChecksum = md5(local.fileString)
}

source "null" "git-plugin-test-tree" {
  communicator = "none"
}

build {
  sources = ["source.null.git-plugin-test-tree"]

  provisioner "shell-local" {
    inline = [
      "echo 'numFiles: ${local.numFiles}'",
      "echo 'files: ${local.fileString}'",
      "echo 'checksum: ${local.fileChecksum}'",
    ]
  }
}
