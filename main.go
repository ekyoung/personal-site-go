package main

import (
    "github.com/ekyoung/personal-site-go/lib/trips"

    "github.com/ekyoung/personal-site-go/server"
)

func main() {
    tripRepo := trips.NewHardcodedTripRepository()

    server := &server.Server{
        TripRepository: tripRepo,
    }

    server.Run()
}
