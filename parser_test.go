package metro2

import "testing"

func TestParseFixedHeader(t *testing.T) {
	file := "426invalidheader"

	_, err := parseFixedHeader(file)

	if err == nil {
		t.Fatalf("Failed parsing a header with incorrect RDW: %s", err)
	}
}
