// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// OBJFileLoader represents a babylon.js OBJFileLoader.
// OBJ file type loader.
// This is a babylon scene loader plugin.
type OBJFileLoader struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (o *OBJFileLoader) JSObject() js.Value { return o.p }

// OBJFileLoader returns a OBJFileLoader JavaScript class.
func (ba *Babylon) OBJFileLoader() *OBJFileLoader {
	p := ba.ctx.Get("OBJFileLoader")
	return OBJFileLoaderFromJSObject(p)
}

// OBJFileLoaderFromJSObject returns a wrapped OBJFileLoader JavaScript class.
func OBJFileLoaderFromJSObject(p js.Value) *OBJFileLoader {
	return &OBJFileLoader{p: p}
}

// NewOBJFileLoaderOpts contains optional parameters for NewOBJFileLoader.
type NewOBJFileLoaderOpts struct {
	MeshLoadOptions *JSValue
}

// NewOBJFileLoader returns a new OBJFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.objfileloader
func (ba *Babylon) NewOBJFileLoader(opts *NewOBJFileLoaderOpts) *OBJFileLoader {
	if opts == nil {
		opts = &NewOBJFileLoaderOpts{}
	}

	p := ba.ctx.Get("OBJFileLoader").New(opts.MeshLoadOptions.JSObject())
	return OBJFileLoaderFromJSObject(p)
}

// TODO: methods
