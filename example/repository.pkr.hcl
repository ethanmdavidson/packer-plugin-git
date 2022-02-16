packer {}

data "git-repository" "test" {}

source "null" "git-plugin-test-reference" {
  communicator = "none"
}

build {
  sources = [
    "source.null.git-plugin-test-reference",
  ]
  provisioner "shell-local" {
    inline = [
      "echo 'HEAD is: ${data.git-repository.test.head}'"
    ]
  }
}
