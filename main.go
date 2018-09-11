package main

import (
	"os"

	"github.com/timpwbaker/mocking_go/deps"
	"github.com/timpwbaker/mocking_go/posts"
)

func main() {
	post := posts.Post{ID: "fg123sa", Name: "Testing with third parties in Go"}

	os.Setenv("AUDITOR_URL", "https://postman-echo.com/post")
	deps := deps.Resolve("production")

	posts.ValidatePost(&post, deps.Auditor)
}
