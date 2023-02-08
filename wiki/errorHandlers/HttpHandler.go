package errorHandlers

import "net/http"

func HttpHandler(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)

	return true
}