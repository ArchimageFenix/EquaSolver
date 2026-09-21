// File: parse.go
// Purpose: turn user text into finite numbers using strict decimal syntax.
// Receives: raw text typed by the user, read from a file or sent by a web form.
// Previous stage: cli (typed data, file, arguments) and web/formcheck.go.
// Next stage: model.FromAugmentedRows (rows of numbers) in the calling adapter.
// Restrictions: accepts only plain decimals with a dot (no commas, letters, hex,
//               NaN or Inf); it does not know about systems, CLI or HTTP.

package numeric

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var decimalNumber = regexp.MustCompile(`^[+-]?([0-9]+(\.[0-9]*)?|\.[0-9]+)([eE][+-]?[0-9]+)?$`)

// ParseNumber converts text into a finite real number.
func ParseNumber(text string) (float64, error) {
	if !decimalNumber.MatchString(text) {
		return 0, fmt.Errorf("%q no es un número válido (use punto decimal, sin letras ni símbolos)", text)
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil || !IsFinite(value) {
		return 0, fmt.Errorf("%q está fuera del rango admitido", text)
	}
	return value, nil
}

// ParseRow converts a line of numbers separated by spaces into exactly want values.
func ParseRow(line string, want int) ([]float64, error) {
	fields := strings.Fields(line)
	if len(fields) != want {
		return nil, fmt.Errorf("se esperaban %d números y se recibieron %d", want, len(fields))
	}
	values := make([]float64, want)
	for i, field := range fields {
		value, err := ParseNumber(field)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}
	return values, nil
}

// ParseRows converts each line into a row of n+1 numbers; errors name the equation.
func ParseRows(lines []string, n int) ([][]float64, error) {
	rows := make([][]float64, 0, len(lines))
	for i, line := range lines {
		row, err := ParseRow(line, n+1)
		if err != nil {
			return nil, fmt.Errorf("ecuación %d: %w", i+1, err)
		}
		rows = append(rows, row)
	}
	return rows, nil
}

// SplitRows splits text into non-empty, trimmed rows separated by line breaks or ';'.
func SplitRows(text string) []string {
	normalized := strings.NewReplacer("\r\n", "\n", "\r", "\n", ";", "\n").Replace(text)
	rows := []string{}
	for _, line := range strings.Split(normalized, "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			rows = append(rows, trimmed)
		}
	}
	return rows
}
