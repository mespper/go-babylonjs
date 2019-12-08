// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Animatable represents a babylon.js Animatable.
// Class used to store an actual running animation
type Animatable struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *Animatable) JSObject() js.Value { return a.p }

// Animatable returns a Animatable JavaScript class.
func (ba *Babylon) Animatable() *Animatable {
	p := ba.ctx.Get("Animatable")
	return AnimatableFromJSObject(p, ba.ctx)
}

// AnimatableFromJSObject returns a wrapped Animatable JavaScript class.
func AnimatableFromJSObject(p js.Value, ctx js.Value) *Animatable {
	return &Animatable{p: p, ctx: ctx}
}

// AnimatableArrayToJSArray returns a JavaScript Array for the wrapped array.
func AnimatableArrayToJSArray(array []*Animatable) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAnimatableOpts contains optional parameters for NewAnimatable.
type NewAnimatableOpts struct {
	FromFrame       *float64
	ToFrame         *float64
	LoopAnimation   *bool
	SpeedRatio      *float64
	OnAnimationEnd  JSFunc
	Animations      []*Animation
	OnAnimationLoop JSFunc
}

// NewAnimatable returns a new Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable
func (ba *Babylon) NewAnimatable(scene *Scene, target interface{}, opts *NewAnimatableOpts) *Animatable {
	if opts == nil {
		opts = &NewAnimatableOpts{}
	}

	args := make([]interface{}, 0, 2+7)

	args = append(args, scene.JSObject())
	args = append(args, target)

	if opts.FromFrame == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FromFrame)
	}
	if opts.ToFrame == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ToFrame)
	}
	if opts.LoopAnimation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LoopAnimation)
	}
	if opts.SpeedRatio == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SpeedRatio)
	}
	if opts.OnAnimationEnd == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnAnimationEnd) /* never freed! */)
	}
	if opts.Animations == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, AnimationArrayToJSArray(opts.Animations))
	}
	if opts.OnAnimationLoop == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnAnimationLoop) /* never freed! */)
	}

	p := ba.ctx.Get("Animatable").New(args...)
	return AnimatableFromJSObject(p, ba.ctx)
}

// AppendAnimations calls the AppendAnimations method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#appendanimations
func (a *Animatable) AppendAnimations(target interface{}, animations []*Animation) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, target)
	args = append(args, AnimationArrayToJSArray(animations))

	a.p.Call("appendAnimations", args...)
}

// DisableBlending calls the DisableBlending method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#disableblending
func (a *Animatable) DisableBlending() {

	a.p.Call("disableBlending")
}

// EnableBlending calls the EnableBlending method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#enableblending
func (a *Animatable) EnableBlending(blendingSpeed float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, blendingSpeed)

	a.p.Call("enableBlending", args...)
}

// GetAnimationByTargetProperty calls the GetAnimationByTargetProperty method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#getanimationbytargetproperty
func (a *Animatable) GetAnimationByTargetProperty(property string) *Animation {

	args := make([]interface{}, 0, 1+0)

	args = append(args, property)

	retVal := a.p.Call("getAnimationByTargetProperty", args...)
	return AnimationFromJSObject(retVal, a.ctx)
}

// GetAnimations calls the GetAnimations method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#getanimations
func (a *Animatable) GetAnimations() []*RuntimeAnimation {

	retVal := a.p.Call("getAnimations")
	result := []*RuntimeAnimation{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, RuntimeAnimationFromJSObject(retVal.Index(ri), a.ctx))
	}
	return result
}

// GetRuntimeAnimationByTargetProperty calls the GetRuntimeAnimationByTargetProperty method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#getruntimeanimationbytargetproperty
func (a *Animatable) GetRuntimeAnimationByTargetProperty(property string) *RuntimeAnimation {

	args := make([]interface{}, 0, 1+0)

	args = append(args, property)

	retVal := a.p.Call("getRuntimeAnimationByTargetProperty", args...)
	return RuntimeAnimationFromJSObject(retVal, a.ctx)
}

