package controllers

import (
	"fmt"
	"gowiki/errorHandlers"
	"gowiki/helpers"
	"gowiki/models"
	"net/http"
)

type PageController struct {}


func (PageController) Welcome(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Fprintf(w, "welcome to my fucking code")
}

func (PageController) ShowHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := models.LoadPage(title)

	if errorHandlers.HttpHandler(w, err) {
		return
	}

	helpers.RenderTemplate(w, "views/show", page)
}

func (PageController) EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := models.LoadPage(title)

	if errorHandlers.HttpHandler(w, err) {
		return
	}
	
	helpers.RenderTemplate(w, "views/edit", page)
}

func (PageController) SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	page := models.MakePage(title, []byte(body))
	page.Save()

	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}
