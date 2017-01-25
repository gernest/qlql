package common

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Table struct {
	Name    string
	Columns []Column
	Active  bool
}

func (i Table) Render() *vecty.HTML {
	return nil
}

type Column struct {
	Name       string
	Type       string
	NotNull    bool
	Constraint string
	Default    string
	Active     bool
}

func (i Column) Render() *vecty.HTML {
	return nil
}

type Index struct {
	Name           string
	Table          string
	Column         string
	Unique         bool
	ExpressionList []string
	Active         bool
}

func (i Index) Render() *vecty.HTML {
	return elem.Div(
		prop.Class("nav-group-item"),
		elem.Span(
			prop.Class("icon icon-star"),
		),
		vecty.Text(i.Name),
	)
}

type DBInfo struct {
	Name    string
	Tables  []Table
	Indices []Index
}

func (i DBInfo) Render() *vecty.HTML {
	return nil
}
