package main

import (
	"github.com/Calebbuffleben/api-with-golang/router"
	"github.com/Calebbuffleben/api-with-golang/config"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}