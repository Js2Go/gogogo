package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDbConn(user, password, host, port, db string) (*sql.DB, error) {
	return initDB(user, password, host, port, db)
}

func initDB(user, password, host, port, db string) (Db *sql.DB, err error) {
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db
	
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	err = Db.Ping()
	if err != nil {
		return nil, err
	}
	return
}
