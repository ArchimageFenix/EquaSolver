// File: tolerance_test.go
// Purpose: check the relative tolerance formula.
// Receives: hand-made systems.
// Previous stage: tolerance.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package tolerance

import (
	"math"
	"testing"

	"gauss/internal/model"
)

func TestComputeUsesLargestMagnitudeOfAugmentedMatrix(t *testing.T) {
	system := model.FromAugmentedRows([][]float64{{1, 2, 5}, {3, 4, -6}})
	want := 2 * machineEpsilon * 6
	if got := Compute(system); math.Abs(got-want) > 1e-28 {
		t.Errorf("got %g, want %g", got, want)
	}
}

func TestComputeScalesWithTheData(t *testing.T) {
	small := model.FromAugmentedRows([][]float64{{1, 2, 3}, {4, 5, 6}})
	big := model.FromAugmentedRows([][]float64{{1e6, 2e6, 3e6}, {4e6, 5e6, 6e6}})
	ratio := Compute(big) / Compute(small)
	if math.Abs(ratio-1e6) > 1 {
		t.Errorf("tolerance should scale with the data, ratio %g", ratio)
	}
}

func TestComputeIsZeroForZeroMatrix(t *testing.T) {
	system := model.FromAugmentedRows([][]float64{{0, 0}})
	if got := Compute(system); got != 0 {
		t.Errorf("got %g, want 0", got)
	}
}
