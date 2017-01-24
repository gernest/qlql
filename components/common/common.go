package common

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name       string
	Type       string
	NotNull    bool
	Constraint string
	Default    string
}
type Index struct {
	Name           string
	Table          string
	Column         string
	Unique         bool
	ExpressionList []string
}

type DBInfo struct {
	Name    string
	Tables  []Table
	Indices []Index
}
