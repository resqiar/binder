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
	code := c.FormValue("ext-code")
	yt := c.FormValue("ext-yt")

	if title == "" {
		return c.String(http.StatusOK, views.SendErrorAlert("Title is required"))
	}

	if len(images) > 5 {
		return c.String(http.StatusOK, views.SendErrorAlert("Max images are only 5"))
	}

	if len(code) > 10000 {
		return c.String(http.StatusOK, views.SendErrorAlert("The length of code is exceeding 10K characters"))
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
		Code:        code,
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

func GetAllExtensionService(c echo.Context) error {
	userID := c.Get("userID").(string)

	exts, err := repos.GetAllExts(userID)
	if err != nil {
		return c.String(http.StatusOK, views.SendErrorAlert("Failed to get your extensions, please try again later"))
	}

	return c.JSON(http.StatusOK, exts)
}

func GetExtensionService(c echo.Context) error {
	userID := c.Get("userID").(string)
	slug := c.Param("slug")

	exts, err := repos.GetExt(userID, slug)
	if err != nil {
		return c.String(http.StatusOK, views.SendErrorAlert("Failed to get your extension, please try again later"))
	}

	return c.JSON(http.StatusOK, exts)
}

func SearchExtensionService(c echo.Context) error {
	userID := c.Get("userID").(string)
	keyword := c.QueryParam("search-keyword")

	// if the keyword is empty, the user assumed to clear
	// the result. Hence, query all extensions.
	if keyword == "" {
		exts, err := repos.GetAllExts(userID)
		if err != nil {
			return c.String(http.StatusOK, views.SendErrorAlert("Failed to get your extensions, please try again later"))
		}

		return c.String(http.StatusOK, views.SendSearchCard(exts))
	}

	exts, err := repos.SearchExt(userID, keyword)
	if err != nil || len(exts) == 0 {
		return c.String(http.StatusOK, views.SendSearchNotFound(fmt.Sprintf("Keyword for \"%s\" not found, try another...", keyword)))
	}

	return c.String(http.StatusOK, views.SendSearchCard(exts))
}

func DeleteExtensionService(c echo.Context) error {
	userID := c.Get("userID").(string)
	extSlug := c.Param("slug")

	if extSlug == "" {
		return c.String(http.StatusNotFound, views.SendErrorAlert("You are requesting to unrecognized URL"))
	}

	// Delete extension with its attachments
	err := repos.DeleteExt(userID, extSlug)
	if err != nil {
		return c.String(http.StatusOK, views.SendErrorAlert("Failed to delete extension, please try again later"))
	}

	// Redirect back to dashboard while in success
	c.Response().Header().Add("HX-Redirect", "/dashboard")
	return c.String(http.StatusOK, "OK")
}

func DeleteExtensionImageService(c echo.Context) error {
	userID := c.Get("userID").(string)
	extSlug := c.Param("slug")
	imageId := c.Param("imageId")

	if extSlug == "" || imageId == "" {
		return c.String(http.StatusNotFound, views.SendErrorAlert("You are requesting to unrecognized URL"))
	}

	// Delete image from an extension
	err := repos.DeleteExtImage(userID, extSlug, imageId)
	if err != nil {
		return c.String(http.StatusOK, views.SendErrorAlert("Failed to delete image, please try again later"))
	}

	// Redirect back (refresh) while in success attempt
	c.Response().Header().Add("HX-Redirect", fmt.Sprintf("/edit/%s", extSlug))
	return c.String(http.StatusOK, "OK")
}
