// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EquiRectangularCubeTextureAssetTask represents a babylon.js EquiRectangularCubeTextureAssetTask.
// Define a task used by AssetsManager to load Equirectangular cube textures
type EquiRectangularCubeTextureAssetTask struct {
	*AbstractAssetTask
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EquiRectangularCubeTextureAssetTask) JSObject() js.Value { return e.p }

// EquiRectangularCubeTextureAssetTask returns a EquiRectangularCubeTextureAssetTask JavaScript class.
func (ba *Babylon) EquiRectangularCubeTextureAssetTask() *EquiRectangularCubeTextureAssetTask {
	p := ba.ctx.Get("EquiRectangularCubeTextureAssetTask")
	return EquiRectangularCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// EquiRectangularCubeTextureAssetTaskFromJSObject returns a wrapped EquiRectangularCubeTextureAssetTask JavaScript class.
func EquiRectangularCubeTextureAssetTaskFromJSObject(p js.Value, ctx js.Value) *EquiRectangularCubeTextureAssetTask {
	return &EquiRectangularCubeTextureAssetTask{AbstractAssetTask: AbstractAssetTaskFromJSObject(p, ctx), ctx: ctx}
}

// EquiRectangularCubeTextureAssetTaskArrayToJSArray returns a JavaScript Array for the wrapped array.
func EquiRectangularCubeTextureAssetTaskArrayToJSArray(array []*EquiRectangularCubeTextureAssetTask) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewEquiRectangularCubeTextureAssetTaskOpts contains optional parameters for NewEquiRectangularCubeTextureAssetTask.
type NewEquiRectangularCubeTextureAssetTaskOpts struct {
	NoMipmap   *bool
	GammaSpace *bool
}

// NewEquiRectangularCubeTextureAssetTask returns a new EquiRectangularCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask
func (ba *Babylon) NewEquiRectangularCubeTextureAssetTask(name string, url string, size float64, opts *NewEquiRectangularCubeTextureAssetTaskOpts) *EquiRectangularCubeTextureAssetTask {
	if opts == nil {
		opts = &NewEquiRectangularCubeTextureAssetTaskOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, name)
	args = append(args, url)
	args = append(args, size)

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.GammaSpace == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GammaSpace)
	}

	p := ba.ctx.Get("EquiRectangularCubeTextureAssetTask").New(args...)
	return EquiRectangularCubeTextureAssetTaskFromJSObject(p, ba.ctx)
}

// RunTask calls the RunTask method on the EquiRectangularCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#runtask
func (e *EquiRectangularCubeTextureAssetTask) RunTask(scene *Scene, onSuccess func(), onError func()) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, scene.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onSuccess(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))

	e.p.Call("runTask", args...)
}

// GammaSpace returns the GammaSpace property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#gammaspace
func (e *EquiRectangularCubeTextureAssetTask) GammaSpace() bool {
	retVal := e.p.Get("gammaSpace")
	return retVal.Bool()
}

// SetGammaSpace sets the GammaSpace property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#gammaspace
func (e *EquiRectangularCubeTextureAssetTask) SetGammaSpace(gammaSpace bool) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("gammaSpace", gammaSpace)
	return e
}

// Name returns the Name property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#name
func (e *EquiRectangularCubeTextureAssetTask) Name() string {
	retVal := e.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#name
func (e *EquiRectangularCubeTextureAssetTask) SetName(name string) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("name", name)
	return e
}

// NoMipmap returns the NoMipmap property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#nomipmap
func (e *EquiRectangularCubeTextureAssetTask) NoMipmap() bool {
	retVal := e.p.Get("noMipmap")
	return retVal.Bool()
}

// SetNoMipmap sets the NoMipmap property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#nomipmap
func (e *EquiRectangularCubeTextureAssetTask) SetNoMipmap(noMipmap bool) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("noMipmap", noMipmap)
	return e
}

// OnError returns the OnError property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#onerror
func (e *EquiRectangularCubeTextureAssetTask) OnError() js.Value {
	retVal := e.p.Get("onError")
	return retVal
}

// SetOnError sets the OnError property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#onerror
func (e *EquiRectangularCubeTextureAssetTask) SetOnError(onError func()) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("onError", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))
	return e
}

// OnSuccess returns the OnSuccess property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#onsuccess
func (e *EquiRectangularCubeTextureAssetTask) OnSuccess() js.Value {
	retVal := e.p.Get("onSuccess")
	return retVal
}

// SetOnSuccess sets the OnSuccess property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#onsuccess
func (e *EquiRectangularCubeTextureAssetTask) SetOnSuccess(onSuccess func()) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("onSuccess", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onSuccess(); return nil }))
	return e
}

// Size returns the Size property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#size
func (e *EquiRectangularCubeTextureAssetTask) Size() float64 {
	retVal := e.p.Get("size")
	return retVal.Float()
}

// SetSize sets the Size property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#size
func (e *EquiRectangularCubeTextureAssetTask) SetSize(size float64) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("size", size)
	return e
}

// Texture returns the Texture property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#texture
func (e *EquiRectangularCubeTextureAssetTask) Texture() *EquiRectangularCubeTexture {
	retVal := e.p.Get("texture")
	return EquiRectangularCubeTextureFromJSObject(retVal, e.ctx)
}

// SetTexture sets the Texture property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#texture
func (e *EquiRectangularCubeTextureAssetTask) SetTexture(texture *EquiRectangularCubeTexture) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("texture", texture.JSObject())
	return e
}

// Url returns the Url property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#url
func (e *EquiRectangularCubeTextureAssetTask) Url() string {
	retVal := e.p.Get("url")
	return retVal.String()
}

// SetUrl sets the Url property of class EquiRectangularCubeTextureAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetextureassettask#url
func (e *EquiRectangularCubeTextureAssetTask) SetUrl(url string) *EquiRectangularCubeTextureAssetTask {
	e.p.Set("url", url)
	return e
}
