package main

import (
	"github.com/gernest/qlql/components/qlql"
	"github.com/gopherjs/vecty"
)

func main() {
	vecty.SetTitle("my way")
	i := &qlql.Index{}
	i.AddTables("users", "languages", "email")
	vecty.RenderBody(i)
}
