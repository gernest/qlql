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
func Window() vecty.Markup {
	return prop.Class(window)
}

//WindowContent applies window-content class
func WindowContent() vecty.Markup {
	return prop.Class(wContent)
}

//BtnGroup applies btn-group class
func BtnGroup() vecty.Markup {
	return prop.Class(btnGroup)
}

//Btn applies btn class
func Btn() vecty.Markup {
	return prop.Class(btn)
}

//Toolbar applies toolbar class
func Toolbar() vecty.Markup {
	return prop.Class(toolbar)
}

//ToolbarHeader applies toolbar-header class
func ToolbarHeader() vecty.Markup {
	return prop.Class(toolbarHeader)
}

//ToolbarFooter applies toolbar-footer class
func ToolbarFooter() vecty.Markup {
	return prop.Class(toolbarFooter)
}

//Title applies title class
func Title() vecty.Markup {
	return prop.Class(title)
}

//ToolbarBorderless applies toolbar-borderless class
func ToolbarBorderless() vecty.Markup {
	return prop.Class(toolbarBorderless)
}

//ToolbarActions applies toolbar-actions class
func ToolbarActions() vecty.Markup {
	return prop.Class(toolbarActions)
}
