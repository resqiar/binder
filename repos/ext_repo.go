package repos

import (
	"binder/configs"
	"binder/dtos"
	"binder/entities"
	"binder/utils"
	"context"
	"database/sql"
	"log"
	"strings"

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
	var images sql.NullString

	SQL := `
		SELECT e.id, e.slug, e.title, e.description, e.code, e.youtube_url, e.created_at, e.updated_at, STRING_AGG(a.url, ',')
		FROM extensions e
		LEFT JOIN image_attachments a
		ON a.extension_id = e.id
		WHERE e.author_id = $1 AND e.slug = $2
		GROUP BY e.id, e.slug, e.title, e.description, e.code, e.youtube_url, e.created_at, e.updated_at
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
		&images,
	); err != nil {
		log.Println(err)
		return nil, err
	}

	// update picture urls from aggregated string
	if images.String != "" {
		splitted_images := strings.Split(images.String, ",")
		ext.ImageUrls = splitted_images
	}

	return &ext, nil
}

func SearchExt(userID string, keyword string) ([]entities.Extension, error) {
	var exts []entities.Extension

	// format keyword from "search keyword" to "search & keyword"
	formatted_keywords := utils.FormatSearch(keyword)

	SQL := `
		SELECT id, slug, title, description, created_at, updated_at 
		FROM extensions
		WHERE search_vector @@ to_tsquery('english', $2) AND author_id = $1
	`

	rows, err := configs.DB_POOL.Query(context.Background(), SQL, userID, formatted_keywords)
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

func DeleteExt(userID string, extSlug string) error {
	GET_EXT_SQL := `
		SELECT e.id, STRING_AGG(a.id, ',') AS ImageIds
		FROM extensions e
		LEFT JOIN image_attachments a
		ON a.extension_id = e.id
		WHERE e.author_id = $1 AND e.slug = $2
		GROUP BY e.id
	`
	DELETE_ATTACHMENT_SQL := "DELETE FROM image_attachments WHERE id = $1 AND extension_id = $2"
	DELETE_EXT_SQL := "DELETE FROM extensions WHERE slug = $1 AND author_id = $2"

	tx, err := configs.DB_POOL.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	type TemporaryExt struct {
		ID       string
		ImageIds []string
	}

	var temp TemporaryExt
	var imageIds sql.NullString

	// Step 1: get extension detail
	row := tx.QueryRow(
		context.Background(),
		GET_EXT_SQL,
		userID,
		extSlug,
	)
	if err := row.Scan(
		&temp.ID,
		&imageIds,
	); err != nil {
		log.Println("Failed to get extension detail:", err)
		return err
	}

	// bind picture ids from aggregated string
	if imageIds.String != "" {
		splitted_images := strings.Split(imageIds.String, ",")
		temp.ImageIds = splitted_images
	}

	for _, imageId := range temp.ImageIds {
		// Step 2: delete image from imagekit
		err = utils.DeleteImage(imageId)
		if err != nil {
			log.Println("[Skipping...] Failed deleting imagekit file:", err)
		}

		// Step 3: delete attachments using the ids
		_, err = tx.Exec(
			context.Background(),
			DELETE_ATTACHMENT_SQL,
			imageId,
			temp.ID,
		)
		if err != nil {
			log.Println("Failed deleting attachments:", err)
			return err
		}
	}

	// Step 4: delete extension
	_, err = tx.Exec(
		context.Background(),
		DELETE_EXT_SQL,
		extSlug,
		userID,
	)
	if err != nil {
		log.Println("Failed deleting extension:", err)
		return err
	}

	// Step 5: commit tx
	if err := tx.Commit(context.Background()); err != nil {
		log.Println("Failed to commit delete ext transaction:", err)
		return err
	}

	return nil
}
