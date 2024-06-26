package skeleton

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/scharissis/go-server-skeleton/skeleton/numbers"
)

type server struct {
	urlPrefix    string
	router       *mux.Router
	numberClient numbers.Client
}

func NewServer(urlPrefix string, nClient numbers.Client) *server {
	s := &server{urlPrefix: urlPrefix, numberClient: nClient}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("failed to encode response: %v", err)
			http.Error(w, "Something went wrong with JSON marshalling.", http.StatusInternalServerError)
		}
	}
}

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
