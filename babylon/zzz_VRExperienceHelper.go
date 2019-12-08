// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRExperienceHelper represents a babylon.js VRExperienceHelper.
// Helps to quickly add VR support to an existing scene.
// See <a href="http://doc.babylonjs.com/how_to/webvr_helper">http://doc.babylonjs.com/how_to/webvr_helper</a>
type VRExperienceHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VRExperienceHelper) JSObject() js.Value { return v.p }

// VRExperienceHelper returns a VRExperienceHelper JavaScript class.
func (ba *Babylon) VRExperienceHelper() *VRExperienceHelper {
	p := ba.ctx.Get("VRExperienceHelper")
	return VRExperienceHelperFromJSObject(p, ba.ctx)
}

// VRExperienceHelperFromJSObject returns a wrapped VRExperienceHelper JavaScript class.
func VRExperienceHelperFromJSObject(p js.Value, ctx js.Value) *VRExperienceHelper {
	return &VRExperienceHelper{p: p, ctx: ctx}
}

// VRExperienceHelperArrayToJSArray returns a JavaScript Array for the wrapped array.
func VRExperienceHelperArrayToJSArray(array []*VRExperienceHelper) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVRExperienceHelperOpts contains optional parameters for NewVRExperienceHelper.
type NewVRExperienceHelperOpts struct {
	WebVROptions js.Value
}

// NewVRExperienceHelper returns a new VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper
func (ba *Babylon) NewVRExperienceHelper(scene *Scene, opts *NewVRExperienceHelperOpts) *VRExperienceHelper {
	if opts == nil {
		opts = &NewVRExperienceHelperOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	args = append(args, opts.WebVROptions)

	p := ba.ctx.Get("VRExperienceHelper").New(args...)
	return VRExperienceHelperFromJSObject(p, ba.ctx)
}

// AddFloorMesh calls the AddFloorMesh method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#addfloormesh
func (v *VRExperienceHelper) AddFloorMesh(floorMesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	if floorMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, floorMesh.JSObject())
	}

	v.p.Call("addFloorMesh", args...)
}

// ChangeGazeColor calls the ChangeGazeColor method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#changegazecolor
func (v *VRExperienceHelper) ChangeGazeColor(color *Color3) {

	args := make([]interface{}, 0, 1+0)

	if color == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, color.JSObject())
	}

	v.p.Call("changeGazeColor", args...)
}

// ChangeLaserColor calls the ChangeLaserColor method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#changelasercolor
func (v *VRExperienceHelper) ChangeLaserColor(color *Color3) {

	args := make([]interface{}, 0, 1+0)

	if color == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, color.JSObject())
	}

	v.p.Call("changeLaserColor", args...)
}

// Dispose calls the Dispose method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#dispose
func (v *VRExperienceHelper) Dispose() {

	v.p.Call("dispose")
}

// EnableInteractions calls the EnableInteractions method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#enableinteractions
func (v *VRExperienceHelper) EnableInteractions() {

	v.p.Call("enableInteractions")
}

// VRExperienceHelperEnableTeleportationOpts contains optional parameters for VRExperienceHelper.EnableTeleportation.
type VRExperienceHelperEnableTeleportationOpts struct {
	VrTeleportationOptions *VRTeleportationOptions
}

// EnableTeleportation calls the EnableTeleportation method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#enableteleportation
func (v *VRExperienceHelper) EnableTeleportation(opts *VRExperienceHelperEnableTeleportationOpts) {
	if opts == nil {
		opts = &VRExperienceHelperEnableTeleportationOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.VrTeleportationOptions == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.VrTeleportationOptions.JSObject())
	}

	v.p.Call("enableTeleportation", args...)
}

// EnterVR calls the EnterVR method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#entervr
func (v *VRExperienceHelper) EnterVR() {

	v.p.Call("enterVR")
}

// ExitVR calls the ExitVR method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#exitvr
func (v *VRExperienceHelper) ExitVR() {

	v.p.Call("exitVR")
}

// GetClassName calls the GetClassName method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#getclassname
func (v *VRExperienceHelper) GetClassName() string {

	retVal := v.p.Call("getClassName")
	return retVal.String()
}

