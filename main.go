package main

import (
	"github.com/Koliras/go-server/api"
	"github.com/Koliras/go-server/config"
)

func main() {
    db := config.NewDB()
    server := api.NewServer(":8080", db)

    err := server.Start()
    if err != nil {
        panic(err)
    }
}

