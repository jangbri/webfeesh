package collection

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/jangbri/webfeesh/internal/feed"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) GetAll(ctx context.Context) ([]Collection, error) {
	stmt := `
		SELECT id, name
		FROM collections
	`

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		slog.Error("failed to query the collections table")
		return nil, err
	}
	defer rows.Close()

	collections := []Collection{}

	for rows.Next() {
		var c Collection

		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			slog.Error("failed to retrieve collection", "collection", c.Name)
		}

		collections = append(collections, c)
	}

	if err = rows.Err(); err != nil {
		slog.Error("errors when retrieving rows from collection table")
	}

	return collections, nil
}

func (r *SQLiteRepository) Create(ctx context.Context, collection Collection) error {
	stmt := `
		INSERT INTO collections(name)
		VALUES (?)
		ON CONFLICT (name)
		DO UPDATE SET name = excluded.name
	`

	_, err := r.db.ExecContext(ctx, stmt, collection.Name)
	if err != nil {
		slog.Error("failed to create new collection", "collection", collection.Name)
		return err
	}

	return nil
}

func (r *SQLiteRepository) Update(ctx context.Context, collection Collection) error {
	stmt := `
		UPDATE collections
		SET name = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, stmt, collection.Name, collection.ID)
	if err != nil {
		slog.Error("failed to update collection", "collection", collection.Name)
		return err
	}

	return nil
}

func (r *SQLiteRepository) Delete(ctx context.Context, collection Collection) error {
	stmt := `
		DELETE FROM collections
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, stmt, collection.ID)
	if err != nil {
		slog.Error("failed to delete collection", "collection", collection.Name)
		return err
	}

	return nil
}

func (r *SQLiteRepository) GetCollectionFeeds(
	ctx context.Context,
	collection Collection,
) ([]feed.Feed, error) {
	stmt := `
		SELECT id, title, link
		FROM feeds
		WHERE collection_id = ?
		ORDER BY title ASC
	`

	rows, err := r.db.QueryContext(ctx, stmt, collection.ID)
	if err != nil {
		slog.Error("failed to retrieve collection's feeds", "collection", collection.Name)
		return nil, err
	}
	defer rows.Close()

	feeds := []feed.Feed{}

	for rows.Next() {
		var feed feed.Feed
		feed.CollectionID = collection.ID

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Link)
		if err != nil {
			slog.Error("failed to retrieve collection's feed", "feed", feed.ID)
		}

		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		slog.Error("errors when retrieving collection's feeds from table")
	}

	return feeds, nil
}
