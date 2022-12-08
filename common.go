package dbs

import (
	"context"
	"database/sql"
)

// Open
//  @Description:
//  @param driverName
//  @param dsn
//  @param immedia bool check the connection immediately if true
//  @return *DB
//  @return error
func Open(driverName, dsn string, immediate bool) (*DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if immediate {
		if err = db.Ping(); err != nil {
			return nil, err
		}
	}

	return &DB{
		driver: driverName,
		ctx:    context.Background(),
		db:     db,
	}, nil
}

// TestOpen
//  @Description: helper function , will test the connection and return the error if failed,or nil if success
//  @param driverName
//  @param dsn
//  @return error
func TestOpen(driverName, dsn string) error {
	if _, err := Open(driverName, dsn, true); err != nil {
		return err
	}
	return nil
}

// Close
//  @Description: close the database connection.
//  it is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
//  @param d
//  @return error
func Close(d *DB) error {
	return d.db.Close()
}
