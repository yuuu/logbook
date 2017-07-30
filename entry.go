package main

import (
	"database/sql"
	"errors"
	"fmt"
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

const DATE_FORMAT = "2006-01-02"

type Entry struct {
	id   int64
	date *time.Time
	text *Text
}

func CreateEntry(db *sql.DB, text *Text, date *time.Time) (*Entry, error) {
	var err error

	insertQuery := fmt.Sprintf(INSERT_ENTRY_QUERY, text, date.Format(DATE_FORMAT))

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

func SearchEntryWithDate(db *sql.DB, date *time.Time) ([]Entry, error) {
	return nil, nil
}

func SearchEntryWithID(db *sql.DB, id int64) (*Entry, error) {
	var err error
	var rows *sql.Rows

	rows, err = db.Query(fmt.Sprintf(SELECT_ENTRY_QUERY_WITH_ID, id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var text string
	var date time.Time
	if rows.Next() {
		rows.Scan(&text, &date)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("rows.Next() == false")
	}

	var entry Entry
	entry.text = CreateText(text)
	entry.date = &date

	return &entry, nil
}

func (self *Entry) Edit() error {
	return nil
}

func (self *Entry) Delete() error {
	return nil
}

func (self *Entry) Print() {
	fmt.Print("DATE: ", self.date.Format(DATE_FORMAT))
	fmt.Print("TEXT: ")
	fmt.Print(self.text.Text())
}
