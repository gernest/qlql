package photon

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/prop"
)

const (
	// base
	window            = "window"
	wContent          = "window-content"
	toolbar           = "toolbar"
	toolbarHeader     = "toolbar-header"
	toolbarFooter     = "toolbar-footer"
	toolbarBorderless = "toolbar-borderless"
	title             = "title"
	toolbarActions    = "toolbar-actions"

	// buttons
	btnGroup = "btn-group"
	btn      = "btn"

	//forms
	formCtrl       = "form-control"
	formActions    = "form-actions"
	checkBox       = "checkbox"
	checkboxInline = "ccheckbox-inline"

	///grid
	paneGroup     = "pane-group"
	pane          = "pane"
	paneSM        = "pane-sm"
	paneMini      = "pane-mini"
	paneOneFourth = "pane-one-fourth"
	paneOneThird  = "pane-one-third"

	//image
	imgCircle  = "img-circle"
	imgRounded = "img-rounded"

	//list
	listGroup       = "list-group"
	listGroupItem   = "list-group-item"
	listGroupHeader = "list-group-header"
	mediaObject     = "media-object"
	mediaBody       = "media-body"
)

func ListGroup(append ...bool) vecty.Markup {
	return addClass(listGroup, append...)
}
func ListGroupItem(append ...bool) vecty.Markup {
	return addClass(listGroupItem, append...)
}
func ListGroupHeader(append ...bool) vecty.Markup {
	return addClass(listGroupHeader, append...)
}
func ImgCircle(append ...bool) vecty.Markup {
	return addClass(imgCircle, append...)
}
func MediaObject(append ...bool) vecty.Markup {
	return addClass(mediaObject, append...)
}
func MediaBody(append ...bool) vecty.Markup {
	return addClass(mediaBody, append...)
}
func ImgRounded(append ...bool) vecty.Markup {
	return addClass(imgRounded, append...)
}
func Pane(append ...bool) vecty.Markup {
	return addClass(pane, append...)
}
func PaneGroup(append ...bool) vecty.Markup {
	return addClass(paneGroup, append...)
}
func PaneSM(append ...bool) vecty.Markup {
	return addClass(paneSM, append...)
}

func PaneMimi(append ...bool) vecty.Markup {
	return addClass(paneMini, append...)
}
func PaneOneFourth(append ...bool) vecty.Markup {
	return addClass(paneOneFourth, append...)
}

func PaneOneThird(append ...bool) vecty.Markup {
	return addClass(paneOneThird, append...)
}

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