// RemoveFloorMesh calls the RemoveFloorMesh method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#removefloormesh
func (v *VRExperienceHelper) RemoveFloorMesh(floorMesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	if floorMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, floorMesh.JSObject())
	}

	v.p.Call("removeFloorMesh", args...)
}

// TeleportCamera calls the TeleportCamera method on the VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportcamera
func (v *VRExperienceHelper) TeleportCamera(location *Vector3) {

	args := make([]interface{}, 0, 1+0)

	if location == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, location.JSObject())
	}

	v.p.Call("teleportCamera", args...)
}

// CurrentVRCamera returns the CurrentVRCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#currentvrcamera
func (v *VRExperienceHelper) CurrentVRCamera() *Camera {
	retVal := v.p.Get("currentVRCamera")
	return CameraFromJSObject(retVal, v.ctx)
}

// SetCurrentVRCamera sets the CurrentVRCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#currentvrcamera
func (v *VRExperienceHelper) SetCurrentVRCamera(currentVRCamera *Camera) *VRExperienceHelper {
	v.p.Set("currentVRCamera", currentVRCamera.JSObject())
	return v
}

// DeviceOrientationCamera returns the DeviceOrientationCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#deviceorientationcamera
func (v *VRExperienceHelper) DeviceOrientationCamera() *DeviceOrientationCamera {
	retVal := v.p.Get("deviceOrientationCamera")
	return DeviceOrientationCameraFromJSObject(retVal, v.ctx)
}

// SetDeviceOrientationCamera sets the DeviceOrientationCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#deviceorientationcamera
func (v *VRExperienceHelper) SetDeviceOrientationCamera(deviceOrientationCamera *DeviceOrientationCamera) *VRExperienceHelper {
	v.p.Set("deviceOrientationCamera", deviceOrientationCamera.JSObject())
	return v
}

// DisplayGaze returns the DisplayGaze property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#displaygaze
func (v *VRExperienceHelper) DisplayGaze() bool {
	retVal := v.p.Get("displayGaze")
	return retVal.Bool()
}

// SetDisplayGaze sets the DisplayGaze property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#displaygaze
func (v *VRExperienceHelper) SetDisplayGaze(displayGaze bool) *VRExperienceHelper {
	v.p.Set("displayGaze", displayGaze)
	return v
}

// DisplayLaserPointer returns the DisplayLaserPointer property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#displaylaserpointer
func (v *VRExperienceHelper) DisplayLaserPointer() bool {
	retVal := v.p.Get("displayLaserPointer")
	return retVal.Bool()
}

// SetDisplayLaserPointer sets the DisplayLaserPointer property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#displaylaserpointer
func (v *VRExperienceHelper) SetDisplayLaserPointer(displayLaserPointer bool) *VRExperienceHelper {
	v.p.Set("displayLaserPointer", displayLaserPointer)
	return v
}

// EnableGazeEvenWhenNoPointerLock returns the EnableGazeEvenWhenNoPointerLock property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#enablegazeevenwhennopointerlock
func (v *VRExperienceHelper) EnableGazeEvenWhenNoPointerLock() bool {
	retVal := v.p.Get("enableGazeEvenWhenNoPointerLock")
	return retVal.Bool()
}

// SetEnableGazeEvenWhenNoPointerLock sets the EnableGazeEvenWhenNoPointerLock property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#enablegazeevenwhennopointerlock
func (v *VRExperienceHelper) SetEnableGazeEvenWhenNoPointerLock(enableGazeEvenWhenNoPointerLock bool) *VRExperienceHelper {
	v.p.Set("enableGazeEvenWhenNoPointerLock", enableGazeEvenWhenNoPointerLock)
	return v
}

// ExitVROnDoubleTap returns the ExitVROnDoubleTap property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#exitvrondoubletap
func (v *VRExperienceHelper) ExitVROnDoubleTap() bool {
	retVal := v.p.Get("exitVROnDoubleTap")
	return retVal.Bool()
}

