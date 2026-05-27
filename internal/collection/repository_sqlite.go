package collection

import (
	"context"
	"database/sql"

	"github.com/jangbri/webfeesh-be/internal/feed"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) GetAll(ctx context.Context) ([]*Collection, error) {
	stmt := `
		SELECT id, name
		FROM collections
	`

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := []*Collection{}

	for rows.Next() {
		var collection Collection

		err = rows.Scan(&collection.ID, &collection.Name)
		if err != nil {
		}

		collections = append(collections, &collection)
	}

	if err = rows.Err(); err != nil {
	}

	return collections, nil
}

func (r *SQLiteRepository) Create(ctx context.Context, collection *Collection) error {
	stmt := `
		INSERT INTO collections(name)
		VALUES (?)
		ON CONFLICT (name)
		DO UPDATE SET name = excluded.name
	`

	_, err := r.db.ExecContext(ctx, stmt, collection.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) Update(ctx context.Context, collection *Collection) error {
	stmt := `
		UPDATE collections
		SET name = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, stmt, collection.Name, collection.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) Delete(ctx context.Context, collection *Collection) error {
	stmt := `
		DELETE FROM collections
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, stmt, collection.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) GetCollectionFeeds(
	ctx context.Context,
	collection *Collection,
) ([]*feed.Feed, error) {
	stmt := `
		SELECT id, title, link
		FROM feeds
		WHERE collection_id = ?
	`

	rows, err := r.db.QueryContext(ctx, stmt, collection.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := []*feed.Feed{}

	for rows.Next() {
		var feed feed.Feed
		feed.CollectionID = collection.ID

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Link)
		if err != nil {
		}

		feeds = append(feeds, &feed)
	}

	if err = rows.Err(); err != nil {
	}

	return feeds, nil
}
