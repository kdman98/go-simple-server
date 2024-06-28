// Handler unit tests.
package skeleton

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/scharissis/go-server-skeleton/skeleton/numbers"
	"github.com/stretchr/testify/assert"
)

func TestAnswerHandler(t *testing.T) {
	srv := NewServer(``, numbers.NewMockClient())

	t.Run("GET /searchMatches", func(t *testing.T) {
		got := struct {
			Result string
		}{}
		expected := "Hello! Your lucky number is 42."
		request := httptest.NewRequest(http.MethodGet, "/searchMatches", nil)
		response := httptest.NewRecorder()
		srv.searchMatches()(response, request)

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")
		assert.Equal(t, expected, got.Result)
	})

	t.Run("POST /searchMatches", func(t *testing.T) {
		got := struct {
			Result string
		}{}
		sent := AsReader(t, struct {
			Name string
		}{
			Name: `Stefano`,
		})
		expected := "Hello, Stefano! Your lucky number is 42."
		request := httptest.NewRequest(http.MethodPost, "/searchMatches", sent)
		response := httptest.NewRecorder()
		srv.searchMatches()(response, request)

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")
		assert.Equal(t, expected, got.Result)
	})
}

func AsReader(t *testing.T, s interface{}) io.Reader {
	structBytes, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	return bytes.NewReader(structBytes)
}
