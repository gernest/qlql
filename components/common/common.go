package common

type Table struct {
	Name    string
	Columns []Column
	Active  bool
}

type Column struct {
	Name       string
	Type       string
	NotNull    bool
	Constraint string
	Default    string
	Active     bool
}

type Index struct {
	Name           string
	Table          string
	Column         string
	Unique         bool
	ExpressionList []string
	Active         bool
}

type DBInfo struct {
	Name    string
	Tables  []Table
	Indices []Index
}

type ExecReq struct {
	DB    string
	Query string
	Tx    bool
}

type ExecRes struct {
	Error   string
	Results []Record
}

type Record struct {
	Fields  []string
	Results []string
}
