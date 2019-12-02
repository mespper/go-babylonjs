// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRExperienceHelperOptions represents a babylon.js VRExperienceHelperOptions.
// Options to modify the vr experience helper&#39;s behavior.
type VRExperienceHelperOptions struct {
	*WebVROptions
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VRExperienceHelperOptions) JSObject() js.Value { return v.p }

// VRExperienceHelperOptions returns a VRExperienceHelperOptions JavaScript class.
func (ba *Babylon) VRExperienceHelperOptions() *VRExperienceHelperOptions {
	p := ba.ctx.Get("VRExperienceHelperOptions")
	return VRExperienceHelperOptionsFromJSObject(p, ba.ctx)
}

// VRExperienceHelperOptionsFromJSObject returns a wrapped VRExperienceHelperOptions JavaScript class.
func VRExperienceHelperOptionsFromJSObject(p js.Value, ctx js.Value) *VRExperienceHelperOptions {
	return &VRExperienceHelperOptions{WebVROptions: WebVROptionsFromJSObject(p, ctx), ctx: ctx}
}

// VRExperienceHelperOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func VRExperienceHelperOptionsArrayToJSArray(array []*VRExperienceHelperOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CreateDeviceOrientationCamera returns the CreateDeviceOrientationCamera property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#createdeviceorientationcamera
func (v *VRExperienceHelperOptions) CreateDeviceOrientationCamera() bool {
	retVal := v.p.Get("createDeviceOrientationCamera")
	return retVal.Bool()
}

// SetCreateDeviceOrientationCamera sets the CreateDeviceOrientationCamera property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#createdeviceorientationcamera
func (v *VRExperienceHelperOptions) SetCreateDeviceOrientationCamera(createDeviceOrientationCamera bool) *VRExperienceHelperOptions {
	v.p.Set("createDeviceOrientationCamera", createDeviceOrientationCamera)
	return v
}

// CreateFallbackVRDeviceOrientationFreeCamera returns the CreateFallbackVRDeviceOrientationFreeCamera property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#createfallbackvrdeviceorientationfreecamera
func (v *VRExperienceHelperOptions) CreateFallbackVRDeviceOrientationFreeCamera() bool {
	retVal := v.p.Get("createFallbackVRDeviceOrientationFreeCamera")
	return retVal.Bool()
}

// SetCreateFallbackVRDeviceOrientationFreeCamera sets the CreateFallbackVRDeviceOrientationFreeCamera property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#createfallbackvrdeviceorientationfreecamera
func (v *VRExperienceHelperOptions) SetCreateFallbackVRDeviceOrientationFreeCamera(createFallbackVRDeviceOrientationFreeCamera bool) *VRExperienceHelperOptions {
	v.p.Set("createFallbackVRDeviceOrientationFreeCamera", createFallbackVRDeviceOrientationFreeCamera)
	return v
}

// FloorMeshes returns the FloorMeshes property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#floormeshes
func (v *VRExperienceHelperOptions) FloorMeshes() *Mesh {
	retVal := v.p.Get("floorMeshes")
	return MeshFromJSObject(retVal, v.ctx)
}

// SetFloorMeshes sets the FloorMeshes property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#floormeshes
func (v *VRExperienceHelperOptions) SetFloorMeshes(floorMeshes *Mesh) *VRExperienceHelperOptions {
	v.p.Set("floorMeshes", floorMeshes.JSObject())
	return v
}

// LaserToggle returns the LaserToggle property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#lasertoggle
func (v *VRExperienceHelperOptions) LaserToggle() bool {
	retVal := v.p.Get("laserToggle")
	return retVal.Bool()
}

// SetLaserToggle sets the LaserToggle property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#lasertoggle
func (v *VRExperienceHelperOptions) SetLaserToggle(laserToggle bool) *VRExperienceHelperOptions {
	v.p.Set("laserToggle", laserToggle)
	return v
}

// UseXR returns the UseXR property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#usexr
func (v *VRExperienceHelperOptions) UseXR() bool {
	retVal := v.p.Get("useXR")
	return retVal.Bool()
}

// SetUseXR sets the UseXR property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#usexr
func (v *VRExperienceHelperOptions) SetUseXR(useXR bool) *VRExperienceHelperOptions {
	v.p.Set("useXR", useXR)
	return v
}

// VrDeviceOrientationCameraMetrics returns the VrDeviceOrientationCameraMetrics property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#vrdeviceorientationcamerametrics
func (v *VRExperienceHelperOptions) VrDeviceOrientationCameraMetrics() *VRCameraMetrics {
	retVal := v.p.Get("vrDeviceOrientationCameraMetrics")
	return VRCameraMetricsFromJSObject(retVal, v.ctx)
}

// SetVrDeviceOrientationCameraMetrics sets the VrDeviceOrientationCameraMetrics property of class VRExperienceHelperOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelperoptions#vrdeviceorientationcamerametrics
func (v *VRExperienceHelperOptions) SetVrDeviceOrientationCameraMetrics(vrDeviceOrientationCameraMetrics *VRCameraMetrics) *VRExperienceHelperOptions {
	v.p.Set("vrDeviceOrientationCameraMetrics", vrDeviceOrientationCameraMetrics.JSObject())
	return v
}
