// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NormalMaterial represents a babylon.js NormalMaterial.
//
type NormalMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NormalMaterial) JSObject() js.Value { return n.p }

// NormalMaterial returns a NormalMaterial JavaScript class.
func (ba *Babylon) NormalMaterial() *NormalMaterial {
	p := ba.ctx.Get("NormalMaterial")
	return NormalMaterialFromJSObject(p, ba.ctx)
}

// NormalMaterialFromJSObject returns a wrapped NormalMaterial JavaScript class.
func NormalMaterialFromJSObject(p js.Value, ctx js.Value) *NormalMaterial {
	return &NormalMaterial{p: p, ctx: ctx}
}

// NormalMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func NormalMaterialArrayToJSArray(array []*NormalMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNormalMaterial returns a new NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial
func (ba *Babylon) NewNormalMaterial(name string, scene *Scene) *NormalMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("NormalMaterial").New(args...)
	return NormalMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#bindforsubmesh
func (n *NormalMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

	args := make([]interface{}, 0, 3+0)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	n.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#clone
func (n *NormalMaterial) Clone(name string) *NormalMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := n.p.Call("clone", args...)
	return NormalMaterialFromJSObject(retVal, n.ctx)
}

// NormalMaterialDisposeOpts contains optional parameters for NormalMaterial.Dispose.
type NormalMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#dispose
func (n *NormalMaterial) Dispose(opts *NormalMaterialDisposeOpts) {
	if opts == nil {
		opts = &NormalMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	n.p.Call("dispose", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#getactivetextures
func (n *NormalMaterial) GetActiveTextures() []*BaseTexture {

	retVal := n.p.Call("getActiveTextures")
	result := []*BaseTexture{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BaseTextureFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#getalphatesttexture
func (n *NormalMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := n.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, n.ctx)
}

// GetAnimatables calls the GetAnimatables method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#getanimatables
func (n *NormalMaterial) GetAnimatables() []*IAnimatable {

	retVal := n.p.Call("getAnimatables")
	result := []*IAnimatable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IAnimatableFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#getclassname
func (n *NormalMaterial) GetClassName() string {

	retVal := n.p.Call("getClassName")
	return retVal.String()
}

// HasTexture calls the HasTexture method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#hastexture
func (n *NormalMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	retVal := n.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// NormalMaterialIsReadyForSubMeshOpts contains optional parameters for NormalMaterial.IsReadyForSubMesh.
type NormalMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#isreadyforsubmesh
func (n *NormalMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *NormalMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &NormalMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := n.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#needalphablending
func (n *NormalMaterial) NeedAlphaBlending() bool {

	retVal := n.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaBlendingForMesh calls the NeedAlphaBlendingForMesh method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#needalphablendingformesh
func (n *NormalMaterial) NeedAlphaBlendingForMesh(mesh *AbstractMesh) bool {

	args := make([]interface{}, 0, 1+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	retVal := n.p.Call("needAlphaBlendingForMesh", args...)
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#needalphatesting
func (n *NormalMaterial) NeedAlphaTesting() bool {

	retVal := n.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Parse calls the Parse method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#parse
func (n *NormalMaterial) Parse(source JSObject, scene *Scene, rootUrl string) *NormalMaterial {

	args := make([]interface{}, 0, 3+0)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	retVal := n.p.Call("Parse", args...)
	return NormalMaterialFromJSObject(retVal, n.ctx)
}

// Serialize calls the Serialize method on the NormalMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#serialize
func (n *NormalMaterial) Serialize() js.Value {

	retVal := n.p.Call("serialize")
	return retVal
}

// DiffuseColor returns the DiffuseColor property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#diffusecolor
func (n *NormalMaterial) DiffuseColor() *Color3 {
	retVal := n.p.Get("diffuseColor")
	return Color3FromJSObject(retVal, n.ctx)
}

// SetDiffuseColor sets the DiffuseColor property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#diffusecolor
func (n *NormalMaterial) SetDiffuseColor(diffuseColor *Color3) *NormalMaterial {
	n.p.Set("diffuseColor", diffuseColor.JSObject())
	return n
}

// DiffuseTexture returns the DiffuseTexture property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#diffusetexture
func (n *NormalMaterial) DiffuseTexture() *BaseTexture {
	retVal := n.p.Get("diffuseTexture")
	return BaseTextureFromJSObject(retVal, n.ctx)
}

// SetDiffuseTexture sets the DiffuseTexture property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#diffusetexture
func (n *NormalMaterial) SetDiffuseTexture(diffuseTexture *BaseTexture) *NormalMaterial {
	n.p.Set("diffuseTexture", diffuseTexture.JSObject())
	return n
}

// DisableLighting returns the DisableLighting property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#disablelighting
func (n *NormalMaterial) DisableLighting() bool {
	retVal := n.p.Get("disableLighting")
	return retVal.Bool()
}

// SetDisableLighting sets the DisableLighting property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#disablelighting
func (n *NormalMaterial) SetDisableLighting(disableLighting bool) *NormalMaterial {
	n.p.Set("disableLighting", disableLighting)
	return n
}

// MaxSimultaneousLights returns the MaxSimultaneousLights property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#maxsimultaneouslights
func (n *NormalMaterial) MaxSimultaneousLights() float64 {
	retVal := n.p.Get("maxSimultaneousLights")
	return retVal.Float()
}

// SetMaxSimultaneousLights sets the MaxSimultaneousLights property of class NormalMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmaterial#maxsimultaneouslights
func (n *NormalMaterial) SetMaxSimultaneousLights(maxSimultaneousLights float64) *NormalMaterial {
	n.p.Set("maxSimultaneousLights", maxSimultaneousLights)
	return n
}
