package metro2

import (
	"io/ioutil"
	"testing"
)

func TestInvalidLength(t *testing.T) {
	file := "0426invalidheader"

	_, err := parseFixed(file)

	if err == nil {
		t.Fatal("File with invalid length should have failed to parse: instead, no error was returned")
	}
	t.Log(err)
}

func TestBadFormatLength(t *testing.T) {
	file := "15invalidheader"

	_, err := parseFixed(file)

	if err == nil {
		t.Fatalf("File with invalid numeric field did not return an error: %s", err)
	}
	t.Log(err)
}

func TestValidFile(t *testing.T) {
	file, _ := ioutil.ReadFile("data/header.txt")

	header, err := parseFixed(string(file)[:len(file)-1])

	if err != nil {
		t.Fatalf("Valid file failed to parse: %s", err)
	}

	t.Log(header)
}

func TestThereAndBackAgain(t *testing.T) {
	file, _ := ioutil.ReadFile("data/header.txt")

	header, _ := parseFixed(string(file)[:len(file)-1])

	formatted := header.Metro(426)

	if string(file) != formatted {
		t.Fatalf("Formatted header file was not the same as the original file.")
	}
}
