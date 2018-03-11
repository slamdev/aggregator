package aggregator

import (
	"testing"
	"os"
	"io/ioutil"
)

func TestParseArgsToReturnErrorWhenNoArgumentsPassed(t *testing.T) {
	assignArgs()
	appArgs, err := ParseArgs()
	assertEmptyAppArgs(appArgs, t)
	assertErrorMessage(err, "path to a file should be provided", t)
}

func TestParseArgsToReturnErrorWhenMoreThanOneArgumentPassed(t *testing.T) {
	assignArgs("1", "2")
	appArgs, err := ParseArgs()
	assertEmptyAppArgs(appArgs, t)
	assertErrorMessage(err, "only one argument is supported", t)
}

func TestParseArgsToReturnErrorWhenPassedPathIsNotExist(t *testing.T) {
	assignArgs("not-existing-file")
	appArgs, err := ParseArgs()
	assertEmptyAppArgs(appArgs, t)
	assertErrorMessage(err, "stat not-existing-file: no such file or directory", t)
}

func TestParseArgsToReturnErrorWhenPassedPathIsNotAFile(t *testing.T) {
	assignArgs(".")
	appArgs, err := ParseArgs()
	assertEmptyAppArgs(appArgs, t)
	assertErrorMessage(err, "provided path is not a file", t)
}

func TestParseArgsToReturnFileByPath(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	defer os.Remove(file.Name())
	assignArgs(file.Name())
	appArgs, err := ParseArgs()
	assertNil(err, t)
	assertEqual(file.Name(), appArgs.Filename, t)
}

func assignArgs(args ...string) {
	newArgs := []string{"cmd"}
	for _, arg := range args {
		newArgs = append(newArgs, arg)
	}
	os.Args = newArgs
}
