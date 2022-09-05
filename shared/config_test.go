package shared

import (
	"fmt"
	"testing"
)

const checkKey = "x"
const checkVal = "y"

var tempJSON = fmt.Sprintf(`{"%s": "%s"}`, checkKey, checkVal)

func Test_NewConfigFileReader_returns_err_on_invalid_file(t *testing.T) {
	_, err := LoadConfig("/bad/file/path.zyy")
	if err == nil {
		t.Error("expected an error with bad file path and did not get one.")
	}
	return
}

func Test_NewConfig_reads_file_ok(t *testing.T) {
	fmt.Println("path: " + GetConfigPath())
	_, err := LoadConfig(GetConfigPath())
	if err != nil {
		t.Error(err)
	}
}
