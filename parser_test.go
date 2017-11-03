package metro2

import (
	"io/ioutil"
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

func TestValidLength(t *testing.T) {
	file := "0015validheader"

	_, err := parseFixedHeader(file)

	if err != nil {
		t.Fatalf("File with valid length failed to parse: %s", err)
	}
}

func TestValidFile(t *testing.T) {
	file, _ := ioutil.ReadFile("header.txt")

	t.Log(len(file))

}