// SetExitVROnDoubleTap sets the ExitVROnDoubleTap property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#exitvrondoubletap
func (v *VRExperienceHelper) SetExitVROnDoubleTap(exitVROnDoubleTap bool) *VRExperienceHelper {
	v.p.Set("exitVROnDoubleTap", exitVROnDoubleTap)
	return v
}

// GazeTrackerMesh returns the GazeTrackerMesh property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#gazetrackermesh
func (v *VRExperienceHelper) GazeTrackerMesh() *Mesh {
	retVal := v.p.Get("gazeTrackerMesh")
	return MeshFromJSObject(retVal, v.ctx)
}

// SetGazeTrackerMesh sets the GazeTrackerMesh property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#gazetrackermesh
func (v *VRExperienceHelper) SetGazeTrackerMesh(gazeTrackerMesh *Mesh) *VRExperienceHelper {
	v.p.Set("gazeTrackerMesh", gazeTrackerMesh.JSObject())
	return v
}

// IsInVRMode returns the IsInVRMode property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#isinvrmode
func (v *VRExperienceHelper) IsInVRMode() bool {
	retVal := v.p.Get("isInVRMode")
	return retVal.Bool()
}

// SetIsInVRMode sets the IsInVRMode property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#isinvrmode
func (v *VRExperienceHelper) SetIsInVRMode(isInVRMode bool) *VRExperienceHelper {
	v.p.Set("isInVRMode", isInVRMode)
	return v
}

// LeftControllerGazeTrackerMesh returns the LeftControllerGazeTrackerMesh property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#leftcontrollergazetrackermesh
func (v *VRExperienceHelper) LeftControllerGazeTrackerMesh() *Mesh {
	retVal := v.p.Get("leftControllerGazeTrackerMesh")
	return MeshFromJSObject(retVal, v.ctx)
}

// SetLeftControllerGazeTrackerMesh sets the LeftControllerGazeTrackerMesh property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#leftcontrollergazetrackermesh
func (v *VRExperienceHelper) SetLeftControllerGazeTrackerMesh(leftControllerGazeTrackerMesh *Mesh) *VRExperienceHelper {
	v.p.Set("leftControllerGazeTrackerMesh", leftControllerGazeTrackerMesh.JSObject())
	return v
}

// MeshSelectionPredicate returns the MeshSelectionPredicate property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#meshselectionpredicate
func (v *VRExperienceHelper) MeshSelectionPredicate() js.Value {
	retVal := v.p.Get("meshSelectionPredicate")
	return retVal
}

// SetMeshSelectionPredicate sets the MeshSelectionPredicate property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#meshselectionpredicate
func (v *VRExperienceHelper) SetMeshSelectionPredicate(meshSelectionPredicate JSFunc) *VRExperienceHelper {
	v.p.Set("meshSelectionPredicate", js.FuncOf(meshSelectionPredicate))
	return v
}

