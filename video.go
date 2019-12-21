package pixels

import "github.com/google/go-querystring/query"

// Video is the response from pixels video APIs
type Video struct {
	ID       int           `json:"id"`
	Width    int           `json:"width"`
	Height   int           `json:"height"`
	URL      string        `json:"url"`
	Image    string        `json:"image"`
	FullRes  interface{}   `json:"full_res"`
	Tags     []interface{} `json:"tags"`
	Duration int           `json:"duration"`
	User     struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"user"`
	VideoFiles []struct {
		ID       int    `json:"id"`
		Quality  string `json:"quality"`
		FileType string `json:"file_type"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		Link     string `json:"link"`
	} `json:"video_files"`
	VideoPictures []struct {
		ID      int    `json:"id"`
		Picture string `json:"picture"`
		Nr      int    `json:"nr"`
	} `json:"video_pictures"`
}

// VideoIndex is the response that is used Search videos API and Popular videos
type VideoIndex struct {
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	TotalResults int     `json:"total_results"`
	URL          string  `json:"url"`
	Videos       []Video `json:"videos"`
}

// VideoOptions is define that search video API's query options
type VideoOptions struct {
	Query       string `url:"query"`
	PerPage     int    `url:"per_page"`
	Page        int    `url:"page"`
	MinWidth    string `url:"min_width"`
	MaxWidth    string `url:"max_width"`
	MinDuration string `url:"min_duration"`
	MaxDuration string `url:"max_duration"`
}

// ToQuery builds query string from VideoOptions
func (vo VideoOptions) ToQuery() (q string, err error) {
	if vo.PerPage == 0 {
		vo.PerPage = minPerPage
	}
	if vo.Page < 1 {
		vo.Page = 1
	}
	v, err := query.Values(vo)
	if err != nil {
		return
	}
	return v.Encode(), nil
}
