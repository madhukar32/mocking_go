package deps

import (
	"os"

	"github.com/timpwbaker/mocking_go/auditor"
)

var Test string = "test"

type Dependencies struct {
	Auditor auditor.Client
}

func Resolve(env string) *Dependencies {
	deps := new(Dependencies)
	if env == Test {
		deps.Auditor = auditor.LoadMock()
	} else {
		requestURL := os.Getenv("AUDITOR_URL")
		deps.Auditor = auditor.LoadClient(requestURL)
	}

	return deps
}
