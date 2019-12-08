// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AdvancedDynamicTexture represents a babylon.js AdvancedDynamicTexture.
// Class used to create texture to support 2D GUI elements
//
// See: http://doc.babylonjs.com/how_to/gui
type AdvancedDynamicTexture struct {
	*DynamicTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AdvancedDynamicTexture) JSObject() js.Value { return a.p }

// AdvancedDynamicTexture returns a AdvancedDynamicTexture JavaScript class.
func (gui *GUI) AdvancedDynamicTexture() *AdvancedDynamicTexture {
	p := gui.ctx.Get("AdvancedDynamicTexture")
	return AdvancedDynamicTextureFromJSObject(p, gui.ctx)
}

// AdvancedDynamicTextureFromJSObject returns a wrapped AdvancedDynamicTexture JavaScript class.
func AdvancedDynamicTextureFromJSObject(p js.Value, ctx js.Value) *AdvancedDynamicTexture {
	return &AdvancedDynamicTexture{DynamicTexture: DynamicTextureFromJSObject(p, ctx), ctx: ctx}
}

// AdvancedDynamicTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func AdvancedDynamicTextureArrayToJSArray(array []*AdvancedDynamicTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAdvancedDynamicTextureOpts contains optional parameters for NewAdvancedDynamicTexture.
type NewAdvancedDynamicTextureOpts struct {
	GenerateMipMaps *bool
	SamplingMode    *float64
}

// NewAdvancedDynamicTexture returns a new AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture
func (gui *GUI) NewAdvancedDynamicTexture(name string, width float64, height float64, scene *Scene, opts *NewAdvancedDynamicTextureOpts) *AdvancedDynamicTexture {
	if opts == nil {
		opts = &NewAdvancedDynamicTextureOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, name)
	args = append(args, width)
	args = append(args, height)
	args = append(args, scene.JSObject())

	if opts.GenerateMipMaps == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateMipMaps)
	}
	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}

	p := gui.ctx.Get("AdvancedDynamicTexture").New(args...)
	return AdvancedDynamicTextureFromJSObject(p, gui.ctx)
}

// AddControl calls the AddControl method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#addcontrol
func (a *AdvancedDynamicTexture) AddControl(control *Control) *AdvancedDynamicTexture {

	args := make([]interface{}, 0, 1+0)

	if control == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, control.JSObject())
	}

	retVal := a.p.Call("addControl", args...)
	return AdvancedDynamicTextureFromJSObject(retVal, a.ctx)
}

// Attach calls the Attach method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#attach
func (a *AdvancedDynamicTexture) Attach() {

	a.p.Call("attach")
}

// AdvancedDynamicTextureAttachToMeshOpts contains optional parameters for AdvancedDynamicTexture.AttachToMesh.
type AdvancedDynamicTextureAttachToMeshOpts struct {
	SupportPointerMove *bool
}

// AttachToMesh calls the AttachToMesh method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#attachtomesh
func (a *AdvancedDynamicTexture) AttachToMesh(mesh *AbstractMesh, opts *AdvancedDynamicTextureAttachToMeshOpts) {
	if opts == nil {
		opts = &AdvancedDynamicTextureAttachToMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.SupportPointerMove == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SupportPointerMove)
	}

	a.p.Call("attachToMesh", args...)
}

// AdvancedDynamicTextureCreateForMeshOpts contains optional parameters for AdvancedDynamicTexture.CreateForMesh.
type AdvancedDynamicTextureCreateForMeshOpts struct {
	Width              *float64
	Height             *float64
	SupportPointerMove *bool
	OnlyAlphaTesting   *bool
}

