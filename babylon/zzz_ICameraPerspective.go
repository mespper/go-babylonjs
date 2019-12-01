// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ICameraPerspective represents a babylon.js ICameraPerspective.
// A perspective camera containing properties to create a perspective projection matrix
type ICameraPerspective struct {
	*IProperty
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ICameraPerspective) JSObject() js.Value { return i.p }

// ICameraPerspective returns a ICameraPerspective JavaScript class.
func (ba *Babylon) ICameraPerspective() *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective")
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// ICameraPerspectiveFromJSObject returns a wrapped ICameraPerspective JavaScript class.
func ICameraPerspectiveFromJSObject(p js.Value, ctx js.Value) *ICameraPerspective {
	return &ICameraPerspective{IProperty: IPropertyFromJSObject(p, ctx), ctx: ctx}
}

// ICameraPerspectiveArrayToJSArray returns a JavaScript Array for the wrapped array.
func ICameraPerspectiveArrayToJSArray(array []*ICameraPerspective) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// AspectRatio returns the AspectRatio property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#aspectratio
func (i *ICameraPerspective) AspectRatio(aspectRatio float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(aspectRatio)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#aspectratio
func (i *ICameraPerspective) SetAspectRatio(aspectRatio float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(aspectRatio)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// Extensions returns the Extensions property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#extensions
func (i *ICameraPerspective) Extensions(extensions js.Value) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(extensions)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// SetExtensions sets the Extensions property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#extensions
func (i *ICameraPerspective) SetExtensions(extensions js.Value) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(extensions)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// Extras returns the Extras property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#extras
func (i *ICameraPerspective) Extras(extras interface{}) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(extras)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// SetExtras sets the Extras property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#extras
func (i *ICameraPerspective) SetExtras(extras interface{}) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(extras)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// Yfov returns the Yfov property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#yfov
func (i *ICameraPerspective) Yfov(yfov float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(yfov)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// SetYfov sets the Yfov property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#yfov
func (i *ICameraPerspective) SetYfov(yfov float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(yfov)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// Zfar returns the Zfar property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#zfar
func (i *ICameraPerspective) Zfar(zfar float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(zfar)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// SetZfar sets the Zfar property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#zfar
func (i *ICameraPerspective) SetZfar(zfar float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(zfar)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// Znear returns the Znear property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#znear
func (i *ICameraPerspective) Znear(znear float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(znear)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

// SetZnear sets the Znear property of class ICameraPerspective.
//
// https://doc.babylonjs.com/api/classes/babylon.icameraperspective#znear
func (i *ICameraPerspective) SetZnear(znear float64) *ICameraPerspective {
	p := ba.ctx.Get("ICameraPerspective").New(znear)
	return ICameraPerspectiveFromJSObject(p, ba.ctx)
}

*/