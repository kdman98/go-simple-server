package skeleton

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) routes() {
	s.router = mux.NewRouter()
	s.router.HandleFunc(s.urlPrefix+"/match-statistics", restrictMethods(s.matchStatistics(), http.MethodGet, http.MethodPost))
}