// CreateForMesh calls the CreateForMesh method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#createformesh
func (a *AdvancedDynamicTexture) CreateForMesh(mesh *AbstractMesh, opts *AdvancedDynamicTextureCreateForMeshOpts) *AdvancedDynamicTexture {
	if opts == nil {
		opts = &AdvancedDynamicTextureCreateForMeshOpts{}
	}

	args := make([]interface{}, 0, 1+4)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if opts.Width == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Width)
	}
	if opts.Height == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Height)
	}
	if opts.SupportPointerMove == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SupportPointerMove)
	}
	if opts.OnlyAlphaTesting == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.OnlyAlphaTesting)
	}

	retVal := a.p.Call("CreateForMesh", args...)
	return AdvancedDynamicTextureFromJSObject(retVal, a.ctx)
}

// AdvancedDynamicTextureCreateFullscreenUIOpts contains optional parameters for AdvancedDynamicTexture.CreateFullscreenUI.
type AdvancedDynamicTextureCreateFullscreenUIOpts struct {
	Foreground *bool
	Scene      *Scene
	Sampling   *float64
}

// CreateFullscreenUI calls the CreateFullscreenUI method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#createfullscreenui
func (a *AdvancedDynamicTexture) CreateFullscreenUI(name string, opts *AdvancedDynamicTextureCreateFullscreenUIOpts) *AdvancedDynamicTexture {
	if opts == nil {
		opts = &AdvancedDynamicTextureCreateFullscreenUIOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, name)

	if opts.Foreground == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Foreground)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.Sampling == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Sampling)
	}

	retVal := a.p.Call("CreateFullscreenUI", args...)
	return AdvancedDynamicTextureFromJSObject(retVal, a.ctx)
}

// CreateStyle calls the CreateStyle method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#createstyle
func (a *AdvancedDynamicTexture) CreateStyle() *Style {

	retVal := a.p.Call("createStyle")
	return StyleFromJSObject(retVal, a.ctx)
}

// Dispose calls the Dispose method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#dispose
func (a *AdvancedDynamicTexture) Dispose() {

	a.p.Call("dispose")
}

// AdvancedDynamicTextureExecuteOnAllControlsOpts contains optional parameters for AdvancedDynamicTexture.ExecuteOnAllControls.
type AdvancedDynamicTextureExecuteOnAllControlsOpts struct {
	Container *Container
}

// ExecuteOnAllControls calls the ExecuteOnAllControls method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#executeonallcontrols
func (a *AdvancedDynamicTexture) ExecuteOnAllControls(jsFunc JSFunc, opts *AdvancedDynamicTextureExecuteOnAllControlsOpts) {
	if opts == nil {
		opts = &AdvancedDynamicTextureExecuteOnAllControlsOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, js.FuncOf(jsFunc))

	if opts.Container == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Container.JSObject())
	}

	a.p.Call("executeOnAllControls", args...)
}

// GetChildren calls the GetChildren method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#getchildren
func (a *AdvancedDynamicTexture) GetChildren() []*Container {

	retVal := a.p.Call("getChildren")
	result := []*Container{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ContainerFromJSObject(retVal.Index(ri), a.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#getclassname
func (a *AdvancedDynamicTexture) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

// AdvancedDynamicTextureGetDescendantsOpts contains optional parameters for AdvancedDynamicTexture.GetDescendants.
type AdvancedDynamicTextureGetDescendantsOpts struct {
	DirectDescendantsOnly *bool
	Predicate             JSFunc
}

// GetDescendants calls the GetDescendants method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#getdescendants
func (a *AdvancedDynamicTexture) GetDescendants(opts *AdvancedDynamicTextureGetDescendantsOpts) []*Control {
	if opts == nil {
		opts = &AdvancedDynamicTextureGetDescendantsOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.DirectDescendantsOnly == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DirectDescendantsOnly)
	}
	if opts.Predicate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.Predicate) /* never freed! */)
	}

	retVal := a.p.Call("getDescendants", args...)
	result := []*Control{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ControlFromJSObject(retVal.Index(ri), a.ctx))
	}
	return result
}

// GetProjectedPosition calls the GetProjectedPosition method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#getprojectedposition
func (a *AdvancedDynamicTexture) GetProjectedPosition(position *Vector3, worldMatrix *Matrix) *Vector2 {

	args := make([]interface{}, 0, 2+0)

	if position == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, position.JSObject())
	}

	if worldMatrix == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, worldMatrix.JSObject())
	}

	retVal := a.p.Call("getProjectedPosition", args...)
	return Vector2FromJSObject(retVal, a.ctx)
}

