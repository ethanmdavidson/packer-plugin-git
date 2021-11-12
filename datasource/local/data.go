//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package local

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
	"github.com/go-git/go-git/v5"
)

type Config struct {
	Directory string `mapstructure:"directory"`
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	CommitSha string `mapstructure:"commit_sha"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}
	if d.config.Directory == "" {
	    d.config.Directory = "."
	}
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
    output := DatasourceOutput{}
    emptyOutput := hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec())

    repo, err := git.PlainOpen(d.config.Directory)
    if err != nil {
        return emptyOutput, err
    }
    hash, err := repo.ResolveRevision("HEAD")
    if err != nil {
        return emptyOutput, err
    }

	output.CommitSha = hash.String()
	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
