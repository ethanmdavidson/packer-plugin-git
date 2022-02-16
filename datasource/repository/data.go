//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package repository

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	Path string `mapstructure:"path"`
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	Head     string   `mapstructure:"head"`
	Branches []string `mapstructure:"branches"`
	Tags     []string `mapstructure:"tags"`
	IsClean  bool     `mapstructure:"is_clean"`
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
	head, err := repo.Head()
	if err != nil {
		return emptyOutput, err
	}
	worktree, err := repo.Worktree()
	if err != nil {
		return emptyOutput, err
	}
	status, err := worktree.Status()
	if err != nil {
		return emptyOutput, err
	}
	branchIter, err := repo.Branches()
	if err != nil {
		return emptyOutput, err
	}
	tagIter, err := repo.Tags()
	if err != nil {
		return emptyOutput, err
	}

	output.Head = head.Name().Short()
	output.IsClean = status.IsClean()
	output.Branches = make([]string, 0)
	_ = branchIter.ForEach(func(reference *plumbing.Reference) error {
		output.Branches = append(output.Branches, reference.Name().Short())
		return nil
	})
	output.Tags = make([]string, 0)
	_ = tagIter.ForEach(func(reference *plumbing.Reference) error {
		output.Tags = append(output.Tags, reference.Name().Short())
		return nil
	})

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
