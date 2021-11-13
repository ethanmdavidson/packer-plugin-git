package local

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

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/git/data_acc_test.go  -timeout=120m
func TestAccScaffoldingDatasource(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "git_local_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "git-local",
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

			fooLog := "null.basic-example: hash: [0-9a-f]{5,40}"

			if matched, _ := regexp.MatchString(fooLog+".*", logsString); !matched {
				t.Fatalf("logs doesn't contain expected hash value %q", logsString)
			}
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
