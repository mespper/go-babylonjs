// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensFlareSystemSceneComponent represents a babylon.js LensFlareSystemSceneComponent.
// Defines the lens flare scene component responsible to manage any lens flares
// in a given scene.
type LensFlareSystemSceneComponent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LensFlareSystemSceneComponent) JSObject() js.Value { return l.p }

// LensFlareSystemSceneComponent returns a LensFlareSystemSceneComponent JavaScript class.
func (ba *Babylon) LensFlareSystemSceneComponent() *LensFlareSystemSceneComponent {
	p := ba.ctx.Get("LensFlareSystemSceneComponent")
	return LensFlareSystemSceneComponentFromJSObject(p, ba.ctx)
}

// LensFlareSystemSceneComponentFromJSObject returns a wrapped LensFlareSystemSceneComponent JavaScript class.
func LensFlareSystemSceneComponentFromJSObject(p js.Value, ctx js.Value) *LensFlareSystemSceneComponent {
	return &LensFlareSystemSceneComponent{p: p, ctx: ctx}
}

// LensFlareSystemSceneComponentArrayToJSArray returns a JavaScript Array for the wrapped array.
func LensFlareSystemSceneComponentArrayToJSArray(array []*LensFlareSystemSceneComponent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewLensFlareSystemSceneComponent returns a new LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent
func (ba *Babylon) NewLensFlareSystemSceneComponent(scene *Scene) *LensFlareSystemSceneComponent {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("LensFlareSystemSceneComponent").New(args...)
	return LensFlareSystemSceneComponentFromJSObject(p, ba.ctx)
}

// AddFromContainer calls the AddFromContainer method on the LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#addfromcontainer
func (l *LensFlareSystemSceneComponent) AddFromContainer(container *AbstractScene) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, container.JSObject())

	l.p.Call("addFromContainer", args...)
}

// Dispose calls the Dispose method on the LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#dispose
func (l *LensFlareSystemSceneComponent) Dispose() {

	l.p.Call("dispose")
}

// Rebuild calls the Rebuild method on the LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#rebuild
func (l *LensFlareSystemSceneComponent) Rebuild() {

	l.p.Call("rebuild")
}

// Register calls the Register method on the LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#register
func (l *LensFlareSystemSceneComponent) Register() {

	l.p.Call("register")
}

// LensFlareSystemSceneComponentRemoveFromContainerOpts contains optional parameters for LensFlareSystemSceneComponent.RemoveFromContainer.
type LensFlareSystemSceneComponentRemoveFromContainerOpts struct {
	Dispose *bool
}

// RemoveFromContainer calls the RemoveFromContainer method on the LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#removefromcontainer
func (l *LensFlareSystemSceneComponent) RemoveFromContainer(container *AbstractScene, opts *LensFlareSystemSceneComponentRemoveFromContainerOpts) {
	if opts == nil {
		opts = &LensFlareSystemSceneComponentRemoveFromContainerOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, container.JSObject())

	if opts.Dispose == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Dispose)
	}

	l.p.Call("removeFromContainer", args...)
}

// Serialize calls the Serialize method on the LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#serialize
func (l *LensFlareSystemSceneComponent) Serialize(serializationObject interface{}) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, serializationObject)

	l.p.Call("serialize", args...)
}

// Name returns the Name property of class LensFlareSystemSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#name
func (l *LensFlareSystemSceneComponent) Name() string {
	retVal := l.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class LensFlareSystemSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#name
func (l *LensFlareSystemSceneComponent) SetName(name string) *LensFlareSystemSceneComponent {
	l.p.Set("name", name)
	return l
}

// Scene returns the Scene property of class LensFlareSystemSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#scene
func (l *LensFlareSystemSceneComponent) Scene() *Scene {
	retVal := l.p.Get("scene")
	return SceneFromJSObject(retVal, l.ctx)
}

// SetScene sets the Scene property of class LensFlareSystemSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent#scene
func (l *LensFlareSystemSceneComponent) SetScene(scene *Scene) *LensFlareSystemSceneComponent {
	l.p.Set("scene", scene.JSObject())
	return l
}
