package main

import (
	"gowiki/controllers"
	"gowiki/router"
	"log"
	"net/http"
)

func main() {
	/* p1 := MakePage("first", []byte("this is the first page"))
	p1.Save()

	loadedPage, err := loadPage("first")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(string(loadedPage.Body)) */

	pageController := &controllers.PageController{}

	http.HandleFunc("/", router.MakeHttpHandler(pageController.Welcome))

	http.HandleFunc("/view/", router.MakeHttpHandler(pageController.ShowHandler))

	http.HandleFunc("/edit/", router.MakeHttpHandler(pageController.EditHandler))

	http.HandleFunc("/save/", router.MakeHttpHandler(pageController.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
