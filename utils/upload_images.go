package utils

import (
	"binder/dtos"
	"context"
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

func UploadImages(images []*multipart.FileHeader) ([]dtos.CreateImagesInput, error) {
	ik, err := imagekit.New()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var imgs []dtos.CreateImagesInput

	for _, v := range images {
		file, err := v.Open()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer file.Close()

		// validate image before proceeding further
		err = validateImage(v)
		if err != nil {
			return nil, err
		}

		// upload image binary to ImageKit
		resp, err := ik.Uploader.Upload(context.Background(), file, uploader.UploadParam{
			FileName: v.Filename,
		})
		if err != nil {
			log.Println(err)
			return nil, err
		}

		img := dtos.CreateImagesInput{
			ID:  resp.Data.FileId,
			URL: resp.Data.Url,
		}

		imgs = append(imgs, img)
	}

	return imgs, nil
}

func validateImage(fileHeader *multipart.FileHeader) error {
	fileType := fileHeader.Header.Get("Content-Type")
	fileSize := fileHeader.Size
	const maxFileSize = 1 * 1024 * 1024 // 1 MB

	if !strings.HasPrefix(fileType, "image/") {
		return errors.New("Invalid file type")
	}

	if fileSize > int64(maxFileSize) {
		return errors.New("Image size is too big")
	}

	return nil
}
