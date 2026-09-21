// File: cli_test.go
// Purpose: check the CLI flow: arguments, menu, file input, retries and output.
// Receives: simulated arguments and console input.
// Previous stage: run.go and its helpers.
// Next stage: none (test result).
// Restrictions: tests only; files are created in temporary directories.

package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func runCLI(args []string, input string) (int, string) {
	var out bytes.Buffer
	code := Run(args, strings.NewReader(input), &out)
	return code, out.String()
}

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "system.txt")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestRunWithMatrixArgument(t *testing.T) {
	code, out := runCLI([]string{"-m", "1 1 1 6; 2 -1 1 3; 1 2 -1 2"}, "")
	if code != exitOK || !strings.Contains(out, "SOLUCIÓN ÚNICA") || !strings.Contains(out, "x3 = 3.00") {
		t.Errorf("code %d, output:\n%s", code, out)
	}
}

func TestRunWithFileArgument(t *testing.T) {
	path := writeTempFile(t, "3\n1 1 1 6\n2 -1 1 3\n1 2 -1 2\n")
	code, out := runCLI([]string{"-f", path}, "")
	if code != exitOK || !strings.Contains(out, "x1 = 1.00") {
		t.Errorf("code %d, output:\n%s", code, out)
	}
}

func TestRunTypedThroughMenu(t *testing.T) {
	code, out := runCLI(nil, "1\n1\n2 4\n")
	if code != exitOK || !strings.Contains(out, "Forma de ingreso") || !strings.Contains(out, "x1 = 2.00") {
		t.Errorf("code %d, output:\n%s", code, out)
	}
}

func TestRunAsksAgainAfterInvalidInput(t *testing.T) {
	code, out := runCLI(nil, "1\n1\nabc\n2 4\n")
	if code != exitOK || !strings.Contains(out, "No válido") || !strings.Contains(out, "x1 = 2.00") {
		t.Errorf("code %d, output:\n%s", code, out)
	}
}

func TestRunFileThroughMenu(t *testing.T) {
	path := writeTempFile(t, "1\n2 4\n")
	code, out := runCLI(nil, "2\n"+path+"\n")
	if code != exitOK || !strings.Contains(out, "x1 = 2.00") {
		t.Errorf("code %d, output:\n%s", code, out)
	}
}

func TestRunShowsContradictoryRow(t *testing.T) {
	code, out := runCLI([]string{"-m", "1 1 2; 1 1 3"}, "")
	if code != exitOK || !strings.Contains(out, "SIN SOLUCIÓN") || !strings.Contains(out, "Fila contradictoria") {
		t.Errorf("code %d, output:\n%s", code, out)
	}
}

func TestRunRejectsBothArguments(t *testing.T) {
	code, _ := runCLI([]string{"-f", "a.txt", "-m", "1 2"}, "")
	if code != exitUsage {
		t.Errorf("code %d, want %d", code, exitUsage)
	}
}

func TestRunReportsUnreadableFile(t *testing.T) {
	code, out := runCLI([]string{"-f", filepath.Join(t.TempDir(), "missing.txt")}, "")
	if code != exitInput || !strings.Contains(out, "Error") {
		t.Errorf("code %d, output:\n%s", code, out)
	}
}

func TestRunStopsWhenInputEnds(t *testing.T) {
	code, _ := runCLI(nil, "")
	if code != exitInput {
		t.Errorf("code %d, want %d", code, exitInput)
	}
}

func TestParseCount(t *testing.T) {
	for _, text := range []string{"0", "-1", "x", "", "2.5"} {
		if _, err := parseCount(text); err == nil {
			t.Errorf("%q should be rejected", text)
		}
	}
	if n, err := parseCount("3"); err != nil || n != 3 {
		t.Errorf("got %d, %v", n, err)
	}
}

func TestSystemFromLinesChecksTheEquationCount(t *testing.T) {
	if _, err := systemFromLines([]string{"2", "1 2 3"}); err == nil {
		t.Error("a missing equation should be rejected")
	}
	if _, err := systemFromLines(nil); err == nil {
		t.Error("an empty file should be rejected")
	}
}
