package auditor_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/timpwbaker/mocking_go/auditor"
	"github.com/timpwbaker/mocking_go/pkg/httputil"
)

var posted m

type m map[string]interface{}

func TestAudit(t *testing.T) {
	testServer := httptest.NewServer(

		// NewServer takes a handler.
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Inside the handler we define our canned responses,
			// switching on URL and then http method
			switch r.URL.Path {
			case "/reports":
				switch r.Method {
				case "POST":

					// Here we read the body that has been posted to our test server
					// and save it to a variable, we can assert against this variable later.
					body, _ := ioutil.ReadAll(r.Body)
					must(t, json.Unmarshal(body, &posted))
					httputil.SendJSON(w, 200, m{})

					// Finally provide some defaults, I generally just use a 404
				default:
					httputil.SendJSON(w, 404, nil)
				}
			default:
				httputil.SendJSON(w, 404, nil)
			}
		}),
	)
	defer testServer.Close()

	// Build your auditorclient configured for your test by using the testServer.
	auditorclient := auditor.LoadClient(testServer.URL + "/reports")

	// Call the Audit function on the actual instance of *RealClient
	auditorclient.Audit("Validate Post", "74561")

	// The client will hit our test web server defined above, and save the payload as a map in the variable 'posted'. We can then assert against it.
	if posted["event"] != "Validate Post" {
		t.Error("Posted event should be 'Validate Post'")
	}

	// We know that we've made a real http request, albeit to a test server, and that our auditorclient package is behaving as it is expected to.
	if posted["user_id"] != "74561" {
		t.Error("Posted user_id should be '74561'")
	}
}

func TestAuditAuthenticated(t *testing.T) {
	// Set some environment variables that the client will use to make
	// authenticated requests.
	os.Setenv("AUDITOR_USERNAME", "foobar")
	os.Setenv("AUDITOR_PASSWORD", "baz")

	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/reports":
				switch r.Method {
				case "POST":

					// Here you can assert that the http basic auth credentials passed
					// in to the test third party service are what you expect them to be,
					// if they're not you can fail the test.
					user, password, ok := r.BasicAuth()
					if ok != true {
						t.Error("Invalid http basic auth credentials")
					}
					if user != "foobar" {
						t.Error("Incorrect http basic username")
					}
					if password != "baz" {
						t.Error("Incorrect http basic password")
					}

					// Beyond this the test is exactly the same as above
					body, _ := ioutil.ReadAll(r.Body)
					must(t, json.Unmarshal(body, &posted))
					httputil.SendJSON(w, 200, m{})
				default:
					httputil.SendJSON(w, 404, nil)
				}
			default:
				httputil.SendJSON(w, 404, nil)
			}
		}),
	)
	defer testServer.Close()

	auditorclient := auditor.LoadClient(testServer.URL + "/reports")
	auditorclient.AuditAuthenticated("Validate Post", "74561")

	if posted["event"] != "Validate Post" {
		t.Error("Posted event should be 'Validate Post'")
	}
	if posted["user_id"] != "74561" {
		t.Error("Posted user_id should be '74561'")
	}
}

func must(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
