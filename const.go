package dbs

const (
	// SQL KEYWORD

	SQL_Like       = "LIKE"
	SQL_NOT_LIKE   = "NO LIKE"
	SQL_IsNull     = "IS NULL"
	SQL_IsNotNull  = "IS NOT NULL"
	SQL_From       = "FROM"
	SQL_Where      = "WHERE"
	SQL_And        = "AND"
	SQL_Or         = "OR"
	SQL_Limit      = "LIMIT"
	SQL_Offset     = "OFFSET"
	SQL_GroupBy    = "GROUP BY"
	SQL_Having     = "HAVING"
	SQL_OrderBy    = "ORDER BY"
	SQL_Select     = "SELECT"
	SQL_Insert     = "INSERT"
	SQL_Update     = "UPDATE"
	SQL_Delete     = "DELETE"
	SQL_Values     = "VALUES"
	SQL_Set        = "SET"
	SQL_In         = "IN"
	SQL_NotIn      = "NOT IN"
	SQL_Between    = "BETWEEN"
	SQL_NotBetween = "NOT BETWEEN"
	SQL_Exists     = "EXISTS"
	SQL_NotExists  = "NOT EXISTS"
	SQL_As         = "AS"
	SQL_On         = "ON"
	SQL_Join       = "JOIN"
	SQL_LeftJoin   = "LEFT JOIN"
	SQL_RightJoin  = "RIGHT JOIN"
	SQL_FullJoin   = "FULL JOIN"
	SQL_Union      = "UNION"
	SQL_UnionAll   = "UNION ALL"
	SQL_ASC        = "ASC"
	SQL_DESC       = "DESC"
)

// Space function add space before and after the string
// Just for sql builder
func Space(s string) string {
	return " " + s + " "
}
