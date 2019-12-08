// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PoseEnabledController represents a babylon.js PoseEnabledController.
// Defines the PoseEnabledController object that contains state of a vr capable controller
type PoseEnabledController struct {
	*Gamepad
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PoseEnabledController) JSObject() js.Value { return p.p }

// PoseEnabledController returns a PoseEnabledController JavaScript class.
func (ba *Babylon) PoseEnabledController() *PoseEnabledController {
	p := ba.ctx.Get("PoseEnabledController")
	return PoseEnabledControllerFromJSObject(p, ba.ctx)
}

// PoseEnabledControllerFromJSObject returns a wrapped PoseEnabledController JavaScript class.
func PoseEnabledControllerFromJSObject(p js.Value, ctx js.Value) *PoseEnabledController {
	return &PoseEnabledController{Gamepad: GamepadFromJSObject(p, ctx), ctx: ctx}
}

// PoseEnabledControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func PoseEnabledControllerArrayToJSArray(array []*PoseEnabledController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPoseEnabledController returns a new PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller
func (ba *Babylon) NewPoseEnabledController(browserGamepad JSObject) *PoseEnabledController {

	args := make([]interface{}, 0, 1+0)

	args = append(args, browserGamepad.JSObject())

	p := ba.ctx.Get("PoseEnabledController").New(args...)
	return PoseEnabledControllerFromJSObject(p, ba.ctx)
}

// AttachToMesh calls the AttachToMesh method on the PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#attachtomesh
func (p *PoseEnabledController) AttachToMesh(mesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	p.p.Call("attachToMesh", args...)
}

// AttachToPoseControlledCamera calls the AttachToPoseControlledCamera method on the PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#attachtoposecontrolledcamera
func (p *PoseEnabledController) AttachToPoseControlledCamera(camera *TargetCamera) {

	args := make([]interface{}, 0, 1+0)

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	p.p.Call("attachToPoseControlledCamera", args...)
}

// Dispose calls the Dispose method on the PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#dispose
func (p *PoseEnabledController) Dispose() {

	p.p.Call("dispose")
}

// PoseEnabledControllerGetForwardRayOpts contains optional parameters for PoseEnabledController.GetForwardRay.
type PoseEnabledControllerGetForwardRayOpts struct {
	Length *float64
}

// GetForwardRay calls the GetForwardRay method on the PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#getforwardray
func (p *PoseEnabledController) GetForwardRay(opts *PoseEnabledControllerGetForwardRayOpts) *Ray {
	if opts == nil {
		opts = &PoseEnabledControllerGetForwardRayOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Length == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Length)
	}

	retVal := p.p.Call("getForwardRay", args...)
	return RayFromJSObject(retVal, p.ctx)
}

// Update calls the Update method on the PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#update
func (p *PoseEnabledController) Update() {

	p.p.Call("update")
}

// UpdateFromDevice calls the UpdateFromDevice method on the PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#updatefromdevice
func (p *PoseEnabledController) UpdateFromDevice(poseData js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, poseData)

	p.p.Call("updateFromDevice", args...)
}

// ControllerType returns the ControllerType property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#controllertype
func (p *PoseEnabledController) ControllerType() js.Value {
	retVal := p.p.Get("controllerType")
	return retVal
}

// SetControllerType sets the ControllerType property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#controllertype
func (p *PoseEnabledController) SetControllerType(controllerType js.Value) *PoseEnabledController {
	p.p.Set("controllerType", controllerType)
	return p
}

