package doohan

import (
	"time"
	"github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type Entry struct {
	ID          int
	Start       time.Time
	Stop        pq.NullTime
	Running     bool
	Description string
}

var db *sqlx.DB

func DB() *sqlx.DB {
	if db != nil {
		return db
	}
	var err error
	db, err = sqlx.Open("postgres", "user=doohan password=doohan dbname=doohan sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func FetchEntries() []Entry {
	db := DB()
	entries := []Entry{}
	err := db.Select(&entries, `SELECT id, start, stop, running, description FROM entries_convenient LIMIT 10`)
	if err != nil {
		panic(err)
	}
	return entries
}

func AddEntry(description string) Entry {
	db := DB()
	var entry Entry
	err := db.QueryRowx(`INSERT INTO entries (description) VALUES ($1) RETURNING id, start, description`, description).StructScan(&entry)
	if err != nil {
		panic(err)
	}
	entry.Running = true
	return entry
}
