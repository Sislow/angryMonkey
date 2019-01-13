package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

/*
	pretty simplified testing here.
	Need some more TLC for learning the testing suite in Golang
*/

func TestHandleShutdown(t *testing.T) {
	req, err := http.NewRequest("GET", "/shutdown", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleShutdown)

	handler.ServeHTTP(resRecorder, req)

	if status := resRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandleStats(t *testing.T) {
	s.count = 5
	c.runTime, _ = time.ParseDuration("10s")
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleStats)

	handler.ServeHTTP(resRecorder, req)

	if status := resRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetHandleHash(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleHash)

	handler.ServeHTTP(resRecorder, req)

	if status := resRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// only check for accepting post method, util_test for encryption
func TestPostHandleHash(t *testing.T) {
	postedPass := strings.NewReader("angryMonkey")
	req, err := http.NewRequest("POST", "/hash", postedPass)
	if err != nil {
		t.Fatal(err)
	}

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleHash)

	handler.ServeHTTP(resRecorder, req)

	if status := resRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
