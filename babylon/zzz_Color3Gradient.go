// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Color3Gradient represents a babylon.js Color3Gradient.
// Class used to store color 3 gradient
type Color3Gradient struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (c *Color3Gradient) JSObject() js.Value { return c.p }

// Color3Gradient returns a Color3Gradient JavaScript class.
func (ba *Babylon) Color3Gradient() *Color3Gradient {
	p := ba.ctx.Get("Color3Gradient")
	return Color3GradientFromJSObject(p)
}

// Color3GradientFromJSObject returns a wrapped Color3Gradient JavaScript class.
func Color3GradientFromJSObject(p js.Value) *Color3Gradient {
	return &Color3Gradient{p: p}
}

// TODO: methods
