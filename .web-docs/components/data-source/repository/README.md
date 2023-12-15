Type: `git-repository`

The repository data source is used to fetch information about a git repository.
It needs to be run inside an existing git repo.


## Required

There are no required configuration fields.


## Optional

- `path` (string) - The path to a directory inside your git repo. The plugin will
search for a git repo, starting in this directory and walking up through
parent directories. Defaults to '.' (the directory packer was executed in).


## Output

- `head` (string) - The short name of HEAD's current location.
- `branches` (list[string]) - The list of branches in the repository.
- `tags` (list[string]) - The list of tags in the repository.
- `is_clean` (bool) - `true` if the working tree is clean, `false` otherwise.


## Example Usage

This example shows how a a suffix can be added to the version number
for any AMI built outside the main branch.

```hcl
data "git-repository" "cwd" {}

variable version {
  type = string
}

locals {
  onMain = data.git-repository.cwd.head == "main"
  version  = onMain ? "${var.version}" : "${var.version}-SNAPSHOT"
}

source "amazon-ebs" "ami1" {
  ami_description = "AMI1"
  ami_name        = "ami1-${local.version}"
}
```
