package photon

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/prop"
)

const (
	window            = "window"
	wContent          = "window-content"
	btnGroup          = "btn-group"
	btn               = "btn"
	toolbar           = "toolbar"
	toolbarHeader     = "toolbar-header"
	toolbarFooter     = "toolbar-footer"
	toolbarBorderless = "toolbar-borderless"
	title             = "title"
	toolbarActions    = "toolbar-actions"
)

//Window applies window class
func Window(append ...bool) vecty.Markup {
	return addClass(window, append...)
}

//WindowContent applies window-content class
func WindowContent(append ...bool) vecty.Markup {
	return addClass(wContent, append...)
}

//BtnGroup applies btn-group class
func BtnGroup(append ...bool) vecty.Markup {
	return addClass(btnGroup, append...)
}

//Btn applies btn class
func Btn(append ...bool) vecty.Markup {
	return addClass(btn, append...)
}

//Toolbar applies toolbar class
func Toolbar(append ...bool) vecty.Markup {
	return addClass(toolbar, append...)
}

//ToolbarHeader applies toolbar-header class
func ToolbarHeader(append ...bool) vecty.Markup {
	if len(append) > 0 {
		if ok := append[0]; ok {
			return appendClass(toolbarHeader)
		}
	}
	return prop.Class(toolbarHeader)
}

//ToolbarFooter applies toolbar-footer class
func ToolbarFooter(append ...bool) vecty.Markup {
	return addClass(toolbarFooter, append...)
}

//Title applies title class
func Title(append ...bool) vecty.Markup {
	return addClass(title, append...)
}

//ToolbarBorderless applies toolbar-borderless class
func ToolbarBorderless(append ...bool) vecty.Markup {
	return addClass(toolbarBorderless, append...)
}

//ToolbarActions applies toolbar-actions class
func ToolbarActions(append ...bool) vecty.Markup {
	return addClass(toolbarActions, append...)
}

func appendClass(class string) vecty.Markup {
	return vecty.AppendProperty("className", class)
}

func addClass(class string, append ...bool) vecty.Markup {
	if len(append) > 0 {
		if ok := append[0]; ok {
			return appendClass(class)
		}
	}
	return prop.Class(class)
}
