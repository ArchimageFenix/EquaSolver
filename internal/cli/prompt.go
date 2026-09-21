// File: prompt.go
// Purpose: print a question, read one line, and repeat until the answer is valid.
// Receives: an input reader, an output writer and a parsing function per question.
// Previous stage: menu.go, interactive.go and file.go (they define each question).
// Next stage: the same callers, which receive the parsed answer.
// Restrictions: no knowledge of systems or of the solver; only asks and reads.
//               An invalid answer prints a short reason and asks again.

package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var errInputEnded = errors.New("no hay más datos de entrada")

type prompter struct {
	in  *bufio.Reader
	out io.Writer
}

func newPrompter(in io.Reader, out io.Writer) prompter {
	return prompter{in: bufio.NewReader(in), out: out}
}

// ask prints the question and returns the trimmed line typed by the user.
func (p prompter) ask(question string) (string, error) {
	fmt.Fprint(p.out, question+" ")
	line, err := p.in.ReadString('\n')
	if err != nil && line == "" {
		if errors.Is(err, io.EOF) {
			return "", errInputEnded
		}
		return "", err
	}
	return strings.TrimSpace(line), nil
}

// tell prints one line of text.
func (p prompter) tell(message string) {
	fmt.Fprintln(p.out, message)
}

// askUntilValid repeats the question until parse accepts the answer.
func askUntilValid[T any](p prompter, question string, parse func(string) (T, error)) (T, error) {
	for {
		text, err := p.ask(question)
		if err != nil {
			var zero T
			return zero, err
		}
		value, parseErr := parse(text)
		if parseErr == nil {
			return value, nil
		}
		p.tell("No válido: " + parseErr.Error())
	}
}
