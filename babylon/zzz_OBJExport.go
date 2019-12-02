// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// OBJExport represents a babylon.js OBJExport.
// Class for generating OBJ data from a Babylon scene.
type OBJExport struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (o *OBJExport) JSObject() js.Value { return o.p }

// OBJExport returns a OBJExport JavaScript class.
func (ba *Babylon) OBJExport() *OBJExport {
	p := ba.ctx.Get("OBJExport")
	return OBJExportFromJSObject(p, ba.ctx)
}

// OBJExportFromJSObject returns a wrapped OBJExport JavaScript class.
func OBJExportFromJSObject(p js.Value, ctx js.Value) *OBJExport {
	return &OBJExport{p: p, ctx: ctx}
}

// OBJExportArrayToJSArray returns a JavaScript Array for the wrapped array.
func OBJExportArrayToJSArray(array []*OBJExport) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// MTL calls the MTL method on the OBJExport object.
//
// https://doc.babylonjs.com/api/classes/babylon.objexport#mtl
func (o *OBJExport) MTL(mesh *Mesh) string {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	retVal := o.p.Call("MTL", args...)
	return retVal.String()
}

// OBJExportOBJOpts contains optional parameters for OBJExport.OBJ.
type OBJExportOBJOpts struct {
	Materials      *bool
	Matlibname     *string
	Globalposition *bool
}

// OBJ calls the OBJ method on the OBJExport object.
//
// https://doc.babylonjs.com/api/classes/babylon.objexport#obj
func (o *OBJExport) OBJ(mesh *Mesh, opts *OBJExportOBJOpts) string {
	if opts == nil {
		opts = &OBJExportOBJOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, mesh.JSObject())

	if opts.Materials == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Materials)
	}
	if opts.Matlibname == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Matlibname)
	}
	if opts.Globalposition == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Globalposition)
	}

	retVal := o.p.Call("OBJ", args...)
	return retVal.String()
}
