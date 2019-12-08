// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GenericController represents a babylon.js GenericController.
// Generic Controller
type GenericController struct {
	*WebVRController
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GenericController) JSObject() js.Value { return g.p }

// GenericController returns a GenericController JavaScript class.
func (ba *Babylon) GenericController() *GenericController {
	p := ba.ctx.Get("GenericController")
	return GenericControllerFromJSObject(p, ba.ctx)
}

// GenericControllerFromJSObject returns a wrapped GenericController JavaScript class.
func GenericControllerFromJSObject(p js.Value, ctx js.Value) *GenericController {
	return &GenericController{WebVRController: WebVRControllerFromJSObject(p, ctx), ctx: ctx}
}

// GenericControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func GenericControllerArrayToJSArray(array []*GenericController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGenericController returns a new GenericController object.
//
// https://doc.babylonjs.com/api/classes/babylon.genericcontroller
func (ba *Babylon) NewGenericController(vrGamepad interface{}) *GenericController {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vrGamepad)

	p := ba.ctx.Get("GenericController").New(args...)
	return GenericControllerFromJSObject(p, ba.ctx)
}

// GenericControllerInitControllerMeshOpts contains optional parameters for GenericController.InitControllerMesh.
type GenericControllerInitControllerMeshOpts struct {
	MeshLoaded JSFunc
}

// InitControllerMesh calls the InitControllerMesh method on the GenericController object.
//
// https://doc.babylonjs.com/api/classes/babylon.genericcontroller#initcontrollermesh
func (g *GenericController) InitControllerMesh(scene *Scene, opts *GenericControllerInitControllerMeshOpts) {
	if opts == nil {
		opts = &GenericControllerInitControllerMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.MeshLoaded == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.MeshLoaded) /* never freed! */)
	}

	g.p.Call("initControllerMesh", args...)
}

// MODEL_BASE_URL returns the MODEL_BASE_URL property of class GenericController.
//
// https://doc.babylonjs.com/api/classes/babylon.genericcontroller#model_base_url
func (g *GenericController) MODEL_BASE_URL() string {
	retVal := g.p.Get("MODEL_BASE_URL")
	return retVal.String()
}

// SetMODEL_BASE_URL sets the MODEL_BASE_URL property of class GenericController.
//
// https://doc.babylonjs.com/api/classes/babylon.genericcontroller#model_base_url
func (g *GenericController) SetMODEL_BASE_URL(MODEL_BASE_URL string) *GenericController {
	g.p.Set("MODEL_BASE_URL", MODEL_BASE_URL)
	return g
}

// MODEL_FILENAME returns the MODEL_FILENAME property of class GenericController.
//
// https://doc.babylonjs.com/api/classes/babylon.genericcontroller#model_filename
func (g *GenericController) MODEL_FILENAME() string {
	retVal := g.p.Get("MODEL_FILENAME")
	return retVal.String()
}

// SetMODEL_FILENAME sets the MODEL_FILENAME property of class GenericController.
//
// https://doc.babylonjs.com/api/classes/babylon.genericcontroller#model_filename
func (g *GenericController) SetMODEL_FILENAME(MODEL_FILENAME string) *GenericController {
	g.p.Set("MODEL_FILENAME", MODEL_FILENAME)
	return g
}
