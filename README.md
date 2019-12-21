# Go-Pixels

Pixels API client for Gopher

## Install

```bash
go get -u github.com/usk81/go-pixels
```

## Usage

``` go
// create new instance
s, _ := pixels.New(http.DefaultClient, "your pixels api key")

// Get a photo
result, err := s.GetPhoto(406014)

// Search photos
result, err := s.SearchPhotos(pixels.PhotoOptions{
		Query: "dog",
	})

// Get curated photos
result, err := s.CuratedPhotos(pixels.PhotoOptions{
		PerPage: 15,
		Page:    1,
	})

// Get a video
result, err := s.GetVideo(2035391)

// Search videos
result, err := s.SearchVideos(pixels.VideoOptions{
		Query: "cat",
	})

// Get popular videos
result, err := s.PopularVideos(pixels.VideoOptions{
		PerPage: 15,
		Page:    1,
	})
```