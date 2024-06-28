package skeleton

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type nexonOuidResponse struct {
	Ouid string `json:"ouid"`
}

const ACCOUNT_INFO_REQUEST_URL = "https://open.api.nexon.com/fconline/v1/id"
const MATCH_LIST_REQUEST_URL = "https://open.api.nexon.com/fconline/v1/user/match"

// Handler function for answering requests
func (s *Server) searchMatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nickname := r.URL.Query().Get("nickname")
		if nickname != "" {
			ouidResponse, isErrorOccurred := s.makeAccountInfoRequestAndCheckError(w, r, nickname)
			if isErrorOccurred {
				s.respond(w, r, ouidResponse, http.StatusInternalServerError)
				return
			}
			log.Println(ouidResponse.Ouid)
			s.respond(w, r, ouidResponse, http.StatusOK)
		} else {
			s.respond(w, r, "nickname not provided", http.StatusBadRequest)
		}
	}
}

func (s *Server) makeAccountInfoRequestAndCheckError(w http.ResponseWriter, r *http.Request, nickname string) (nexonOuidResponse, bool) {
	apiKey := s.nexonApiKey
	// Make API request to Nexon with the nickname
	nexonURL := ACCOUNT_INFO_REQUEST_URL + "?nickname=" + nickname // TODO: refactor
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
