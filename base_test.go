package metro2

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestValidBaseFile(t *testing.T) {
	file, _ := ioutil.ReadFile("data/base.txt")

	base, err := parseFixed(string(file)[:len(file)-1])

	if err != nil {
		t.Fatalf("Valid file failed to parse: %s", err)
	}

	json, _ := json.MarshalIndent(base, "", "  ")
	t.Log(string(json))
}

func TestInvalidBaseFile(t *testing.T) {
	file, _ := ioutil.ReadFile("data/bad_base.txt")

	base, err := parseFixed(string(file)[:len(file)-1])

	if err == nil {
		t.Fatalf("Invalid file failed to throw an error: %s", err)
	}
	t.Log(err)

	json, _ := json.MarshalIndent(base, "", "  ")
	t.Log(string(json))
}
