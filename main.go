package main

import (
	"fmt"
	"github.com/ethanmdavidson/packer-plugin-git/datasource/commit"
	"github.com/ethanmdavidson/packer-plugin-git/datasource/repository"
	"github.com/ethanmdavidson/packer-plugin-git/datasource/tree"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
)

var (
	// Version is the main version number that is being run at the moment.
	Version = "0.6.4"

	// PluginVersion is used by the plugin set to allow Packer to recognize
	// what version this plugin is.
	PluginVersion = version.NewRawVersion(Version)
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("commit", new(commit.Datasource))
	pps.RegisterDatasource("repository", new(repository.Datasource))
	pps.RegisterDatasource("tree", new(tree.Datasource))
	pps.SetVersion(PluginVersion)
	err := pps.Run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
