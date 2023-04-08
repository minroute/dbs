package dbs

import (
	"database/sql"

	structure "github.com/minroute/structure/set"
)

type Repo struct {
	WhereExp  string
	WhereArgs []structure.AnySetCore

	Having     string
	HavingArgs []structure.AnySetCore

	SelectExp  string
	SelectArgs []structure.AnySetCore
}

type Executer interface {
	Exec() (sql.Result, error)

	// Add Colunmns

}
