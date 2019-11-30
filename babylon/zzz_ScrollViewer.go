// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScrollViewer represents a babylon.js ScrollViewer.
// Class used to hold a viewer window and sliders in a grid
type ScrollViewer struct{ *Rectangle }

// JSObject returns the underlying js.Value.
func (s *ScrollViewer) JSObject() js.Value { return s.p }

// ScrollViewer returns a ScrollViewer JavaScript class.
func (ba *Babylon) ScrollViewer() *ScrollViewer {
	p := ba.ctx.Get("ScrollViewer")
	return ScrollViewerFromJSObject(p)
}

// ScrollViewerFromJSObject returns a wrapped ScrollViewer JavaScript class.
func ScrollViewerFromJSObject(p js.Value) *ScrollViewer {
	return &ScrollViewer{RectangleFromJSObject(p)}
}

// NewScrollViewerOpts contains optional parameters for NewScrollViewer.
type NewScrollViewerOpts struct {
	Name *JSString

	IsImageBased *JSBool
}

// NewScrollViewer returns a new ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer
func (ba *Babylon) NewScrollViewer(opts *NewScrollViewerOpts) *ScrollViewer {
	if opts == nil {
		opts = &NewScrollViewerOpts{}
	}

	p := ba.ctx.Get("ScrollViewer").New(opts.Name, opts.IsImageBased)
	return ScrollViewerFromJSObject(p)
}

// TODO: methods
