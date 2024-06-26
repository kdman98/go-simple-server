package skeleton

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/scharissis/go-server-skeleton/skeleton/numbers"
)

func TestAnswer(t *testing.T) {
	srv := NewServer(`/api`, numbers.NewMockClient())

	t.Run("GET /answer", func(t *testing.T) {
		got := struct {
			Result string
		}{}
		expected := "Hello! Your lucky number is 42."
		request, err := http.NewRequest(http.MethodGet, "http://localhost:8000/api/answer", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")

		err = json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, got.Result)
	})

	t.Run("POST /answer", func(t *testing.T) {
		got := struct {
			Result string
		}{}
		sent := AsReader(t, struct {
			Name string
		}{
			Name: `Stefano`,
		})
		expected := "Hello, Stefano! Your lucky number is 42."
		request, err := http.NewRequest(http.MethodPost, "http://localhost:8000/api/answer", sent)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")

		err = json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, got.Result)
	})
}
