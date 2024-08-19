package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/foo", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	// create a test server
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	pathToTemplates = "./../../templates"

	// range through test data
	for _, entry := range theTests {
		res, err := testServer.Client().Get(testServer.URL + entry.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if res.StatusCode != entry.expectedStatusCode {
			t.Errorf("for %s: expected status %d, but got %d", entry.name, entry.expectedStatusCode, res.StatusCode)
		}
	}
}
