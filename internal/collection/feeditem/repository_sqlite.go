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

func (r *SQLiteRepository) GetLatest(
	ctx context.Context,
	cID int64,
) ([]*CollectionFeedItem, error) {
	stmt := `
		SELECT *
		FROM feed_items
		WHERE collection_id = ?
		ORDER BY created_time DESC
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
			&item.Description,
		); err != nil {
			slog.Error("unable to read specific feed item")
		}

		items = append(items, &item)
	}

	return items, nil
}
