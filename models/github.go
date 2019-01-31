package models

type RepoLink struct {
	Link  string
	Label string
}

type GitHubProfile struct {
	Repos []RepoLink
	Name  string
	Login string
}
