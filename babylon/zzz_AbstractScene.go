// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AbstractScene represents a babylon.js AbstractScene.
// Base class of the scene acting as a container for the different elements composing a scene.
// This class is dynamically extended by the different components of the scene increasing
// flexibility and reducing coupling
type AbstractScene struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AbstractScene) JSObject() js.Value { return a.p }

// AbstractScene returns a AbstractScene JavaScript class.
func (ba *Babylon) AbstractScene() *AbstractScene {
	p := ba.ctx.Get("AbstractScene")
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// AbstractSceneFromJSObject returns a wrapped AbstractScene JavaScript class.
func AbstractSceneFromJSObject(p js.Value, ctx js.Value) *AbstractScene {
	return &AbstractScene{p: p, ctx: ctx}
}

// AbstractSceneArrayToJSArray returns a JavaScript Array for the wrapped array.
func AbstractSceneArrayToJSArray(array []*AbstractScene) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddEffectLayer calls the AddEffectLayer method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#addeffectlayer
func (a *AbstractScene) AddEffectLayer(newEffectLayer *EffectLayer) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newEffectLayer.JSObject())

	a.p.Call("addEffectLayer", args...)
}

// AddIndividualParser calls the AddIndividualParser method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#addindividualparser
func (a *AbstractScene) AddIndividualParser(name string, parser js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, parser)

	a.p.Call("AddIndividualParser", args...)
}

// AddLensFlareSystem calls the AddLensFlareSystem method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#addlensflaresystem
func (a *AbstractScene) AddLensFlareSystem(newLensFlareSystem *LensFlareSystem) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newLensFlareSystem.JSObject())

	a.p.Call("addLensFlareSystem", args...)
}

// AddParser calls the AddParser method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#addparser
func (a *AbstractScene) AddParser(name string, parser js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, parser)

	a.p.Call("AddParser", args...)
}

// AddReflectionProbe calls the AddReflectionProbe method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#addreflectionprobe
func (a *AbstractScene) AddReflectionProbe(newReflectionProbe *ReflectionProbe) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newReflectionProbe.JSObject())

	a.p.Call("addReflectionProbe", args...)
}

// GetGlowLayerByName calls the GetGlowLayerByName method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#getglowlayerbyname
func (a *AbstractScene) GetGlowLayerByName(name string) *GlowLayer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := a.p.Call("getGlowLayerByName", args...)
	return GlowLayerFromJSObject(retVal, a.ctx)
}

// GetHighlightLayerByName calls the GetHighlightLayerByName method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#gethighlightlayerbyname
func (a *AbstractScene) GetHighlightLayerByName(name string) *HighlightLayer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := a.p.Call("getHighlightLayerByName", args...)
	return HighlightLayerFromJSObject(retVal, a.ctx)
}

// GetIndividualParser calls the GetIndividualParser method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#getindividualparser
func (a *AbstractScene) GetIndividualParser(name string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := a.p.Call("GetIndividualParser", args...)
	return retVal
}

// GetLensFlareSystemByID calls the GetLensFlareSystemByID method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#getlensflaresystembyid
func (a *AbstractScene) GetLensFlareSystemByID(id string) *LensFlareSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, id)

	retVal := a.p.Call("getLensFlareSystemByID", args...)
	return LensFlareSystemFromJSObject(retVal, a.ctx)
}

// GetLensFlareSystemByName calls the GetLensFlareSystemByName method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#getlensflaresystembyname
func (a *AbstractScene) GetLensFlareSystemByName(name string) *LensFlareSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := a.p.Call("getLensFlareSystemByName", args...)
	return LensFlareSystemFromJSObject(retVal, a.ctx)
}

// GetParser calls the GetParser method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#getparser
func (a *AbstractScene) GetParser(name string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := a.p.Call("GetParser", args...)
	return retVal
}

// Parse calls the Parse method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#parse
func (a *AbstractScene) Parse(jsonData interface{}, scene *Scene, container *AssetContainer, rootUrl string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, jsonData)
	args = append(args, scene.JSObject())
	args = append(args, container.JSObject())
	args = append(args, rootUrl)

	a.p.Call("Parse", args...)
}

