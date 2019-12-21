package pixels

import (
	"github.com/google/go-querystring/query"
)

const (
	minPerPage = 15
)

// Photo is the response from pixels photo APIs
type Photo struct {
	ID              int    `json:"id"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	URL             string `json:"url"`
	Photographer    string `json:"photographer"`
	PhotographerURL string `json:"photographer_url"`
	PhotographerID  int    `json:"photographer_id"`
	Src             struct {
		Original  string `json:"original"`
		Large2X   string `json:"large2x"`
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Small     string `json:"small"`
		Portrait  string `json:"portrait"`
		Landscape string `json:"landscape"`
		Tiny      string `json:"tiny"`
	} `json:"src"`
	Liked bool `json:"liked"`
}

// PhotoIndex is the response that is used Search photos API and Curated photos API
type PhotoIndex struct {
	TotalResults int     `json:"total_results"`
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	Photos       []Photo `json:"photos"`
	NextPage     string  `json:"next_page"`
}

// PhotoOptions is define that search photos API's query options
type PhotoOptions struct {
	Query   string `url:"query"`
	PerPage int    `url:"per_page"`
	Page    int    `url:"page"`
}

// ToQuery builds query string from PhotoOptions
func (po PhotoOptions) ToQuery() (q string, err error) {
	if po.PerPage == 0 {
		po.PerPage = minPerPage
	}
	if po.Page < 1 {
		po.Page = 1
	}
	v, err := query.Values(po)
	if err != nil {
		return
	}
	return v.Encode(), nil
}
