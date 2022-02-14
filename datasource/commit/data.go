//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package commit

import (
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
	branches, err := repo.Branches()
	if err != nil {
		return emptyOutput, err
	}
	branchesAtResolvedCommit := make([]string, 0)
	_ = branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Hash().String() == hash.String() {
			branchesAtResolvedCommit = append(branchesAtResolvedCommit, ref.Name().Short())
		}
		return nil
	})
	commit, err := repo.CommitObject(*hash)
	if err != nil {
		return emptyOutput, err
	}

	output.Hash = hash.String()
	output.Branches = branchesAtResolvedCommit
	output.Author = commit.Author.String()
	output.Committer = commit.Committer.String()
	output.PGPSignature = commit.PGPSignature
	output.Message = commit.Message
	output.TreeHash = commit.TreeHash.String()
	for _, parent := range commit.ParentHashes {
		output.ParentHashes = append(output.ParentHashes, parent.String())
	}

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
