package app

import (
	"database/sql"
	"fmt"
	"log/slog"
)

var migrations = []func(*sql.Tx) error{
	m01_initial,
}

var latestVersion = int64(len(migrations))

func migrateToLatest(db *sql.DB) error {
	var dbVersion int64
	if err := db.QueryRow(`PRAGMA user_version`).Scan(&dbVersion); err != nil {
		return err
	}

	if dbVersion >= latestVersion {
		return nil
	}

	slog.Info(
		fmt.Sprintf("migrating from db version %d to db version %d", dbVersion, latestVersion),
	)

	for v := dbVersion; v < latestVersion; v++ {
		slog.Info("migration - start", "version", v)

		err := migrateToVersion(v, db)
		if err != nil {
			return err
		}

		slog.Info("migration - end", "version", v)
	}

	return nil
}

func migrateToVersion(v int64, db *sql.DB) error {
	var err error
	var tx *sql.Tx
	migrateFunction := migrations[v]

	if tx, err = db.Begin(); err != nil {
		slog.Error("migration transaction start failure", "version", v)
		return err
	}

	if err = migrateFunction(tx); err != nil {
		slog.Error("migration failure during transaction", "version", v)
		slog.Error(err.Error())
		tx.Rollback()
		return err
	}

	stmt := fmt.Sprintf("PRAGMA user_version = %d", v+1)
	if _, err = tx.Exec(stmt); err != nil {
		slog.Error("migration failure during version increment", "version", v)
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		slog.Error("migration failure during db commit", "version", v)
		return err
	}

	return nil
}

func m01_initial(tx *sql.Tx) error {
	sql := `
		PRAGMA foreign_keys = ON;

		CREATE TABLE IF NOT EXISTS collections (
			id   INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS feeds (
			id            INTEGER PRIMARY KEY AUTOINCREMENT,
			collection_id INTEGER NOT NULL,
			title         TEXT NOT NULL,
			link          TEXT NOT NULL,

			FOREIGN KEY (collection_id)
				REFERENCES collections(id)
				ON DELETE CASCADE
		);

		CREATE INDEX IF NOT EXISTS idx_feed_collection_id ON feeds(collection_id);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_feed_link ON feeds(link);

		CREATE TABLE IF NOT EXISTS items (
			id           INTEGER PRIMARY KEY AUTOINCREMENT,
			feed_id      INTEGER NOT NULL,
			guid         TEXT NOT NULL,
			link         TEXT,
			title        TEXT,
			date_fetched DATETIME NOT NULL,
			date_updated DATETIME NOT NULL,

			FOREIGN KEY (feed_id)
				REFERENCES feeds(id)
				ON DELETE CASCADE
		);

		CREATE INDEX IF NOT EXISTS idx_item_feed_id ON items(feed_id);
		CREATE INDEX IF NOT EXISTS idx_item_date_fetched ON items(date_fetched);
		CREATE UNIQUE INDEX IF NOT EXISTS idx_item_guid ON items(feed_id, guid);

		CREATE TRIGGER cleanup_old_items
		AFTER INSERT ON items
		BEGIN
			DELETE FROM items
			WHERE date_fetched < DATETIME('now', '-3 months');
		END;
	`

	// following tx.Exec gets committed in func migrateToVersion
	_, err := tx.Exec(sql)
	return err
}
