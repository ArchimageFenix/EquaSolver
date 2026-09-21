// File: validator.go
// Purpose: check that a System is well formed and holds only finite numbers.
// Receives: model.System.
// Previous stage: solver (first thing it does with any system, from CLI or web).
// Next stage: solver (gets nil or a descriptive error).
// Restrictions: never modifies the system; no I/O; knows nothing about CLI or web.
//               It is the last line of defence and always runs.

package validation

import (
	"errors"
	"fmt"

	"gauss/internal/model"
	"gauss/internal/numeric"
)

// Validate returns nil when the system can be solved, or an error saying what is wrong.
func Validate(system model.System) error {
	if err := checkSize(system); err != nil {
		return err
	}
	return checkValues(system)
}

func checkSize(system model.System) error {
	if system.N < 1 {
		return errors.New("n debe ser un entero mayor que 0")
	}
	if len(system.A) != system.N || len(system.B) != system.N {
		return fmt.Errorf("el sistema debe tener %d ecuaciones con su término independiente", system.N)
	}
	for i, row := range system.A {
		if len(row) != system.N {
			return fmt.Errorf("la ecuación %d debe tener %d coeficientes", i+1, system.N)
		}
	}
	return nil
}

func checkValues(system model.System) error {
	for i, row := range system.A {
		for j, value := range row {
			if !numeric.IsFinite(value) {
				return fmt.Errorf("el coeficiente %d de la ecuación %d no es un número finito", j+1, i+1)
			}
		}
		if !numeric.IsFinite(system.B[i]) {
			return fmt.Errorf("el término independiente de la ecuación %d no es un número finito", i+1)
		}
	}
	return nil
}
