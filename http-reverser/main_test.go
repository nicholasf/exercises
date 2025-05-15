package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverse(t *testing.T) {
	r := Reverse("abcd")

	if r != "dcba" {
		t.Errorf("Expected 'dcba', got '%s'", r)
	}
}

func TestArgValidator(t *testing.T) {
	tests := []struct {
		req    *http.Request
		rr     *httptest.ResponseRecorder
		status int
	}{
		{
			req:    httptest.NewRequest("GET", "/reverser?arg=abcd", nil),
			rr:     httptest.NewRecorder(),
			status: http.StatusOK,
		},
		{
			req:    httptest.NewRequest("GET", "/reverser", nil),
			rr:     httptest.NewRecorder(),
			status: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.req.URL.String(), func(t *testing.T) {
			rr := test.rr
			handler := ArgValidator((http.HandlerFunc(reverser)))

			handler.ServeHTTP(rr, test.req)

			if status := rr.Code; status != test.status {
				t.Errorf("Expected status code %d, got %d", test.status, status)
			}
		})
	}
}
