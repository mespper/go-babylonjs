// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneLoaderProgressEvent represents a babylon.js SceneLoaderProgressEvent.
// Class used to represent data loading progression
type SceneLoaderProgressEvent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SceneLoaderProgressEvent) JSObject() js.Value { return s.p }

// SceneLoaderProgressEvent returns a SceneLoaderProgressEvent JavaScript class.
func (ba *Babylon) SceneLoaderProgressEvent() *SceneLoaderProgressEvent {
	p := ba.ctx.Get("SceneLoaderProgressEvent")
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

// SceneLoaderProgressEventFromJSObject returns a wrapped SceneLoaderProgressEvent JavaScript class.
func SceneLoaderProgressEventFromJSObject(p js.Value, ctx js.Value) *SceneLoaderProgressEvent {
	return &SceneLoaderProgressEvent{p: p, ctx: ctx}
}

// SceneLoaderProgressEventArrayToJSArray returns a JavaScript Array for the wrapped array.
func SceneLoaderProgressEventArrayToJSArray(array []*SceneLoaderProgressEvent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSceneLoaderProgressEvent returns a new SceneLoaderProgressEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent
func (ba *Babylon) NewSceneLoaderProgressEvent(lengthComputable bool, loaded float64, total float64) *SceneLoaderProgressEvent {

	args := make([]interface{}, 0, 3+0)

	args = append(args, lengthComputable)
	args = append(args, loaded)
	args = append(args, total)

	p := ba.ctx.Get("SceneLoaderProgressEvent").New(args...)
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

// FromProgressEvent calls the FromProgressEvent method on the SceneLoaderProgressEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent#fromprogressevent
func (s *SceneLoaderProgressEvent) FromProgressEvent(event js.Value) *SceneLoaderProgressEvent {

	args := make([]interface{}, 0, 1+0)

	args = append(args, event)

	retVal := s.p.Call("FromProgressEvent", args...)
	return SceneLoaderProgressEventFromJSObject(retVal, s.ctx)
}

/*

// LengthComputable returns the LengthComputable property of class SceneLoaderProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent#lengthcomputable
func (s *SceneLoaderProgressEvent) LengthComputable(lengthComputable bool) *SceneLoaderProgressEvent {
	p := ba.ctx.Get("SceneLoaderProgressEvent").New(lengthComputable)
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

// SetLengthComputable sets the LengthComputable property of class SceneLoaderProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent#lengthcomputable
func (s *SceneLoaderProgressEvent) SetLengthComputable(lengthComputable bool) *SceneLoaderProgressEvent {
	p := ba.ctx.Get("SceneLoaderProgressEvent").New(lengthComputable)
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

// Loaded returns the Loaded property of class SceneLoaderProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent#loaded
func (s *SceneLoaderProgressEvent) Loaded(loaded float64) *SceneLoaderProgressEvent {
	p := ba.ctx.Get("SceneLoaderProgressEvent").New(loaded)
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

// SetLoaded sets the Loaded property of class SceneLoaderProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent#loaded
func (s *SceneLoaderProgressEvent) SetLoaded(loaded float64) *SceneLoaderProgressEvent {
	p := ba.ctx.Get("SceneLoaderProgressEvent").New(loaded)
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

// Total returns the Total property of class SceneLoaderProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent#total
func (s *SceneLoaderProgressEvent) Total(total float64) *SceneLoaderProgressEvent {
	p := ba.ctx.Get("SceneLoaderProgressEvent").New(total)
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

// SetTotal sets the Total property of class SceneLoaderProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneloaderprogressevent#total
func (s *SceneLoaderProgressEvent) SetTotal(total float64) *SceneLoaderProgressEvent {
	p := ba.ctx.Get("SceneLoaderProgressEvent").New(total)
	return SceneLoaderProgressEventFromJSObject(p, ba.ctx)
}

*/
