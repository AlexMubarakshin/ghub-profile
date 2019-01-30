package utils

func BuildGithubUrlByUsername(username string) (profileUrl string) {
	const githubHostUrl = "https://github.com/"
	const queryUrl = "?tab=repositories"
	return githubHostUrl + username + queryUrl
}
