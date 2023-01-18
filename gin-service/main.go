package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums []Album = []Album{
	{ID: "1", Title: "first one", Artist: "mamad", Price: 12000},
	{ID: "2", Title: "second one", Artist: "ali", Price: 32000},
	{ID: "3", Title: "third one", Artist: "sina", Price: 14000},
}

func handleList(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusAccepted, albums)
}

func handleShow(ctx *gin.Context) {
	for _, album := range albums {
		if album.ID == ctx.Param("id") {
			ctx.IndentedJSON(http.StatusAccepted, album)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "the album not found",
	})
}

func handlePost(ctx *gin.Context) {
	var newAlbum Album

	if err := ctx.BindJSON(newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	r := gin.Default()

	r.GET("/album", handleList)
	r.GET("/album/:id", handleShow)
	r.POST("/album", handlePost)

	r.Run()
}
