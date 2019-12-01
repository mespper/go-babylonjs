// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IAnimationKey represents a babylon.js IAnimationKey.
// Defines an interface which represents an animation key frame
type IAnimationKey struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IAnimationKey) JSObject() js.Value { return i.p }

// IAnimationKey returns a IAnimationKey JavaScript class.
func (ba *Babylon) IAnimationKey() *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey")
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// IAnimationKeyFromJSObject returns a wrapped IAnimationKey JavaScript class.
func IAnimationKeyFromJSObject(p js.Value, ctx js.Value) *IAnimationKey {
	return &IAnimationKey{p: p, ctx: ctx}
}

// IAnimationKeyArrayToJSArray returns a JavaScript Array for the wrapped array.
func IAnimationKeyArrayToJSArray(array []*IAnimationKey) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Frame returns the Frame property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#frame
func (i *IAnimationKey) Frame(frame float64) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(frame)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// SetFrame sets the Frame property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#frame
func (i *IAnimationKey) SetFrame(frame float64) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(frame)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// InTangent returns the InTangent property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#intangent
func (i *IAnimationKey) InTangent(inTangent interface{}) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(inTangent)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// SetInTangent sets the InTangent property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#intangent
func (i *IAnimationKey) SetInTangent(inTangent interface{}) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(inTangent)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// Interpolation returns the Interpolation property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#interpolation
func (i *IAnimationKey) Interpolation(interpolation *AnimationKeyInterpolation) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(interpolation.JSObject())
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// SetInterpolation sets the Interpolation property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#interpolation
func (i *IAnimationKey) SetInterpolation(interpolation *AnimationKeyInterpolation) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(interpolation.JSObject())
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// OutTangent returns the OutTangent property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#outtangent
func (i *IAnimationKey) OutTangent(outTangent interface{}) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(outTangent)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// SetOutTangent sets the OutTangent property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#outtangent
func (i *IAnimationKey) SetOutTangent(outTangent interface{}) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(outTangent)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// Value returns the Value property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#value
func (i *IAnimationKey) Value(value interface{}) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(value)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class IAnimationKey.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationkey#value
func (i *IAnimationKey) SetValue(value interface{}) *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey").New(value)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

*/
