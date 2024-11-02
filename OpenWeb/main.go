package main

import (
	"TIL_GoWeb/OpenWeb/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())

}
