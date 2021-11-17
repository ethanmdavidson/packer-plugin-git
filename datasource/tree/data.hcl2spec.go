// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package tree

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	Path      *string `mapstructure:"path" cty:"path" hcl:"path"`
	CommitIsh *string `mapstructure:"commit_ish" cty:"commit_ish" hcl:"commit_ish"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"path":       &hcldec.AttrSpec{Name: "path", Type: cty.String, Required: false},
		"commit_ish": &hcldec.AttrSpec{Name: "commit_ish", Type: cty.String, Required: false},
	}
	return s
}

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	Hash    *string             `mapstructure:"hash" cty:"hash" hcl:"hash"`
	Entries []map[string]string `mapstructure:"entries" cty:"entries" hcl:"entries"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"hash":    &hcldec.AttrSpec{Name: "hash", Type: cty.String, Required: false},
		"entries": &hcldec.AttrSpec{Name: "entries", Type: cty.List(cty.Map(cty.String)), Required: false},
	}
	return s
}
