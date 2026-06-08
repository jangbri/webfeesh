package feeditem

import (
	"context"
	"database/sql"
	"log/slog"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Create(ctx context.Context, c *CollectionFeedItem) error {
	stmt := `
		INSERT INTO feed_items(collection_id, title, description, time_created)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(
		ctx,
		stmt,
		c.CollectionID,
		c.Title,
		c.Description.Text,
		c.TimeCreated,
	)
	if err != nil {
		slog.Error("failed to create new collection feed item", "error", err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) GetLatest(
	ctx context.Context,
	cID int64,
) ([]*CollectionFeedItem, error) {
	stmt := `
		SELECT id, collection_id, title, description, time_created
		FROM feed_items
		WHERE collection_id = ?
		ORDER BY time_created DESC
		LIMIT 5;
	`

	rows, err := r.db.QueryContext(ctx, stmt, cID)
	if err != nil {
		slog.Error("unable to retrieve collection's feed items to create webfeed")
		return nil, err
	}
	defer rows.Close()

	items := []*CollectionFeedItem{}

	for rows.Next() {
		var item CollectionFeedItem

		if err := rows.Scan(
			&item.ID,
			&item.CollectionID,
			&item.Title,
			&item.Description.Text,
			&item.TimeCreated,
		); err != nil {
			slog.Error("unable to read specific feed item")
		}

		items = append(items, &item)
	}

	return items, nil
}
