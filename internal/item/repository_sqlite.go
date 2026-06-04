package item

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

func (r *SQLiteRepository) UpsertMany(ctx context.Context, items []*Item) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("failed to begin transaction for upserting items")
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO items (
			feed_id,
			guid, link, title,
			date_fetched, date_updated
		)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT (feed_id, guid) DO NOTHING
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		slog.Error("failed to prepare context to ready for transaction")
		return err
	}
	defer stmt.Close()

	for _, item := range items {
		_, err := stmt.ExecContext(
			ctx,
			item.FeedID,
			item.GUID, item.Link, item.Title,
			item.DateFetched, item.DateUpdated,
		)
		if err != nil {
			slog.Error("failed to insert into items", "item", item.Link)
			return err
		}
	}

	return tx.Commit()
}
