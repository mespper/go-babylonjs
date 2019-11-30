// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextBlock represents a babylon.js TextBlock.
// Class used to create text block control
type TextBlock struct{ *Control }

// JSObject returns the underlying js.Value.
func (t *TextBlock) JSObject() js.Value { return t.p }

// TextBlock returns a TextBlock JavaScript class.
func (ba *Babylon) TextBlock() *TextBlock {
	p := ba.ctx.Get("TextBlock")
	return TextBlockFromJSObject(p)
}

// TextBlockFromJSObject returns a wrapped TextBlock JavaScript class.
func TextBlockFromJSObject(p js.Value) *TextBlock {
	return &TextBlock{ControlFromJSObject(p)}
}

// NewTextBlockOpts contains optional parameters for NewTextBlock.
type NewTextBlockOpts struct {
	Name *JSString

	Text *JSString
}

// NewTextBlock returns a new TextBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textblock
func (ba *Babylon) NewTextBlock(opts *NewTextBlockOpts) *TextBlock {
	if opts == nil {
		opts = &NewTextBlockOpts{}
	}

	p := ba.ctx.Get("TextBlock").New(opts.Name.JSObject(), opts.Text.JSObject())
	return TextBlockFromJSObject(p)
}

// TODO: methods
