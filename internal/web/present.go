// File: present.go
// Purpose: build the view models and render the page templates.
// Receives: a pageView (input text, error message or a model.Result).
// Previous stage: handlers.go.
// Next stage: the browser (HTML response).
// Restrictions: no calculations; every number and text goes through package format,
//               so the web shows the same as the CLI; templates hold no domain logic.

package web

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"gauss/internal/format"
	"gauss/internal/model"
)

//go:embed templates/*.html
var templateFiles embed.FS

var pages = template.Must(template.ParseFS(templateFiles, "templates/*.html"))

type pageView struct {
	Input  string
	Error  string
	Result *resultView
}

type resultView struct {
	Original [][]string
	Steps    []stepView
	Echelon  [][]string
	Title    string
	Detail   []string
	Solution []string
}

type stepView struct {
	Title      string
	Operations []string
	Matrix     [][]string
}

func renderPage(w http.ResponseWriter, status int, view pageView) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := pages.ExecuteTemplate(w, "index.html", view); err != nil {
		log.Printf("render error: %v", err)
	}
}

func newResultView(result model.Result) *resultView {
	return &resultView{
		Original: format.Matrix(result.Original),
		Steps:    newStepViews(result.Steps),
		Echelon:  format.Matrix(result.Echelon),
		Title:    format.CaseTitle(result.Case),
		Detail:   format.VerdictLines(result),
		Solution: format.SolutionLines(result.Solution),
	}
}

func newStepViews(steps []model.Step) []stepView {
	views := make([]stepView, len(steps))
	for i, step := range steps {
		views[i] = stepView{
			Title:      format.StepTitle(i+1, step),
			Operations: operationLines(step),
			Matrix:     format.Matrix(step.Matrix),
		}
	}
	return views
}

func operationLines(step model.Step) []string {
	lines := make([]string, len(step.Operations))
	for i, operation := range step.Operations {
		lines[i] = format.OperationText(operation, step.PivotRow)
	}
	return lines
}
