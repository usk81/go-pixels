package pixels_test

import (
	"fmt"
	"os"

	"github.com/usk81/go-pixels"
)

func ExampleGetPhoto() {
	token := os.Getenv("PIXELS_API_KEY")

	s, err := pixels.New(nil, token)
	if err != nil {
		panic(err)
	}

	p, err := s.GetPhoto(406014)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

func ExampleSearchPhotos() {
	token := os.Getenv("PIXELS_API_KEY")

	s, err := pixels.New(nil, token)
	if err != nil {
		panic(err)
	}
	ps, err := s.SearchPhotos(pixels.PhotoOptions{
		Query: "dog",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(ps)
}

func ExampleCuratedPhotos() {
	token := os.Getenv("PIXELS_API_KEY")

	s, err := pixels.New(nil, token)
	if err != nil {
		panic(err)
	}
	c, err := s.CuratedPhotos(pixels.PhotoOptions{
		PerPage: 15,
		Page:    1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}

func ExampleGetVideo() {
	token := os.Getenv("PIXELS_API_KEY")

	s, err := pixels.New(nil, token)
	if err != nil {
		panic(err)
	}

	v, err := s.GetVideo(2035391)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func ExampleSearchVideos() {
	token := os.Getenv("PIXELS_API_KEY")

	s, err := pixels.New(nil, token)
	if err != nil {
		panic(err)
	}

	vs, err := s.SearchVideos(pixels.VideoOptions{
		Query: "cat",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(vs)
}

func ExamplePopularVideos() {
	token := os.Getenv("PIXELS_API_KEY")

	s, err := pixels.New(nil, token)
	if err != nil {
		panic(err)
	}

	pv, err := s.PopularVideos(pixels.VideoOptions{
		PerPage: 15,
		Page:    1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(pv)
}
