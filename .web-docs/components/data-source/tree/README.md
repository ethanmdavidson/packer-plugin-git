Type: `git-tree`

The tree data source is used to fetch the 'tree' or list of files
from a specific commit. It needs to be run inside an existing git
repo.


## Required

There are no required configuration fields.


## Optional

- `path` (string) - The path to a directory inside your git repo. The
plugin will search for a git repo, starting in this directory and walking
up through parent directories. Defaults to '.' (the directory packer
was execued in).

- `commit_ish` (string) - A [Commit-Ish value](https://git-scm.com/docs/gitglossary#Documentation/gitglossary.txt-aiddefcommit-ishacommit-ishalsocommittish)
(e.g. tag) pointing to the target commit object.
See [go-git ResolveRevision](https://pkg.go.dev/github.com/go-git/go-git/v5#Repository.ResolveRevision)
for the list of supported values. Defaults to 'HEAD'.


## Output

- `hash` (string) - The SHA1 checksum or "hash" value of the selected commit.
- `files` (list[string]) - The list of files present at this commit.

## Example Usage

This example shows how to get the checksum of the files tracked by git. 

```hcl
data "git-tree" "cwd-head" { }

locals {
  checksum = md5(join(",", sort(data.git-tree.cwd-head.files)))
}
```
