package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"time"
)

const CREATE_TABLE_QUERY = `
CREATE TABLE entry (
id INTEGER PRIMARY KEY AUTOINCREMENT, 
text TEXT NOT NULL, 
date TIMESTAMP NOT NULL)
`
const INSERT_ENTRY_QUERY = `
INSERT INTO entry (text, date) 
VALUES ('%s', '%s')
`

const SELECT_ENTRY_QUERY_WITH_ID = `
SELECT text, date from entry where id = %d
`
const SELECT_ENTRY_QUERY_WITH_DATE = `
SELECT text, date from entry where date = '%s'
`

const SELECT_MOST_RECENT_ENTRY_QUERY = `
SELECT text, date from entry ORDER BY id DESC LIMIT %d OFFSET %d
`

const DATE_FORMAT = "2006-01-02"

type Entry struct {
	id   int64
	date *time.Time
	text *Text
}

func CreateEntry(db *sql.DB, text *Text, date *time.Time) (*Entry, error) {
	var err error

	insertQuery := fmt.Sprintf(INSERT_ENTRY_QUERY, text.Text(), date.Format(DATE_FORMAT))

	var ret sql.Result
	ret, err = db.Exec(insertQuery)
	if err != nil {
		ret, err = db.Exec(CREATE_TABLE_QUERY)
		if err != nil {
			return nil, err
		}
		ret, err = db.Exec(insertQuery)
		if err != nil {
			return nil, err
		}
	}

	var id int64
	id, err = ret.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Entry{id, date, text}, nil
}

func createSearchResultEntry(rows *sql.Rows) ([]*Entry, error) {
	var err error
	var entries []*Entry

	for rows.Next() {
		var text string
		var date time.Time
		rows.Scan(&text, &date)
		if err != nil {
			return nil, err
		}
		var entry Entry
		entry.text = CreateText(text)
		entry.date = &date
		entries = append(entries, &entry)
	}
	if len(entries) == 0 {
		return nil, errors.New("rows.Next() == false")
	}

	return entries, nil
}

func SearchMostRecentEntry(db *sql.DB, limit int, offset int) ([]*Entry, error) {
	var err error
	var rows *sql.Rows

	rows, err = db.Query(fmt.Sprintf(SELECT_MOST_RECENT_ENTRY_QUERY, limit, offset))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*Entry
	entries, err = createSearchResultEntry(rows)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func SearchEntryWithDate(db *sql.DB, targetDate *time.Time) (*Entry, error) {
	var err error
	var rows *sql.Rows

	rows, err = db.Query(fmt.Sprintf(SELECT_ENTRY_QUERY_WITH_DATE, targetDate.Format(DATE_FORMAT)))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*Entry
	entries, err = createSearchResultEntry(rows)
	if err != nil {
		return nil, err
	}

	return entries[0], nil
}

func SearchEntryWithID(db *sql.DB, id int64) (*Entry, error) {
	var err error
	var rows *sql.Rows

	rows, err = db.Query(fmt.Sprintf(SELECT_ENTRY_QUERY_WITH_ID, id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*Entry
	entries, err = createSearchResultEntry(rows)
	if err != nil {
		return nil, err
	}

	return entries[0], nil
}

func (self *Entry) Edit() error {
	return nil
}

func (self *Entry) Delete() error {
	return nil
}

func (self *Entry) Print(writer io.Writer) {
	fmt.Fprintln(writer, "================================")
	fmt.Fprintln(writer, ("DATE: " + self.date.Format(DATE_FORMAT)))
	fmt.Fprintln(writer, self.text.Text())
	fmt.Fprintln(writer, "")
}
