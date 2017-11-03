package metro2

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestValidBaseFile(t *testing.T) {
	file, _ := ioutil.ReadFile("base.txt")

	base, err := parseFixed(string(file))

	if err != nil {
		t.Fatalf("Valid file failed to parse: %s", err)
	}

	json, _ := json.MarshalIndent(base, "", "  ")
	t.Log(string(json))
}
