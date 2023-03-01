package main

import (
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestTesting(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(GetUsers))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
}

//need to add more tests
