// File: format_test.go
// Purpose: check number formatting and the shared texts.
// Receives: sample values and results.
// Previous stage: number.go and text.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package format

import (
	"strings"
	"testing"

	"gauss/internal/model"
)

func TestNumberUsesTwoDecimals(t *testing.T) {
	cases := map[float64]string{2.5: "2.50", -1: "-1.00", 1.005: "1.00", 0: "0.00"}
	for value, want := range cases {
		if got := Number(value); got != want {
			t.Errorf("Number(%g) = %q, want %q", value, got, want)
		}
	}
}

func TestNumberNeverShowsNegativeZero(t *testing.T) {
	if got := Number(-0.004); got != "0.00" {
		t.Errorf("got %q, want 0.00", got)
	}
}

func TestNumberUsesScientificNotationForExtremeValues(t *testing.T) {
	if got := Number(1e20); got != "1.00e+20" {
		t.Errorf("got %q", got)
	}
}

func TestStepTitle(t *testing.T) {
	step := model.Step{Column: 0, PivotRow: 0, SwappedWith: 1, Pivot: 2}
	want := "Paso 1 (columna 1): intercambiar fila 1 con fila 2 (pivote 2.00)"
	if got := StepTitle(1, step); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	skipped := model.Step{Column: 2, NoPivot: true, SwappedWith: -1}
	if got := StepTitle(3, skipped); !strings.Contains(got, "sin pivote") {
		t.Errorf("got %q", got)
	}
}

func TestOperationTextUsesPlusForNegativeFactor(t *testing.T) {
	got := OperationText(model.RowOperation{Row: 2, Factor: -0.5}, 0)
	if got != "fila 3 = fila 3 + 0.50 × fila 1" {
		t.Errorf("got %q", got)
	}
}

func TestVerdictLinesForNoSolutionPointToContradiction(t *testing.T) {
	result := model.Result{
		Case:             model.NoSolution,
		Echelon:          [][]float64{{1, 1, 2}, {0, 0, -1}},
		RankCoefficients: 1,
		RankAugmented:    2,
		ContradictoryRow: 1,
	}
	lines := VerdictLines(result)
	if len(lines) != 2 || !strings.Contains(lines[1], "Fila contradictoria: 2") {
		t.Errorf("unexpected lines: %q", lines)
	}
}

func TestSolutionLines(t *testing.T) {
	lines := SolutionLines([]float64{1, 2.5})
	if len(lines) != 2 || lines[0] != "x1 = 1.00" || lines[1] != "x2 = 2.50" {
		t.Errorf("unexpected lines: %q", lines)
	}
}
