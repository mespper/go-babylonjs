// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// XmlLoader represents a babylon.js XmlLoader.
// Class used to load GUI via XML.
type XmlLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (x *XmlLoader) JSObject() js.Value { return x.p }

// XmlLoader returns a XmlLoader JavaScript class.
func (ba *Babylon) XmlLoader() *XmlLoader {
	p := ba.ctx.Get("XmlLoader")
	return XmlLoaderFromJSObject(p, ba.ctx)
}

// XmlLoaderFromJSObject returns a wrapped XmlLoader JavaScript class.
func XmlLoaderFromJSObject(p js.Value, ctx js.Value) *XmlLoader {
	return &XmlLoader{p: p, ctx: ctx}
}

// XmlLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func XmlLoaderArrayToJSArray(array []*XmlLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewXmlLoaderOpts contains optional parameters for NewXmlLoader.
type NewXmlLoaderOpts struct {
	ParentClass js.Value
}

// NewXmlLoader returns a new XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.xmlloader
func (ba *Babylon) NewXmlLoader(opts *NewXmlLoaderOpts) *XmlLoader {
	if opts == nil {
		opts = &NewXmlLoaderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	args = append(args, opts.ParentClass)

	p := ba.ctx.Get("XmlLoader").New(args...)
	return XmlLoaderFromJSObject(p, ba.ctx)
}

// GetNodeById calls the GetNodeById method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.xmlloader#getnodebyid
func (x *XmlLoader) GetNodeById(id string) interface{} {

	args := make([]interface{}, 0, 1+0)

	args = append(args, id)

	retVal := x.p.Call("getNodeById", args...)
	return retVal
}

// GetNodes calls the GetNodes method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.xmlloader#getnodes
func (x *XmlLoader) GetNodes() interface{} {

	retVal := x.p.Call("getNodes")
	return retVal
}

// IsLoaded calls the IsLoaded method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.xmlloader#isloaded
func (x *XmlLoader) IsLoaded() bool {

	retVal := x.p.Call("isLoaded")
	return retVal.Bool()
}

// LoadLayout calls the LoadLayout method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.xmlloader#loadlayout
func (x *XmlLoader) LoadLayout(xmlFile interface{}, rootNode interface{}, callback interface{}) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, xmlFile)
	args = append(args, rootNode)
	args = append(args, callback)

	x.p.Call("loadLayout", args...)
}
