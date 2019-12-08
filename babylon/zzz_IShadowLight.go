// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IShadowLight represents a babylon.js IShadowLight.
// Interface describing all the common properties and methods a shadow light needs to implement.
// This helps both the shadow generator and materials to genrate the corresponding shadow maps
// as well as binding the different shadow properties to the effects.
type IShadowLight struct {
	*Light
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IShadowLight) JSObject() js.Value { return i.p }

// IShadowLight returns a IShadowLight JavaScript class.
func (ba *Babylon) IShadowLight() *IShadowLight {
	p := ba.ctx.Get("IShadowLight")
	return IShadowLightFromJSObject(p, ba.ctx)
}

// IShadowLightFromJSObject returns a wrapped IShadowLight JavaScript class.
func IShadowLightFromJSObject(p js.Value, ctx js.Value) *IShadowLight {
	return &IShadowLight{Light: LightFromJSObject(p, ctx), ctx: ctx}
}

// IShadowLightArrayToJSArray returns a JavaScript Array for the wrapped array.
func IShadowLightArrayToJSArray(array []*IShadowLight) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewIShadowLight returns a new IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight
func (ba *Babylon) NewIShadowLight(name string, scene *Scene) *IShadowLight {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("IShadowLight").New(args...)
	return IShadowLightFromJSObject(p, ba.ctx)
}

// ComputeTransformedInformation calls the ComputeTransformedInformation method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#computetransformedinformation
func (i *IShadowLight) ComputeTransformedInformation() bool {

	retVal := i.p.Call("computeTransformedInformation")
	return retVal.Bool()
}

// ForceProjectionMatrixCompute calls the ForceProjectionMatrixCompute method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#forceprojectionmatrixcompute
func (i *IShadowLight) ForceProjectionMatrixCompute() {

	i.p.Call("forceProjectionMatrixCompute")
}

// GetDepthMaxZ calls the GetDepthMaxZ method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#getdepthmaxz
func (i *IShadowLight) GetDepthMaxZ(activeCamera *Camera) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, activeCamera.JSObject())

	retVal := i.p.Call("getDepthMaxZ", args...)
	return retVal.Float()
}

// GetDepthMinZ calls the GetDepthMinZ method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#getdepthminz
func (i *IShadowLight) GetDepthMinZ(activeCamera *Camera) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, activeCamera.JSObject())

	retVal := i.p.Call("getDepthMinZ", args...)
	return retVal.Float()
}

// GetDepthScale calls the GetDepthScale method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#getdepthscale
func (i *IShadowLight) GetDepthScale() float64 {

	retVal := i.p.Call("getDepthScale")
	return retVal.Float()
}

// GetScene calls the GetScene method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#getscene
func (i *IShadowLight) GetScene() *Scene {

	retVal := i.p.Call("getScene")
	return SceneFromJSObject(retVal, i.ctx)
}

// IShadowLightGetShadowDirectionOpts contains optional parameters for IShadowLight.GetShadowDirection.
type IShadowLightGetShadowDirectionOpts struct {
	FaceIndex *float64
}

// GetShadowDirection calls the GetShadowDirection method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#getshadowdirection
func (i *IShadowLight) GetShadowDirection(opts *IShadowLightGetShadowDirectionOpts) *Vector3 {
	if opts == nil {
		opts = &IShadowLightGetShadowDirectionOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.FaceIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FaceIndex)
	}

	retVal := i.p.Call("getShadowDirection", args...)
	return Vector3FromJSObject(retVal, i.ctx)
}

// NeedCube calls the NeedCube method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#needcube
func (i *IShadowLight) NeedCube() bool {

	retVal := i.p.Call("needCube")
	return retVal.Bool()
}

// NeedProjectionMatrixCompute calls the NeedProjectionMatrixCompute method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#needprojectionmatrixcompute
func (i *IShadowLight) NeedProjectionMatrixCompute() bool {

	retVal := i.p.Call("needProjectionMatrixCompute")
	return retVal.Bool()
}

// SetShadowProjectionMatrix calls the SetShadowProjectionMatrix method on the IShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#setshadowprojectionmatrix
func (i *IShadowLight) SetShadowProjectionMatrix(matrix *Matrix, viewMatrix *Matrix, renderList []*AbstractMesh) *IShadowLight {

	args := make([]interface{}, 0, 3+0)

	args = append(args, matrix.JSObject())
	args = append(args, viewMatrix.JSObject())
	args = append(args, AbstractMeshArrayToJSArray(renderList))

	retVal := i.p.Call("setShadowProjectionMatrix", args...)
	return IShadowLightFromJSObject(retVal, i.ctx)
}