// RemoveEffectLayer calls the RemoveEffectLayer method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#removeeffectlayer
func (a *AbstractScene) RemoveEffectLayer(toRemove *EffectLayer) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, toRemove.JSObject())

	retVal := a.p.Call("removeEffectLayer", args...)
	return retVal.Float()
}

// RemoveLensFlareSystem calls the RemoveLensFlareSystem method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#removelensflaresystem
func (a *AbstractScene) RemoveLensFlareSystem(toRemove *LensFlareSystem) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, toRemove.JSObject())

	retVal := a.p.Call("removeLensFlareSystem", args...)
	return retVal.Float()
}

// RemoveReflectionProbe calls the RemoveReflectionProbe method on the AbstractScene object.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#removereflectionprobe
func (a *AbstractScene) RemoveReflectionProbe(toRemove *ReflectionProbe) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, toRemove.JSObject())

	retVal := a.p.Call("removeReflectionProbe", args...)
	return retVal.Float()
}

/*

// ActionManagers returns the ActionManagers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#actionmanagers
func (a *AbstractScene) ActionManagers(actionManagers *AbstractActionManager) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(actionManagers.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetActionManagers sets the ActionManagers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#actionmanagers
func (a *AbstractScene) SetActionManagers(actionManagers *AbstractActionManager) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(actionManagers.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// AnimationGroups returns the AnimationGroups property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#animationgroups
func (a *AbstractScene) AnimationGroups(animationGroups *AnimationGroup) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(animationGroups.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetAnimationGroups sets the AnimationGroups property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#animationgroups
func (a *AbstractScene) SetAnimationGroups(animationGroups *AnimationGroup) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(animationGroups.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#animations
func (a *AbstractScene) Animations(animations *Animation) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(animations.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#animations
func (a *AbstractScene) SetAnimations(animations *Animation) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(animations.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Cameras returns the Cameras property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#cameras
func (a *AbstractScene) Cameras(cameras *Camera) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(cameras.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetCameras sets the Cameras property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#cameras
func (a *AbstractScene) SetCameras(cameras *Camera) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(cameras.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// EffectLayers returns the EffectLayers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#effectlayers
func (a *AbstractScene) EffectLayers(effectLayers []*EffectLayer) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(effectLayers)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetEffectLayers sets the EffectLayers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#effectlayers
func (a *AbstractScene) SetEffectLayers(effectLayers []*EffectLayer) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(effectLayers)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// EnvironmentTexture returns the EnvironmentTexture property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#environmenttexture
func (a *AbstractScene) EnvironmentTexture(environmentTexture *BaseTexture) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(environmentTexture.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetEnvironmentTexture sets the EnvironmentTexture property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#environmenttexture
func (a *AbstractScene) SetEnvironmentTexture(environmentTexture *BaseTexture) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(environmentTexture.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Geometries returns the Geometries property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#geometries
func (a *AbstractScene) Geometries(geometries *Geometry) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(geometries.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetGeometries sets the Geometries property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#geometries
func (a *AbstractScene) SetGeometries(geometries *Geometry) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(geometries.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Layers returns the Layers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#layers
func (a *AbstractScene) Layers(layers []*Layer) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(layers)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetLayers sets the Layers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#layers
func (a *AbstractScene) SetLayers(layers []*Layer) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(layers)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// LensFlareSystems returns the LensFlareSystems property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#lensflaresystems
func (a *AbstractScene) LensFlareSystems(lensFlareSystems []*LensFlareSystem) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(lensFlareSystems)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetLensFlareSystems sets the LensFlareSystems property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#lensflaresystems
func (a *AbstractScene) SetLensFlareSystems(lensFlareSystems []*LensFlareSystem) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(lensFlareSystems)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Lights returns the Lights property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#lights
func (a *AbstractScene) Lights(lights *Light) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(lights.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetLights sets the Lights property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#lights
func (a *AbstractScene) SetLights(lights *Light) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(lights.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Materials returns the Materials property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#materials
func (a *AbstractScene) Materials(materials *Material) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(materials.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetMaterials sets the Materials property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#materials
func (a *AbstractScene) SetMaterials(materials *Material) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(materials.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Meshes returns the Meshes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#meshes
func (a *AbstractScene) Meshes(meshes *AbstractMesh) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(meshes.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetMeshes sets the Meshes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#meshes
func (a *AbstractScene) SetMeshes(meshes *AbstractMesh) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(meshes.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// MorphTargetManagers returns the MorphTargetManagers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#morphtargetmanagers
func (a *AbstractScene) MorphTargetManagers(morphTargetManagers *MorphTargetManager) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(morphTargetManagers.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetMorphTargetManagers sets the MorphTargetManagers property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#morphtargetmanagers
func (a *AbstractScene) SetMorphTargetManagers(morphTargetManagers *MorphTargetManager) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(morphTargetManagers.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// MultiMaterials returns the MultiMaterials property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#multimaterials
func (a *AbstractScene) MultiMaterials(multiMaterials *MultiMaterial) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(multiMaterials.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetMultiMaterials sets the MultiMaterials property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#multimaterials
func (a *AbstractScene) SetMultiMaterials(multiMaterials *MultiMaterial) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(multiMaterials.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// ParticleSystems returns the ParticleSystems property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#particlesystems
func (a *AbstractScene) ParticleSystems(particleSystems *IParticleSystem) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(particleSystems.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetParticleSystems sets the ParticleSystems property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#particlesystems
func (a *AbstractScene) SetParticleSystems(particleSystems *IParticleSystem) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(particleSystems.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// ProceduralTextures returns the ProceduralTextures property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#proceduraltextures
func (a *AbstractScene) ProceduralTextures(proceduralTextures []*ProceduralTexture) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(proceduralTextures)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetProceduralTextures sets the ProceduralTextures property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#proceduraltextures
func (a *AbstractScene) SetProceduralTextures(proceduralTextures []*ProceduralTexture) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(proceduralTextures)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// ReflectionProbes returns the ReflectionProbes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#reflectionprobes
func (a *AbstractScene) ReflectionProbes(reflectionProbes []*ReflectionProbe) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(reflectionProbes)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetReflectionProbes sets the ReflectionProbes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#reflectionprobes
func (a *AbstractScene) SetReflectionProbes(reflectionProbes []*ReflectionProbe) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(reflectionProbes)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// RootNodes returns the RootNodes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#rootnodes
func (a *AbstractScene) RootNodes(rootNodes *Node) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(rootNodes.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetRootNodes sets the RootNodes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#rootnodes
func (a *AbstractScene) SetRootNodes(rootNodes *Node) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(rootNodes.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Skeletons returns the Skeletons property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#skeletons
func (a *AbstractScene) Skeletons(skeletons *Skeleton) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(skeletons.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetSkeletons sets the Skeletons property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#skeletons
func (a *AbstractScene) SetSkeletons(skeletons *Skeleton) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(skeletons.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Sounds returns the Sounds property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#sounds
func (a *AbstractScene) Sounds(sounds []*Sound) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(sounds)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetSounds sets the Sounds property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#sounds
func (a *AbstractScene) SetSounds(sounds []*Sound) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(sounds)
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// Textures returns the Textures property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#textures
func (a *AbstractScene) Textures(textures *BaseTexture) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(textures.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetTextures sets the Textures property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#textures
func (a *AbstractScene) SetTextures(textures *BaseTexture) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(textures.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// TransformNodes returns the TransformNodes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#transformnodes
func (a *AbstractScene) TransformNodes(transformNodes *TransformNode) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(transformNodes.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

// SetTransformNodes sets the TransformNodes property of class AbstractScene.
//
// https://doc.babylonjs.com/api/classes/babylon.abstractscene#transformnodes
func (a *AbstractScene) SetTransformNodes(transformNodes *TransformNode) *AbstractScene {
	p := ba.ctx.Get("AbstractScene").New(transformNodes.JSObject())
	return AbstractSceneFromJSObject(p, ba.ctx)
}

*/
