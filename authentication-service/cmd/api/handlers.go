
package main

import (
	"net/http"
)

func (app *Config) TEST(w http.ResponseWriter,r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "Hit the auth",
	}
	_= app.writeJSON(w,http.StatusOK,payload)
	
}