package skeleton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Respond method to write the response
func (s *Server) respondNexon(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Handler function for answering requests
func (s *Server) answer() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	type response struct {
		Result string `json:"result"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}

		if r.Method == http.MethodGet {
			nickname := r.URL.Query().Get("nickname")
			apiKey := s.apiKey
			if nickname != "" && apiKey != "" {
				// Make API request to Nexon with the nickname
				nexonURL := fmt.Sprintf("https://open.api.nexon.com/fconline/v1/id?nickname=%s", nickname)
				client := &http.Client{Timeout: 10 * time.Second}
				apiReq, err := http.NewRequest("GET", nexonURL, nil)
				if err != nil {
					s.respond(w, r, nil, http.StatusInternalServerError)
					log.Println(err)
					return
				}
				apiReq.Header.Set("x-nxopen-api-key", apiKey)

				resp, err := client.Do(apiReq)
				if err != nil {
					s.respond(w, r, nil, http.StatusInternalServerError)
					log.Println(err)
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					s.respond(w, r, resp.Body, resp.StatusCode)
					log.Println(resp)
					return
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					s.respond(w, r, nil, http.StatusInternalServerError)
					log.Println(err)
					return
				}

				var apiResponse map[string]interface{}
				if err := json.Unmarshal(body, &apiResponse); err != nil {
					s.respond(w, r, nil, http.StatusInternalServerError)
					log.Println(err)
					return
				}
				//log.Println(apiResponse)
				s.respond(w, r, apiResponse, http.StatusOK)
			} else {
				s.respond(w, r, response{Result: "Nickname or apiKey not provided."}, http.StatusBadRequest)
			}
		} else {
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				s.respond(w, r, nil, http.StatusBadRequest)
				return
			}

			resp := response{
				Result: fmt.Sprintf("Hello, %s!", req.Name),
			}
			s.respond(w, r, resp, http.StatusOK)
		}
	}
}
