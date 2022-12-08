package dbs

import (
	"database/sql"

	kit "github.com/xiiapp/kit/sctucture/set"
)

type Repo struct {
	WhereExp  string
	WhereArgs []kit.AnySetCore

	Having     string
	HavingArgs []kit.AnySetCore

	SelectExp  string
	SelectArgs []kit.AnySetCore
}

type Executer interface {
	Exec() (sql.Result, error)

	// Add Colunmns

}
