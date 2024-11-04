package main

import (
	"TIL_GoWeb/RESTfulAPI/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
