package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthZ(t *testing.T) {
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/api/v1/healthz", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthZ)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
