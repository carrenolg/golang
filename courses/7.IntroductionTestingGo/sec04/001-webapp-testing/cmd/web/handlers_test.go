package main

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var tests = []struct {
		name                    string
		url                     string
		expectedStatusCode      int
		expectedURL             string
		expectedFirstStatusCode int
	}{
		{"home", "/", http.StatusOK, "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound, "/fish", http.StatusNotFound},
		{"profile", "/user/profile", http.StatusOK, "/", http.StatusTemporaryRedirect},
	}

	// app
	routes := app.routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	pathToTemplates = "./../../templates/"

	// range through test data
	for _, e := range tests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("from %s: expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
		if resp.Request.URL.Path != e.expectedURL {
			t.Errorf("from %s: expected path %s, but got %s", e.name, e.expectedURL, resp.Request.URL.Path)
		}

		resp2, _ := client.Get(ts.URL + e.url)
		if resp2.StatusCode != e.expectedFirstStatusCode {
			t.Errorf("from %s: expected first returned status %d, but got %d", e.name, e.expectedFirstStatusCode, resp2.StatusCode)
		}
	}
}

func TestAppHome(t *testing.T) {
	var test = []struct {
		name         string
		putInSession string
		expectedHtml string
	}{
		{"first visit", "", "<small>From Session:"},
		{"second visit", "hello, world!", "<small>From Session: hello, world!"},
	}

	for _, e := range test {

		// create a request
		req, _ := http.NewRequest("GET", "/", nil)

		// add context and session to request
		req = addContextAndSessionToRequest(req, app)

		// destroy
		_ = app.Session.Destroy(req.Context())

		// put in session
		if e.putInSession != "" {
			app.Session.Put(req.Context(), "test", e.putInSession)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.Home)

		handler.ServeHTTP(rr, req)
		// check status code
		if rr.Code != http.StatusOK {
			t.Errorf("TestAppHome expected status %d, but got %d", http.StatusOK, rr.Code)
		}

		body, _ := io.ReadAll(rr.Body)
		if !strings.Contains(string(body), e.expectedHtml) {
			t.Errorf("%s: did not find %s in response body", e.name, e.expectedHtml)
		}
	}
}

func TestApp_renderWithBadTemplate(t *testing.T) {
	pathToTemplates = "./testdata/"

	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	rr := httptest.NewRecorder()

	err := app.render(rr, req, "bad.page.gohtml", nil)
	if err == nil {
		t.Error("expected error from bad template, but did not get one")
	}

	pathToTemplates = "./../../templates/"
}

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))

	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}

func Test_app_Login(t *testing.T) {
	var test = []struct {
		name               string
		poastedData        url.Values
		expectedStatusCode int
		expectedLocation   string
	}{
		{
			name: "valid credentials",
			poastedData: url.Values{
				"email":    {"admin@example.com"},
				"password": {"secret"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/user/profile",
		},
		{
			name: "missing data",
			poastedData: url.Values{
				"email":    {""},
				"password": {""},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/",
		},
		{
			name: "user not found",
			poastedData: url.Values{
				"email":    {"you@there.com"},
				"password": {"password"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/",
		},
		{
			name: "no auth",
			poastedData: url.Values{
				"email":    {"admin@example.com"},
				"password": {"password"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLocation:   "/",
		},
	}

	// loop
	for _, e := range test {
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(e.poastedData.Encode()))
		req = addContextAndSessionToRequest(req, app)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Login)
		handler.ServeHTTP(rr, req)

		// check status code
		if rr.Code != e.expectedStatusCode {
			t.Errorf("Test_app_Login expected status %d, but got %d", e.expectedStatusCode, rr.Code)
		}

		// check location
		actualLocation, err := rr.Result().Location()
		if err == nil {
			if actualLocation.String() != e.expectedLocation {
				t.Errorf("Test_app_Login expected location %s, but got %s", e.expectedLocation, actualLocation.String())
			}
		} else {
			t.Errorf("Test_app_Login no location header set %s,", e.name)
		}

	}
}
