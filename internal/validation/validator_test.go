// File: validator_test.go
// Purpose: check the validation rules of the core.
// Receives: hand-made systems.
// Previous stage: validator.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package validation

import (
	"math"
	"testing"

	"gauss/internal/model"
)

func validSystem() model.System {
	return model.FromAugmentedRows([][]float64{{1, 2, 3}, {4, 5, 6}})
}

func TestValidateAcceptsWellFormedSystem(t *testing.T) {
	if err := Validate(validSystem()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateRejectsNonPositiveN(t *testing.T) {
	for _, n := range []int{0, -2} {
		if err := Validate(model.System{N: n}); err == nil {
			t.Errorf("n = %d should be rejected", n)
		}
	}
}

func TestValidateRejectsShapeMismatch(t *testing.T) {
	system := validSystem()
	system.A[1] = system.A[1][:1]
	if err := Validate(system); err == nil {
		t.Error("a short row should be rejected")
	}
	system = validSystem()
	system.B = system.B[:1]
	if err := Validate(system); err == nil {
		t.Error("a short independent vector should be rejected")
	}
}

func TestValidateRejectsNonFiniteValues(t *testing.T) {
	for _, bad := range []float64{math.NaN(), math.Inf(1), math.Inf(-1)} {
		system := validSystem()
		system.A[0][1] = bad
		if err := Validate(system); err == nil {
			t.Errorf("%v in A should be rejected", bad)
		}
		system = validSystem()
		system.B[1] = bad
		if err := Validate(system); err == nil {
			t.Errorf("%v in B should be rejected", bad)
		}
	}
}
