package models

import "os"

type Page struct {
	Title string
	Body  []byte
}

func MakePage(title string, body []byte) *Page {
	return &Page{Title: title, Body: body}
}

func LoadPage(title string) (*Page, error) {
	body, err := os.ReadFile(title + ".txt")

	return MakePage(title, body), err
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	err := os.WriteFile(filename, p.Body, 0600)

	return err
}