// InvalidateRect calls the InvalidateRect method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#invalidaterect
func (a *AdvancedDynamicTexture) InvalidateRect(invalidMinX float64, invalidMinY float64, invalidMaxX float64, invalidMaxY float64) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, invalidMinX)

	args = append(args, invalidMinY)

	args = append(args, invalidMaxX)

	args = append(args, invalidMaxY)

	a.p.Call("invalidateRect", args...)
}

// MarkAsDirty calls the MarkAsDirty method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#markasdirty
func (a *AdvancedDynamicTexture) MarkAsDirty() {

	a.p.Call("markAsDirty")
}

// MoveFocusToControl calls the MoveFocusToControl method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#movefocustocontrol
func (a *AdvancedDynamicTexture) MoveFocusToControl(control *IFocusableControl) {

	args := make([]interface{}, 0, 1+0)

	if control == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, control.JSObject())
	}

	a.p.Call("moveFocusToControl", args...)
}

// RegisterClipboardEvents calls the RegisterClipboardEvents method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#registerclipboardevents
func (a *AdvancedDynamicTexture) RegisterClipboardEvents() {

	a.p.Call("registerClipboardEvents")
}

// RemoveControl calls the RemoveControl method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#removecontrol
func (a *AdvancedDynamicTexture) RemoveControl(control *Control) *AdvancedDynamicTexture {

	args := make([]interface{}, 0, 1+0)

	if control == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, control.JSObject())
	}

	retVal := a.p.Call("removeControl", args...)
	return AdvancedDynamicTextureFromJSObject(retVal, a.ctx)
}

// UnRegisterClipboardEvents calls the UnRegisterClipboardEvents method on the AdvancedDynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#unregisterclipboardevents
func (a *AdvancedDynamicTexture) UnRegisterClipboardEvents() {

	a.p.Call("unRegisterClipboardEvents")
}

// Background returns the Background property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#background
func (a *AdvancedDynamicTexture) Background() string {
	retVal := a.p.Get("background")
	return retVal.String()
}

// SetBackground sets the Background property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#background
func (a *AdvancedDynamicTexture) SetBackground(background string) *AdvancedDynamicTexture {
	a.p.Set("background", background)
	return a
}

// ClipboardData returns the ClipboardData property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#clipboarddata
func (a *AdvancedDynamicTexture) ClipboardData() string {
	retVal := a.p.Get("clipboardData")
	return retVal.String()
}

// SetClipboardData sets the ClipboardData property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#clipboarddata
func (a *AdvancedDynamicTexture) SetClipboardData(clipboardData string) *AdvancedDynamicTexture {
	a.p.Set("clipboardData", clipboardData)
	return a
}

// FocusedControl returns the FocusedControl property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#focusedcontrol
func (a *AdvancedDynamicTexture) FocusedControl() *IFocusableControl {
	retVal := a.p.Get("focusedControl")
	return IFocusableControlFromJSObject(retVal, a.ctx)
}

// SetFocusedControl sets the FocusedControl property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#focusedcontrol
func (a *AdvancedDynamicTexture) SetFocusedControl(focusedControl *IFocusableControl) *AdvancedDynamicTexture {
	a.p.Set("focusedControl", focusedControl.JSObject())
	return a
}

// IdealHeight returns the IdealHeight property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#idealheight
func (a *AdvancedDynamicTexture) IdealHeight() float64 {
	retVal := a.p.Get("idealHeight")
	return retVal.Float()
}

// SetIdealHeight sets the IdealHeight property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#idealheight
func (a *AdvancedDynamicTexture) SetIdealHeight(idealHeight float64) *AdvancedDynamicTexture {
	a.p.Set("idealHeight", idealHeight)
	return a
}

