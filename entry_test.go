package main

import (
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const TEST_DB = "./test.db"

func Test_CreateEntry(t *testing.T) {
	var db *sql.DB
	//var stmt *sql.Stmt
	var err error
	var entry *Entry
	//var nowTime time.Time

	db, err = sql.Open("sqlite3", TEST_DB)
	if err != nil {
		t.Fatal("sql.Open is failed. (", err.Error(), ")")
	}
	defer db.Close()
	defer os.Remove(TEST_DB)

	tim, _ := time.Parse(DATE_FORMAT, "2017-07-27")
	entry, err = CreateEntry(db, CreateText("abcd"), &tim)
	if err != nil {
		t.Fatal("Create Entry is failed. (", err.Error(), ")")
	}
	if entry.text.Text() != "abcd" {
		t.Fatal("entry.text is incorrect. (", entry.text.Text(), entry.date, entry.id, ")")
	}
	if entry.date.Format("2006-01-02") != "2017-07-27" {
		t.Fatal("entry.date is incorrect. (", entry.text.Text(), entry.date, entry.id, ")")
	}

	/*
		stmt, err = db.Prepare("INSERT INTO entry (text, date) VALUES ('abc', 2017-7-24)")
		if err != nil {
			t.Fatal("db.Prepare is failed. (", err.Error(), ")")
		}

		nowTime = time.Now()
		entry, err = CreateEntry(db, stmt, &nowTime)
		if err != nil {
			t.Fatal("CreateEntry is failed.")
		}
	*/

	entry.Edit()
}