// GoToFrame calls the GoToFrame method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#gotoframe
func (a *Animatable) GoToFrame(frame float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, frame)

	a.p.Call("goToFrame", args...)
}

// Pause calls the Pause method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#pause
func (a *Animatable) Pause() {

	a.p.Call("pause")
}

// Reset calls the Reset method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#reset
func (a *Animatable) Reset() {

	a.p.Call("reset")
}

// Restart calls the Restart method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#restart
func (a *Animatable) Restart() {

	a.p.Call("restart")
}

// AnimatableStopOpts contains optional parameters for Animatable.Stop.
type AnimatableStopOpts struct {
	AnimationName *string
	TargetMask    JSFunc
}

// Stop calls the Stop method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#stop
func (a *Animatable) Stop(opts *AnimatableStopOpts) {
	if opts == nil {
		opts = &AnimatableStopOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.AnimationName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AnimationName)
	}
	if opts.TargetMask == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.TargetMask) /* never freed! */)
	}

	a.p.Call("stop", args...)
}

// SyncWith calls the SyncWith method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#syncwith
func (a *Animatable) SyncWith(root *Animatable) *Animatable {

	args := make([]interface{}, 0, 1+0)

	args = append(args, root.JSObject())

	retVal := a.p.Call("syncWith", args...)
	return AnimatableFromJSObject(retVal, a.ctx)
}

// WaitAsync calls the WaitAsync method on the Animatable object.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#waitasync
func (a *Animatable) WaitAsync() *Promise {

	retVal := a.p.Call("waitAsync")
	return PromiseFromJSObject(retVal, a.ctx)
}

// AnimationStarted returns the AnimationStarted property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#animationstarted
func (a *Animatable) AnimationStarted() bool {
	retVal := a.p.Get("animationStarted")
	return retVal.Bool()
}

// SetAnimationStarted sets the AnimationStarted property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#animationstarted
func (a *Animatable) SetAnimationStarted(animationStarted bool) *Animatable {
	a.p.Set("animationStarted", animationStarted)
	return a
}

// DisposeOnEnd returns the DisposeOnEnd property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#disposeonend
func (a *Animatable) DisposeOnEnd() bool {
	retVal := a.p.Get("disposeOnEnd")
	return retVal.Bool()
}

// SetDisposeOnEnd sets the DisposeOnEnd property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#disposeonend
func (a *Animatable) SetDisposeOnEnd(disposeOnEnd bool) *Animatable {
	a.p.Set("disposeOnEnd", disposeOnEnd)
	return a
}

// FromFrame returns the FromFrame property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#fromframe
func (a *Animatable) FromFrame() float64 {
	retVal := a.p.Get("fromFrame")
	return retVal.Float()
}

// SetFromFrame sets the FromFrame property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#fromframe
func (a *Animatable) SetFromFrame(fromFrame float64) *Animatable {
	a.p.Set("fromFrame", fromFrame)
	return a
}

// LoopAnimation returns the LoopAnimation property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#loopanimation
func (a *Animatable) LoopAnimation() bool {
	retVal := a.p.Get("loopAnimation")
	return retVal.Bool()
}

// SetLoopAnimation sets the LoopAnimation property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#loopanimation
func (a *Animatable) SetLoopAnimation(loopAnimation bool) *Animatable {
	a.p.Set("loopAnimation", loopAnimation)
	return a
}

// MasterFrame returns the MasterFrame property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#masterframe
func (a *Animatable) MasterFrame() float64 {
	retVal := a.p.Get("masterFrame")
	return retVal.Float()
}

// SetMasterFrame sets the MasterFrame property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#masterframe
func (a *Animatable) SetMasterFrame(masterFrame float64) *Animatable {
	a.p.Set("masterFrame", masterFrame)
	return a
}