// IdealWidth returns the IdealWidth property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#idealwidth
func (a *AdvancedDynamicTexture) IdealWidth() float64 {
	retVal := a.p.Get("idealWidth")
	return retVal.Float()
}

// SetIdealWidth sets the IdealWidth property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#idealwidth
func (a *AdvancedDynamicTexture) SetIdealWidth(idealWidth float64) *AdvancedDynamicTexture {
	a.p.Set("idealWidth", idealWidth)
	return a
}

// IsForeground returns the IsForeground property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#isforeground
func (a *AdvancedDynamicTexture) IsForeground() bool {
	retVal := a.p.Get("isForeground")
	return retVal.Bool()
}

// SetIsForeground sets the IsForeground property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#isforeground
func (a *AdvancedDynamicTexture) SetIsForeground(isForeground bool) *AdvancedDynamicTexture {
	a.p.Set("isForeground", isForeground)
	return a
}

// Layer returns the Layer property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#layer
func (a *AdvancedDynamicTexture) Layer() *Layer {
	retVal := a.p.Get("layer")
	return LayerFromJSObject(retVal, a.ctx)
}

// SetLayer sets the Layer property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#layer
func (a *AdvancedDynamicTexture) SetLayer(layer *Layer) *AdvancedDynamicTexture {
	a.p.Set("layer", layer.JSObject())
	return a
}

