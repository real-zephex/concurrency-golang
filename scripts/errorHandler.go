package scripts

import (
	"fmt"
	"net/http"
)

func handleError(err error, msg string, w http.ResponseWriter) {
	fmt.Printf("An error has occured here: %s \nError: %v\n", msg, err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%s", msg)
}
