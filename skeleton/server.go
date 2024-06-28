package skeleton

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	urlPrefix   string
	router      *mux.Router
	nexonApiKey string
}

func NewServer(urlPrefix string) *Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get API key from environment variables
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}
	s := &Server{urlPrefix: urlPrefix, nexonApiKey: apiKey}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("failed to encode nexonOuidResponse: %v", err)
			http.Error(w, "Something went wrong with JSON marshalling.", http.StatusInternalServerError)
		}
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
