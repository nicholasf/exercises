package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverse(t *testing.T) {
	r := reverse("abcd")

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

func TestReverser(t *testing.T) {
	req := httptest.NewRequest("GET", "/reverser?arg=abcd", nil)
	rr := httptest.NewRecorder()

	reverser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	if rr.Body.String() != "dcba" {
		t.Errorf("Expected body 'dcba', got '%s'", rr.Body.String())
	}
}
