package repository

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

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/repository/data_acc_test.go  -timeout=120m
func TestAccGitRepositoryDatasource(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "git_repository_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "git-repository",
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

			headLog := "null.basic-example: head: .*"
			isCleanLog := "null.basic-example: is_clean: [true|false]"
			branchesLog := "null.basic-example: num_branches: [0-9]*"
			tagsLog := "null.basic-example: num_tags: [0-9]*"

			checkMatch(t, logsString, "head", headLog)
			checkMatch(t, logsString, "clean", isCleanLog)
			checkMatch(t, logsString, "branches", branchesLog)
			checkMatch(t, logsString, "tags", tagsLog)

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
