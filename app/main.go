package main

import (
	"github.com/gernest/qlql/components/qlql"
	"github.com/gopherjs/vecty"
)

func main() {
	vecty.SetTitle("my way")
	vecty.RenderBody(&qlql.Index{})
}
