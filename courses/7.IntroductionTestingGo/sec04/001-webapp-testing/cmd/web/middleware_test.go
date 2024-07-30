package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"webapp/pkg/data"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "ip:port", false},
		{"", "", "", false},
	}

	// create a dummy handler that we'll use to check de context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Errorf("(%s) not present", contextUserKey)
		}

		// make sure we got a string back
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})

	for _, e := range tests {
		// create the handler to test
		handleToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://127.0.0.1:8080/home", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handleToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	// get a context
	var ctx = context.Background()

	// put something in the context
	ctx = context.WithValue(ctx, contextUserKey, "whatever")

	// call the function
	ip := app.ipFromContext(ctx)

	// perform the test
	if !strings.EqualFold("whatever", ip) {
		t.Error("wrong value returned from context")
	}
}

func Test_application_auth(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	var tests = []struct {
		name   string
		isAuth bool
	}{
		{"logged in", true},
		{"not logged in", false},
	}

	for _, e := range tests {
		handler := app.auth(next)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)

		if e.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if e.isAuth && rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		}

		if !e.isAuth && rr.Code != http.StatusTemporaryRedirect {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusTemporaryRedirect)
		}

	}
}
