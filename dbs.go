package dbs

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// 安全事务-计数器方式
// 每次开启事务时，都会将db.txCount+1，每次提交事务时，都会将db.txCount-1，
// 当DB.txCount=1(即txCommitNum)时，才会真正提交事务。
const txCommitNum = 1

type (

	// LangFunc I18n Support, default is enlish
	// The default function is a wrapper of  fmt.Sprintf function ,then you can use it like fmt.Sprintf:
	//      LangFunc("Fortmat String %v", "value"。。。。。)
	//      LangFunc("Single¸String")
	LangFunc func(v ...any) string

	// LogFunc logs a message for each SQL statement being executed.
	// This method takes one or multiple parameters. If a single parameter
	// is provided, it will be treated as the log message. If multiple parameters
	// are provided, they will be passed to fmt.Sprintf() to generate the log message.
	LogFunc func(format string, v ...any)

	// QueryOrExecuteLogFunc will be called for each query or execute statement.
	// op is either "query" or "execute".
	// t is the time taken to execute the statement.
	// sql is the SQL statement which you run
	// rows is the result of the query or the number of affected rows.
	// err is the error if any.
	QueryOrExecuteLogFunc = func(ctx context.Context, op string, t time.Duration, sql string, rows *sql.Rows, err error)

	// DB  enhances sql.DB by providing a set of DB-agnostic query building methods.
	// DB allows easier query building and population of data into Go variables.
	DB struct {
		TableName             TableNameFunc
		FieldMapper           FieldMapFunc
		LogFunc               LogFunc
		LangFunc              LangFunc
		QueryOrExecuteLogFunc QueryOrExecuteLogFunc

		driver string
		ctx    context.Context

		// db is the underlying sql.DB object.
		db *sql.DB

		// Transaction
		tx *sql.Tx

		// Nested Transaction Counter,
		txCounter int

		// Nested Transaction mode, True is savepoint mode,False is counter mode
		txModeIsSavePoint bool

		// Rollback flag，for counter mode
		isRollback bool

		// For Transaction with savepoint
		txNext *DB

		// For Transaction with savepoint
		txResolved bool
	}
)

// Ping
//  @Description:  check the connection
//  @receiver d
//  @return error
func (d *DB) Ping() error {
	if err := d.db.Ping(); err != nil {
		return err
	}
	return nil
}

// Close
//  @Description: close the database connection
//  @Notes:It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
//  @receiver d
//  @return error
func (d *DB) Close() error {
	return d.db.Close()
}

// Lang
//  @Description: get the language translation function
//  @todo: add i18n support
func (d *DB) Lang(v ...any) string {
	if d.LangFunc != nil {
		return d.LangFunc(v...)
	}
	if len(v) > 1 {
		return fmt.Sprintf(v[0].(string), v...)
	}
	return fmt.Sprintf("%v", v[0])
}
