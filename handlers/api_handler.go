package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Nikhils-179/redis-caching/services"
	"github.com/Nikhils-179/redis-caching/utils"
	"github.com/redis/go-redis/v9"
)

func GetPhotos(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cacheKey := "photos"

		// caching layer
		cachedPhotos, err := utils.GetCache(ctx, rdb, cacheKey)
		if err != nil {
			log.Printf("Error getting cache: %v", err)
			http.Error(w, "Failed to get cache", http.StatusInternalServerError)
			return
		}

		if cachedPhotos != nil {
			log.Println("Cache hit")
			w.Header().Set("Content-Type", "application/json")
			w.Write(cachedPhotos)
			return
		}
		log.Println("Cache miss")

		// fetch data from 3rd party API
		photos, err := services.FetchPhotos()
		if err != nil {
			log.Println("Failed to fetch photos")
			http.Error(w, "Failed to fetch photos", http.StatusInternalServerError)
			return
		}

		// Marshal the photos data to JSON
		photosJSON, err := json.Marshal(photos)
		if err != nil {
			http.Error(w, "Failed to marshal photos", http.StatusInternalServerError)
			return
		}

		// Store the data into the Redis cache
		err = utils.SetCache(ctx, rdb, cacheKey, photosJSON, time.Hour)
		if err != nil {
			log.Printf("Failed to set cache: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(photosJSON)
	}
}
