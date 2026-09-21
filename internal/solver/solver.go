// File: solver.go
// Purpose: orchestrate the core: validate, compute tolerance, eliminate, classify,
//          substitute, and assemble the Result.
// Receives: a model.System from the CLI or web adapters.
// Previous stage: cli/run.go and web/handlers.go.
// Next stage: cli/present.go and web/present.go (they receive the model.Result).
// Restrictions: no formulas of its own, no reading and no printing; it only chains
//               the core stages. A validation failure is returned as an error, while
//               the three cases and a numeric failure are normal Results.

package solver

import (
	"gauss/internal/classification"
	"gauss/internal/elimination"
	"gauss/internal/model"
	"gauss/internal/substitution"
	"gauss/internal/tolerance"
	"gauss/internal/validation"
)

// Solve runs the whole core over the system.
func Solve(system model.System) (model.Result, error) {
	if err := validation.Validate(system); err != nil {
		return model.Result{}, err
	}
	limit := tolerance.Compute(system)
	outcome := elimination.Run(system, limit)
	if outcome.Failed {
		return markFailure(baseResult(outcome)), nil
	}
	class := classification.Classify(outcome.Echelon, outcome.Rank, limit)
	return completeResult(baseResult(outcome), class), nil
}

func baseResult(outcome elimination.Outcome) model.Result {
	return model.Result{
		Original:         outcome.Original,
		Steps:            outcome.Steps,
		Echelon:          outcome.Echelon,
		RankCoefficients: outcome.Rank,
		ContradictoryRow: -1,
	}
}

func completeResult(result model.Result, class classification.Outcome) model.Result {
	result.Case = class.Case
	result.RankAugmented = class.RankAugmented
	result.FreeVariables = class.FreeVariables
	result.ContradictoryRow = class.ContradictoryRow
	if class.Case == model.UniqueSolution {
		return attachSolution(result)
	}
	return result
}

func attachSolution(result model.Result) model.Result {
	solution, ok := substitution.Solve(result.Echelon)
	if !ok {
		return markFailure(result)
	}
	result.Solution = solution
	return result
}

// markFailure turns a result into a controlled numeric failure.
func markFailure(result model.Result) model.Result {
	result.Case = model.NumericFailure
	result.Echelon = nil
	result.RankCoefficients = 0
	result.RankAugmented = 0
	result.FreeVariables = 0
	result.ContradictoryRow = -1
	return result
}
