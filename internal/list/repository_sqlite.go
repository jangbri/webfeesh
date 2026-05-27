package list

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

func (r *SQLiteRepository) GetAll(ctx context.Context) ([]*List, error) {
	stmt := `
		SELECT id, name
		FROM lists
	`

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := []*List{}

	for rows.Next() {
		var list List

		err = rows.Scan(&list.ID, &list.Name)
		if err != nil {
		}

		lists = append(lists, &list)
	}

	if err = rows.Err(); err != nil {
	}

	return lists, nil
}

func (r *SQLiteRepository) Create(ctx context.Context, list *List) error {
	stmt := `
		INSERT INTO lists (name)
		VALUES (?)
		ON CONFLICT (name)
		DO UPDATE SET name = excluded.name
	`

	_, err := r.db.ExecContext(ctx, stmt, list.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) Update(ctx context.Context, list *List) error {
	stmt := `
		UPDATE lists
		SET name = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, stmt, list.Name, list.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) Delete(ctx context.Context, list *List) error {
	stmt := `
		DELETE FROM lists
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, stmt, list.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) GetListFeeds(ctx context.Context, list *List) ([]*feed.Feed, error) {
	stmt := `
		SELECT id, title, link
		FROM feeds
		WHERE list_id = ?
	`

	rows, err := r.db.QueryContext(ctx, stmt, list.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := []*feed.Feed{}

	for rows.Next() {
		var feed feed.Feed
		feed.ListID = list.ID

		err = rows.Scan(&feed.ID, &feed.Title, &feed.Link)
		if err != nil {
		}

		feeds = append(feeds, &feed)
	}

	if err = rows.Err(); err != nil {
	}

	return feeds, nil
}
