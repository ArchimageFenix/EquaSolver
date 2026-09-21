// File: menu.go
// Purpose: show the menu of input forms and route to the chosen one.
// Receives: a prompter (console input and output).
// Previous stage: run.go (only when no file or matrix argument was given).
// Next stage: interactive.go (typed data) or file.go (data from a file).
// Restrictions: does not read the system itself and does not calculate; it is
//               skipped when the data arrive through arguments.

package cli

import (
	"errors"

	"gauss/internal/model"
)

type inputMode int

const (
	modeTyped inputMode = iota + 1
	modeFile
)

// readInteractively shows the menu and reads the system in the chosen way.
func readInteractively(p prompter) (model.System, error) {
	mode, err := askMode(p)
	if err != nil {
		return model.System{}, err
	}
	if mode == modeFile {
		return readSystemFromPrompt(p)
	}
	return readSystemTyped(p)
}

func askMode(p prompter) (inputMode, error) {
	for _, line := range menuLines {
		p.tell(line)
	}
	return askUntilValid(p, modeQuestion, parseMode)
}

func parseMode(text string) (inputMode, error) {
	switch text {
	case "1":
		return modeTyped, nil
	case "2":
		return modeFile, nil
	}
	return 0, errors.New("elija 1 o 2")
}
