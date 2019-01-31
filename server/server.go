package server

import (
	"fmt"
	"net/http"
	"text/template"

	"../models"
)

var server http.Server

func Init(githubProfile models.GitHubProfile) {
	server = http.Server{
		Addr: ":8080",
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("templates/index.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		template.ExecuteTemplate(w, "home", githubProfile)
	})

	fmt.Println("Server started at :8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
