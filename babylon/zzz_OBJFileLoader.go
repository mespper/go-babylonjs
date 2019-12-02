// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// OBJFileLoader represents a babylon.js OBJFileLoader.
// OBJ file type loader.
// This is a babylon scene loader plugin.
type OBJFileLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (o *OBJFileLoader) JSObject() js.Value { return o.p }

// OBJFileLoader returns a OBJFileLoader JavaScript class.
func (ba *Babylon) OBJFileLoader() *OBJFileLoader {
	p := ba.ctx.Get("OBJFileLoader")
	return OBJFileLoaderFromJSObject(p, ba.ctx)
}

// OBJFileLoaderFromJSObject returns a wrapped OBJFileLoader JavaScript class.
func OBJFileLoaderFromJSObject(p js.Value, ctx js.Value) *OBJFileLoader {
	return &OBJFileLoader{p: p, ctx: ctx}
}

// OBJFileLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func OBJFileLoaderArrayToJSArray(array []*OBJFileLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewOBJFileLoaderOpts contains optional parameters for NewOBJFileLoader.
type NewOBJFileLoaderOpts struct {
	MeshLoadOptions js.Value
}

// NewOBJFileLoader returns a new OBJFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader
func (ba *Babylon) NewOBJFileLoader(opts *NewOBJFileLoaderOpts) *OBJFileLoader {
	if opts == nil {
		opts = &NewOBJFileLoaderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	args = append(args, opts.MeshLoadOptions)

	p := ba.ctx.Get("OBJFileLoader").New(args...)
	return OBJFileLoaderFromJSObject(p, ba.ctx)
}

// CanDirectLoad calls the CanDirectLoad method on the OBJFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#candirectload
func (o *OBJFileLoader) CanDirectLoad(data string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	retVal := o.p.Call("canDirectLoad", args...)
	return retVal.Bool()
}

// CreatePlugin calls the CreatePlugin method on the OBJFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#createplugin
func (o *OBJFileLoader) CreatePlugin() *ISceneLoaderPluginAsync {

	retVal := o.p.Call("createPlugin")
	return ISceneLoaderPluginAsyncFromJSObject(retVal, o.ctx)
}

// OBJFileLoaderImportMeshAsyncOpts contains optional parameters for OBJFileLoader.ImportMeshAsync.
type OBJFileLoaderImportMeshAsyncOpts struct {
	OnProgress func()
	FileName   *string
}

// ImportMeshAsync calls the ImportMeshAsync method on the OBJFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#importmeshasync
func (o *OBJFileLoader) ImportMeshAsync(meshesNames interface{}, scene *Scene, data interface{}, rootUrl string, opts *OBJFileLoaderImportMeshAsyncOpts) *Promise {
	if opts == nil {
		opts = &OBJFileLoaderImportMeshAsyncOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, meshesNames)
	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := o.p.Call("importMeshAsync", args...)
	return PromiseFromJSObject(retVal, o.ctx)
}

// OBJFileLoaderLoadAssetContainerAsyncOpts contains optional parameters for OBJFileLoader.LoadAssetContainerAsync.
type OBJFileLoaderLoadAssetContainerAsyncOpts struct {
	OnProgress func()
	FileName   *string
}

// LoadAssetContainerAsync calls the LoadAssetContainerAsync method on the OBJFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#loadassetcontainerasync
func (o *OBJFileLoader) LoadAssetContainerAsync(scene *Scene, data string, rootUrl string, opts *OBJFileLoaderLoadAssetContainerAsyncOpts) *Promise {
	if opts == nil {
		opts = &OBJFileLoaderLoadAssetContainerAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := o.p.Call("loadAssetContainerAsync", args...)
	return PromiseFromJSObject(retVal, o.ctx)
}

// OBJFileLoaderLoadAsyncOpts contains optional parameters for OBJFileLoader.LoadAsync.
type OBJFileLoaderLoadAsyncOpts struct {
	OnProgress func()
	FileName   *string
}

// LoadAsync calls the LoadAsync method on the OBJFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#loadasync
func (o *OBJFileLoader) LoadAsync(scene *Scene, data string, rootUrl string, opts *OBJFileLoaderLoadAsyncOpts) *Promise {
	if opts == nil {
		opts = &OBJFileLoaderLoadAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := o.p.Call("loadAsync", args...)
	return PromiseFromJSObject(retVal, o.ctx)
}

// COMPUTE_NORMALS returns the COMPUTE_NORMALS property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#compute_normals
func (o *OBJFileLoader) COMPUTE_NORMALS() bool {
	retVal := o.p.Get("COMPUTE_NORMALS")
	return retVal.Bool()
}

// SetCOMPUTE_NORMALS sets the COMPUTE_NORMALS property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#compute_normals
func (o *OBJFileLoader) SetCOMPUTE_NORMALS(COMPUTE_NORMALS bool) *OBJFileLoader {
	o.p.Set("COMPUTE_NORMALS", COMPUTE_NORMALS)
	return o
}

// Extensions returns the Extensions property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#extensions
func (o *OBJFileLoader) Extensions() string {
	retVal := o.p.Get("extensions")
	return retVal.String()
}

// SetExtensions sets the Extensions property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#extensions
func (o *OBJFileLoader) SetExtensions(extensions string) *OBJFileLoader {
	o.p.Set("extensions", extensions)
	return o
}

// IMPORT_VERTEX_COLORS returns the IMPORT_VERTEX_COLORS property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#import_vertex_colors
func (o *OBJFileLoader) IMPORT_VERTEX_COLORS() bool {
	retVal := o.p.Get("IMPORT_VERTEX_COLORS")
	return retVal.Bool()
}

// SetIMPORT_VERTEX_COLORS sets the IMPORT_VERTEX_COLORS property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#import_vertex_colors
func (o *OBJFileLoader) SetIMPORT_VERTEX_COLORS(IMPORT_VERTEX_COLORS bool) *OBJFileLoader {
	o.p.Set("IMPORT_VERTEX_COLORS", IMPORT_VERTEX_COLORS)
	return o
}

// INVERT_TEXTURE_Y returns the INVERT_TEXTURE_Y property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#invert_texture_y
func (o *OBJFileLoader) INVERT_TEXTURE_Y() bool {
	retVal := o.p.Get("INVERT_TEXTURE_Y")
	return retVal.Bool()
}

// SetINVERT_TEXTURE_Y sets the INVERT_TEXTURE_Y property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#invert_texture_y
func (o *OBJFileLoader) SetINVERT_TEXTURE_Y(INVERT_TEXTURE_Y bool) *OBJFileLoader {
	o.p.Set("INVERT_TEXTURE_Y", INVERT_TEXTURE_Y)
	return o
}

// INVERT_Y returns the INVERT_Y property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#invert_y
func (o *OBJFileLoader) INVERT_Y() bool {
	retVal := o.p.Get("INVERT_Y")
	return retVal.Bool()
}

// SetINVERT_Y sets the INVERT_Y property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#invert_y
func (o *OBJFileLoader) SetINVERT_Y(INVERT_Y bool) *OBJFileLoader {
	o.p.Set("INVERT_Y", INVERT_Y)
	return o
}

// MATERIAL_LOADING_FAILS_SILENTLY returns the MATERIAL_LOADING_FAILS_SILENTLY property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#material_loading_fails_silently
func (o *OBJFileLoader) MATERIAL_LOADING_FAILS_SILENTLY() bool {
	retVal := o.p.Get("MATERIAL_LOADING_FAILS_SILENTLY")
	return retVal.Bool()
}

// SetMATERIAL_LOADING_FAILS_SILENTLY sets the MATERIAL_LOADING_FAILS_SILENTLY property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#material_loading_fails_silently
func (o *OBJFileLoader) SetMATERIAL_LOADING_FAILS_SILENTLY(MATERIAL_LOADING_FAILS_SILENTLY bool) *OBJFileLoader {
	o.p.Set("MATERIAL_LOADING_FAILS_SILENTLY", MATERIAL_LOADING_FAILS_SILENTLY)
	return o
}

// Name returns the Name property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#name
func (o *OBJFileLoader) Name() string {
	retVal := o.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#name
func (o *OBJFileLoader) SetName(name string) *OBJFileLoader {
	o.p.Set("name", name)
	return o
}

// OPTIMIZE_WITH_UV returns the OPTIMIZE_WITH_UV property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#optimize_with_uv
func (o *OBJFileLoader) OPTIMIZE_WITH_UV() bool {
	retVal := o.p.Get("OPTIMIZE_WITH_UV")
	return retVal.Bool()
}

// SetOPTIMIZE_WITH_UV sets the OPTIMIZE_WITH_UV property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#optimize_with_uv
func (o *OBJFileLoader) SetOPTIMIZE_WITH_UV(OPTIMIZE_WITH_UV bool) *OBJFileLoader {
	o.p.Set("OPTIMIZE_WITH_UV", OPTIMIZE_WITH_UV)
	return o
}

// SKIP_MATERIALS returns the SKIP_MATERIALS property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#skip_materials
func (o *OBJFileLoader) SKIP_MATERIALS() bool {
	retVal := o.p.Get("SKIP_MATERIALS")
	return retVal.Bool()
}

// SetSKIP_MATERIALS sets the SKIP_MATERIALS property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#skip_materials
func (o *OBJFileLoader) SetSKIP_MATERIALS(SKIP_MATERIALS bool) *OBJFileLoader {
	o.p.Set("SKIP_MATERIALS", SKIP_MATERIALS)
	return o
}

// UV_SCALING returns the UV_SCALING property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#uv_scaling
func (o *OBJFileLoader) UV_SCALING() *Vector2 {
	retVal := o.p.Get("UV_SCALING")
	return Vector2FromJSObject(retVal, o.ctx)
}

// SetUV_SCALING sets the UV_SCALING property of class OBJFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader#uv_scaling
func (o *OBJFileLoader) SetUV_SCALING(UV_SCALING *Vector2) *OBJFileLoader {
	o.p.Set("UV_SCALING", UV_SCALING.JSObject())
	return o
}
