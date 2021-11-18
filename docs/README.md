# Git Components

The Git plugin is able to interact with Git repos through Packer.
Currently, it comes with the following components:

### Data Sources

- [commit](/docs/datasources/commit.mdx) - Retrieve information
    about a specific commit, e.g. the commit hash.
- [tree](/docs/datasources/tree.mdx) - Retrieve the list of 
    files present in a specific commit, similar to `git ls-tree -r`.

