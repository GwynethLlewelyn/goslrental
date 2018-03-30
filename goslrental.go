package main

import (
	"fmt"
	"database/sql"
	"log"
	"github.com/cznic/ql"
	"runtime"
	"path/filepath"
	"time"
)

// setUp tries to create a table on the QL database for testing purposes.
func setUp(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS note (
	  id INT 
	  ,title STRING
	  ,body STRING
	  ,created_at STRING
	  ,updated_at STRING
	);
	`)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// main starts here.
func main() {
	ql.RegisterDriver()	// this should allow us to use the 'normal' SQL Go bindings to use QL.
	db, err := sql.Open("ql", "db.ql")
	defer db.Close()
	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}
	
	if err = setUp(db); err != nil {
		log.Fatalf("failed to create table: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %s", err)
	}
	
	// now insert some stuff and read from the database.
	tx, err := db.Begin()
	checkErr(err)
	
	stmt, err := tx.Prepare("INSERT INTO note (id, title, body, created_at, updated_at) VALUES (?1,?2,?3,?4,?5)");
	checkErrPanic(err)

	defer stmt.Close()
	
	curTime := fmt.Sprintf("%s", time.Now())
	
	_, err = stmt.Exec(1, "this is my note", "blah blah", curTime, curTime)
	checkErr(err)

	err = tx.Commit()	
	checkErr(err)
}

// checkErrPanic logs a fatal error and panics.
func checkErrPanic(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		log.Panic(filepath.Base(file), ":", line, ":", pc, ok, " - panic:", err)
	}
}

// checkErr checks if there is an error, and if yes, it logs it out and continues.
//  this is for 'normal' situations when we want to get a log if something goes wrong but do not need to panic
func checkErr(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		log.Print(filepath.Base(file), ":", line, ":", pc, ok, " - error:", err)
	}
}