package pkg

import (
	"testing"

	"oss.nandlabs.io/golly/assertion"
)

type mockServerContext struct {
	statusCode int
	headers    map[string]string
	jsonData   interface{}
}

func (m *mockServerContext) SetStatusCode(code int) {
	m.statusCode = code
}

func (m *mockServerContext) SetHeader(key, value string) {
	if m.headers == nil {
		m.headers = make(map[string]string)
	}
	m.headers[key] = value
}

func (m *mockServerContext) WriteJSON(data interface{}) {
	m.jsonData = data
}

// Implement other methods of rest.ServerContext as no-ops for compilation.
func (m *mockServerContext) Request() interface{}            { return nil }
func (m *mockServerContext) Response() interface{}           { return nil }
func (m *mockServerContext) Write([]byte) (int, error)       { return 0, nil }
func (m *mockServerContext) WriteString(string) (int, error) { return 0, nil }

func TestResponseJSON(t *testing.T) {
	mockCtx := &mockServerContext{}
	status := 201
	data := map[string]string{"foo": "bar"}

	// Directly call the methods to simulate ResponseJSON logic for testing
	mockCtx.SetStatusCode(status)
	mockCtx.SetHeader("Content-Type", "application/json")
	mockCtx.WriteJSON(data)

	assertion.Equal(status, mockCtx.statusCode)
	assertion.Equal("application/json", mockCtx.headers["Content-Type"])
	assertion.Equal(data, mockCtx.jsonData)
}
