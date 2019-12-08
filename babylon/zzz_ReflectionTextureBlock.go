// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ReflectionTextureBlock represents a babylon.js ReflectionTextureBlock.
// Block used to read a reflection texture from a sampler
type ReflectionTextureBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *ReflectionTextureBlock) JSObject() js.Value { return r.p }

// ReflectionTextureBlock returns a ReflectionTextureBlock JavaScript class.
func (ba *Babylon) ReflectionTextureBlock() *ReflectionTextureBlock {
	p := ba.ctx.Get("ReflectionTextureBlock")
	return ReflectionTextureBlockFromJSObject(p, ba.ctx)
}

// ReflectionTextureBlockFromJSObject returns a wrapped ReflectionTextureBlock JavaScript class.
func ReflectionTextureBlockFromJSObject(p js.Value, ctx js.Value) *ReflectionTextureBlock {
	return &ReflectionTextureBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// ReflectionTextureBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func ReflectionTextureBlockArrayToJSArray(array []*ReflectionTextureBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewReflectionTextureBlock returns a new ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock
func (ba *Babylon) NewReflectionTextureBlock(name string) *ReflectionTextureBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ReflectionTextureBlock").New(args...)
	return ReflectionTextureBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#autoconfigure
func (r *ReflectionTextureBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	if material == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, material.JSObject())
	}

	r.p.Call("autoConfigure", args...)
}

// ReflectionTextureBlockBindOpts contains optional parameters for ReflectionTextureBlock.Bind.
type ReflectionTextureBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#bind
func (r *ReflectionTextureBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *ReflectionTextureBlockBindOpts) {
	if opts == nil {
		opts = &ReflectionTextureBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	if nodeMaterial == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, nodeMaterial.JSObject())
	}

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	r.p.Call("bind", args...)
}

// GetClassName calls the GetClassName method on the ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#getclassname
func (r *ReflectionTextureBlock) GetClassName() string {

	retVal := r.p.Call("getClassName")
	return retVal.String()
}

// IsReady calls the IsReady method on the ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#isready
func (r *ReflectionTextureBlock) IsReady() bool {

	retVal := r.p.Call("isReady")
	return retVal.Bool()
}

// PrepareDefines calls the PrepareDefines method on the ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#preparedefines
func (r *ReflectionTextureBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value) {

	args := make([]interface{}, 0, 3+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if nodeMaterial == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, nodeMaterial.JSObject())
	}

	args = append(args, defines)

	r.p.Call("prepareDefines", args...)
}

// Serialize calls the Serialize method on the ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#serialize
func (r *ReflectionTextureBlock) Serialize() js.Value {

	retVal := r.p.Call("serialize")
	return retVal
}

// _deserialize calls the _deserialize method on the ReflectionTextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#_deserialize
func (r *ReflectionTextureBlock) _deserialize(serializationObject JSObject, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	r.p.Call("_deserialize", args...)
}

// B returns the B property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#b
func (r *ReflectionTextureBlock) B() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("b")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetB sets the B property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#b
func (r *ReflectionTextureBlock) SetB(b *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("b", b.JSObject())
	return r
}

// CameraPosition returns the CameraPosition property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#cameraposition
func (r *ReflectionTextureBlock) CameraPosition() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("cameraPosition")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetCameraPosition sets the CameraPosition property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#cameraposition
func (r *ReflectionTextureBlock) SetCameraPosition(cameraPosition *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("cameraPosition", cameraPosition.JSObject())
	return r
}

// G returns the G property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#g
func (r *ReflectionTextureBlock) G() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("g")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetG sets the G property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#g
func (r *ReflectionTextureBlock) SetG(g *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("g", g.JSObject())
	return r
}

// Position returns the Position property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#position
func (r *ReflectionTextureBlock) Position() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("position")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetPosition sets the Position property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#position
func (r *ReflectionTextureBlock) SetPosition(position *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("position", position.JSObject())
	return r
}

// R returns the R property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#r
func (r *ReflectionTextureBlock) R() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("r")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetR sets the R property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#r
func (r *ReflectionTextureBlock) SetR(rr *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("r", rr.JSObject())
	return r
}

// Rgb returns the Rgb property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#rgb
func (r *ReflectionTextureBlock) Rgb() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("rgb")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetRgb sets the Rgb property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#rgb
func (r *ReflectionTextureBlock) SetRgb(rgb *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("rgb", rgb.JSObject())
	return r
}

// Texture returns the Texture property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#texture
func (r *ReflectionTextureBlock) Texture() *BaseTexture {
	retVal := r.p.Get("texture")
	return BaseTextureFromJSObject(retVal, r.ctx)
}

// SetTexture sets the Texture property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#texture
func (r *ReflectionTextureBlock) SetTexture(texture *BaseTexture) *ReflectionTextureBlock {
	r.p.Set("texture", texture.JSObject())
	return r
}

// View returns the View property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#view
func (r *ReflectionTextureBlock) View() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("view")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetView sets the View property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#view
func (r *ReflectionTextureBlock) SetView(view *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("view", view.JSObject())
	return r
}

// World returns the World property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#world
func (r *ReflectionTextureBlock) World() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("world")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetWorld sets the World property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#world
func (r *ReflectionTextureBlock) SetWorld(world *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("world", world.JSObject())
	return r
}

// WorldNormal returns the WorldNormal property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#worldnormal
func (r *ReflectionTextureBlock) WorldNormal() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("worldNormal")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetWorldNormal sets the WorldNormal property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#worldnormal
func (r *ReflectionTextureBlock) SetWorldNormal(worldNormal *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("worldNormal", worldNormal.JSObject())
	return r
}

// WorldPosition returns the WorldPosition property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#worldposition
func (r *ReflectionTextureBlock) WorldPosition() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("worldPosition")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetWorldPosition sets the WorldPosition property of class ReflectionTextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectiontextureblock#worldposition
func (r *ReflectionTextureBlock) SetWorldPosition(worldPosition *NodeMaterialConnectionPoint) *ReflectionTextureBlock {
	r.p.Set("worldPosition", worldPosition.JSObject())
	return r
}
