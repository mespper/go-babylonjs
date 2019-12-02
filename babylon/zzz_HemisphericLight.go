// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HemisphericLight represents a babylon.js HemisphericLight.
// The HemisphericLight simulates the ambient environment light,
// so the passed direction is the light reflection direction, not the incoming direction.
type HemisphericLight struct {
	*Light
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HemisphericLight) JSObject() js.Value { return h.p }

// HemisphericLight returns a HemisphericLight JavaScript class.
func (ba *Babylon) HemisphericLight() *HemisphericLight {
	p := ba.ctx.Get("HemisphericLight")
	return HemisphericLightFromJSObject(p, ba.ctx)
}

// HemisphericLightFromJSObject returns a wrapped HemisphericLight JavaScript class.
func HemisphericLightFromJSObject(p js.Value, ctx js.Value) *HemisphericLight {
	return &HemisphericLight{Light: LightFromJSObject(p, ctx), ctx: ctx}
}

// HemisphericLightArrayToJSArray returns a JavaScript Array for the wrapped array.
func HemisphericLightArrayToJSArray(array []*HemisphericLight) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHemisphericLight returns a new HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight
func (ba *Babylon) NewHemisphericLight(name string, direction *Vector3, scene *Scene) *HemisphericLight {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, direction.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("HemisphericLight").New(args...)
	return HemisphericLightFromJSObject(p, ba.ctx)
}

// ComputeWorldMatrix calls the ComputeWorldMatrix method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#computeworldmatrix
func (h *HemisphericLight) ComputeWorldMatrix() *Matrix {

	retVal := h.p.Call("computeWorldMatrix")
	return MatrixFromJSObject(retVal, h.ctx)
}

// GetClassName calls the GetClassName method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#getclassname
func (h *HemisphericLight) GetClassName() string {

	retVal := h.p.Call("getClassName")
	return retVal.String()
}

// GetShadowGenerator calls the GetShadowGenerator method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#getshadowgenerator
func (h *HemisphericLight) GetShadowGenerator() *IShadowGenerator {

	retVal := h.p.Call("getShadowGenerator")
	return IShadowGeneratorFromJSObject(retVal, h.ctx)
}

// GetTypeID calls the GetTypeID method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#gettypeid
func (h *HemisphericLight) GetTypeID() float64 {

	retVal := h.p.Call("getTypeID")
	return retVal.Float()
}

// PrepareLightSpecificDefines calls the PrepareLightSpecificDefines method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#preparelightspecificdefines
func (h *HemisphericLight) PrepareLightSpecificDefines(defines interface{}, lightIndex float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, defines)
	args = append(args, lightIndex)

	h.p.Call("prepareLightSpecificDefines", args...)
}

// SetDirectionToTarget calls the SetDirectionToTarget method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#setdirectiontotarget
func (h *HemisphericLight) SetDirectionToTarget(target *Vector3) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, target.JSObject())

	retVal := h.p.Call("setDirectionToTarget", args...)
	return Vector3FromJSObject(retVal, h.ctx)
}

// TransferToEffect calls the TransferToEffect method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#transfertoeffect
func (h *HemisphericLight) TransferToEffect(effect *Effect, lightIndex string) *HemisphericLight {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, lightIndex)

	retVal := h.p.Call("transferToEffect", args...)
	return HemisphericLightFromJSObject(retVal, h.ctx)
}

// TransferToNodeMaterialEffect calls the TransferToNodeMaterialEffect method on the HemisphericLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#transfertonodematerialeffect
func (h *HemisphericLight) TransferToNodeMaterialEffect(effect *Effect, lightDataUniformName string) *HemisphericLight {

	args := make([]interface{}, 0, 2+0)

	args = append(args, effect.JSObject())
	args = append(args, lightDataUniformName)

	retVal := h.p.Call("transferToNodeMaterialEffect", args...)
	return HemisphericLightFromJSObject(retVal, h.ctx)
}

// Direction returns the Direction property of class HemisphericLight.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#direction
func (h *HemisphericLight) Direction() *Vector3 {
	retVal := h.p.Get("direction")
	return Vector3FromJSObject(retVal, h.ctx)
}

// SetDirection sets the Direction property of class HemisphericLight.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#direction
func (h *HemisphericLight) SetDirection(direction *Vector3) *HemisphericLight {
	h.p.Set("direction", direction.JSObject())
	return h
}

// GroundColor returns the GroundColor property of class HemisphericLight.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#groundcolor
func (h *HemisphericLight) GroundColor() *Color3 {
	retVal := h.p.Get("groundColor")
	return Color3FromJSObject(retVal, h.ctx)
}

// SetGroundColor sets the GroundColor property of class HemisphericLight.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericlight#groundcolor
func (h *HemisphericLight) SetGroundColor(groundColor *Color3) *HemisphericLight {
	h.p.Set("groundColor", groundColor.JSObject())
	return h
}
