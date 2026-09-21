// File: args.go
// Purpose: interpret the execution arguments (-f file, -m matrix).
// Receives: the arguments after the program name.
// Previous stage: run.go.
// Next stage: run.go, which picks the source of the data from the result.
// Restrictions: does not read files or matrices and does not print; -f and -m
//               cannot be used together.

package cli

import (
	"errors"
	"flag"
	"io"
)

type arguments struct {
	file   string
	matrix string
}

func parseArgs(args []string) (arguments, error) {
	flags := flag.NewFlagSet("gauss", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	var parsed arguments
	flags.StringVar(&parsed.file, "f", "", "")
	flags.StringVar(&parsed.matrix, "m", "", "")
	if err := flags.Parse(args); err != nil || flags.NArg() > 0 {
		return arguments{}, errors.New(usageText)
	}
	if parsed.file != "" && parsed.matrix != "" {
		return arguments{}, errors.New("use -f o -m, no ambos.\n" + usageText)
	}
	return parsed, nil
}
