package skeleton

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) routes() {
	s.router = mux.NewRouter()
	s.router.HandleFunc(s.urlPrefix+"/answer", restrictMethods(s.answer(), http.MethodGet, http.MethodPost))
}
