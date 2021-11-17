package commit

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/commit/data_acc_test.go  -timeout=120m
func TestAccGitCommitDatasource(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "git_commit_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "git-commit",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			if buildCommand.ProcessState != nil {
				if buildCommand.ProcessState.ExitCode() != 0 {
					return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
				}
			}

			logs, err := os.Open(logfile)
			if err != nil {
				return fmt.Errorf("Unable find %s", logfile)
			}
			defer logs.Close()

			logsBytes, err := ioutil.ReadAll(logs)
			if err != nil {
				return fmt.Errorf("Unable to read %s", logfile)
			}
			logsString := string(logsBytes)

			hashLog := "null.basic-example: hash: [0-9a-f]{5,40}"
			authorLog := "null.basic-example: author: [^\\n]*<[^\\n]*>"
			committerLog := "null.basic-example: committer: [^\\n]*<[^\\n]*>"
			//Can't test pgp_signature since that isn't set on most of my commits
			messageLog := "null.basic-example: message: .*"
			treeHashLog := "null.basic-example: tree_hash: [0-9a-f]{5,40}"
			parentLog := "null.basic-example: first_parent: [0-9a-f]{5,40}"

			checkMatch(t, logsString, "hash", hashLog)
			checkMatch(t, logsString, "author", authorLog)
			checkMatch(t, logsString, "committer", committerLog)
			checkMatch(t, logsString, "message", messageLog)
			checkMatch(t, logsString, "tree_hash", treeHashLog)
			checkMatch(t, logsString, "first_parent", parentLog)
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}

func checkMatch(test *testing.T, logs string, checkName string, regex string) {
	if matched, _ := regexp.MatchString(regex, logs); !matched {
		test.Fatalf("logs don't contain expected %s value", checkName)
	}
}