// CustomProjectionMatrixBuilder returns the CustomProjectionMatrixBuilder property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#customprojectionmatrixbuilder
func (i *IShadowLight) CustomProjectionMatrixBuilder() js.Value {
	retVal := i.p.Get("customProjectionMatrixBuilder")
	return retVal
}

// SetCustomProjectionMatrixBuilder sets the CustomProjectionMatrixBuilder property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#customprojectionmatrixbuilder
func (i *IShadowLight) SetCustomProjectionMatrixBuilder(customProjectionMatrixBuilder JSFunc) *IShadowLight {
	i.p.Set("customProjectionMatrixBuilder", js.FuncOf(customProjectionMatrixBuilder))
	return i
}

// Direction returns the Direction property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#direction
func (i *IShadowLight) Direction() *Vector3 {
	retVal := i.p.Get("direction")
	return Vector3FromJSObject(retVal, i.ctx)
}

// SetDirection sets the Direction property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#direction
func (i *IShadowLight) SetDirection(direction *Vector3) *IShadowLight {
	i.p.Set("direction", direction.JSObject())
	return i
}

// Id returns the Id property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#id
func (i *IShadowLight) Id() string {
	retVal := i.p.Get("id")
	return retVal.String()
}

// SetId sets the Id property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#id
func (i *IShadowLight) SetId(id string) *IShadowLight {
	i.p.Set("id", id)
	return i
}

// Name returns the Name property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#name
func (i *IShadowLight) Name() string {
	retVal := i.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#name
func (i *IShadowLight) SetName(name string) *IShadowLight {
	i.p.Set("name", name)
	return i
}

// Position returns the Position property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#position
func (i *IShadowLight) Position() *Vector3 {
	retVal := i.p.Get("position")
	return Vector3FromJSObject(retVal, i.ctx)
}

// SetPosition sets the Position property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#position
func (i *IShadowLight) SetPosition(position *Vector3) *IShadowLight {
	i.p.Set("position", position.JSObject())
	return i
}

// ShadowMaxZ returns the ShadowMaxZ property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#shadowmaxz
func (i *IShadowLight) ShadowMaxZ() float64 {
	retVal := i.p.Get("shadowMaxZ")
	return retVal.Float()
}

// SetShadowMaxZ sets the ShadowMaxZ property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#shadowmaxz
func (i *IShadowLight) SetShadowMaxZ(shadowMaxZ float64) *IShadowLight {
	i.p.Set("shadowMaxZ", shadowMaxZ)
	return i
}

// ShadowMinZ returns the ShadowMinZ property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#shadowminz
func (i *IShadowLight) ShadowMinZ() float64 {
	retVal := i.p.Get("shadowMinZ")
	return retVal.Float()
}

// SetShadowMinZ sets the ShadowMinZ property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#shadowminz
func (i *IShadowLight) SetShadowMinZ(shadowMinZ float64) *IShadowLight {
	i.p.Set("shadowMinZ", shadowMinZ)
	return i
}

// TransformedDirection returns the TransformedDirection property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#transformeddirection
func (i *IShadowLight) TransformedDirection() *Vector3 {
	retVal := i.p.Get("transformedDirection")
	return Vector3FromJSObject(retVal, i.ctx)
}

// SetTransformedDirection sets the TransformedDirection property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#transformeddirection
func (i *IShadowLight) SetTransformedDirection(transformedDirection *Vector3) *IShadowLight {
	i.p.Set("transformedDirection", transformedDirection.JSObject())
	return i
}

// TransformedPosition returns the TransformedPosition property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#transformedposition
func (i *IShadowLight) TransformedPosition() *Vector3 {
	retVal := i.p.Get("transformedPosition")
	return Vector3FromJSObject(retVal, i.ctx)
}

// SetTransformedPosition sets the TransformedPosition property of class IShadowLight.
//
// https://doc.babylonjs.com/api/classes/babylon.ishadowlight#transformedposition
func (i *IShadowLight) SetTransformedPosition(transformedPosition *Vector3) *IShadowLight {
	i.p.Set("transformedPosition", transformedPosition.JSObject())
	return i
}
