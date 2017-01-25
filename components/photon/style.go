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
	return AddClass(listGroup, append...)
}
func ListGroupItem(append ...bool) vecty.Markup {
	return AddClass(listGroupItem, append...)
}
func ListGroupHeader(append ...bool) vecty.Markup {
	return AddClass(listGroupHeader, append...)
}
func ImgCircle(append ...bool) vecty.Markup {
	return AddClass(imgCircle, append...)
}
func MediaObject(append ...bool) vecty.Markup {
	return AddClass(mediaObject, append...)
}
func MediaBody(append ...bool) vecty.Markup {
	return AddClass(mediaBody, append...)
}
func ImgRounded(append ...bool) vecty.Markup {
	return AddClass(imgRounded, append...)
}
func Pane(append ...bool) vecty.Markup {
	return AddClass(pane, append...)
}
func PaneGroup(append ...bool) vecty.Markup {
	return AddClass(paneGroup, append...)
}
func PaneSM(append ...bool) vecty.Markup {
	return AddClass(paneSM, append...)
}

func PaneMiii(append ...bool) vecty.Markup {
	return AddClass(paneMini, append...)
}
func PaneOneFourth(append ...bool) vecty.Markup {
	return AddClass(paneOneFourth, append...)
}

func PaneOneThird(append ...bool) vecty.Markup {
	return AddClass(paneOneThird, append...)
}

//Window applies window class
func Window(append ...bool) vecty.Markup {
	return AddClass(window, append...)
}

//WindowContent applies window-content class
func WindowContent(append ...bool) vecty.Markup {
	return AddClass(wContent, append...)
}

//BtnGroup applies btn-group class
func BtnGroup(append ...bool) vecty.Markup {
	return AddClass(btnGroup, append...)
}

//Btn applies btn class
func Btn(append ...bool) vecty.Markup {
	return AddClass(btn, append...)
}

//Toolbar applies toolbar class
func Toolbar(append ...bool) vecty.Markup {
	return AddClass(toolbar, append...)
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
	return AddClass(toolbarFooter, append...)
}

//Title applies title class
func Title(append ...bool) vecty.Markup {
	return AddClass(title, append...)
}

//ToolbarBorderless applies toolbar-borderless class
func ToolbarBorderless(append ...bool) vecty.Markup {
	return AddClass(toolbarBorderless, append...)
}

//ToolbarActions applies toolbar-actions class
func ToolbarActions(append ...bool) vecty.Markup {
	return AddClass(toolbarActions, append...)
}

func appendClass(class string) vecty.Markup {
	return vecty.AppendProperty("className", class)
}

func AddClass(class string, append ...bool) vecty.Markup {
	if len(append) > 0 {
		if ok := append[0]; ok {
			return appendClass(class)
		}
	}
	return prop.Class(class)
}
