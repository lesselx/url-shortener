package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	generateLink = "/api/short_urls"
	redirectLink = "/api/short_urls/:alias"
)

func (app *application) routes() *gin.Engine {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	r.POST(generateLink, app.GenerateLink)

	r.GET(redirectLink, app.RedirectLink)

	r.HEAD(redirectLink, app.RedirectLink)

	return r

}
