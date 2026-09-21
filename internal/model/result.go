// File: result.go
// Purpose: define the neutral Result returned by the solver and its four cases.
// Receives: nothing (data definition only).
// Previous stage: solver (builds it from elimination and classification).
// Next stage: cli/present.go and web/present.go.
// Restrictions: only data; no logic and no presentation text.

package model

// Case is the classification of a system.
type Case int

const (
	UniqueSolution Case = iota
	InfiniteSolutions
	NoSolution
	NumericFailure
)

// Result is everything the interfaces need to show. ContradictoryRow is 0-based
// and -1 when it does not apply. Solution is set only for UniqueSolution.
// Echelon is nil for NumericFailure.
type Result struct {
	Case             Case
	Original         [][]float64
	Steps            []Step
	Echelon          [][]float64
	RankCoefficients int
	RankAugmented    int
	FreeVariables    int
	Solution         []float64
	ContradictoryRow int
}