// OnAnimationEnd returns the OnAnimationEnd property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationend
func (a *Animatable) OnAnimationEnd() js.Value {
	retVal := a.p.Get("onAnimationEnd")
	return retVal
}

// SetOnAnimationEnd sets the OnAnimationEnd property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationend
func (a *Animatable) SetOnAnimationEnd(onAnimationEnd JSFunc) *Animatable {
	a.p.Set("onAnimationEnd", js.FuncOf(onAnimationEnd))
	return a
}

// OnAnimationEndObservable returns the OnAnimationEndObservable property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationendobservable
func (a *Animatable) OnAnimationEndObservable() *Observable {
	retVal := a.p.Get("onAnimationEndObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnAnimationEndObservable sets the OnAnimationEndObservable property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationendobservable
func (a *Animatable) SetOnAnimationEndObservable(onAnimationEndObservable *Observable) *Animatable {
	a.p.Set("onAnimationEndObservable", onAnimationEndObservable.JSObject())
	return a
}

// OnAnimationLoop returns the OnAnimationLoop property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationloop
func (a *Animatable) OnAnimationLoop() js.Value {
	retVal := a.p.Get("onAnimationLoop")
	return retVal
}

// SetOnAnimationLoop sets the OnAnimationLoop property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationloop
func (a *Animatable) SetOnAnimationLoop(onAnimationLoop JSFunc) *Animatable {
	a.p.Set("onAnimationLoop", js.FuncOf(onAnimationLoop))
	return a
}

// OnAnimationLoopObservable returns the OnAnimationLoopObservable property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationloopobservable
func (a *Animatable) OnAnimationLoopObservable() *Observable {
	retVal := a.p.Get("onAnimationLoopObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnAnimationLoopObservable sets the OnAnimationLoopObservable property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#onanimationloopobservable
func (a *Animatable) SetOnAnimationLoopObservable(onAnimationLoopObservable *Observable) *Animatable {
	a.p.Set("onAnimationLoopObservable", onAnimationLoopObservable.JSObject())
	return a
}

// SpeedRatio returns the SpeedRatio property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#speedratio
func (a *Animatable) SpeedRatio() float64 {
	retVal := a.p.Get("speedRatio")
	return retVal.Float()
}

// SetSpeedRatio sets the SpeedRatio property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#speedratio
func (a *Animatable) SetSpeedRatio(speedRatio float64) *Animatable {
	a.p.Set("speedRatio", speedRatio)
	return a
}

// SyncRoot returns the SyncRoot property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#syncroot
func (a *Animatable) SyncRoot() *Animatable {
	retVal := a.p.Get("syncRoot")
	return AnimatableFromJSObject(retVal, a.ctx)
}

// SetSyncRoot sets the SyncRoot property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#syncroot
func (a *Animatable) SetSyncRoot(syncRoot *Animatable) *Animatable {
	a.p.Set("syncRoot", syncRoot.JSObject())
	return a
}

// Target returns the Target property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#target
func (a *Animatable) Target() interface{} {
	retVal := a.p.Get("target")
	return retVal
}

// SetTarget sets the Target property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#target
func (a *Animatable) SetTarget(target interface{}) *Animatable {
	a.p.Set("target", target)
	return a
}

// ToFrame returns the ToFrame property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#toframe
func (a *Animatable) ToFrame() float64 {
	retVal := a.p.Get("toFrame")
	return retVal.Float()
}

// SetToFrame sets the ToFrame property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#toframe
func (a *Animatable) SetToFrame(toFrame float64) *Animatable {
	a.p.Set("toFrame", toFrame)
	return a
}

// Weight returns the Weight property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#weight
func (a *Animatable) Weight() float64 {
	retVal := a.p.Get("weight")
	return retVal.Float()
}

// SetWeight sets the Weight property of class Animatable.
//
// https://doc.babylonjs.com/api/classes/babylon.animatable#weight
func (a *Animatable) SetWeight(weight float64) *Animatable {
	a.p.Set("weight", weight)
	return a
}
