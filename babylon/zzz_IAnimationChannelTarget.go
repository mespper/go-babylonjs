// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IAnimationChannelTarget represents a babylon.js IAnimationChannelTarget.
// The index of the node and TRS property that an animation channel targets
type IAnimationChannelTarget struct {
	*IProperty
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IAnimationChannelTarget) JSObject() js.Value { return i.p }

// IAnimationChannelTarget returns a IAnimationChannelTarget JavaScript class.
func (ba *Babylon) IAnimationChannelTarget() *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget")
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// IAnimationChannelTargetFromJSObject returns a wrapped IAnimationChannelTarget JavaScript class.
func IAnimationChannelTargetFromJSObject(p js.Value, ctx js.Value) *IAnimationChannelTarget {
	return &IAnimationChannelTarget{IProperty: IPropertyFromJSObject(p, ctx), ctx: ctx}
}

// IAnimationChannelTargetArrayToJSArray returns a JavaScript Array for the wrapped array.
func IAnimationChannelTargetArrayToJSArray(array []*IAnimationChannelTarget) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Extensions returns the Extensions property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#extensions
func (i *IAnimationChannelTarget) Extensions(extensions js.Value) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(extensions)
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// SetExtensions sets the Extensions property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#extensions
func (i *IAnimationChannelTarget) SetExtensions(extensions js.Value) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(extensions)
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// Extras returns the Extras property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#extras
func (i *IAnimationChannelTarget) Extras(extras interface{}) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(extras)
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// SetExtras sets the Extras property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#extras
func (i *IAnimationChannelTarget) SetExtras(extras interface{}) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(extras)
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// Node returns the Node property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#node
func (i *IAnimationChannelTarget) Node(node float64) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(node)
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// SetNode sets the Node property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#node
func (i *IAnimationChannelTarget) SetNode(node float64) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(node)
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// Path returns the Path property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#path
func (i *IAnimationChannelTarget) Path(path *AnimationChannelTargetPath) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(path.JSObject())
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

// SetPath sets the Path property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#path
func (i *IAnimationChannelTarget) SetPath(path *AnimationChannelTargetPath) *IAnimationChannelTarget {
	p := ba.ctx.Get("IAnimationChannelTarget").New(path.JSObject())
	return IAnimationChannelTargetFromJSObject(p, ba.ctx)
}

*/