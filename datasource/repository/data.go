// Package repository contains logic for providing repo data to Packer
//
//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package repository

import (
	"github.com/ethanmdavidson/packer-plugin-git/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
	"log"
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
	log.Println("Starting execution")
	output := DatasourceOutput{}
	emptyOutput := hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec())

	common.PrintOpeningRepo(d.config.Path)
	openOptions := &git.PlainOpenOptions{DetectDotGit: true}
	repo, err := git.PlainOpenWithOptions(d.config.Path, openOptions)
	if err != nil {
		return emptyOutput, err
	}
	log.Println("Repo opened")

	head, err := repo.Head()
	if err != nil {
		return emptyOutput, err
	}
	log.Printf("Head found: '%s'\n", head.String())

	worktree, err := repo.Worktree()
	if err != nil {
		return emptyOutput, err
	}
	log.Println("Worktree found")

	status, err := worktree.Status()
	if err != nil {
		return emptyOutput, err
	}
	log.Printf("Worktree status found: '%s'\n", status.String())

	branchIter, err := repo.Branches()
	if err != nil {
		return emptyOutput, err
	}
	log.Println("Branches found")

	tagIter, err := repo.Tags()
	if err != nil {
		return emptyOutput, err
	}
	log.Println("Tags found")

	output.Head = head.Name().Short()
	log.Printf("output.Head: '%s'\n", output.Head)

	output.IsClean = status.IsClean()
	log.Printf("output.IsClean: '%t'\n", output.IsClean)

	output.Branches = make([]string, 0)
	_ = branchIter.ForEach(func(reference *plumbing.Reference) error {
		log.Printf("Adding branch: '%s'\n", reference.Name().Short())
		output.Branches = append(output.Branches, reference.Name().Short())
		return nil
	})
	log.Printf("len(output.Branches): '%d'\n", len(output.Branches))

	output.Tags = make([]string, 0)
	_ = tagIter.ForEach(func(reference *plumbing.Reference) error {
		log.Printf("Adding tag: '%s'\n", reference.Name().Short())
		output.Tags = append(output.Tags, reference.Name().Short())
		return nil
	})
	log.Printf("len(output.Tags): '%d'\n", len(output.Tags))

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
