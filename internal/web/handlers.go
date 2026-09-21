// File: handlers.go
// Purpose: receive the form, delegate to formcheck and solver, and return the page.
// Receives: HTTP requests from the browser.
// Previous stage: server.go (routing).
// Next stage: formcheck.go (input check), solver.Solve (core) and present.go (page).
// Restrictions: does not calculate and does not apply domain rules by itself; data
//               that fails the check never reaches the solver; the request body is
//               limited in size.

package web

import (
	"errors"
	"net/http"

	"gauss/internal/solver"
)

// maxBodyBytes limits the size of a request so a huge input cannot exhaust memory.
const maxBodyBytes = 1 << 20

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderPage(w, http.StatusOK, pageView{})
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}
	text, err := readMatrixField(w, r)
	if err != nil {
		renderError(w, text, err)
		return
	}
	system, err := checkForm(text)
	if err != nil {
		renderError(w, text, err)
		return
	}
	result, err := solver.Solve(system)
	if err != nil {
		renderError(w, text, err)
		return
	}
	renderPage(w, http.StatusOK, pageView{Input: text, Result: newResultView(result)})
}

func readMatrixField(w http.ResponseWriter, r *http.Request) (string, error) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	if err := r.ParseForm(); err != nil {
		return "", errors.New("los datos son demasiado grandes o no se pudieron leer")
	}
	return r.PostFormValue("matrix"), nil
}

func renderError(w http.ResponseWriter, text string, err error) {
	renderPage(w, http.StatusBadRequest, pageView{Input: text, Error: err.Error()})
}
