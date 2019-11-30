// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StackPanel represents a babylon.js StackPanel.
// Class used to create a 2D stack panel container
type StackPanel struct{ *Container }

// JSObject returns the underlying js.Value.
func (s *StackPanel) JSObject() js.Value { return s.p }

// StackPanel returns a StackPanel JavaScript class.
func (ba *Babylon) StackPanel() *StackPanel {
	p := ba.ctx.Get("StackPanel")
	return StackPanelFromJSObject(p)
}

// StackPanelFromJSObject returns a wrapped StackPanel JavaScript class.
func StackPanelFromJSObject(p js.Value) *StackPanel {
	return &StackPanel{ContainerFromJSObject(p)}
}

// NewStackPanelOpts contains optional parameters for NewStackPanel.
type NewStackPanelOpts struct {
	Name *JSString
}

// NewStackPanel returns a new StackPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel
func (ba *Babylon) NewStackPanel(opts *NewStackPanelOpts) *StackPanel {
	if opts == nil {
		opts = &NewStackPanelOpts{}
	}

	p := ba.ctx.Get("StackPanel").New(opts.Name)
	return StackPanelFromJSObject(p)
}

// TODO: methods
