package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nikhils-179/redis-caching/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	var addr = flag.String("addr", ":8081", "The address of the application")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Panic("Unable to load env files")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	fmt.Println(rdb)


	router := mux.NewRouter()
	router.HandleFunc("/photos", handlers.GetPhotos(ctx, rdb)).Methods("GET")

	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, router); err != nil {
		log.Fatal("Listen and serve:", err)
	}


}
