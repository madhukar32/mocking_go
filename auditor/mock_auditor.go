package auditor

import "log"

type MockClient struct {
	RequestURL string
}

func (rc *MockClient) Audit(event string, userID string) error {
	log.Printf("Successfully mocked the Audit function")
	return nil
}

func (rc *MockClient) AuditAuthenticated(event string, userID string) error {
	log.Printf("Successfully mocked the AuditAuthenticated function")
	return nil
}

func LoadMock() *MockClient {
	return &MockClient{}
}
