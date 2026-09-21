// File: parse_test.go
// Purpose: check strict number parsing and row splitting.
// Receives: sample texts.
// Previous stage: parse.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package numeric

import "testing"

func TestParseNumberAcceptsPlainDecimals(t *testing.T) {
	for _, text := range []string{"0", "-3", "+2.5", ".5", "4.", "1e3", "-2.5E-2"} {
		if _, err := ParseNumber(text); err != nil {
			t.Errorf("%q should be valid: %v", text, err)
		}
	}
}

func TestParseNumberRejectsInvalidText(t *testing.T) {
	for _, text := range []string{"", "abc", "2,5", "1_000", "0x10", "NaN", "Inf", "1e999", "--1", "3$"} {
		if _, err := ParseNumber(text); err == nil {
			t.Errorf("%q should be rejected", text)
		}
	}
}

func TestParseRowChecksCount(t *testing.T) {
	if _, err := ParseRow("1 2 3", 4); err == nil {
		t.Error("a row with the wrong count should be rejected")
	}
	row, err := ParseRow("1 -2 3.5", 3)
	if err != nil || len(row) != 3 || row[1] != -2 {
		t.Errorf("unexpected result %v, %v", row, err)
	}
}

func TestParseRowsNamesTheFailingEquation(t *testing.T) {
	if _, err := ParseRows([]string{"1 2 3", "1 x 3"}, 2); err == nil {
		t.Fatal("expected an error")
	}
}

func TestSplitRows(t *testing.T) {
	rows := SplitRows("1 2\n3 4;5 6\r\n\n")
	if len(rows) != 3 || rows[0] != "1 2" || rows[2] != "5 6" {
		t.Errorf("unexpected rows %q", rows)
	}
}
