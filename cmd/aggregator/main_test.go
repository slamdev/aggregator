package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestMainToAggregateDataFromFile(t *testing.T) {
	os.Args = []string{"cmd", "../../test/sampleInput.json"}
	output := captureOutput(func() {
		main()
	})
	expected, _ := ioutil.ReadFile("../../test/expectedOutput.json")
	assertEqual(string(expected), output, t)
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func assertEqual(expected, actual string, t *testing.T) {
	var o1 interface{}
	var o2 interface{}
	json.Unmarshal([]byte(expected), &o1)
	json.Unmarshal([]byte(actual), &o2)
	if !reflect.DeepEqual(o1, o2) {
		t.Fatalf("Expected %v be equal to %v", expected, actual)
	}
}
