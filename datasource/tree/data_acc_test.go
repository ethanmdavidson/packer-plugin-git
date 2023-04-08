package tree

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/tree/data_acc_test.go  -timeout=120m
func TestAccGitTreeDatasource(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "git_tree_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "git-tree",
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
			defer func(logs *os.File) {
				_ = logs.Close()
			}(logs)

			logsBytes, err := io.ReadAll(logs)
			if err != nil {
				return fmt.Errorf("Unable to read %s", logfile)
			}
			logsString := string(logsBytes)

			hashLog := "null.basic-example: hash: [0-9a-f]{5,40}"
			countLog := "null.basic-example: fileCount: [0-9]*"
			namesLog := "null.basic-example: files: [^\\n]*README.*"

			checkMatch(t, logsString, "hash", hashLog)
			checkMatch(t, logsString, "fileCount", countLog)
			checkMatch(t, logsString, "files", namesLog)
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
