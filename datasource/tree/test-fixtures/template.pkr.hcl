data "git-commit" "test" {}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo 'hash: ${data.git-commit.test.files}'",
    ]
  }
}
