package auditor

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Client interface {
	Audit(event string, userID string) error
	AuditAuthenticated(event string, userID string) error
}

type RealClient struct {
	RequestURL string
}

func (rc *RealClient) Audit(event string, userID string) error {
	payload := buildPayload(event, userID)

	req, err := http.NewRequest("POST", rc.RequestURL, payload)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Printf("successful request. resp: %v", resp)

	return nil
}

func (rc *RealClient) AuditAuthenticated(event string, userID string) error {
	payload := buildPayload(event, userID)

	req, err := http.NewRequest("POST", rc.RequestURL, payload)
	username := os.Getenv("AUDITOR_USERNAME")
	password := os.Getenv("AUDITOR_PASSWORD")
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Printf("successful request. resp: %v", resp)

	return nil
}

func buildPayload(event string, userID string) *bytes.Buffer {
	payload := map[string]string{
		"event":   event,
		"user_id": userID,
	}

	json, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshalling json: %v", err)
	}

	return bytes.NewBuffer(json)
}

func LoadClient(requestURL string) *RealClient {
	return &RealClient{
		RequestURL: requestURL,
	}
}
