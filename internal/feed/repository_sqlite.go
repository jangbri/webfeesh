package feed

import (
	"context"
	"database/sql"

	"github.com/jangbri/webfeesh-be/internal/item"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Create(ctx context.Context, feed *Feed) error {
	if feed.Title == "" {
		feed.Title = feed.Link
	}

	stmt := `
		INSERT INTO feeds (list_id, title, link)
		VALUES (?, ?, ?)
		ON CONFLICT (link)
		DO UPDATE SET list_id = excluded.list_id
	`

	_, err := r.db.ExecContext(ctx, stmt, feed.ListID, feed.Title, feed.Link)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) Update(ctx context.Context, feed *Feed) error {
	stmt := `
		UPDATE feeds
		SET list_id = ?, title = ?, link = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(
		ctx, stmt,
		feed.ListID, feed.Title, feed.Link,
		feed.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) Delete(ctx context.Context, feed *Feed) error {
	stmt := `
		DELETE FROM feeds
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, stmt, feed.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows != 1 {
		return err
	}

	return nil
}

func (r *SQLiteRepository) GetFeedItems(ctx context.Context, feed *Feed) ([]*item.Item, error) {
	stmt := `
		SELECT
			id, feed_id, guid,
			title, link,
			date_fetched, date_updated
		FROM items
		WHERE feed_id = ?
	`

	rows, err := r.db.QueryContext(ctx, stmt, feed.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*item.Item{}

	for rows.Next() {
		var i item.Item

		err = rows.Scan(
			&i.ID,
			&i.FeedID,
			&i.GUID,
			&i.Title,
			&i.Link,
			&i.DateFetched,
			&i.DateUpdated,
		)
		if err != nil {
		}

		items = append(items, &i)
	}

	if err = rows.Err(); err != nil {
	}

	return items, nil
}
