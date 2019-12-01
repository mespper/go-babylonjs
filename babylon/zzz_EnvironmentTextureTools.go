// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EnvironmentTextureTools represents a babylon.js EnvironmentTextureTools.
// Sets of helpers addressing the serialization and deserialization of environment texture
// stored in a BabylonJS env file.
// Those files are usually stored as .env files.
type EnvironmentTextureTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EnvironmentTextureTools) JSObject() js.Value { return e.p }

// EnvironmentTextureTools returns a EnvironmentTextureTools JavaScript class.
func (ba *Babylon) EnvironmentTextureTools() *EnvironmentTextureTools {
	p := ba.ctx.Get("EnvironmentTextureTools")
	return EnvironmentTextureToolsFromJSObject(p, ba.ctx)
}

// EnvironmentTextureToolsFromJSObject returns a wrapped EnvironmentTextureTools JavaScript class.
func EnvironmentTextureToolsFromJSObject(p js.Value, ctx js.Value) *EnvironmentTextureTools {
	return &EnvironmentTextureTools{p: p, ctx: ctx}
}

// EnvironmentTextureToolsArrayToJSArray returns a JavaScript Array for the wrapped array.
func EnvironmentTextureToolsArrayToJSArray(array []*EnvironmentTextureTools) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CreateEnvTextureAsync calls the CreateEnvTextureAsync method on the EnvironmentTextureTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenttexturetools#createenvtextureasync
func (e *EnvironmentTextureTools) CreateEnvTextureAsync(texture *CubeTexture) *Promise {

	args := make([]interface{}, 0, 1+0)

	args = append(args, texture.JSObject())

	retVal := e.p.Call("CreateEnvTextureAsync", args...)
	return PromiseFromJSObject(retVal, e.ctx)
}

// CreateImageDataArrayBufferViews calls the CreateImageDataArrayBufferViews method on the EnvironmentTextureTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenttexturetools#createimagedataarraybufferviews
func (e *EnvironmentTextureTools) CreateImageDataArrayBufferViews(arrayBuffer interface{}, info js.Value) js.Value /* [][]ArrayBufferView */ {

	args := make([]interface{}, 0, 2+0)

	args = append(args, arrayBuffer)
	args = append(args, info)

	retVal := e.p.Call("CreateImageDataArrayBufferViews", args...)
	return retVal
}

// GetEnvInfo calls the GetEnvInfo method on the EnvironmentTextureTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenttexturetools#getenvinfo
func (e *EnvironmentTextureTools) GetEnvInfo(data js.Value) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	retVal := e.p.Call("GetEnvInfo", args...)
	return retVal
}

// UploadEnvLevelsAsync calls the UploadEnvLevelsAsync method on the EnvironmentTextureTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenttexturetools#uploadenvlevelsasync
func (e *EnvironmentTextureTools) UploadEnvLevelsAsync(texture *InternalTexture, arrayBuffer interface{}, info js.Value) *Promise {

	args := make([]interface{}, 0, 3+0)

	args = append(args, texture.JSObject())
	args = append(args, arrayBuffer)
	args = append(args, info)

	retVal := e.p.Call("UploadEnvLevelsAsync", args...)
	return PromiseFromJSObject(retVal, e.ctx)
}

// UploadEnvSpherical calls the UploadEnvSpherical method on the EnvironmentTextureTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenttexturetools#uploadenvspherical
func (e *EnvironmentTextureTools) UploadEnvSpherical(texture *InternalTexture, info js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, texture.JSObject())
	args = append(args, info)

	e.p.Call("UploadEnvSpherical", args...)
}

// UploadLevelsAsync calls the UploadLevelsAsync method on the EnvironmentTextureTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenttexturetools#uploadlevelsasync
func (e *EnvironmentTextureTools) UploadLevelsAsync(texture *InternalTexture, imageData js.Value) *Promise {

	args := make([]interface{}, 0, 2+0)

	args = append(args, texture.JSObject())
	args = append(args, imageData)

	retVal := e.p.Call("UploadLevelsAsync", args...)
	return PromiseFromJSObject(retVal, e.ctx)
}

/*

 */
