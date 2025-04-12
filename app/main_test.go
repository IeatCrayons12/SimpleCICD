package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	payload := `{"name":"Alice","age":25}`
	req := httptest.NewRequest(http.MethodPost, "/submit", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handle(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 but got %v", status)
	}

	expected := `{"message":"Hello Alice, age 25!"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("expected body %s but got %s", expected, rr.Body.String())
	}
}
