package repos

import (
	"binder/configs"
	"binder/dtos"
	"binder/entities"
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

func GetAllExts(userID string) ([]entities.Extension, error) {
	var exts []entities.Extension

	SQL := "SELECT id, slug, title, description, created_at, updated_at FROM extensions WHERE author_id = $1"

	rows, err := configs.DB_POOL.Query(context.Background(), SQL, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ext entities.Extension

		if err = rows.Scan(
			&ext.ID,
			&ext.Slug,
			&ext.Title,
			&ext.Description,
			&ext.CreatedAt,
			&ext.UpdatedAt,
		); err != nil {
			log.Println(err)
			return nil, err
		}

		exts = append(exts, ext)
	}

	return exts, nil
}

func GetExt(userID string, slug string) (*entities.Extension, error) {
	var ext entities.Extension

	SQL := `
		SELECT e.id, e.slug, e.title, e.description, e.code, e.youtube_url, e.created_at, e.updated_at 
		FROM extensions e
		LEFT JOIN image_attachments a
		ON a.extension_id = e.id
		WHERE e.author_id = $1 AND e.slug = $2
	`

	row := configs.DB_POOL.QueryRow(context.Background(), SQL, userID, slug)
	if err := row.Scan(
		&ext.ID,
		&ext.Slug,
		&ext.Title,
		&ext.Description,
		&ext.Code,
		&ext.YoutubeURL,
		&ext.CreatedAt,
		&ext.UpdatedAt,
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return &ext, nil
}
