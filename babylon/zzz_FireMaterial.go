// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FireMaterial represents a babylon.js FireMaterial.
//
type FireMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FireMaterial) JSObject() js.Value { return f.p }

// FireMaterial returns a FireMaterial JavaScript class.
func (ba *Babylon) FireMaterial() *FireMaterial {
	p := ba.ctx.Get("FireMaterial")
	return FireMaterialFromJSObject(p, ba.ctx)
}

// FireMaterialFromJSObject returns a wrapped FireMaterial JavaScript class.
func FireMaterialFromJSObject(p js.Value, ctx js.Value) *FireMaterial {
	return &FireMaterial{p: p, ctx: ctx}
}

// FireMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func FireMaterialArrayToJSArray(array []*FireMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewFireMaterial returns a new FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial
func (ba *Babylon) NewFireMaterial(name string, scene *Scene) *FireMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("FireMaterial").New(args...)
	return FireMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#bindforsubmesh
func (f *FireMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

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

	f.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#clone
func (f *FireMaterial) Clone(name string) *FireMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := f.p.Call("clone", args...)
	return FireMaterialFromJSObject(retVal, f.ctx)
}

// FireMaterialDisposeOpts contains optional parameters for FireMaterial.Dispose.
type FireMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#dispose
func (f *FireMaterial) Dispose(opts *FireMaterialDisposeOpts) {
	if opts == nil {
		opts = &FireMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	f.p.Call("dispose", args...)
}

// GetActiveTextures calls the GetActiveTextures method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#getactivetextures
func (f *FireMaterial) GetActiveTextures() []*BaseTexture {

	retVal := f.p.Call("getActiveTextures")
	result := []*BaseTexture{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BaseTextureFromJSObject(retVal.Index(ri), f.ctx))
	}
	return result
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#getalphatesttexture
func (f *FireMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := f.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, f.ctx)
}

// GetAnimatables calls the GetAnimatables method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#getanimatables
func (f *FireMaterial) GetAnimatables() []*IAnimatable {

	retVal := f.p.Call("getAnimatables")
	result := []*IAnimatable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IAnimatableFromJSObject(retVal.Index(ri), f.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#getclassname
func (f *FireMaterial) GetClassName() string {

	retVal := f.p.Call("getClassName")
	return retVal.String()
}

// HasTexture calls the HasTexture method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#hastexture
func (f *FireMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	if texture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, texture.JSObject())
	}

	retVal := f.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// FireMaterialIsReadyForSubMeshOpts contains optional parameters for FireMaterial.IsReadyForSubMesh.
type FireMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#isreadyforsubmesh
func (f *FireMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *FireMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &FireMaterialIsReadyForSubMeshOpts{}
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

	retVal := f.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#needalphablending
func (f *FireMaterial) NeedAlphaBlending() bool {

	retVal := f.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#needalphatesting
func (f *FireMaterial) NeedAlphaTesting() bool {

	retVal := f.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Parse calls the Parse method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#parse
func (f *FireMaterial) Parse(source JSObject, scene *Scene, rootUrl string) *FireMaterial {

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

	retVal := f.p.Call("Parse", args...)
	return FireMaterialFromJSObject(retVal, f.ctx)
}

// Serialize calls the Serialize method on the FireMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#serialize
func (f *FireMaterial) Serialize() js.Value {

	retVal := f.p.Call("serialize")
	return retVal
}

// DiffuseColor returns the DiffuseColor property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#diffusecolor
func (f *FireMaterial) DiffuseColor() *Color3 {
	retVal := f.p.Get("diffuseColor")
	return Color3FromJSObject(retVal, f.ctx)
}

// SetDiffuseColor sets the DiffuseColor property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#diffusecolor
func (f *FireMaterial) SetDiffuseColor(diffuseColor *Color3) *FireMaterial {
	f.p.Set("diffuseColor", diffuseColor.JSObject())
	return f
}

// DiffuseTexture returns the DiffuseTexture property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#diffusetexture
func (f *FireMaterial) DiffuseTexture() *BaseTexture {
	retVal := f.p.Get("diffuseTexture")
	return BaseTextureFromJSObject(retVal, f.ctx)
}

// SetDiffuseTexture sets the DiffuseTexture property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#diffusetexture
func (f *FireMaterial) SetDiffuseTexture(diffuseTexture *BaseTexture) *FireMaterial {
	f.p.Set("diffuseTexture", diffuseTexture.JSObject())
	return f
}

// DistortionTexture returns the DistortionTexture property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#distortiontexture
func (f *FireMaterial) DistortionTexture() *BaseTexture {
	retVal := f.p.Get("distortionTexture")
	return BaseTextureFromJSObject(retVal, f.ctx)
}

// SetDistortionTexture sets the DistortionTexture property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#distortiontexture
func (f *FireMaterial) SetDistortionTexture(distortionTexture *BaseTexture) *FireMaterial {
	f.p.Set("distortionTexture", distortionTexture.JSObject())
	return f
}

// OpacityTexture returns the OpacityTexture property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#opacitytexture
func (f *FireMaterial) OpacityTexture() *BaseTexture {
	retVal := f.p.Get("opacityTexture")
	return BaseTextureFromJSObject(retVal, f.ctx)
}

// SetOpacityTexture sets the OpacityTexture property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#opacitytexture
func (f *FireMaterial) SetOpacityTexture(opacityTexture *BaseTexture) *FireMaterial {
	f.p.Set("opacityTexture", opacityTexture.JSObject())
	return f
}

// Speed returns the Speed property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#speed
func (f *FireMaterial) Speed() float64 {
	retVal := f.p.Get("speed")
	return retVal.Float()
}

// SetSpeed sets the Speed property of class FireMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.firematerial#speed
func (f *FireMaterial) SetSpeed(speed float64) *FireMaterial {
	f.p.Set("speed", speed)
	return f
}
