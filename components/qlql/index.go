package qlql

import (
	"github.com/cathalgarvey/fmtless/encoding/json"
	"github.com/gernest/qlql/components/common"
	"github.com/gernest/qlql/components/photon"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"honnef.co/go/js/xhr"
)

type Index struct {
	vecty.Core
	baseURL  string
	infoChan chan common.DBInfo
	info     *common.DBInfo
	children chan common.DBInfo
}

func NewIdex(baseURL string) *Index {
	i := &Index{
		baseURL:  baseURL,
		infoChan: make(chan common.DBInfo),
		children: make(chan common.DBInfo),
	}
	return i.Start()
}
func (i *Index) Start() *Index {
	go func() {
		for {
			select {
			case v := <-i.infoChan:
				i.info = &v
				i.children <- v
				println("dispatced info")
			}
		}
	}()
	return i
}

func (i *Index) Render() *vecty.HTML {
	return elem.Body(
		elem.Div(
			photon.Window(),
			elem.Div(
				photon.WindowContent(),
				elem.Div(
					photon.PaneGroup(),
					elem.Div(
						photon.Pane(true),
						photon.PaneSM(true),
						NewTableNav(i.children),
					),
					elem.Div(
						photon.Pane(true),
						photon.PaneOneThird(true),
						NewDBConnect(i.infoChan, i.baseURL),
					),
				),
			),
		),
	)
}

type TableNav struct {
	vecty.Core
	Tables   vecty.List
	info     *common.DBInfo
	infoChan chan common.DBInfo
}

func NewTableNav(ch chan common.DBInfo) *TableNav {
	n := &TableNav{infoChan: ch}
	return n.Start()
}
func (t *TableNav) Start() *TableNav {
	go func() {
		println("started listening")
		for {
			select {
			case v := <-t.infoChan:
				t.info = &v
				vecty.Rerender(t)
			}
		}
	}()
	return t
}

func (t *TableNav) Render() *vecty.HTML {
	var name string
	if t.info != nil {
		name = t.info.Name
		println(name)
		var tables vecty.List
		for _, v := range t.info.Tables {
			println(v.Name)
			if len(v.Name) > 2 {
				if v.Name[0] == '_' && v.Name[1] == '_' {
					continue
				}
			}
			tables = append(tables, &wrapTable{t: v})
		}
		t.Tables = tables
	} else {
		name = "Tables"
		t.Tables = append(t.Tables, elem.Div())
	}
	return elem.Navigation(
		prop.Class("nav-group"),
		elem.Header5(
			prop.Class("nav-group-title"),
			vecty.Text(name),
		),
		t.Tables,
	)
}

type DBConnect struct {
	vecty.Core
	infoChan  chan common.DBInfo
	input     string
	baseURL   string
	available []string
}

func NewDBConnect(ch chan common.DBInfo, base string) *DBConnect {
	return &DBConnect{
		infoChan: ch,
		baseURL:  base,
	}
}

func (d *DBConnect) Render() *vecty.HTML {
	all := d.availableDB()
	var opts vecty.List
	for _, v := range all {
		opts = append(opts, elem.Option(
			vecty.Text(v),
			prop.Value(v),
		))
	}
	return elem.Div(
		photon.WindowContent(),
		elem.Form(
			prop.Class("padded"),
			elem.Select(
				elem.Option(
					vecty.Text("select available databases"),
				),
				opts,
				event.Change(func(e *vecty.Event) {
					d.input = e.Target.Get("value").String()
					vecty.Rerender(d)
				}),
			),
			elem.Input(
				prop.Value(d.input),
			),
			elem.Button(
				vecty.Text("connect"),
			),
			event.Submit(func(e *vecty.Event) {
				go d.connect(e.Target.Index(1).Get("value").String())
			}).PreventDefault(),
		),
	)
}

func (d *DBConnect) availableDB() []string {
	if d.available != nil {
		return d.available
	}
	b, err := xhr.Send("GET", d.baseURL+"/all", nil)
	if err != nil {
		println(err.Error())
		return []string{}
	}
	var out []string
	err = json.Unmarshal(b, &out)
	if err != nil {
		println(err.Error())
		return []string{}
	}
	d.available = out
	return out
}
func (d *DBConnect) connect(db string) {
	b, err := xhr.Send("GET", d.baseURL+"/info?db="+db, nil)
	if err != nil {
		println(err.Error())
		return
	}
	//println(string(b))
	i := common.DBInfo{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		println(err.Error())
		return
	}
	d.infoChan <- i
}

type wrapTable struct {
	vecty.Core
	t common.Table
}

func (w *wrapTable) Render() *vecty.HTML {
	var state string
	var display string
	if w.t.Active {
		state = "icon icon-minus"
		display = "block"
	} else {
		display = "none"
		state = "icon icon-plus"
	}
	var cols vecty.List
	for _, v := range w.t.Columns {
		cols = append(cols, &wrapCol{t: v})
	}

	return elem.Span(
		prop.Class("nav-group-item"),
		elem.Span(
			prop.Class(state),
			event.Click(func(e *vecty.Event) {
				w.t.Active = !w.t.Active
				vecty.Rerender(w)
			}),
		),
		vecty.Text(w.t.Name),
		elem.Navigation(
			prop.Class("nav-group"),
			vecty.Style("display", display),
			elem.Header5(
				prop.Class("nav-group-title"),
				vecty.Text("columns"),
			),
			cols,
		),
	)
}

type wrapCol struct {
	vecty.Core
	t common.Column
}

func keyValue(key, value string) *vecty.HTML {
	return elem.Span(
		prop.Class("nav-group-item"),
		vecty.Text(key+":"+value),
	)
}

func (w *wrapCol) Render() *vecty.HTML {
	var state string
	var display string
	if w.t.Active {
		state = "icon icon-minus"
		display = "block"
	} else {
		display = "none"
		state = "icon icon-plus"
	}
	var cols vecty.List
	cols = append(cols, keyValue("Type ", w.t.Type))

	return elem.Span(
		prop.Class("nav-group-item"),
		elem.Span(
			prop.Class(state),
			event.Click(func(e *vecty.Event) {
				w.t.Active = !w.t.Active
				vecty.Rerender(w)
			}),
		),
		vecty.Text(w.t.Name),
		elem.Navigation(
			prop.Class("nav-group"),
			vecty.Style("display", display),
			elem.Header5(
				prop.Class("nav-group-title"),
				vecty.Text("properties"),
			),
			cols,
		),
	)
}
