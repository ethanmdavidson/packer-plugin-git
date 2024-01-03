// Package commit contains logic for providing commit data to Packer
//
//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package commit

import (
	"log"
	"time"

	"github.com/ethanmdavidson/packer-plugin-git/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	Path      string `mapstructure:"path"`
	CommitIsh string `mapstructure:"commit_ish"`
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	Hash         string   `mapstructure:"hash"`
	Branches     []string `mapstructure:"branches"`
	Author       string   `mapstructure:"author"`
	Committer    string   `mapstructure:"committer"`
	Timestamp    string   `mapstructure:"timestamp"`
	PGPSignature string   `mapstructure:"pgp_signature"`
	Message      string   `mapstructure:"message"`
	TreeHash     string   `mapstructure:"tree_hash"`
	ParentHashes []string `mapstructure:"parent_hashes"`
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
	log.Println("Starting execution")
	output := DatasourceOutput{}
	emptyOutput := hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec())

	common.PrintOpeningRepo(d.config.Path)
	openOptions := &git.PlainOpenOptions{
		DetectDotGit:          true,
		EnableDotGitCommonDir: true,
	}
	repo, err := git.PlainOpenWithOptions(d.config.Path, openOptions)
	if err != nil {
		return emptyOutput, err
	}
	log.Println("Repo opened")

	hash, err := repo.ResolveRevision(plumbing.Revision(d.config.CommitIsh))
	if err != nil {
		return emptyOutput, err
	}
	log.Printf("Hash found: '%s'\n", hash.String())

	branchIter, err := repo.Branches()
	if err != nil {
		return emptyOutput, err
	}
	log.Println("Branches found")

	commit, err := repo.CommitObject(*hash)
	if err != nil {
		return emptyOutput, err
	}
	log.Printf("Commit found: '%s'\n", commit.String())

	output.Hash = hash.String()
	log.Printf("output.Hash: '%s'\n", output.Hash)

	output.Branches = make([]string, 0)
	_ = branchIter.ForEach(func(ref *plumbing.Reference) error {
		if ref.Hash().String() == hash.String() {
			log.Printf("Adding branch '%s'\n", ref.Name().Short())
			output.Branches = append(output.Branches, ref.Name().Short())
		} else {
			log.Printf("Not in branch '%s' (%s)\n", ref.Name().Short(), ref.Hash().String())
		}
		return nil
	})
	log.Printf("len(output.Branches): '%d'\n", len(output.Branches))

	output.Author = commit.Author.String()
	log.Printf("output.Author: '%s'\n", output.Author)

	output.Committer = commit.Committer.String()
	log.Printf("output.Committer: '%s'\n", output.Committer)

	output.Timestamp = commit.Committer.When.UTC().Format(time.RFC3339)
	log.Printf("output.Timestamp: '%s'\n", output.Timestamp)

	output.PGPSignature = commit.PGPSignature
	log.Printf("len(output.PGPSignature): '%d'\n", len(output.PGPSignature))

	output.Message = commit.Message
	log.Printf("len(output.Message): '%d'\n", len(output.Message))

	output.TreeHash = commit.TreeHash.String()
	log.Printf("output.TreeHash: '%s'\n", output.TreeHash)

	output.ParentHashes = make([]string, 0)
	for _, parent := range commit.ParentHashes {
		log.Printf("Adding parent hash: '%s'\n", parent.String())
		output.ParentHashes = append(output.ParentHashes, parent.String())
	}
	log.Printf("len(output.ParentHashes): '%d'\n", len(output.ParentHashes))

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
