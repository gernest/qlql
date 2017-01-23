package qlql

import (
	"github.com/gernest/qlql/components/photon"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type Index struct {
	vecty.Core
	tables []string
}

func (i *Index) Render() *vecty.HTML {
	return elem.Body(
		elem.Header(
			photon.Toolbar(true),
			photon.ToolbarHeader(true),
			elem.Header1(
				photon.Title(),
				vecty.Text("HEADER"),
			),
		),
		NewTableNave(i.tables...),
		elem.Footer(
			photon.Toolbar(),
			photon.ToolbarFooter(),
			elem.Header1(
				photon.Title(),
				vecty.Text("footer"),
			),
		),
	)
}

func (i *Index) AddTables(names ...string) {
	i.tables = append(i.tables, names...)
}

type TableNav struct {
	vecty.Core
	Tables vecty.List
}

func NewTableNave(tables ...string) *TableNav {
	n := &TableNav{}
	for _, v := range tables {
		n.Tables = append(n.Tables, &Item{Name: v})
	}
	return n
}

type Item struct {
	vecty.Core
	Name   string
	Active bool
}

func (i *Item) Render() *vecty.HTML {
	var state string
	if i.Active {
		state = "icon icon-minus"
	} else {
		state = "icon icon-plus"
	}
	return elem.Span(
		prop.Class("nav-group-item"),
		elem.Span(
			prop.Class(state),
			event.Click(func(e *vecty.Event) {
				i.Active = !i.Active
				vecty.Rerender(i)
			}),
		),
		vecty.Text(i.Name),
	)
}

func (t *TableNav) Render() *vecty.HTML {
	return elem.Navigation(
		prop.Class("nav-group"),
		elem.Header5(
			prop.Class("nav-group-title"),
			vecty.Text("Tables"),
		),
		t.Tables,
	)
}
