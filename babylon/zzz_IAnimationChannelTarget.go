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

// Node returns the Node property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#node
func (i *IAnimationChannelTarget) Node() float64 {
	retVal := i.p.Get("node")
	return retVal.Float()
}

// SetNode sets the Node property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#node
func (i *IAnimationChannelTarget) SetNode(node float64) *IAnimationChannelTarget {
	i.p.Set("node", node)
	return i
}

// Path returns the Path property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#path
func (i *IAnimationChannelTarget) Path() js.Value {
	retVal := i.p.Get("path")
	return retVal
}

// SetPath sets the Path property of class IAnimationChannelTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.ianimationchanneltarget#path
func (i *IAnimationChannelTarget) SetPath(path js.Value) *IAnimationChannelTarget {
	i.p.Set("path", path)
	return i
}
