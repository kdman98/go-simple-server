package skeleton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) answer() http.HandlerFunc {
	type request struct {
		Name string
	}
	type response struct {
		Result string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}

		// Acquire a lucky number
		luckyNum := s.numberClient.Get()

		if r.Method == http.MethodGet {
			resp := response{
				Result: fmt.Sprintf("Hello! Your lucky number is %d.", luckyNum),
			}
			s.respond(w, r, resp, http.StatusOK)
		} else {
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				s.respond(w, r, nil, http.StatusBadRequest)
				return
			}

			resp := response{
				Result: fmt.Sprintf("Hello, %s! Your lucky number is %d.", req.Name, luckyNum),
			}
			s.respond(w, r, resp, http.StatusOK)
		}
	}
}
