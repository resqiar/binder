package repos

import (
	"binder/configs"
	"binder/dtos"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateExt(extInput *dtos.CreateExtInput, imgInput []dtos.CreateImagesInput) (*string, error) {
	var extID string
	var slug string

	SQL := "INSERT INTO extensions(slug, title, description, code, youtube_url, author_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, slug"

	tx, err := configs.DB_POOL.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	// Step 1: create extension
	err = tx.QueryRow(
		context.Background(),
		SQL,
		extInput.Slug,
		extInput.Title,
		extInput.Description,
		extInput.Code,
		extInput.Youtube_url,
		extInput.Author_id,
	).Scan(&extID, &slug)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Step 2: create attachments and use ext id as FK
	if err := createAttachments(tx, extID, imgInput); err != nil {
		return nil, err
	}

	// Step 3: commit tx
	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return &slug, nil
}

func createAttachments(tx pgx.Tx, extID string, urls []dtos.CreateImagesInput) error {
	var ids []string

	SQL := "INSERT INTO image_attachments(id, url, extension_id) VALUES ($1, $2, $3) RETURNING id"

	for _, v := range urls {
		var id string

		err := tx.QueryRow(
			context.Background(),
			SQL,
			v.ID,
			v.URL,
			extID,
		).Scan(&id)
		if err != nil {
			log.Println(err)
			return err
		}

		ids = append(ids, id)
	}

	return nil
}
