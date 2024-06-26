package services

import (
	"binder/dtos"
	"binder/entities"
	"binder/repos"
	"binder/utils"
)

func CreateUser(payload *dtos.GooglePayload) (*entities.User, error) {
	// format the given name from the provider
	formattedName := utils.FormatUsername(payload.GivenName)

	newUser := &dtos.CreateUserInput{
		Username:   formattedName,
		Email:      payload.Email,
		Provider:   "google",
		PictureURL: payload.Picture,
	}

	result, err := repos.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}
