package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type RepoLink struct {
	Link  string
	Label string
}

type GitHubProfile struct {
	Repos []RepoLink
	Name  string
	Login string
}

const repoSrc = "https://github.com/AlexMubarakshin?tab=repositories"

func main() {

	profile, err := parseGithubProfile(repoSrc)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("templates/index.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		template.ExecuteTemplate(w, "home", profile)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func parseGithubProfile(url string) (profile GitHubProfile, err error) {
	var name string
	var login string
	var repos []RepoLink

	if doc, err := goquery.NewDocument(url); err != nil {
		return profile, err
	} else {
		doc.Find(".vcard-names-container .vcard-names").Each(func(i int, spans *goquery.Selection) {
			name = spans.Find("span").Get(0).FirstChild.Data
			login = spans.Find("span").Get(1).FirstChild.Data
		})

		doc.Find("#user-repositories-list ul").Each(func(i int, projects *goquery.Selection) {
			projects.Find("h3 a").Each(func(i int, s *goquery.Selection) {
				link, _ := s.Attr("href")
				text := s.Text()
				repos = append(repos, RepoLink{Link: "https://github.com" + link, Label: text})
			})
		})
	}

	profile.Name = name
	profile.Repos = repos
	profile.Login = login

	return profile, nil
}
