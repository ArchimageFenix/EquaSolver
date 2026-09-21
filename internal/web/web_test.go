// File: web_test.go
// Purpose: check the form validation, the handlers and that invalid data never
//          reaches the core.
// Receives: simulated HTTP requests.
// Previous stage: formcheck.go, handlers.go and present.go.
// Next stage: none (test result).
// Restrictions: tests only; no real network.

package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func postMatrix(text string) *httptest.ResponseRecorder {
	body := url.Values{"matrix": {text}}.Encode()
	request := httptest.NewRequest(http.MethodPost, "/solve", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	solveHandler(recorder, request)
	return recorder
}

func TestCheckFormAcceptsWellFormedText(t *testing.T) {
	system, err := checkForm("1 1 1 6\n2 -1 1 3\n1 2 -1 2")
	if err != nil || system.N != 3 {
		t.Errorf("got n=%d, err=%v", system.N, err)
	}
}

func TestCheckFormRejectsInvalidText(t *testing.T) {
	for _, text := range []string{"", "1 2 x\n3 4 5", "1,5 2 3\n4 5 6", "1 2\n3 4 5", "1 2 3 4"} {
		if _, err := checkForm(text); err == nil {
			t.Errorf("%q should be rejected", text)
		}
	}
}

func TestIndexShowsTheForm(t *testing.T) {
	recorder := httptest.NewRecorder()
	indexHandler(recorder, httptest.NewRequest(http.MethodGet, "/", nil))
	if recorder.Code != http.StatusOK || !strings.Contains(recorder.Body.String(), "<form") {
		t.Errorf("status %d", recorder.Code)
	}
}

func TestIndexUnknownPathIsNotFound(t *testing.T) {
	recorder := httptest.NewRecorder()
	indexHandler(recorder, httptest.NewRequest(http.MethodGet, "/nope", nil))
	if recorder.Code != http.StatusNotFound {
		t.Errorf("status %d", recorder.Code)
	}
}

func TestSolveShowsTheResult(t *testing.T) {
	recorder := postMatrix("1 1 1 6\n2 -1 1 3\n1 2 -1 2")
	body := recorder.Body.String()
	if recorder.Code != http.StatusOK || !strings.Contains(body, "SOLUCIÓN ÚNICA") || !strings.Contains(body, "x1 = 1.00") {
		t.Errorf("status %d, body:\n%s", recorder.Code, body)
	}
}

func TestSolveRejectsInvalidDataWithAMessage(t *testing.T) {
	recorder := postMatrix("1 2 x\n3 4 5")
	body := recorder.Body.String()
	if recorder.Code != http.StatusBadRequest || !strings.Contains(body, "no es un número válido") {
		t.Errorf("status %d, body:\n%s", recorder.Code, body)
	}
	if strings.Contains(body, "Resultado:") {
		t.Error("invalid data must not produce a result")
	}
}

func TestSolveRejectsOversizedBodies(t *testing.T) {
	recorder := postMatrix(strings.Repeat("1", maxBodyBytes+10))
	if recorder.Code != http.StatusBadRequest || !strings.Contains(recorder.Body.String(), "demasiado grandes") {
		t.Errorf("status %d", recorder.Code)
	}
}

func TestSolveOnlyAcceptsPost(t *testing.T) {
	recorder := httptest.NewRecorder()
	solveHandler(recorder, httptest.NewRequest(http.MethodGet, "/solve", nil))
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("status %d", recorder.Code)
	}
}
