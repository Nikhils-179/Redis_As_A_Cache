package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Photo struct {
	AlbumId		 int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func FetchPhotos() ([]Photo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch photos: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var photos []Photo
	err = json.Unmarshal(bodyBytes, &photos)
	if err != nil {
		log.Println("Failed to decode respopnse")
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return photos, nil
}
