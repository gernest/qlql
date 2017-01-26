package main

import (
	"github.com/gernest/qlql/components/qlql"
	"github.com/gopherjs/vecty"
)

func main() {
	vecty.SetTitle("qlql")
	vecty.RenderBody(qlql.NewIdex("http://localhost:8000"))
}
