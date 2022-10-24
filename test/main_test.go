package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sosmed/controllers"
	"sosmed/routes"
	"testing"
)

var baseUrl = "http://localhost:3000/api/v1"

func TestEdnpoints(t *testing.T) {
	r := routes.SetupRouter()
	r.GET(baseUrl, controllers.GetAllShortVideo)
	req, _ := http.NewRequest("GET", baseUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	ioutil.ReadAll(w.Body)
	fmt.Println("responseData")
}