// OnBeginLayoutObservable returns the OnBeginLayoutObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onbeginlayoutobservable
func (a *AdvancedDynamicTexture) OnBeginLayoutObservable() *Observable {
	retVal := a.p.Get("onBeginLayoutObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnBeginLayoutObservable sets the OnBeginLayoutObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onbeginlayoutobservable
func (a *AdvancedDynamicTexture) SetOnBeginLayoutObservable(onBeginLayoutObservable *Observable) *AdvancedDynamicTexture {
	a.p.Set("onBeginLayoutObservable", onBeginLayoutObservable.JSObject())
	return a
}

// OnBeginRenderObservable returns the OnBeginRenderObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onbeginrenderobservable
func (a *AdvancedDynamicTexture) OnBeginRenderObservable() *Observable {
	retVal := a.p.Get("onBeginRenderObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnBeginRenderObservable sets the OnBeginRenderObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onbeginrenderobservable
func (a *AdvancedDynamicTexture) SetOnBeginRenderObservable(onBeginRenderObservable *Observable) *AdvancedDynamicTexture {
	a.p.Set("onBeginRenderObservable", onBeginRenderObservable.JSObject())
	return a
}

// OnClipboardObservable returns the OnClipboardObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onclipboardobservable
func (a *AdvancedDynamicTexture) OnClipboardObservable() *Observable {
	retVal := a.p.Get("onClipboardObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnClipboardObservable sets the OnClipboardObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onclipboardobservable
func (a *AdvancedDynamicTexture) SetOnClipboardObservable(onClipboardObservable *Observable) *AdvancedDynamicTexture {
	a.p.Set("onClipboardObservable", onClipboardObservable.JSObject())
	return a
}

// OnControlPickedObservable returns the OnControlPickedObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#oncontrolpickedobservable
func (a *AdvancedDynamicTexture) OnControlPickedObservable() *Observable {
	retVal := a.p.Get("onControlPickedObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnControlPickedObservable sets the OnControlPickedObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#oncontrolpickedobservable
func (a *AdvancedDynamicTexture) SetOnControlPickedObservable(onControlPickedObservable *Observable) *AdvancedDynamicTexture {
	a.p.Set("onControlPickedObservable", onControlPickedObservable.JSObject())
	return a
}

// OnEndLayoutObservable returns the OnEndLayoutObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onendlayoutobservable
func (a *AdvancedDynamicTexture) OnEndLayoutObservable() *Observable {
	retVal := a.p.Get("onEndLayoutObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnEndLayoutObservable sets the OnEndLayoutObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onendlayoutobservable
func (a *AdvancedDynamicTexture) SetOnEndLayoutObservable(onEndLayoutObservable *Observable) *AdvancedDynamicTexture {
	a.p.Set("onEndLayoutObservable", onEndLayoutObservable.JSObject())
	return a
}

// OnEndRenderObservable returns the OnEndRenderObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onendrenderobservable
func (a *AdvancedDynamicTexture) OnEndRenderObservable() *Observable {
	retVal := a.p.Get("onEndRenderObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnEndRenderObservable sets the OnEndRenderObservable property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#onendrenderobservable
func (a *AdvancedDynamicTexture) SetOnEndRenderObservable(onEndRenderObservable *Observable) *AdvancedDynamicTexture {
	a.p.Set("onEndRenderObservable", onEndRenderObservable.JSObject())
	return a
}

// PremulAlpha returns the PremulAlpha property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#premulalpha
func (a *AdvancedDynamicTexture) PremulAlpha() bool {
	retVal := a.p.Get("premulAlpha")
	return retVal.Bool()
}

// SetPremulAlpha sets the PremulAlpha property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#premulalpha
func (a *AdvancedDynamicTexture) SetPremulAlpha(premulAlpha bool) *AdvancedDynamicTexture {
	a.p.Set("premulAlpha", premulAlpha)
	return a
}

// RenderAtIdealSize returns the RenderAtIdealSize property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#renderatidealsize
func (a *AdvancedDynamicTexture) RenderAtIdealSize() bool {
	retVal := a.p.Get("renderAtIdealSize")
	return retVal.Bool()
}

// SetRenderAtIdealSize sets the RenderAtIdealSize property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#renderatidealsize
func (a *AdvancedDynamicTexture) SetRenderAtIdealSize(renderAtIdealSize bool) *AdvancedDynamicTexture {
	a.p.Set("renderAtIdealSize", renderAtIdealSize)
	return a
}

// RenderScale returns the RenderScale property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#renderscale
func (a *AdvancedDynamicTexture) RenderScale() float64 {
	retVal := a.p.Get("renderScale")
	return retVal.Float()
}

// SetRenderScale sets the RenderScale property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#renderscale
func (a *AdvancedDynamicTexture) SetRenderScale(renderScale float64) *AdvancedDynamicTexture {
	a.p.Set("renderScale", renderScale)
	return a
}

// RootContainer returns the RootContainer property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#rootcontainer
func (a *AdvancedDynamicTexture) RootContainer() *Container {
	retVal := a.p.Get("rootContainer")
	return ContainerFromJSObject(retVal, a.ctx)
}

// SetRootContainer sets the RootContainer property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#rootcontainer
func (a *AdvancedDynamicTexture) SetRootContainer(rootContainer *Container) *AdvancedDynamicTexture {
	a.p.Set("rootContainer", rootContainer.JSObject())
	return a
}

// UseInvalidateRectOptimization returns the UseInvalidateRectOptimization property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#useinvalidaterectoptimization
func (a *AdvancedDynamicTexture) UseInvalidateRectOptimization() bool {
	retVal := a.p.Get("useInvalidateRectOptimization")
	return retVal.Bool()
}

// SetUseInvalidateRectOptimization sets the UseInvalidateRectOptimization property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#useinvalidaterectoptimization
func (a *AdvancedDynamicTexture) SetUseInvalidateRectOptimization(useInvalidateRectOptimization bool) *AdvancedDynamicTexture {
	a.p.Set("useInvalidateRectOptimization", useInvalidateRectOptimization)
	return a
}

// UseSmallestIdeal returns the UseSmallestIdeal property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#usesmallestideal
func (a *AdvancedDynamicTexture) UseSmallestIdeal() bool {
	retVal := a.p.Get("useSmallestIdeal")
	return retVal.Bool()
}

// SetUseSmallestIdeal sets the UseSmallestIdeal property of class AdvancedDynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictexture#usesmallestideal
func (a *AdvancedDynamicTexture) SetUseSmallestIdeal(useSmallestIdeal bool) *AdvancedDynamicTexture {
	a.p.Set("useSmallestIdeal", useSmallestIdeal)
	return a
}
