//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package tree

import (
	"errors"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	Path      string `mapstructure:"path"`
	CommitIsh string `mapstructure:"commit_ish"` //should this be tree-ish instead?
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	Hash  string   `mapstructure:"hash"`
	Files []string `mapstructure:"files"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}
	if d.config.Path == "" {
		d.config.Path = "."
	}
	if d.config.CommitIsh == "" {
		d.config.CommitIsh = "HEAD"
	}
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	output := DatasourceOutput{}
	emptyOutput := hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec())

	openOptions := &git.PlainOpenOptions{DetectDotGit: true}
	repo, err := git.PlainOpenWithOptions(d.config.Path, openOptions)
	if err != nil {
		return emptyOutput, err
	}
	hash, err := repo.ResolveRevision(plumbing.Revision(d.config.CommitIsh))
	if err != nil {
		return emptyOutput, err
	}
	commit, err := repo.CommitObject(*hash)
	if err != nil {
		return emptyOutput, errors.New("couldn't find commit")
	}
	tree, err := commit.Tree()
	if err != nil {
		return emptyOutput, errors.New("couldn't find tree")
	}

	output.Hash = hash.String()
	tree.Files().ForEach(func(file *object.File) error {
		if file != nil {
			output.Files = append(output.Files, file.Name)
		}
		return nil
	})

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