// OnAfterCameraTeleport returns the OnAfterCameraTeleport property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onaftercamerateleport
func (v *VRExperienceHelper) OnAfterCameraTeleport() *Observable {
	retVal := v.p.Get("onAfterCameraTeleport")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnAfterCameraTeleport sets the OnAfterCameraTeleport property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onaftercamerateleport
func (v *VRExperienceHelper) SetOnAfterCameraTeleport(onAfterCameraTeleport *Observable) *VRExperienceHelper {
	v.p.Set("onAfterCameraTeleport", onAfterCameraTeleport.JSObject())
	return v
}

// OnAfterEnteringVRObservable returns the OnAfterEnteringVRObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onafterenteringvrobservable
func (v *VRExperienceHelper) OnAfterEnteringVRObservable() *Observable {
	retVal := v.p.Get("onAfterEnteringVRObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnAfterEnteringVRObservable sets the OnAfterEnteringVRObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onafterenteringvrobservable
func (v *VRExperienceHelper) SetOnAfterEnteringVRObservable(onAfterEnteringVRObservable *Observable) *VRExperienceHelper {
	v.p.Set("onAfterEnteringVRObservable", onAfterEnteringVRObservable.JSObject())
	return v
}

// OnBeforeCameraTeleport returns the OnBeforeCameraTeleport property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onbeforecamerateleport
func (v *VRExperienceHelper) OnBeforeCameraTeleport() *Observable {
	retVal := v.p.Get("onBeforeCameraTeleport")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnBeforeCameraTeleport sets the OnBeforeCameraTeleport property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onbeforecamerateleport
func (v *VRExperienceHelper) SetOnBeforeCameraTeleport(onBeforeCameraTeleport *Observable) *VRExperienceHelper {
	v.p.Set("onBeforeCameraTeleport", onBeforeCameraTeleport.JSObject())
	return v
}

// OnControllerMeshLoaded returns the OnControllerMeshLoaded property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#oncontrollermeshloaded
func (v *VRExperienceHelper) OnControllerMeshLoaded() *Observable {
	retVal := v.p.Get("onControllerMeshLoaded")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnControllerMeshLoaded sets the OnControllerMeshLoaded property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#oncontrollermeshloaded
func (v *VRExperienceHelper) SetOnControllerMeshLoaded(onControllerMeshLoaded *Observable) *VRExperienceHelper {
	v.p.Set("onControllerMeshLoaded", onControllerMeshLoaded.JSObject())
	return v
}

// OnControllerMeshLoadedObservable returns the OnControllerMeshLoadedObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#oncontrollermeshloadedobservable
func (v *VRExperienceHelper) OnControllerMeshLoadedObservable() *Observable {
	retVal := v.p.Get("onControllerMeshLoadedObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnControllerMeshLoadedObservable sets the OnControllerMeshLoadedObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#oncontrollermeshloadedobservable
func (v *VRExperienceHelper) SetOnControllerMeshLoadedObservable(onControllerMeshLoadedObservable *Observable) *VRExperienceHelper {
	v.p.Set("onControllerMeshLoadedObservable", onControllerMeshLoadedObservable.JSObject())
	return v
}

// OnEnteringVR returns the OnEnteringVR property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onenteringvr
func (v *VRExperienceHelper) OnEnteringVR() *Observable {
	retVal := v.p.Get("onEnteringVR")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnEnteringVR sets the OnEnteringVR property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onenteringvr
func (v *VRExperienceHelper) SetOnEnteringVR(onEnteringVR *Observable) *VRExperienceHelper {
	v.p.Set("onEnteringVR", onEnteringVR.JSObject())
	return v
}

// OnEnteringVRObservable returns the OnEnteringVRObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onenteringvrobservable
func (v *VRExperienceHelper) OnEnteringVRObservable() *Observable {
	retVal := v.p.Get("onEnteringVRObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnEnteringVRObservable sets the OnEnteringVRObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onenteringvrobservable
func (v *VRExperienceHelper) SetOnEnteringVRObservable(onEnteringVRObservable *Observable) *VRExperienceHelper {
	v.p.Set("onEnteringVRObservable", onEnteringVRObservable.JSObject())
	return v
}

// OnExitingVR returns the OnExitingVR property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onexitingvr
func (v *VRExperienceHelper) OnExitingVR() *Observable {
	retVal := v.p.Get("onExitingVR")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnExitingVR sets the OnExitingVR property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onexitingvr
func (v *VRExperienceHelper) SetOnExitingVR(onExitingVR *Observable) *VRExperienceHelper {
	v.p.Set("onExitingVR", onExitingVR.JSObject())
	return v
}

// OnExitingVRObservable returns the OnExitingVRObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onexitingvrobservable
func (v *VRExperienceHelper) OnExitingVRObservable() *Observable {
	retVal := v.p.Get("onExitingVRObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnExitingVRObservable sets the OnExitingVRObservable property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onexitingvrobservable
func (v *VRExperienceHelper) SetOnExitingVRObservable(onExitingVRObservable *Observable) *VRExperienceHelper {
	v.p.Set("onExitingVRObservable", onExitingVRObservable.JSObject())
	return v
}

// OnMeshSelectedWithController returns the OnMeshSelectedWithController property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onmeshselectedwithcontroller
func (v *VRExperienceHelper) OnMeshSelectedWithController() *Observable {
	retVal := v.p.Get("onMeshSelectedWithController")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnMeshSelectedWithController sets the OnMeshSelectedWithController property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onmeshselectedwithcontroller
func (v *VRExperienceHelper) SetOnMeshSelectedWithController(onMeshSelectedWithController *Observable) *VRExperienceHelper {
	v.p.Set("onMeshSelectedWithController", onMeshSelectedWithController.JSObject())
	return v
}

// OnNewMeshPicked returns the OnNewMeshPicked property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onnewmeshpicked
func (v *VRExperienceHelper) OnNewMeshPicked() *Observable {
	retVal := v.p.Get("onNewMeshPicked")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnNewMeshPicked sets the OnNewMeshPicked property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onnewmeshpicked
func (v *VRExperienceHelper) SetOnNewMeshPicked(onNewMeshPicked *Observable) *VRExperienceHelper {
	v.p.Set("onNewMeshPicked", onNewMeshPicked.JSObject())
	return v
}

// OnNewMeshSelected returns the OnNewMeshSelected property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onnewmeshselected
func (v *VRExperienceHelper) OnNewMeshSelected() *Observable {
	retVal := v.p.Get("onNewMeshSelected")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnNewMeshSelected sets the OnNewMeshSelected property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onnewmeshselected
func (v *VRExperienceHelper) SetOnNewMeshSelected(onNewMeshSelected *Observable) *VRExperienceHelper {
	v.p.Set("onNewMeshSelected", onNewMeshSelected.JSObject())
	return v
}

// OnSelectedMeshUnselected returns the OnSelectedMeshUnselected property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onselectedmeshunselected
func (v *VRExperienceHelper) OnSelectedMeshUnselected() *Observable {
	retVal := v.p.Get("onSelectedMeshUnselected")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnSelectedMeshUnselected sets the OnSelectedMeshUnselected property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#onselectedmeshunselected
func (v *VRExperienceHelper) SetOnSelectedMeshUnselected(onSelectedMeshUnselected *Observable) *VRExperienceHelper {
	v.p.Set("onSelectedMeshUnselected", onSelectedMeshUnselected.JSObject())
	return v
}

// Position returns the Position property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#position
func (v *VRExperienceHelper) Position() *Vector3 {
	retVal := v.p.Get("position")
	return Vector3FromJSObject(retVal, v.ctx)
}

// SetPosition sets the Position property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#position
func (v *VRExperienceHelper) SetPosition(position *Vector3) *VRExperienceHelper {
	v.p.Set("position", position.JSObject())
	return v
}

// RaySelectionPredicate returns the RaySelectionPredicate property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#rayselectionpredicate
func (v *VRExperienceHelper) RaySelectionPredicate() js.Value {
	retVal := v.p.Get("raySelectionPredicate")
	return retVal
}

// SetRaySelectionPredicate sets the RaySelectionPredicate property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#rayselectionpredicate
func (v *VRExperienceHelper) SetRaySelectionPredicate(raySelectionPredicate JSFunc) *VRExperienceHelper {
	v.p.Set("raySelectionPredicate", js.FuncOf(raySelectionPredicate))
	return v
}

// RequestPointerLockOnFullScreen returns the RequestPointerLockOnFullScreen property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#requestpointerlockonfullscreen
func (v *VRExperienceHelper) RequestPointerLockOnFullScreen() bool {
	retVal := v.p.Get("requestPointerLockOnFullScreen")
	return retVal.Bool()
}

// SetRequestPointerLockOnFullScreen sets the RequestPointerLockOnFullScreen property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#requestpointerlockonfullscreen
func (v *VRExperienceHelper) SetRequestPointerLockOnFullScreen(requestPointerLockOnFullScreen bool) *VRExperienceHelper {
	v.p.Set("requestPointerLockOnFullScreen", requestPointerLockOnFullScreen)
	return v
}

// RightControllerGazeTrackerMesh returns the RightControllerGazeTrackerMesh property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#rightcontrollergazetrackermesh
func (v *VRExperienceHelper) RightControllerGazeTrackerMesh() *Mesh {
	retVal := v.p.Get("rightControllerGazeTrackerMesh")
	return MeshFromJSObject(retVal, v.ctx)
}

// SetRightControllerGazeTrackerMesh sets the RightControllerGazeTrackerMesh property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#rightcontrollergazetrackermesh
func (v *VRExperienceHelper) SetRightControllerGazeTrackerMesh(rightControllerGazeTrackerMesh *Mesh) *VRExperienceHelper {
	v.p.Set("rightControllerGazeTrackerMesh", rightControllerGazeTrackerMesh.JSObject())
	return v
}

// TELEPORTATIONMODE_CONSTANTSPEED returns the TELEPORTATIONMODE_CONSTANTSPEED property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationmode_constantspeed
func (v *VRExperienceHelper) TELEPORTATIONMODE_CONSTANTSPEED() float64 {
	retVal := v.p.Get("TELEPORTATIONMODE_CONSTANTSPEED")
	return retVal.Float()
}

// SetTELEPORTATIONMODE_CONSTANTSPEED sets the TELEPORTATIONMODE_CONSTANTSPEED property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationmode_constantspeed
func (v *VRExperienceHelper) SetTELEPORTATIONMODE_CONSTANTSPEED(TELEPORTATIONMODE_CONSTANTSPEED float64) *VRExperienceHelper {
	v.p.Set("TELEPORTATIONMODE_CONSTANTSPEED", TELEPORTATIONMODE_CONSTANTSPEED)
	return v
}

// TELEPORTATIONMODE_CONSTANTTIME returns the TELEPORTATIONMODE_CONSTANTTIME property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationmode_constanttime
func (v *VRExperienceHelper) TELEPORTATIONMODE_CONSTANTTIME() float64 {
	retVal := v.p.Get("TELEPORTATIONMODE_CONSTANTTIME")
	return retVal.Float()
}

// SetTELEPORTATIONMODE_CONSTANTTIME sets the TELEPORTATIONMODE_CONSTANTTIME property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationmode_constanttime
func (v *VRExperienceHelper) SetTELEPORTATIONMODE_CONSTANTTIME(TELEPORTATIONMODE_CONSTANTTIME float64) *VRExperienceHelper {
	v.p.Set("TELEPORTATIONMODE_CONSTANTTIME", TELEPORTATIONMODE_CONSTANTTIME)
	return v
}

// TeleportationEnabled returns the TeleportationEnabled property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationenabled
func (v *VRExperienceHelper) TeleportationEnabled() bool {
	retVal := v.p.Get("teleportationEnabled")
	return retVal.Bool()
}

// SetTeleportationEnabled sets the TeleportationEnabled property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationenabled
func (v *VRExperienceHelper) SetTeleportationEnabled(teleportationEnabled bool) *VRExperienceHelper {
	v.p.Set("teleportationEnabled", teleportationEnabled)
	return v
}

// TeleportationTarget returns the TeleportationTarget property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationtarget
func (v *VRExperienceHelper) TeleportationTarget() *Mesh {
	retVal := v.p.Get("teleportationTarget")
	return MeshFromJSObject(retVal, v.ctx)
}

// SetTeleportationTarget sets the TeleportationTarget property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#teleportationtarget
func (v *VRExperienceHelper) SetTeleportationTarget(teleportationTarget *Mesh) *VRExperienceHelper {
	v.p.Set("teleportationTarget", teleportationTarget.JSObject())
	return v
}

// UpdateControllerLaserColor returns the UpdateControllerLaserColor property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#updatecontrollerlasercolor
func (v *VRExperienceHelper) UpdateControllerLaserColor() bool {
	retVal := v.p.Get("updateControllerLaserColor")
	return retVal.Bool()
}

// SetUpdateControllerLaserColor sets the UpdateControllerLaserColor property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#updatecontrollerlasercolor
func (v *VRExperienceHelper) SetUpdateControllerLaserColor(updateControllerLaserColor bool) *VRExperienceHelper {
	v.p.Set("updateControllerLaserColor", updateControllerLaserColor)
	return v
}

// UpdateGazeTrackerColor returns the UpdateGazeTrackerColor property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#updategazetrackercolor
func (v *VRExperienceHelper) UpdateGazeTrackerColor() bool {
	retVal := v.p.Get("updateGazeTrackerColor")
	return retVal.Bool()
}

// SetUpdateGazeTrackerColor sets the UpdateGazeTrackerColor property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#updategazetrackercolor
func (v *VRExperienceHelper) SetUpdateGazeTrackerColor(updateGazeTrackerColor bool) *VRExperienceHelper {
	v.p.Set("updateGazeTrackerColor", updateGazeTrackerColor)
	return v
}

// UpdateGazeTrackerScale returns the UpdateGazeTrackerScale property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#updategazetrackerscale
func (v *VRExperienceHelper) UpdateGazeTrackerScale() bool {
	retVal := v.p.Get("updateGazeTrackerScale")
	return retVal.Bool()
}

// SetUpdateGazeTrackerScale sets the UpdateGazeTrackerScale property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#updategazetrackerscale
func (v *VRExperienceHelper) SetUpdateGazeTrackerScale(updateGazeTrackerScale bool) *VRExperienceHelper {
	v.p.Set("updateGazeTrackerScale", updateGazeTrackerScale)
	return v
}

// VrButton returns the VrButton property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#vrbutton
func (v *VRExperienceHelper) VrButton() js.Value {
	retVal := v.p.Get("vrButton")
	return retVal
}

// SetVrButton sets the VrButton property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#vrbutton
func (v *VRExperienceHelper) SetVrButton(vrButton js.Value) *VRExperienceHelper {
	v.p.Set("vrButton", vrButton)
	return v
}

// VrDeviceOrientationCamera returns the VrDeviceOrientationCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#vrdeviceorientationcamera
func (v *VRExperienceHelper) VrDeviceOrientationCamera() *VRDeviceOrientationFreeCamera {
	retVal := v.p.Get("vrDeviceOrientationCamera")
	return VRDeviceOrientationFreeCameraFromJSObject(retVal, v.ctx)
}

// SetVrDeviceOrientationCamera sets the VrDeviceOrientationCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#vrdeviceorientationcamera
func (v *VRExperienceHelper) SetVrDeviceOrientationCamera(vrDeviceOrientationCamera *VRDeviceOrientationFreeCamera) *VRExperienceHelper {
	v.p.Set("vrDeviceOrientationCamera", vrDeviceOrientationCamera.JSObject())
	return v
}

// WebVRCamera returns the WebVRCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#webvrcamera
func (v *VRExperienceHelper) WebVRCamera() *WebVRFreeCamera {
	retVal := v.p.Get("webVRCamera")
	return WebVRFreeCameraFromJSObject(retVal, v.ctx)
}

// SetWebVRCamera sets the WebVRCamera property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#webvrcamera
func (v *VRExperienceHelper) SetWebVRCamera(webVRCamera *WebVRFreeCamera) *VRExperienceHelper {
	v.p.Set("webVRCamera", webVRCamera.JSObject())
	return v
}

// WebVROptions returns the WebVROptions property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#webvroptions
func (v *VRExperienceHelper) WebVROptions() js.Value {
	retVal := v.p.Get("webVROptions")
	return retVal
}

// SetWebVROptions sets the WebVROptions property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#webvroptions
func (v *VRExperienceHelper) SetWebVROptions(webVROptions js.Value) *VRExperienceHelper {
	v.p.Set("webVROptions", webVROptions)
	return v
}

// Xr returns the Xr property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#xr
func (v *VRExperienceHelper) Xr() *WebXRDefaultExperience {
	retVal := v.p.Get("xr")
	return WebXRDefaultExperienceFromJSObject(retVal, v.ctx)
}

// SetXr sets the Xr property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#xr
func (v *VRExperienceHelper) SetXr(xr *WebXRDefaultExperience) *VRExperienceHelper {
	v.p.Set("xr", xr.JSObject())
	return v
}

// XrTestDone returns the XrTestDone property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#xrtestdone
func (v *VRExperienceHelper) XrTestDone() bool {
	retVal := v.p.Get("xrTestDone")
	return retVal.Bool()
}

// SetXrTestDone sets the XrTestDone property of class VRExperienceHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper#xrtestdone
func (v *VRExperienceHelper) SetXrTestDone(xrTestDone bool) *VRExperienceHelper {
	v.p.Set("xrTestDone", xrTestDone)
	return v
}