// DevicePosition returns the DevicePosition property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#deviceposition
func (p *PoseEnabledController) DevicePosition() *Vector3 {
	retVal := p.p.Get("devicePosition")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetDevicePosition sets the DevicePosition property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#deviceposition
func (p *PoseEnabledController) SetDevicePosition(devicePosition *Vector3) *PoseEnabledController {
	p.p.Set("devicePosition", devicePosition.JSObject())
	return p
}

// DeviceRotationQuaternion returns the DeviceRotationQuaternion property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#devicerotationquaternion
func (p *PoseEnabledController) DeviceRotationQuaternion() *Quaternion {
	retVal := p.p.Get("deviceRotationQuaternion")
	return QuaternionFromJSObject(retVal, p.ctx)
}

// SetDeviceRotationQuaternion sets the DeviceRotationQuaternion property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#devicerotationquaternion
func (p *PoseEnabledController) SetDeviceRotationQuaternion(deviceRotationQuaternion *Quaternion) *PoseEnabledController {
	p.p.Set("deviceRotationQuaternion", deviceRotationQuaternion.JSObject())
	return p
}

// DeviceScaleFactor returns the DeviceScaleFactor property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#devicescalefactor
func (p *PoseEnabledController) DeviceScaleFactor() float64 {
	retVal := p.p.Get("deviceScaleFactor")
	return retVal.Float()
}

// SetDeviceScaleFactor sets the DeviceScaleFactor property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#devicescalefactor
func (p *PoseEnabledController) SetDeviceScaleFactor(deviceScaleFactor float64) *PoseEnabledController {
	p.p.Set("deviceScaleFactor", deviceScaleFactor)
	return p
}

// IsXR returns the IsXR property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#isxr
func (p *PoseEnabledController) IsXR() bool {
	retVal := p.p.Get("isXR")
	return retVal.Bool()
}

// SetIsXR sets the IsXR property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#isxr
func (p *PoseEnabledController) SetIsXR(isXR bool) *PoseEnabledController {
	p.p.Set("isXR", isXR)
	return p
}

// Mesh returns the Mesh property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#mesh
func (p *PoseEnabledController) Mesh() *AbstractMesh {
	retVal := p.p.Get("mesh")
	return AbstractMeshFromJSObject(retVal, p.ctx)
}

// SetMesh sets the Mesh property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#mesh
func (p *PoseEnabledController) SetMesh(mesh *AbstractMesh) *PoseEnabledController {
	p.p.Set("mesh", mesh.JSObject())
	return p
}

// POINTING_POSE returns the POINTING_POSE property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#pointing_pose
func (p *PoseEnabledController) POINTING_POSE() string {
	retVal := p.p.Get("POINTING_POSE")
	return retVal.String()
}

// SetPOINTING_POSE sets the POINTING_POSE property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#pointing_pose
func (p *PoseEnabledController) SetPOINTING_POSE(POINTING_POSE string) *PoseEnabledController {
	p.p.Set("POINTING_POSE", POINTING_POSE)
	return p
}

// Position returns the Position property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#position
func (p *PoseEnabledController) Position() *Vector3 {
	retVal := p.p.Get("position")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetPosition sets the Position property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#position
func (p *PoseEnabledController) SetPosition(position *Vector3) *PoseEnabledController {
	p.p.Set("position", position.JSObject())
	return p
}

// RawPose returns the RawPose property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#rawpose
func (p *PoseEnabledController) RawPose() js.Value {
	retVal := p.p.Get("rawPose")
	return retVal
}

// SetRawPose sets the RawPose property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#rawpose
func (p *PoseEnabledController) SetRawPose(rawPose js.Value) *PoseEnabledController {
	p.p.Set("rawPose", rawPose)
	return p
}

// RotationQuaternion returns the RotationQuaternion property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#rotationquaternion
func (p *PoseEnabledController) RotationQuaternion() *Quaternion {
	retVal := p.p.Get("rotationQuaternion")
	return QuaternionFromJSObject(retVal, p.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class PoseEnabledController.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller#rotationquaternion
func (p *PoseEnabledController) SetRotationQuaternion(rotationQuaternion *Quaternion) *PoseEnabledController {
	p.p.Set("rotationQuaternion", rotationQuaternion.JSObject())
	return p
}
