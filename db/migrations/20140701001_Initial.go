package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20140701001(txn *sql.Tx) {
	txn.Exec(`
		CREATE TABLE projects (
			id   serial       PRIMARY KEY,
			name varchar(100)
		)
	`)
	txn.Exec(`
		CREATE TABLE entries (
			id          serial      PRIMARY KEY,
			start       timestamptz DEFAULT now() NOT NULL UNIQUE,
			project     int,
			description text
		)
	`)
	txn.Exec(`
		CREATE VIEW entries_with_stop AS
			SELECT
				id,
				start,
				lead(start) OVER (ORDER BY start) AS stop,
				project,
				description
			FROM
				entries
	`)
	txn.Exec(`
		CREATE VIEW entries_convenient AS
		SELECT
			e.id,
			e.start,
			e.stop,
			e.stop - e.start AS duration,
			CASE
				WHEN e.stop IS NULL THEN true
				ELSE false
			END AS running,
			e.project AS project_id,
			p.name AS project,
			e.description
		FROM
			entries_with_stop e
		LEFT JOIN
			projects p
			ON p.id = e.project
		ORDER BY
			e.start DESC
	`)
}

// Down is executed when this migration is rolled back
func Down_20140701001(txn *sql.Tx) {
	txn.Exec(`DROP VIEW entries_convenient`)
	txn.Exec(`DROP VIEW entries_with_stop`)
	txn.Exec(`DROP TABLE entries`)
	txn.Exec(`DROP TABLE projects`)
}
