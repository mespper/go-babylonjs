// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AutoRotationBehavior represents a babylon.js AutoRotationBehavior.
// The autoRotation behavior (AutoRotationBehavior) is designed to create a smooth rotation of an ArcRotateCamera when there is no user interaction.
//
// See: http://doc.babylonjs.com/how_to/camera_behaviors#autorotation-behavior
type AutoRotationBehavior struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AutoRotationBehavior) JSObject() js.Value { return a.p }

// AutoRotationBehavior returns a AutoRotationBehavior JavaScript class.
func (ba *Babylon) AutoRotationBehavior() *AutoRotationBehavior {
	p := ba.ctx.Get("AutoRotationBehavior")
	return AutoRotationBehaviorFromJSObject(p, ba.ctx)
}

// AutoRotationBehaviorFromJSObject returns a wrapped AutoRotationBehavior JavaScript class.
func AutoRotationBehaviorFromJSObject(p js.Value, ctx js.Value) *AutoRotationBehavior {
	return &AutoRotationBehavior{p: p, ctx: ctx}
}

// AutoRotationBehaviorArrayToJSArray returns a JavaScript Array for the wrapped array.
func AutoRotationBehaviorArrayToJSArray(array []*AutoRotationBehavior) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Attach calls the Attach method on the AutoRotationBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#attach
func (a *AutoRotationBehavior) Attach(camera *ArcRotateCamera) {

	args := make([]interface{}, 0, 1+0)

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	a.p.Call("attach", args...)
}

// Detach calls the Detach method on the AutoRotationBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#detach
func (a *AutoRotationBehavior) Detach() {

	a.p.Call("detach")
}

// Init calls the Init method on the AutoRotationBehavior object.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#init
func (a *AutoRotationBehavior) Init() {

	a.p.Call("init")
}

// IdleRotationSpeed returns the IdleRotationSpeed property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#idlerotationspeed
func (a *AutoRotationBehavior) IdleRotationSpeed() float64 {
	retVal := a.p.Get("idleRotationSpeed")
	return retVal.Float()
}

// SetIdleRotationSpeed sets the IdleRotationSpeed property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#idlerotationspeed
func (a *AutoRotationBehavior) SetIdleRotationSpeed(idleRotationSpeed float64) *AutoRotationBehavior {
	a.p.Set("idleRotationSpeed", idleRotationSpeed)
	return a
}

// IdleRotationSpinupTime returns the IdleRotationSpinupTime property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#idlerotationspinuptime
func (a *AutoRotationBehavior) IdleRotationSpinupTime() float64 {
	retVal := a.p.Get("idleRotationSpinupTime")
	return retVal.Float()
}

// SetIdleRotationSpinupTime sets the IdleRotationSpinupTime property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#idlerotationspinuptime
func (a *AutoRotationBehavior) SetIdleRotationSpinupTime(idleRotationSpinupTime float64) *AutoRotationBehavior {
	a.p.Set("idleRotationSpinupTime", idleRotationSpinupTime)
	return a
}

// IdleRotationWaitTime returns the IdleRotationWaitTime property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#idlerotationwaittime
func (a *AutoRotationBehavior) IdleRotationWaitTime() float64 {
	retVal := a.p.Get("idleRotationWaitTime")
	return retVal.Float()
}

// SetIdleRotationWaitTime sets the IdleRotationWaitTime property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#idlerotationwaittime
func (a *AutoRotationBehavior) SetIdleRotationWaitTime(idleRotationWaitTime float64) *AutoRotationBehavior {
	a.p.Set("idleRotationWaitTime", idleRotationWaitTime)
	return a
}

// Name returns the Name property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#name
func (a *AutoRotationBehavior) Name() string {
	retVal := a.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#name
func (a *AutoRotationBehavior) SetName(name string) *AutoRotationBehavior {
	a.p.Set("name", name)
	return a
}

// RotationInProgress returns the RotationInProgress property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#rotationinprogress
func (a *AutoRotationBehavior) RotationInProgress() bool {
	retVal := a.p.Get("rotationInProgress")
	return retVal.Bool()
}

// SetRotationInProgress sets the RotationInProgress property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#rotationinprogress
func (a *AutoRotationBehavior) SetRotationInProgress(rotationInProgress bool) *AutoRotationBehavior {
	a.p.Set("rotationInProgress", rotationInProgress)
	return a
}

// ZoomStopsAnimation returns the ZoomStopsAnimation property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#zoomstopsanimation
func (a *AutoRotationBehavior) ZoomStopsAnimation() bool {
	retVal := a.p.Get("zoomStopsAnimation")
	return retVal.Bool()
}

// SetZoomStopsAnimation sets the ZoomStopsAnimation property of class AutoRotationBehavior.
//
// https://doc.babylonjs.com/api/classes/babylon.autorotationbehavior#zoomstopsanimation
func (a *AutoRotationBehavior) SetZoomStopsAnimation(zoomStopsAnimation bool) *AutoRotationBehavior {
	a.p.Set("zoomStopsAnimation", zoomStopsAnimation)
	return a
}
