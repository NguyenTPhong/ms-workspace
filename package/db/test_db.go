package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"
	
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetTestDB(url string) (*gorm.DB, func()) {

	if url == "" {
		panic("no test db connection string")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can not connect to database", err.Error())
	}

	databaseName := fmt.Sprintf("test_%v", time.Now().Nanosecond())
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", databaseName))
	if err != nil {
		fmt.Println(err)
		log.Fatal("can not create test database")
	}

	dbTest, connectErr := NewDatabase(fmt.Sprintf("%v dbname=%v ", url, databaseName), 2, 1, 4)
	if connectErr != nil {
		log.Fatal(errors.New("no db connection"))
	}

	return dbTest, func() {
		Close(dbTest)
		_, e := db.Exec(fmt.Sprintf("DROP DATABASE %s;", databaseName))
		if e != nil {
			fmt.Println("fail to delete database", e)
		}
		db.Close()
	}
}
