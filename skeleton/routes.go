package skeleton

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) routes() {
	s.router = mux.NewRouter()
	s.router.HandleFunc(s.urlPrefix+"/search-matches", restrictMethods(s.searchMatches(), http.MethodGet, http.MethodPost))
}
