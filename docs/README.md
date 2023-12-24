The Git plugin is able to interact with Git repos through Packer.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    git = {
      version = ">= 0.5.0"
      source  = "github.com/ethanmdavidson/git"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/ethanmdavidson/git
```

### Manual Installation

You can find pre-built binary releases of the plugin 
[here](https://github.com/ethanmdavidson/packer-plugin-git/releases).

Or, if you prefer to build the plugin from its source
code, clone [the GitHub repository](https://github.com/ethanmdavidson/packer-plugin-git)
locally and run the command `make build` from the root
directory. Upon successful compilation, a `packer-plugin-git` plugin
binary file can be found in the root directory.

Once the binary is downloaded or built, please follow the Packer
documentation on
[installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).

### Components

### Data Sources

- [git-commit](/packer/integrations/ethanmdavidson/git/latest/components/data-source/commit) - Retrieve information
    about a specific commit, e.g. the commit hash.
- [git-repository](/packer/integrations/ethanmdavidson/git/latest/components/data-source/repository) - Retrieve information
    about a repository, e.g. the value of HEAD.
- [git-tree](/packer/integrations/ethanmdavidson/git/latest/components/data-source/tree) - Retrieve the list of
    files present in a specific commit, similar to `git ls-tree -r`.

