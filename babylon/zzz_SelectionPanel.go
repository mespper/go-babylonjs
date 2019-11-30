// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SelectionPanel represents a babylon.js SelectionPanel.
// Class used to hold the controls for the checkboxes, radio buttons and sliders
//
// See: http://doc.babylonjs.com/how_to/selector
type SelectionPanel struct{ *Rectangle }

// JSObject returns the underlying js.Value.
func (s *SelectionPanel) JSObject() js.Value { return s.p }

// SelectionPanel returns a SelectionPanel JavaScript class.
func (ba *Babylon) SelectionPanel() *SelectionPanel {
	p := ba.ctx.Get("SelectionPanel")
	return SelectionPanelFromJSObject(p)
}

// SelectionPanelFromJSObject returns a wrapped SelectionPanel JavaScript class.
func SelectionPanelFromJSObject(p js.Value) *SelectionPanel {
	return &SelectionPanel{RectangleFromJSObject(p)}
}

// NewSelectionPanelOpts contains optional parameters for NewSelectionPanel.
type NewSelectionPanelOpts struct {
	Groups *SelectorGroup
}

// NewSelectionPanel returns a new SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel
func (ba *Babylon) NewSelectionPanel(name string, opts *NewSelectionPanelOpts) *SelectionPanel {
	if opts == nil {
		opts = &NewSelectionPanelOpts{}
	}

	p := ba.ctx.Get("SelectionPanel").New(name, opts.Groups.JSObject())
	return SelectionPanelFromJSObject(p)
}

// TODO: methods
