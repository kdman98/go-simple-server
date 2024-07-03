package skeleton

import (
	"encoding/json"
	"github.com/scharissis/go-server-skeleton/skeleton/enums"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type nexonOuidResponse struct {
	Ouid string `json:"ouid"`
}

// Assuming matchListResponse is just a slice of strings
type matchListResponse []string

const (
	ACCOUNT_INFO_REQUEST_URL = "https://open.api.nexon.com/fconline/v1/id"
	MATCH_LIST_REQUEST_URL   = "https://open.api.nexon.com/fconline/v1/user/match"
)

func (s *Server) matchStatistics() http.HandlerFunc { // TODO: add up functions on demand
	return func(w http.ResponseWriter, r *http.Request) {
		nickname := r.URL.Query().Get("nickname")
		if nickname == "" {
			s.respond(w, r, "nickname not provided", http.StatusBadRequest)
			return
		}

		matchList, err := s.searchMatches(nickname)
		if err != nil {
			log.Println(err)
			s.respond(w, r, "Failed to get match list", http.StatusInternalServerError)
			return
		}

		s.respond(w, r, matchList, http.StatusOK)
	}
}

// Function for searching matches based on the nickname
func (s *Server) searchMatches(nickname string) (matchListResponse, error) {
	ouidResponse, err := s.makeAccountInfoRequest(nickname)
	if err != nil {
		log.Println(err)
		log.Println("searchMatches()")
		return nil, err
	}
	log.Println(ouidResponse.Ouid)

	matchList, err := s.getMatchList(ouidResponse.Ouid, enums.MatchTypeClassic1on1)
	if err != nil {
		log.Println(err)
		log.Println("searchMatches()")
		return nil, err
	}

	return matchList, nil
}

func (s *Server) makeAccountInfoRequest(nickname string) (nexonOuidResponse, error) {
	nexonURL, err := url.Parse(ACCOUNT_INFO_REQUEST_URL)
	if err != nil {
		return nexonOuidResponse{}, err
	}

	// Set query parameters
	params := url.Values{}
	params.Add("nickname", nickname)
	nexonURL.RawQuery = params.Encode()

	resp, err := s.makeAPIRequest(nexonURL.String())
	if err != nil {
		return nexonOuidResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nexonOuidResponse{}, &httpError{resp.StatusCode, "Failed to get account info in API"}
	}

	var apiResponse nexonOuidResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nexonOuidResponse{}, err
	}
	return apiResponse, nil
}

func (s *Server) getMatchList(ouid, matchtype string) (matchListResponse, error) {
	matchListURL, err := url.Parse(MATCH_LIST_REQUEST_URL)
	if err != nil {
		return nil, err
	}

	// Set query parameters
	params := url.Values{}
	params.Add("ouid", ouid)
	params.Add("matchtype", matchtype)
	matchListURL.RawQuery = params.Encode()

	resp, err := s.makeAPIRequest(matchListURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &httpError{resp.StatusCode, "Failed to get match list in API"}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var matchList matchListResponse
	if err := json.Unmarshal(body, &matchList); err != nil {
		return nil, err
	}
	return matchList, nil
}

func (s *Server) makeAPIRequest(url string) (*http.Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	apiReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	apiReq.Header.Set("x-nxopen-api-key", s.nexonApiKey)

	return client.Do(apiReq)
}

type httpError struct {
	StatusCode int
	Message    string
}

func (e *httpError) Error() string {
	return e.Message
}
