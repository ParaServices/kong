package kong

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// setup sets up a test HTTP server along with a findface.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (Client, *http.ServeMux, *httptest.Server) {
	// test server
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	u, _ := url.Parse(server.URL)
	c := NewClient(nil, u)

	return c, mux, server
}

// teardown closes the test HTTP server.
func teardown(s *httptest.Server) {
	s.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func setupHandleFunc(t *testing.T, mux *http.ServeMux, path, verb string, statusCode int, respBody []byte) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, verb)
		w.WriteHeader(statusCode)
		_, writeErr := w.Write(respBody)
		if writeErr != nil {
			t.Error(writeErr)
		}
	})
}
