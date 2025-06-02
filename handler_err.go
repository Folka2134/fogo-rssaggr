package main

import "net/http"

func handleErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 401, "Something went wrong")
}
