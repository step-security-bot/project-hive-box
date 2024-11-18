package handlers

import (
	"fmt"
	"net/http"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "v0.0.1")
}
