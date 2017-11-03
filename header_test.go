package metro2

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

func TestInvalidLength(t *testing.T) {
	file := "0426invalidheader"

	_, err := parseFixedHeader(file)

	if err == nil {
		t.Fatal("File with invalid length should have failed to parse: instead, no error was returned")
	}
}

func TestBadFormatLength(t *testing.T) {
	file := "15invalidheader"

	_, err := parseFixedHeader(file)

	if err == nil {
		t.Fatalf("File with invalid numeric field did not return an error: %s", err)
	}
}

func TestValidFile(t *testing.T) {
	file, _ := ioutil.ReadFile("header.txt")

	header, err := parseFixed(string(file))

	if err != nil {
		t.Fatalf("Valid file failed to parse: %s", err)
	}

	json, _ := json.MarshalIndent(header, "", "  ")
	t.Log(string(json))
}

func TestThereAndBackAgain(t *testing.T) {
	file, _ := ioutil.ReadFile("header.txt")

	header, _ := parseFixed(string(file))

	formatted := formatFixedHeader(header.(*Header), 426)

	t.Log(string(file) == formatted)
	t.Log(strings.Compare(string(file), formatted))
	t.Log(len(string(file)), len(formatted))
	t.Log(string(file))
	t.Log(formatted)
}