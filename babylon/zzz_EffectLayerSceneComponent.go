// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EffectLayerSceneComponent represents a babylon.js EffectLayerSceneComponent.
// Defines the layer scene component responsible to manage any effect layers
// in a given scene.
type EffectLayerSceneComponent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EffectLayerSceneComponent) JSObject() js.Value { return e.p }

// EffectLayerSceneComponent returns a EffectLayerSceneComponent JavaScript class.
func (ba *Babylon) EffectLayerSceneComponent() *EffectLayerSceneComponent {
	p := ba.ctx.Get("EffectLayerSceneComponent")
	return EffectLayerSceneComponentFromJSObject(p, ba.ctx)
}

// EffectLayerSceneComponentFromJSObject returns a wrapped EffectLayerSceneComponent JavaScript class.
func EffectLayerSceneComponentFromJSObject(p js.Value, ctx js.Value) *EffectLayerSceneComponent {
	return &EffectLayerSceneComponent{p: p, ctx: ctx}
}

// EffectLayerSceneComponentArrayToJSArray returns a JavaScript Array for the wrapped array.
func EffectLayerSceneComponentArrayToJSArray(array []*EffectLayerSceneComponent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewEffectLayerSceneComponent returns a new EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent
func (ba *Babylon) NewEffectLayerSceneComponent(scene *Scene) *EffectLayerSceneComponent {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("EffectLayerSceneComponent").New(args...)
	return EffectLayerSceneComponentFromJSObject(p, ba.ctx)
}

// AddFromContainer calls the AddFromContainer method on the EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#addfromcontainer
func (e *EffectLayerSceneComponent) AddFromContainer(container *AbstractScene) {

	args := make([]interface{}, 0, 1+0)

	if container == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, container.JSObject())
	}

	e.p.Call("addFromContainer", args...)
}

// Dispose calls the Dispose method on the EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#dispose
func (e *EffectLayerSceneComponent) Dispose() {

	e.p.Call("dispose")
}

// Rebuild calls the Rebuild method on the EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#rebuild
func (e *EffectLayerSceneComponent) Rebuild() {

	e.p.Call("rebuild")
}

// Register calls the Register method on the EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#register
func (e *EffectLayerSceneComponent) Register() {

	e.p.Call("register")
}

// EffectLayerSceneComponentRemoveFromContainerOpts contains optional parameters for EffectLayerSceneComponent.RemoveFromContainer.
type EffectLayerSceneComponentRemoveFromContainerOpts struct {
	Dispose *bool
}

// RemoveFromContainer calls the RemoveFromContainer method on the EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#removefromcontainer
func (e *EffectLayerSceneComponent) RemoveFromContainer(container *AbstractScene, opts *EffectLayerSceneComponentRemoveFromContainerOpts) {
	if opts == nil {
		opts = &EffectLayerSceneComponentRemoveFromContainerOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if container == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, container.JSObject())
	}

	if opts.Dispose == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Dispose)
	}

	e.p.Call("removeFromContainer", args...)
}

// Serialize calls the Serialize method on the EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#serialize
func (e *EffectLayerSceneComponent) Serialize(serializationObject JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	e.p.Call("serialize", args...)
}

// Name returns the Name property of class EffectLayerSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#name
func (e *EffectLayerSceneComponent) Name() string {
	retVal := e.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class EffectLayerSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#name
func (e *EffectLayerSceneComponent) SetName(name string) *EffectLayerSceneComponent {
	e.p.Set("name", name)
	return e
}

// Scene returns the Scene property of class EffectLayerSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#scene
func (e *EffectLayerSceneComponent) Scene() *Scene {
	retVal := e.p.Get("scene")
	return SceneFromJSObject(retVal, e.ctx)
}

// SetScene sets the Scene property of class EffectLayerSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent#scene
func (e *EffectLayerSceneComponent) SetScene(scene *Scene) *EffectLayerSceneComponent {
	e.p.Set("scene", scene.JSObject())
	return e
}
