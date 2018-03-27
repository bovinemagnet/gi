// Copyright (c) 2018, Randall C. O'Reilly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/rcoreilly/goki/gi"
	_ "github.com/rcoreilly/goki/gi/init"
	"github.com/rcoreilly/goki/gi/units"
	"github.com/rcoreilly/goki/ki"
)

func main() {
	go mainrun()
	gi.RunBackendEventLoop() // this needs to run in main loop
}

func mainrun() {
	width := 800
	height := 800

	recv := ki.Node{}               // receiver for events
	recv.SetThisName(&recv, "recv") // this is essential for root objects not owned by other Ki tree nodes

	win := gi.NewWindow2D("GoGi Widgets Window", width, height)
	win.UpdateStart()

	vp := win.WinViewport2D()
	vp.SetProp("background-color", "#FFF")
	vp.Fill = true

	vlay := vp.AddNewChildNamed(gi.KiT_Frame, "vlay").(*gi.Frame)
	vlay.Lay = gi.LayoutCol

	row1 := vlay.AddNewChildNamed(gi.KiT_Layout, "row1").(*gi.Layout)
	row1.Lay = gi.LayoutRow
	row1.SetProp("align-vert", "middle")
	row1.SetProp("align-horiz", "center")
	row1.SetProp("margin", 4.0)
	row1.SetProp("max-width", -1)

	spc := vlay.AddNewChildNamed(gi.KiT_Space, "spc1").(*gi.Space)
	spc.SetFixedHeight(units.NewValue(2.0, units.Em))

	row1.AddNewChildNamed(gi.KiT_Stretch, "str1")
	lab1 := row1.AddNewChildNamed(gi.KiT_Label, "lab1").(*gi.Label)
	lab1.Text = "This is a demonstration of the various GoGi Widgets"
	lab1.SetProp("max-width", -1)
	lab1.SetProp("text-align", "center")
	row1.AddNewChildNamed(gi.KiT_Stretch, "str2")

	row2 := vlay.AddNewChildNamed(gi.KiT_Layout, "row2").(*gi.Layout)
	row2.Lay = gi.LayoutRow
	row2.SetProp("align-vert", "center")
	row2.SetProp("align-horiz", "left")
	row2.SetProp("margin", 4.0)
	row2.SetProp("max-width", -1)

	mb1 := row2.AddNewChildNamed(gi.KiT_MenuButton, "menubutton1").(*gi.MenuButton)

	mb1.Text = "Menu Button"
	mb1.AddMenuText("Menu Item 1", recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received menu action signal: %v from menu action: %v\n", gi.ActionSignals(sig), send.KiName())
	})

	mb1.AddMenuText("Menu Item 2", recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received menu action signal: %v from menu action: %v\n", gi.ActionSignals(sig), send.KiName())
	})

	mb1.AddMenuText("Menu Item 3", recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received menu action signal: %v from menu action: %v\n", gi.ActionSignals(sig), send.KiName())
	})

	button1 := row2.AddNewChildNamed(gi.KiT_Button, "button1").(*gi.Button)
	button2 := row2.AddNewChildNamed(gi.KiT_Button, "button2").(*gi.Button)
	edit1 := row2.AddNewChildNamed(gi.KiT_TextField, "edit1").(*gi.TextField)

	edit1.Text = "Edit this text"
	edit1.SetProp("min-width", "20em")
	button1.Text = "Button 1"
	button2.Text = "Button 2"

	edit1.TextFieldSig.Connect(recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received line edit signal: %v from edit: %v with data: %v\n", gi.TextFieldSignals(sig), send.KiName(), data)
	})

	button1.ButtonSig.Connect(recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.KiName())
	})

	button2.ButtonSig.Connect(recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.KiName())
	})

	row3 := vlay.AddNewChildNamed(gi.KiT_Layout, "row3").(*gi.Layout)
	row3.Lay = gi.LayoutRow
	row3.SetProp("align-vert", "center")
	row3.SetProp("align-horiz", "left")
	row3.SetProp("margin", 4.0)
	row3.SetProp("max-width", -1) // always stretch width

	slider1 := row3.AddNewChildNamed(gi.KiT_Slider, "slider1").(*gi.Slider)
	slider1.Horiz = true
	slider1.SetProp("width", "20em")
	slider1.Defaults()
	slider1.SetValue(0.5)
	slider1.Snap = true
	slider1.Tracking = true

	slider2 := row3.AddNewChildNamed(gi.KiT_Slider, "slider2").(*gi.Slider)
	slider2.Horiz = false
	slider2.SetProp("height", "10em")
	slider2.Defaults()
	slider2.SetValue(0.5)

	slider1.SliderSig.Connect(recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received slider signal: %v from slider: %v with data: %v\n", gi.SliderSignals(sig), send.KiName(), data)
	})

	slider2.SliderSig.Connect(recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received slider signal: %v from slider: %v with data: %v\n", gi.SliderSignals(sig), send.KiName(), data)
	})

	scrollbar1 := row3.AddNewChildNamed(gi.KiT_ScrollBar, "scrollbar1").(*gi.ScrollBar)
	scrollbar1.Horiz = true
	scrollbar1.SetProp("width", "20em")
	scrollbar1.SetFixedHeight(units.NewValue(20, units.Px))
	scrollbar1.Defaults()
	scrollbar1.SetThumbValue(0.25)
	scrollbar1.SetValue(0.25)
	// scrollbar1.Snap = true
	scrollbar1.Tracking = true

	scrollbar2 := row3.AddNewChildNamed(gi.KiT_ScrollBar, "scrollbar2").(*gi.ScrollBar)
	scrollbar2.Horiz = false
	scrollbar2.SetProp("height", "10em")
	scrollbar2.Defaults()
	scrollbar2.SetThumbValue(0.1)
	scrollbar2.SetValue(0.5)

	scrollbar1.SliderSig.Connect(recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received scrollbar signal: %v from scrollbar: %v with data: %v\n", gi.SliderSignals(sig), send.KiName(), data)
	})

	scrollbar2.SliderSig.Connect(recv.This, func(rec, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received scrollbar signal: %v from scrollbar: %v with data: %v\n", gi.SliderSignals(sig), send.KiName(), data)
	})

	win.UpdateEnd()

	win.StartEventLoop()
}