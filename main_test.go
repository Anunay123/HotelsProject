package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCase1(t *testing.T) {

	r := getRouter()

	r.GET("/hotels/5784721289468963595", getData())

	req, _ := http.NewRequest("GET", "/hotels/5784721289468963595", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		_, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil

		return statusOK && pageOK
	})
}

func TestCase2(t *testing.T) {

	r := getRouter()

	r.GET("/hotels/5604282523447316015?rating=2", getData())

	req, _ := http.NewRequest("GET", "/hotels/5604282523447316015?rating=2", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		_, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil

		return statusOK && pageOK
	})
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}


