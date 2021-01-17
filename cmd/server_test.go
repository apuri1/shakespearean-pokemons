package main

import (
	"net/http"
	"net/http/httptest"
	"shakespeareanpokemons/shakespearean"
	"testing"
)

func TestGoodRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/pokemon/charizard", nil)

	rr := httptest.NewRecorder()

	shakespearean.HandleIncoming(rr, req, incomingMetrics)

	if rr.Code != http.StatusOK {
		if rr.Code != http.StatusTooManyRequests { //website is throttling requests
			t.Errorf("Wrong status code, got %v,expect %v or %v", rr.Code, http.StatusOK, http.StatusTooManyRequests)
		}
	}
}

func TestBadRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/boo/charizard", nil)

	rr := httptest.NewRecorder()

	shakespearean.HandleIncoming(rr, req, incomingMetrics)

	if rr.Code != http.StatusNotFound {
		if rr.Code != http.StatusTooManyRequests { //website is throttling requests
			t.Errorf("Wrong status code, got %v,expect %v or %v", rr.Code, http.StatusOK, http.StatusTooManyRequests)
		}
	}
}

func TestIncompleteRequest(t *testing.T) {

	req, _ := http.NewRequest("GET", "/pokemon", nil)

	rr := httptest.NewRecorder()

	shakespearean.HandleIncoming(rr, req, incomingMetrics)

	if rr.Code != http.StatusNotFound {
		if rr.Code != http.StatusTooManyRequests { //website is throttling requests
			t.Errorf("Wrong status code, got %v,expect %v or %v", rr.Code, http.StatusNotFound, http.StatusTooManyRequests)
		}
	}
}
