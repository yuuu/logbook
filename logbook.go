package main

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const DB_FILE_NAME = "logbook.db"

type Logbook struct {
	db *sql.DB
}

func LoadLogbook(path string) (*Logbook, error) {
	var db *sql.DB
	var err error

	err = os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}

	db, err = sql.Open("sqlite3", filepath.Join(path, DB_FILE_NAME))
	if err != nil {
		return nil, err
	}
	return &Logbook{db}, nil
}

func (self *Logbook) Keep(timeStr string, workpath string) error {
	var err error

	var text *Text
	text, err = InputText(workpath)
	if err != nil {
		return err
	}

	var date time.Time
	date, err = time.Parse(DATE_FORMAT, timeStr)
	if err != nil {
		return err
	}

	_, err = CreateEntry(self.db, text, &date)
	if err != nil {
		return err
	}

	return nil
}

func (self *Logbook) Entry(timeStr string) error {
	var err error

	if timeStr == "" {
		var entries []*Entry
		entries, err = SearchMostRecentEntry(self.db, 100, 0)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			entry.Print()
		}

	} else {
		var date time.Time
		date, err = time.Parse(DATE_FORMAT, timeStr)
		if err != nil {
			return err
		}

		var entry *Entry
		entry, err = SearchEntryWithDate(self.db, &date)
		if err != nil {
			return err
		}

		entry.Print()
	}

	return nil
}

func (self *Logbook) Close() {
	defer self.db.Close()
}
