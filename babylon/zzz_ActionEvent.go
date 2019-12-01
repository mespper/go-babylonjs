// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ActionEvent represents a babylon.js ActionEvent.
// ActionEvent is the event being sent when an action is triggered.
type ActionEvent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ActionEvent) JSObject() js.Value { return a.p }

// ActionEvent returns a ActionEvent JavaScript class.
func (ba *Babylon) ActionEvent() *ActionEvent {
	p := ba.ctx.Get("ActionEvent")
	return ActionEventFromJSObject(p, ba.ctx)
}

// ActionEventFromJSObject returns a wrapped ActionEvent JavaScript class.
func ActionEventFromJSObject(p js.Value, ctx js.Value) *ActionEvent {
	return &ActionEvent{p: p, ctx: ctx}
}

// ActionEventArrayToJSArray returns a JavaScript Array for the wrapped array.
func ActionEventArrayToJSArray(array []*ActionEvent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewActionEventOpts contains optional parameters for NewActionEvent.
type NewActionEventOpts struct {
	SourceEvent    *interface{}
	AdditionalData *interface{}
}

// NewActionEvent returns a new ActionEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent
func (ba *Babylon) NewActionEvent(source interface{}, pointerX float64, pointerY float64, meshUnderPointer *AbstractMesh, opts *NewActionEventOpts) *ActionEvent {
	if opts == nil {
		opts = &NewActionEventOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, source)
	args = append(args, pointerX)
	args = append(args, pointerY)
	args = append(args, meshUnderPointer.JSObject())

	if opts.SourceEvent == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.SourceEvent)
	}
	if opts.AdditionalData == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.AdditionalData)
	}

	p := ba.ctx.Get("ActionEvent").New(args...)
	return ActionEventFromJSObject(p, ba.ctx)
}

// ActionEventCreateNewOpts contains optional parameters for ActionEvent.CreateNew.
type ActionEventCreateNewOpts struct {
	Evt            js.Value
	AdditionalData *interface{}
}

// CreateNew calls the CreateNew method on the ActionEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#createnew
func (a *ActionEvent) CreateNew(source *AbstractMesh, opts *ActionEventCreateNewOpts) *ActionEvent {
	if opts == nil {
		opts = &ActionEventCreateNewOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, source.JSObject())

	args = append(args, opts.Evt)
	if opts.AdditionalData == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.AdditionalData)
	}

	retVal := a.p.Call("CreateNew", args...)
	return ActionEventFromJSObject(retVal, a.ctx)
}

// ActionEventCreateNewFromPrimitiveOpts contains optional parameters for ActionEvent.CreateNewFromPrimitive.
type ActionEventCreateNewFromPrimitiveOpts struct {
	Evt            js.Value
	AdditionalData *interface{}
}

// CreateNewFromPrimitive calls the CreateNewFromPrimitive method on the ActionEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#createnewfromprimitive
func (a *ActionEvent) CreateNewFromPrimitive(prim interface{}, pointerPos *Vector2, opts *ActionEventCreateNewFromPrimitiveOpts) *ActionEvent {
	if opts == nil {
		opts = &ActionEventCreateNewFromPrimitiveOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, prim)
	args = append(args, pointerPos.JSObject())

	args = append(args, opts.Evt)
	if opts.AdditionalData == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.AdditionalData)
	}

	retVal := a.p.Call("CreateNewFromPrimitive", args...)
	return ActionEventFromJSObject(retVal, a.ctx)
}

// CreateNewFromScene calls the CreateNewFromScene method on the ActionEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#createnewfromscene
func (a *ActionEvent) CreateNewFromScene(scene *Scene, evt js.Value) *ActionEvent {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, evt)

	retVal := a.p.Call("CreateNewFromScene", args...)
	return ActionEventFromJSObject(retVal, a.ctx)
}

// ActionEventCreateNewFromSpriteOpts contains optional parameters for ActionEvent.CreateNewFromSprite.
type ActionEventCreateNewFromSpriteOpts struct {
	Evt            js.Value
	AdditionalData *interface{}
}

// CreateNewFromSprite calls the CreateNewFromSprite method on the ActionEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#createnewfromsprite
func (a *ActionEvent) CreateNewFromSprite(source *Sprite, scene *Scene, opts *ActionEventCreateNewFromSpriteOpts) *ActionEvent {
	if opts == nil {
		opts = &ActionEventCreateNewFromSpriteOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, source.JSObject())
	args = append(args, scene.JSObject())

	args = append(args, opts.Evt)
	if opts.AdditionalData == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.AdditionalData)
	}

	retVal := a.p.Call("CreateNewFromSprite", args...)
	return ActionEventFromJSObject(retVal, a.ctx)
}

/*

// AdditionalData returns the AdditionalData property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#additionaldata
func (a *ActionEvent) AdditionalData(additionalData interface{}) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(additionalData)
	return ActionEventFromJSObject(p, ba.ctx)
}

// SetAdditionalData sets the AdditionalData property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#additionaldata
func (a *ActionEvent) SetAdditionalData(additionalData interface{}) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(additionalData)
	return ActionEventFromJSObject(p, ba.ctx)
}

// MeshUnderPointer returns the MeshUnderPointer property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#meshunderpointer
func (a *ActionEvent) MeshUnderPointer(meshUnderPointer *AbstractMesh) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(meshUnderPointer.JSObject())
	return ActionEventFromJSObject(p, ba.ctx)
}

// SetMeshUnderPointer sets the MeshUnderPointer property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#meshunderpointer
func (a *ActionEvent) SetMeshUnderPointer(meshUnderPointer *AbstractMesh) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(meshUnderPointer.JSObject())
	return ActionEventFromJSObject(p, ba.ctx)
}

// PointerX returns the PointerX property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#pointerx
func (a *ActionEvent) PointerX(pointerX float64) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(pointerX)
	return ActionEventFromJSObject(p, ba.ctx)
}

// SetPointerX sets the PointerX property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#pointerx
func (a *ActionEvent) SetPointerX(pointerX float64) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(pointerX)
	return ActionEventFromJSObject(p, ba.ctx)
}

// PointerY returns the PointerY property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#pointery
func (a *ActionEvent) PointerY(pointerY float64) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(pointerY)
	return ActionEventFromJSObject(p, ba.ctx)
}

// SetPointerY sets the PointerY property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#pointery
func (a *ActionEvent) SetPointerY(pointerY float64) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(pointerY)
	return ActionEventFromJSObject(p, ba.ctx)
}

// Source returns the Source property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#source
func (a *ActionEvent) Source(source interface{}) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(source)
	return ActionEventFromJSObject(p, ba.ctx)
}

// SetSource sets the Source property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#source
func (a *ActionEvent) SetSource(source interface{}) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(source)
	return ActionEventFromJSObject(p, ba.ctx)
}

// SourceEvent returns the SourceEvent property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#sourceevent
func (a *ActionEvent) SourceEvent(sourceEvent interface{}) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(sourceEvent)
	return ActionEventFromJSObject(p, ba.ctx)
}

// SetSourceEvent sets the SourceEvent property of class ActionEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.actionevent#sourceevent
func (a *ActionEvent) SetSourceEvent(sourceEvent interface{}) *ActionEvent {
	p := ba.ctx.Get("ActionEvent").New(sourceEvent)
	return ActionEventFromJSObject(p, ba.ctx)
}

*/
