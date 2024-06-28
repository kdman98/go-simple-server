package skeleton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type nexonOuidResponse struct {
	Result string `json:"ouid"`
}

// Handler function for answering requests
func (s *Server) searchMatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nickname := r.URL.Query().Get("nickname")
		if nickname != "" {
			apiResponse, isErrorOccurred := s.makeRequestAndCheckError(w, r, nickname)
			if isErrorOccurred {
				s.respond(w, r, apiResponse, http.StatusInternalServerError)
				return
			}
			log.Println(apiResponse)
			s.respond(w, r, apiResponse, http.StatusOK)
		} else {
			s.respond(w, r, "nickname not provided", http.StatusBadRequest)
		}
	}
}

func (s *Server) makeRequestAndCheckError(w http.ResponseWriter, r *http.Request, nickname string) (nexonOuidResponse, bool) {
	apiKey := s.nexonApiKey
	// Make API request to Nexon with the nickname
	nexonURL := fmt.Sprintf("https://open.api.nexon.com/fconline/v1/id?nickname=%s", nickname)
	client := &http.Client{Timeout: 10 * time.Second}
	apiReq, err := http.NewRequest("GET", nexonURL, nil)
	if err != nil {
		log.Println(err)
		return nexonOuidResponse{}, true
	}
	apiReq.Header.Set("x-nxopen-api-key", apiKey)

	resp, err := client.Do(apiReq)
	if err != nil {
		log.Println(err)
		return nexonOuidResponse{}, true
	}
	defer resp.Body.Close() // TODO: find out

	if resp.StatusCode != http.StatusOK {
		s.respond(w, r, resp.Body, resp.StatusCode)
		log.Println(resp)
		return nexonOuidResponse{}, true
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nexonOuidResponse{}, true
	}

	var apiResponse nexonOuidResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Println(err)
		return nexonOuidResponse{}, true
	}
	return apiResponse, false
}
