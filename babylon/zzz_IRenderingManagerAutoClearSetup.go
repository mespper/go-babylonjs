// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IRenderingManagerAutoClearSetup represents a babylon.js IRenderingManagerAutoClearSetup.
// Interface describing the different options available in the rendering manager
// regarding Auto Clear between groups.
type IRenderingManagerAutoClearSetup struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IRenderingManagerAutoClearSetup) JSObject() js.Value { return i.p }

// IRenderingManagerAutoClearSetup returns a IRenderingManagerAutoClearSetup JavaScript class.
func (ba *Babylon) IRenderingManagerAutoClearSetup() *IRenderingManagerAutoClearSetup {
	p := ba.ctx.Get("IRenderingManagerAutoClearSetup")
	return IRenderingManagerAutoClearSetupFromJSObject(p, ba.ctx)
}

// IRenderingManagerAutoClearSetupFromJSObject returns a wrapped IRenderingManagerAutoClearSetup JavaScript class.
func IRenderingManagerAutoClearSetupFromJSObject(p js.Value, ctx js.Value) *IRenderingManagerAutoClearSetup {
	return &IRenderingManagerAutoClearSetup{p: p, ctx: ctx}
}

// IRenderingManagerAutoClearSetupArrayToJSArray returns a JavaScript Array for the wrapped array.
func IRenderingManagerAutoClearSetupArrayToJSArray(array []*IRenderingManagerAutoClearSetup) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AutoClear returns the AutoClear property of class IRenderingManagerAutoClearSetup.
//
// https://doc.babylonjs.com/api/classes/babylon.irenderingmanagerautoclearsetup#autoclear
func (i *IRenderingManagerAutoClearSetup) AutoClear() bool {
	retVal := i.p.Get("autoClear")
	return retVal.Bool()
}

// SetAutoClear sets the AutoClear property of class IRenderingManagerAutoClearSetup.
//
// https://doc.babylonjs.com/api/classes/babylon.irenderingmanagerautoclearsetup#autoclear
func (i *IRenderingManagerAutoClearSetup) SetAutoClear(autoClear bool) *IRenderingManagerAutoClearSetup {
	i.p.Set("autoClear", autoClear)
	return i
}

// Depth returns the Depth property of class IRenderingManagerAutoClearSetup.
//
// https://doc.babylonjs.com/api/classes/babylon.irenderingmanagerautoclearsetup#depth
func (i *IRenderingManagerAutoClearSetup) Depth() bool {
	retVal := i.p.Get("depth")
	return retVal.Bool()
}

// SetDepth sets the Depth property of class IRenderingManagerAutoClearSetup.
//
// https://doc.babylonjs.com/api/classes/babylon.irenderingmanagerautoclearsetup#depth
func (i *IRenderingManagerAutoClearSetup) SetDepth(depth bool) *IRenderingManagerAutoClearSetup {
	i.p.Set("depth", depth)
	return i
}

// Stencil returns the Stencil property of class IRenderingManagerAutoClearSetup.
//
// https://doc.babylonjs.com/api/classes/babylon.irenderingmanagerautoclearsetup#stencil
func (i *IRenderingManagerAutoClearSetup) Stencil() bool {
	retVal := i.p.Get("stencil")
	return retVal.Bool()
}

// SetStencil sets the Stencil property of class IRenderingManagerAutoClearSetup.
//
// https://doc.babylonjs.com/api/classes/babylon.irenderingmanagerautoclearsetup#stencil
func (i *IRenderingManagerAutoClearSetup) SetStencil(stencil bool) *IRenderingManagerAutoClearSetup {
	i.p.Set("stencil", stencil)
	return i
}
