package router

import (
	"gowiki/errorHandlers"
	"gowiki/validations"
	"net/http"
)

func MakeHttpHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := validations.GetTitle(w, r)

		if errorHandlers.HttpHandler(w, err) {
			return
		}

		fn(w, r, title)
	}
}
