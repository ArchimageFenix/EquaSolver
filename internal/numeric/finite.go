// File: finite.go
// Purpose: tell whether a floating point value is a finite number.
// Receives: a float64 value.
// Previous stage: validation, elimination and substitution (numeric checks).
// Next stage: none (returns a boolean to the caller).
// Restrictions: no state and no I/O; it does not parse, format or repair values.

package numeric

import "math"

// IsFinite reports whether value is neither NaN nor infinite.
func IsFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}
