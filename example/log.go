package example

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/minroute/dbs"
)

func main() {
	dbs, err := dbs.Open("mysql", "user:pass@/example", true)
	if err != nil {
		log.Fatal(err)
	}

	// 执行日志函数
	dbs.LogFunc = log.Printf

	// query 或者execute 时的执行情况
	dbs.QueryOrExecuteLogFunc = func(ctx context.Context, op string, t time.Duration, sql string, rows *sql.Rows, err error) {
		log.Printf("op=%s, t=[%.2fms], sql=%v, err=%s", op, float64(t.Milliseconds()), sql, err)
	}
}
