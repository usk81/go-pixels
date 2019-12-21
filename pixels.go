package pixels

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version        = "0.0.1"
	defaultBaseURL = "https://api.pexels.com/"
	userAgent      = "go-pixels/" + version
)

// Service is Fortnite tracker API instance
type Service struct {
	client    *http.Client
	token     string
	BaseURL   *url.URL // API endpoint base URL
	UserAgent string   // optional additional User-Agent fragment
}

// ServiceOpt are options for New.
type ServiceOpt func(*Service) error

// NewService creates a new Service.
func NewService(client *http.Client, token string) (s *Service, err error) {
	if client == nil {
		client = http.DefaultClient
	}
	if token == "" {
		return nil, errors.New("token is empty")
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Service{
		client:    client,
		token:     token,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}, nil
}

// New creates a new Service instance
func New(client *http.Client, token string, opts ...ServiceOpt) (s *Service, err error) {
	if s, err = NewService(client, token); err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

// SetBaseURL is a service option for setting the base URL.
func SetBaseURL(bu string) ServiceOpt {
	return func(s *Service) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}
		s.BaseURL = u
		return nil
	}
}

// SetUserAgent is a service option for setting the user agent.
func SetUserAgent(ua string) ServiceOpt {
	return func(s *Service) error {
		if ua != "" {
			s.UserAgent = ua
		}
		return nil
	}
}

func (s *Service) Do(urlStr string, v interface{}) (response *http.Response, err error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	u := s.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", s.token)
	req.Header.Add("User-Agent", s.UserAgent)

	response, err = s.client.Do(req)
	if err != nil {
		return
	}
	if response == nil {
		err = errors.New("unknown error")
		return
	}
	defer func() {
		if response.Body != nil {
			if rerr := response.Body.Close(); err == nil {
				err = rerr
			}
		}
	}()

	if err = CheckResponse(response); err != nil {
		return
	}

	if v != nil {
		err = json.NewDecoder(response.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (s *Service) SearchPhotos(o PhotoOptions) (result PhotoIndex, err error) {
	if o.Query == "" {
		err = errors.New("Query parameter is required")
		return
	}
	r := "v1/search"
	q, err := o.ToQuery()
	if err != nil {
		return
	}
	if q != "" {
		r = r + "?" + q
	}
	_, err = s.Do(r, &result)
	return
}

func (s *Service) CuratedPhotos(o PhotoOptions) (result PhotoIndex, err error) {
	r := "v1/curated"
	q, err := o.ToQuery()
	if err != nil {
		return
	}
	if q != "" {
		r = r + "?" + q
	}
	_, err = s.Do(r, &result)
	return
}

func (s *Service) GetPhoto(id int) (result Photo, err error) {
	_, err = s.Do(fmt.Sprintf("v1/photos/%d", id), &result)
	return
}

func (s *Service) SearchVideos(o VideoOptions) (result VideoIndex, err error) {
	if o.Query == "" {
		err = errors.New("Query parameter is required")
		return
	}
	r := "videos/search"
	q, err := o.ToQuery()
	if err != nil {
		return
	}
	if q != "" {
		r = r + "?" + q
	}
	_, err = s.Do(r, &result)
	return
}

func (s *Service) PopularVideos(o VideoOptions) (result VideoIndex, err error) {
	r := "videos/popular"
	q, err := o.ToQuery()
	if err != nil {
		return
	}
	if q != "" {
		r = r + "?" + q
	}
	_, err = s.Do(r, &result)
	return
}

func (s *Service) GetVideo(id int) (result Video, err error) {
	_, err = s.Do(fmt.Sprintf("videos/videos/%d", id), &result)
	return
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Message,
	)
}

// CheckResponse checks the API response for errors, and returns them if present
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
