data "git-tree" "test" {}

source "null" "basic-example" {
  communicator = "none"
}

locals {
  numFiles = length(data.git-tree.test.entries)
  files = join(",", data.git-tree.test.entries[*].name)
  allData = [for e in data.git-tree.test.entries: "${e.name}:${e.mode}:${e.hash}"]
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo 'hash: ${data.git-tree.test.hash}'",
      "echo 'filenames: ${local.files}'",
      "echo 'test: ${local.numFiles}'",
      "echo 'alldata: ${join(",", local.allData)}'",
    ]
  }
}
