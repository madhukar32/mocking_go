package posts_test

import (
	"testing"

	"github.com/timpwbaker/mocking_go/deps"
	"github.com/timpwbaker/mocking_go/posts"
)

func TestValidatePost(t *testing.T) {
	// First we must resolve our dependecies as the mocked implementations.
	deps := deps.Resolve(deps.Test)

	// The deps require an implementation of the auditorclient.Client interface,
	// in this case our resolver returns the mocked implementation defined above.
	auditorclient := deps.Auditor
	post := posts.Post{ID: "123ab9x", Name: "Testing with third parties in Go"}

	// This code path calls auditorclient.Audit, but the client is the mocked version.
	valid := posts.ValidatePost(&post, auditorclient)

	// Using the mocked version of the auditorclient means we can assert
	// against what we care about - that the post is valid, practically
	// ignoring the auditorclient all together.
	if valid != true {
		t.Error("Should be valid")
	}
}
