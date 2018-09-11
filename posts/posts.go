package posts

import (
	"github.com/timpwbaker/mocking_go/auditor"
)

type Post struct {
	ID   string
	Name string
}

func ValidatePost(p *Post, auditorclient auditor.Client) bool {
	auditorclient.Audit("Validate Post", p.ID)

	return validate(p)
}

func validate(p *Post) bool {
	return p.Name != ""
}
