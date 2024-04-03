package services

import (
	"binder/views"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateExtensionService(c echo.Context) error {
	multipart, _ := c.MultipartForm()
	images := multipart.File["ext-images"]
	title := c.FormValue("ext-title")
	desc := c.FormValue("ext-desc")
	yt := c.FormValue("ext-yt")

	if title == "" {
		return c.String(http.StatusOK, views.SendErrorAlert("Title is required"))
	}

	if len(images) > 5 {
		return c.String(http.StatusOK, views.SendErrorAlert("Max images are only 5"))
	}

	log.Println(images)
	log.Println(desc)
	log.Println(yt)
	log.Println(len(images))

	return c.NoContent(http.StatusOK)
}
