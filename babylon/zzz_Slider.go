// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Slider represents a babylon.js Slider.
// Class used to create slider controls
type Slider struct{ *BaseSlider }

// JSObject returns the underlying js.Value.
func (s *Slider) JSObject() js.Value { return s.p }

// Slider returns a Slider JavaScript class.
func (ba *Babylon) Slider() *Slider {
	p := ba.ctx.Get("Slider")
	return SliderFromJSObject(p)
}

// SliderFromJSObject returns a wrapped Slider JavaScript class.
func SliderFromJSObject(p js.Value) *Slider {
	return &Slider{BaseSliderFromJSObject(p)}
}

// NewSliderOpts contains optional parameters for NewSlider.
type NewSliderOpts struct {
	Name *JSString
}

// NewSlider returns a new Slider object.
//
// https://doc.babylonjs.com/api/classes/babylon.slider
func (ba *Babylon) NewSlider(opts *NewSliderOpts) *Slider {
	if opts == nil {
		opts = &NewSliderOpts{}
	}

	p := ba.ctx.Get("Slider").New(opts.Name.JSObject())
	return SliderFromJSObject(p)
}

// TODO: methods
