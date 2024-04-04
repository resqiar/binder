package services

import (
	"binder/dtos"
	"binder/repos"
	"binder/utils"
	"binder/views"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateExtensionService(c echo.Context) error {
	multipart, _ := c.MultipartForm()
	userID := c.Get("userID").(string)
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

	// Upload images to ImageKit
	uploadedImages, err := utils.UploadImages(images)
	if err != nil {
		return c.String(http.StatusOK, views.SendErrorAlert(err.Error()))
	}

	extInput := dtos.CreateExtInput{
		Slug:        utils.GenerateRandomString(8),
		Title:       title,
		Description: desc,
		Youtube_url: yt,
		Author_id:   userID,
	}

	// Create extension with its attachments
	extSlug, err := repos.CreateExt(&extInput, uploadedImages)
	if err != nil {
		return c.String(http.StatusOK, views.SendErrorAlert("Failed to create new extension, please try again later"))
	}

	return c.String(http.StatusOK, views.SendCreateExtSuccessAlert(fmt.Sprintf("/ext/%s", *extSlug)))
}
