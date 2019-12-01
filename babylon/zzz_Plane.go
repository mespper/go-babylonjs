// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Plane represents a babylon.js Plane.
// Represens a plane by the equation ax + by + cz + d = 0
type Plane struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *Plane) JSObject() js.Value { return p.p }

// Plane returns a Plane JavaScript class.
func (ba *Babylon) Plane() *Plane {
	p := ba.ctx.Get("Plane")
	return PlaneFromJSObject(p, ba.ctx)
}

// PlaneFromJSObject returns a wrapped Plane JavaScript class.
func PlaneFromJSObject(p js.Value, ctx js.Value) *Plane {
	return &Plane{p: p, ctx: ctx}
}

// PlaneArrayToJSArray returns a JavaScript Array for the wrapped array.
func PlaneArrayToJSArray(array []*Plane) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPlane returns a new Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane
func (ba *Babylon) NewPlane(a float64, b float64, c float64, d float64) *Plane {

	args := make([]interface{}, 0, 4+0)

	args = append(args, a)
	args = append(args, b)
	args = append(args, c)
	args = append(args, d)

	p := ba.ctx.Get("Plane").New(args...)
	return PlaneFromJSObject(p, ba.ctx)
}

// AsArray calls the AsArray method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#asarray
func (p *Plane) AsArray() float64 {

	retVal := p.p.Call("asArray")
	return retVal.Float()
}

// Clone calls the Clone method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#clone
func (p *Plane) Clone() *Plane {

	retVal := p.p.Call("clone")
	return PlaneFromJSObject(retVal, p.ctx)
}

// CopyFromPoints calls the CopyFromPoints method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#copyfrompoints
func (p *Plane) CopyFromPoints(point1 *Vector3, point2 *Vector3, point3 *Vector3) *Plane {

	args := make([]interface{}, 0, 3+0)

	args = append(args, point1.JSObject())
	args = append(args, point2.JSObject())
	args = append(args, point3.JSObject())

	retVal := p.p.Call("copyFromPoints", args...)
	return PlaneFromJSObject(retVal, p.ctx)
}

// DotCoordinate calls the DotCoordinate method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#dotcoordinate
func (p *Plane) DotCoordinate(point *Vector3) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, point.JSObject())

	retVal := p.p.Call("dotCoordinate", args...)
	return retVal.Float()
}

// FromArray calls the FromArray method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#fromarray
func (p *Plane) FromArray(array js.Value) *Plane {

	args := make([]interface{}, 0, 1+0)

	args = append(args, array)

	retVal := p.p.Call("FromArray", args...)
	return PlaneFromJSObject(retVal, p.ctx)
}

// FromPoints calls the FromPoints method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#frompoints
func (p *Plane) FromPoints(point1 *Vector3, point2 *Vector3, point3 *Vector3) *Plane {

	args := make([]interface{}, 0, 3+0)

	args = append(args, point1.JSObject())
	args = append(args, point2.JSObject())
	args = append(args, point3.JSObject())

	retVal := p.p.Call("FromPoints", args...)
	return PlaneFromJSObject(retVal, p.ctx)
}

// FromPositionAndNormal calls the FromPositionAndNormal method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#frompositionandnormal
func (p *Plane) FromPositionAndNormal(origin *Vector3, normal *Vector3) *Plane {

	args := make([]interface{}, 0, 2+0)

	args = append(args, origin.JSObject())
	args = append(args, normal.JSObject())

	retVal := p.p.Call("FromPositionAndNormal", args...)
	return PlaneFromJSObject(retVal, p.ctx)
}

// GetClassName calls the GetClassName method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#getclassname
func (p *Plane) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// GetHashCode calls the GetHashCode method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#gethashcode
func (p *Plane) GetHashCode() float64 {

	retVal := p.p.Call("getHashCode")
	return retVal.Float()
}

// IsFrontFacingTo calls the IsFrontFacingTo method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#isfrontfacingto
func (p *Plane) IsFrontFacingTo(direction *Vector3, epsilon float64) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, direction.JSObject())
	args = append(args, epsilon)

	retVal := p.p.Call("isFrontFacingTo", args...)
	return retVal.Bool()
}

// Normalize calls the Normalize method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#normalize
func (p *Plane) Normalize() *Plane {

	retVal := p.p.Call("normalize")
	return PlaneFromJSObject(retVal, p.ctx)
}

// SignedDistanceTo calls the SignedDistanceTo method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#signeddistanceto
func (p *Plane) SignedDistanceTo(point *Vector3) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, point.JSObject())

	retVal := p.p.Call("signedDistanceTo", args...)
	return retVal.Float()
}

// SignedDistanceToPlaneFromPositionAndNormal calls the SignedDistanceToPlaneFromPositionAndNormal method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#signeddistancetoplanefrompositionandnormal
func (p *Plane) SignedDistanceToPlaneFromPositionAndNormal(origin *Vector3, normal *Vector3, point *Vector3) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, origin.JSObject())
	args = append(args, normal.JSObject())
	args = append(args, point.JSObject())

	retVal := p.p.Call("SignedDistanceToPlaneFromPositionAndNormal", args...)
	return retVal.Float()
}

// Transform calls the Transform method on the Plane object.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#transform
func (p *Plane) Transform(transformation *Matrix) *Plane {

	args := make([]interface{}, 0, 1+0)

	args = append(args, transformation.JSObject())

	retVal := p.p.Call("transform", args...)
	return PlaneFromJSObject(retVal, p.ctx)
}

/*

// D returns the D property of class Plane.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#d
func (p *Plane) D(d float64) *Plane {
	p := ba.ctx.Get("Plane").New(d)
	return PlaneFromJSObject(p, ba.ctx)
}

// SetD sets the D property of class Plane.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#d
func (p *Plane) SetD(d float64) *Plane {
	p := ba.ctx.Get("Plane").New(d)
	return PlaneFromJSObject(p, ba.ctx)
}

// Normal returns the Normal property of class Plane.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#normal
func (p *Plane) Normal(normal *Vector3) *Plane {
	p := ba.ctx.Get("Plane").New(normal.JSObject())
	return PlaneFromJSObject(p, ba.ctx)
}

// SetNormal sets the Normal property of class Plane.
//
// https://doc.babylonjs.com/api/classes/babylon.plane#normal
func (p *Plane) SetNormal(normal *Vector3) *Plane {
	p := ba.ctx.Get("Plane").New(normal.JSObject())
	return PlaneFromJSObject(p, ba.ctx)
}

*/
