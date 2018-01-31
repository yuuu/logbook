package main

import (
	"database/sql"
	"io"
	"os"
	"os/exec"
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

	var date *time.Time
	date, err = timeParse(timeStr)
	if err != nil {
		return err
	}
	if date == nil {
		nowDate := time.Now()
		date = &nowDate
	}
	_, err = CreateEntry(self.db, text, date)
	if err != nil {
		return err
	}

	return nil
}

func (self *Logbook) Entry(timeStr string) error {
	var err error
	var stdin io.WriteCloser

	cmd := exec.Command("less", "-R")
	cmd.Stdout = os.Stdout
	stdin, err = cmd.StdinPipe()
	if err != nil {
		return err
	}

	var date *time.Time
	date, err = timeParse(timeStr)
	if err != nil {
		return err
	}

	if date == nil {
		var entries []*Entry
		entries, err = SearchMostRecentEntry(self.db, 100, 0)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			entry.Print(stdin)
		}
	} else {
		var entry *Entry
		entry, err = SearchEntryWithDate(self.db, date)
		if err != nil {
			return err
		}

		entry.Print(stdin)
	}

	stdin.Close()
	cmd.Run()

	return nil
}

func (self *Logbook) Close() {
	defer self.db.Close()
}

func timeParse(timeStr string) (*time.Time, error) {
	var err error
	if timeStr == "" {
		return nil, nil
	}
	var date time.Time
	date, err = time.Parse(DATE_FORMAT, timeStr)
	if err != nil {
		return nil, err
	}
	return &date, nil
}
