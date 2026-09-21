// File: tolerance.go
// Purpose: compute the relative tolerance used to decide when a value counts as zero.
// Receives: a validated model.System.
// Previous stage: solver (calls it right after validation).
// Next stage: elimination and classification (through the solver).
// Restrictions: the only place where the tolerance is computed; no I/O, no other duties.
//               Formula: n * epsilon * max|a_ij| over the augmented matrix.

package tolerance

import (
	"math"

	"gauss/internal/model"
)

// machineEpsilon is the float64 machine epsilon (2^-52).
const machineEpsilon = 2.220446049250313e-16

// Compute returns the relative tolerance for the system.
func Compute(system model.System) float64 {
	return float64(system.N) * machineEpsilon * largestMagnitude(system)
}

func largestMagnitude(system model.System) float64 {
	largest := 0.0
	for i, row := range system.A {
		largest = math.Max(largest, maxAbs(row))
		largest = math.Max(largest, math.Abs(system.B[i]))
	}
	return largest
}

func maxAbs(values []float64) float64 {
	largest := 0.0
	for _, value := range values {
		largest = math.Max(largest, math.Abs(value))
	}
	return largest
}
