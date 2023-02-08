package helpers

import (
	"gowiki/errorHandlers"
	"gowiki/models"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, view string, p *models.Page) {
	t, err := template.ParseFiles(view+".html")
	
	if errorHandlers.HttpHandler(w, err) {
		return
	}

	t.Execute(w, p)
}