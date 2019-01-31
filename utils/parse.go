package utils

import (
	"../models"
	"github.com/PuerkitoBio/goquery"
)

func ParseGithubProfile(url string) (profile models.GitHubProfile, err error) {
	var name string
	var login string
	var repos []models.RepoLink

	if doc, err := goquery.NewDocument(url); err != nil {
		return profile, err
	} else {
		doc.Find(".vcard-names-container .vcard-names").Each(func(i int, spans *goquery.Selection) {
			if spans.Find("span").Get(0).FirstChild != nil {
				name = spans.Find("span").Get(0).FirstChild.Data
			}

			login = spans.Find("span").Get(1).FirstChild.Data
		})

		doc.Find("#user-repositories-list ul").Each(func(i int, projects *goquery.Selection) {
			projects.Find("h3 a").Each(func(i int, s *goquery.Selection) {
				link, _ := s.Attr("href")
				text := s.Text()
				repos = append(repos, models.RepoLink{Link: "https://github.com" + link, Label: text})
			})
		})
	}

	profile.Name = name
	profile.Repos = repos
	profile.Login = login

	return profile, nil
}
