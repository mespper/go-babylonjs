// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IInternalTextureLoader represents a babylon.js IInternalTextureLoader.
// This represents the required contract to create a new type of texture loader.
type IInternalTextureLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IInternalTextureLoader) JSObject() js.Value { return i.p }

// IInternalTextureLoader returns a IInternalTextureLoader JavaScript class.
func (ba *Babylon) IInternalTextureLoader() *IInternalTextureLoader {
	p := ba.ctx.Get("IInternalTextureLoader")
	return IInternalTextureLoaderFromJSObject(p, ba.ctx)
}

// IInternalTextureLoaderFromJSObject returns a wrapped IInternalTextureLoader JavaScript class.
func IInternalTextureLoaderFromJSObject(p js.Value, ctx js.Value) *IInternalTextureLoader {
	return &IInternalTextureLoader{p: p, ctx: ctx}
}

// IInternalTextureLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func IInternalTextureLoaderArrayToJSArray(array []*IInternalTextureLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CanLoad calls the CanLoad method on the IInternalTextureLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.iinternaltextureloader#canload
func (i *IInternalTextureLoader) CanLoad(extension string, textureFormatInUse string, fallback *InternalTexture, isBase64 bool, isBuffer bool) bool {

	args := make([]interface{}, 0, 5+0)

	args = append(args, extension)
	args = append(args, textureFormatInUse)
	args = append(args, fallback.JSObject())
	args = append(args, isBase64)
	args = append(args, isBuffer)

	retVal := i.p.Call("canLoad", args...)
	return retVal.Bool()
}

// GetFallbackTextureUrl calls the GetFallbackTextureUrl method on the IInternalTextureLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.iinternaltextureloader#getfallbacktextureurl
func (i *IInternalTextureLoader) GetFallbackTextureUrl(rootUrl string, textureFormatInUse string) string {

	args := make([]interface{}, 0, 2+0)

	args = append(args, rootUrl)
	args = append(args, textureFormatInUse)

	retVal := i.p.Call("getFallbackTextureUrl", args...)
	return retVal.String()
}

// LoadCubeData calls the LoadCubeData method on the IInternalTextureLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.iinternaltextureloader#loadcubedata
func (i *IInternalTextureLoader) LoadCubeData(data string, texture *InternalTexture, createPolynomials bool, onLoad func(), onError func()) {

	args := make([]interface{}, 0, 5+0)

	args = append(args, data)
	args = append(args, texture.JSObject())
	args = append(args, createPolynomials)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onLoad(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))

	i.p.Call("loadCubeData", args...)
}

// LoadData calls the LoadData method on the IInternalTextureLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.iinternaltextureloader#loaddata
func (i *IInternalTextureLoader) LoadData(data js.Value, texture *InternalTexture, callback func()) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, data)
	args = append(args, texture.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	i.p.Call("loadData", args...)
}

// TransformUrl calls the TransformUrl method on the IInternalTextureLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.iinternaltextureloader#transformurl
func (i *IInternalTextureLoader) TransformUrl(rootUrl string, textureFormatInUse string) string {

	args := make([]interface{}, 0, 2+0)

	args = append(args, rootUrl)
	args = append(args, textureFormatInUse)

	retVal := i.p.Call("transformUrl", args...)
	return retVal.String()
}

// SupportCascades returns the SupportCascades property of class IInternalTextureLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.iinternaltextureloader#supportcascades
func (i *IInternalTextureLoader) SupportCascades() bool {
	retVal := i.p.Get("supportCascades")
	return retVal.Bool()
}

// SetSupportCascades sets the SupportCascades property of class IInternalTextureLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.iinternaltextureloader#supportcascades
func (i *IInternalTextureLoader) SetSupportCascades(supportCascades bool) *IInternalTextureLoader {
	i.p.Set("supportCascades", supportCascades)
	return i
}
