package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},

	{"post-search-avai", "/search-availability", "POST", []postData{
		{key: "start", value: "2022-06-23"},
		{key: "end", value: "2022-06-25"},
	}, http.StatusOK},
	{"post-search-avai-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2022-06-23"},
		{key: "end", value: "2022-06-25"},
	}, http.StatusOK},
	{"make-reservation-post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Zi-Yang"},
		{key: "last_name", value: "Huang"},
		{key: "email", value: "a@a.com"},
		{key: "phone", value: "0912345678"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	// Create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url) // call test server URL before calling my own e.url ("/")
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		}
	}
}
