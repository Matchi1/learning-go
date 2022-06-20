package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestGETRequest(t *testing.T) {
    t.Run("testing request", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/players/Peppers", nil)
        response := httptest.NewRecorder() 

        PlayerServer(response, request)
        got := response.Body.String()
        want := "20"

        if want != got {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        got := response.Body.String()
        want := "10"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}
