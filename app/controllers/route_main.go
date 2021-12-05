package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateTHML(w, "hello", "layout", "top")
}
