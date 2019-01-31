package main

import (
	"flag"
	"fmt"

	"./server"
	"./utils"
)

func main() {
	nickname := flag.String("nickname", "AlexMubarakshin", "Github user nickname")
	flag.Parse()

	fmt.Println("Start scraping", *nickname, "profile")

	repositoriesUrl := utils.BuildGithubUrlByUsername(*nickname)

	profile, err := utils.ParseGithubProfile(repositoriesUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.Init(profile)
}
