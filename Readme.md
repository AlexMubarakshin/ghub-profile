# ghub-profile

Scrap my github profile and display personal github projects.

<p align="center">
  <img src="./screenshots/img.png" alt="Screenshot"
       width="50%">
</p>

## Usage
Replace in `ghub.go` `const repoSrc` value to your Github profile URL, for example
```go
# From:
const repoSrc = "https://github.com/AlexMubarakshin?tab=repositories"
# To:
const repoSrc = "https://github.com/gaearon?tab=repositories"
```
Then build application
```
$ go build ./ghub.go
$ ./ghub.exe

# Open localhost:8080
```
