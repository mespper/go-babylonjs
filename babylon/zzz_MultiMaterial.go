// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiMaterial represents a babylon.js MultiMaterial.
// A multi-material is used to apply different materials to different parts of the same object without the need of
// separate meshes. This can be use to improve performances.
//
// See: http://doc.babylonjs.com/how_to/multi_materials
type MultiMaterial struct {
	*Material
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MultiMaterial) JSObject() js.Value { return m.p }

// MultiMaterial returns a MultiMaterial JavaScript class.
func (ba *Babylon) MultiMaterial() *MultiMaterial {
	p := ba.ctx.Get("MultiMaterial")
	return MultiMaterialFromJSObject(p, ba.ctx)
}

// MultiMaterialFromJSObject returns a wrapped MultiMaterial JavaScript class.
func MultiMaterialFromJSObject(p js.Value, ctx js.Value) *MultiMaterial {
	return &MultiMaterial{Material: MaterialFromJSObject(p, ctx), ctx: ctx}
}

// MultiMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func MultiMaterialArrayToJSArray(array []*MultiMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMultiMaterial returns a new MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial
func (ba *Babylon) NewMultiMaterial(name string, scene *Scene) *MultiMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("MultiMaterial").New(args...)
	return MultiMaterialFromJSObject(p, ba.ctx)
}

// MultiMaterialCloneOpts contains optional parameters for MultiMaterial.Clone.
type MultiMaterialCloneOpts struct {
	CloneChildren *bool
}

// Clone calls the Clone method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#clone
func (m *MultiMaterial) Clone(name string, opts *MultiMaterialCloneOpts) *MultiMaterial {
	if opts == nil {
		opts = &MultiMaterialCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, name)

	if opts.CloneChildren == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CloneChildren)
	}

	retVal := m.p.Call("clone", args...)
	return MultiMaterialFromJSObject(retVal, m.ctx)
}

// MultiMaterialDisposeOpts contains optional parameters for MultiMaterial.Dispose.
type MultiMaterialDisposeOpts struct {
	ForceDisposeEffect   *bool
	ForceDisposeTextures *bool
	ForceDisposeChildren *bool
}

// Dispose calls the Dispose method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#dispose
func (m *MultiMaterial) Dispose(opts *MultiMaterialDisposeOpts) {
	if opts == nil {
		opts = &MultiMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}
	if opts.ForceDisposeTextures == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeTextures)
	}
	if opts.ForceDisposeChildren == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeChildren)
	}

	m.p.Call("dispose", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#getactivetextures
func (m *MultiMaterial) GetActiveTextures() *BaseTexture {

	retVal := m.p.Call("getActiveTextures")
	return BaseTextureFromJSObject(retVal, m.ctx)
}

// GetChildren calls the GetChildren method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#getchildren
func (m *MultiMaterial) GetChildren() *Material {

	retVal := m.p.Call("getChildren")
	return MaterialFromJSObject(retVal, m.ctx)
}

// GetClassName calls the GetClassName method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#getclassname
func (m *MultiMaterial) GetClassName() string {

	retVal := m.p.Call("getClassName")
	return retVal.String()
}

// GetSubMaterial calls the GetSubMaterial method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#getsubmaterial
func (m *MultiMaterial) GetSubMaterial(index float64) *Material {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := m.p.Call("getSubMaterial", args...)
	return MaterialFromJSObject(retVal, m.ctx)
}

// MultiMaterialIsReadyForSubMeshOpts contains optional parameters for MultiMaterial.IsReadyForSubMesh.
type MultiMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#isreadyforsubmesh
func (m *MultiMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *BaseSubMesh, opts *MultiMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &MultiMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, mesh.JSObject())
	args = append(args, subMesh.JSObject())

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := m.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// ParseMultiMaterial calls the ParseMultiMaterial method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#parsemultimaterial
func (m *MultiMaterial) ParseMultiMaterial(parsedMultiMaterial interface{}, scene *Scene) *MultiMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, parsedMultiMaterial)
	args = append(args, scene.JSObject())

	retVal := m.p.Call("ParseMultiMaterial", args...)
	return MultiMaterialFromJSObject(retVal, m.ctx)
}

// Serialize calls the Serialize method on the MultiMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#serialize
func (m *MultiMaterial) Serialize() interface{} {

	retVal := m.p.Call("serialize")
	return retVal
}

// SubMaterials returns the SubMaterials property of class MultiMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#submaterials
func (m *MultiMaterial) SubMaterials() *Material {
	retVal := m.p.Get("subMaterials")
	return MaterialFromJSObject(retVal, m.ctx)
}

// SetSubMaterials sets the SubMaterials property of class MultiMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.multimaterial#submaterials
func (m *MultiMaterial) SetSubMaterials(subMaterials *Material) *MultiMaterial {
	m.p.Set("subMaterials", subMaterials.JSObject())
	return m
}
