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
	tables   []string
	baseURL  string
	infoChan chan common.DBInfo
	info     *common.DBInfo
	children []chan common.DBInfo
	tableNav *TableNav
}

func NewIdex(baseURL string) *Index {
	i := &Index{baseURL: baseURL}
	return i.Start()
}

func (i *Index) Start() *Index {
	go func() {
		for {
			select {
			case v := <-i.infoChan:
				i.info = &v
				for _, c := range i.children {
					c <- v
				}
				vecty.Rerender(i)
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
						NewTableNav(i.infoChannel()),
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

func (i Index) infoChannel() <-chan common.DBInfo {
	c := make(chan common.DBInfo)
	i.children = append(i.children, c)
	return c

}

func (i *Index) AddTables(names ...string) {
	i.tables = append(i.tables, names...)
}

type TableNav struct {
	vecty.Core
	Tables   vecty.List
	info     *common.DBInfo
	infoChan <-chan common.DBInfo
}

func (t *TableNav) Start() *TableNav {
	go func() {
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

func NewTableNav(ch <-chan common.DBInfo) *TableNav {
	n := &TableNav{infoChan: ch}
	return n
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

type DBConnect struct {
	vecty.Core
	infoChan chan common.DBInfo
	input    string
	baseURL  string
}

func NewDBConnect(ch chan common.DBInfo, base string) *DBConnect {
	return &DBConnect{
		infoChan: make(chan common.DBInfo),
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
	return elem.Form(
		elem.Select(
			opts,
		),
		elem.Input(
			//photon.AddClass("form-control", true),
			prop.Type(prop.TypeUrl),
			event.Input(func(e *vecty.Event) {
				d.input = e.Target.Get("value").String()
			}),
		),
		elem.Button(
			vecty.Text("open"),
		),
	)
}

func (d *DBConnect) availableDB() []string {
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
	println(string(b))
	return out
}
