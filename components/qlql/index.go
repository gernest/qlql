package qlql

import (
	"github.com/gernest/qlql/components/photon"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Index struct {
	vecty.Core
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
