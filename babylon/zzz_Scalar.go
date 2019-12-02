// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Scalar represents a babylon.js Scalar.
// Scalar computation library
type Scalar struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *Scalar) JSObject() js.Value { return s.p }

// Scalar returns a Scalar JavaScript class.
func (ba *Babylon) Scalar() *Scalar {
	p := ba.ctx.Get("Scalar")
	return ScalarFromJSObject(p, ba.ctx)
}

// ScalarFromJSObject returns a wrapped Scalar JavaScript class.
func ScalarFromJSObject(p js.Value, ctx js.Value) *Scalar {
	return &Scalar{p: p, ctx: ctx}
}

// ScalarArrayToJSArray returns a JavaScript Array for the wrapped array.
func ScalarArrayToJSArray(array []*Scalar) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ScalarClampOpts contains optional parameters for Scalar.Clamp.
type ScalarClampOpts struct {
	Min *float64
	Max *float64
}

// Clamp calls the Clamp method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#clamp
func (s *Scalar) Clamp(value float64, opts *ScalarClampOpts) float64 {
	if opts == nil {
		opts = &ScalarClampOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, value)

	if opts.Min == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Min)
	}
	if opts.Max == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Max)
	}

	retVal := s.p.Call("Clamp", args...)
	return retVal.Float()
}

// DeltaAngle calls the DeltaAngle method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#deltaangle
func (s *Scalar) DeltaAngle(current float64, target float64) float64 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, current)
	args = append(args, target)

	retVal := s.p.Call("DeltaAngle", args...)
	return retVal.Float()
}

// Denormalize calls the Denormalize method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#denormalize
func (s *Scalar) Denormalize(normalized float64, min float64, max float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, normalized)
	args = append(args, min)
	args = append(args, max)

	retVal := s.p.Call("Denormalize", args...)
	return retVal.Float()
}

// Hermite calls the Hermite method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#hermite
func (s *Scalar) Hermite(value1 float64, tangent1 float64, value2 float64, tangent2 float64, amount float64) float64 {

	args := make([]interface{}, 0, 5+0)

	args = append(args, value1)
	args = append(args, tangent1)
	args = append(args, value2)
	args = append(args, tangent2)
	args = append(args, amount)

	retVal := s.p.Call("Hermite", args...)
	return retVal.Float()
}

// InverseLerp calls the InverseLerp method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#inverselerp
func (s *Scalar) InverseLerp(a float64, b float64, value float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, a)
	args = append(args, b)
	args = append(args, value)

	retVal := s.p.Call("InverseLerp", args...)
	return retVal.Float()
}

// Lerp calls the Lerp method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#lerp
func (s *Scalar) Lerp(start float64, end float64, amount float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, start)
	args = append(args, end)
	args = append(args, amount)

	retVal := s.p.Call("Lerp", args...)
	return retVal.Float()
}

// LerpAngle calls the LerpAngle method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#lerpangle
func (s *Scalar) LerpAngle(start float64, end float64, amount float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, start)
	args = append(args, end)
	args = append(args, amount)

	retVal := s.p.Call("LerpAngle", args...)
	return retVal.Float()
}

// Log2 calls the Log2 method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#log2
func (s *Scalar) Log2(value float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value)

	retVal := s.p.Call("Log2", args...)
	return retVal.Float()
}

// MoveTowards calls the MoveTowards method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#movetowards
func (s *Scalar) MoveTowards(current float64, target float64, maxDelta float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, current)
	args = append(args, target)
	args = append(args, maxDelta)

	retVal := s.p.Call("MoveTowards", args...)
	return retVal.Float()
}

// MoveTowardsAngle calls the MoveTowardsAngle method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#movetowardsangle
func (s *Scalar) MoveTowardsAngle(current float64, target float64, maxDelta float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, current)
	args = append(args, target)
	args = append(args, maxDelta)

	retVal := s.p.Call("MoveTowardsAngle", args...)
	return retVal.Float()
}

// Normalize calls the Normalize method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#normalize
func (s *Scalar) Normalize(value float64, min float64, max float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, value)
	args = append(args, min)
	args = append(args, max)

	retVal := s.p.Call("Normalize", args...)
	return retVal.Float()
}

// NormalizeRadians calls the NormalizeRadians method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#normalizeradians
func (s *Scalar) NormalizeRadians(angle float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, angle)

	retVal := s.p.Call("NormalizeRadians", args...)
	return retVal.Float()
}

// PercentToRange calls the PercentToRange method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#percenttorange
func (s *Scalar) PercentToRange(percent float64, min float64, max float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, percent)
	args = append(args, min)
	args = append(args, max)

	retVal := s.p.Call("PercentToRange", args...)
	return retVal.Float()
}

// PingPong calls the PingPong method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#pingpong
func (s *Scalar) PingPong(tx float64, length float64) float64 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, tx)
	args = append(args, length)

	retVal := s.p.Call("PingPong", args...)
	return retVal.Float()
}

// RandomRange calls the RandomRange method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#randomrange
func (s *Scalar) RandomRange(min float64, max float64) float64 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, min)
	args = append(args, max)

	retVal := s.p.Call("RandomRange", args...)
	return retVal.Float()
}

// RangeToPercent calls the RangeToPercent method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#rangetopercent
func (s *Scalar) RangeToPercent(number float64, min float64, max float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, number)
	args = append(args, min)
	args = append(args, max)

	retVal := s.p.Call("RangeToPercent", args...)
	return retVal.Float()
}

// Repeat calls the Repeat method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#repeat
func (s *Scalar) Repeat(value float64, length float64) float64 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, value)
	args = append(args, length)

	retVal := s.p.Call("Repeat", args...)
	return retVal.Float()
}

// Sign calls the Sign method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#sign
func (s *Scalar) Sign(value float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value)

	retVal := s.p.Call("Sign", args...)
	return retVal.Float()
}

// SmoothStep calls the SmoothStep method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#smoothstep
func (s *Scalar) SmoothStep(from float64, to float64, tx float64) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, from)
	args = append(args, to)
	args = append(args, tx)

	retVal := s.p.Call("SmoothStep", args...)
	return retVal.Float()
}

// ToHex calls the ToHex method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#tohex
func (s *Scalar) ToHex(i float64) string {

	args := make([]interface{}, 0, 1+0)

	args = append(args, i)

	retVal := s.p.Call("ToHex", args...)
	return retVal.String()
}

// ScalarWithinEpsilonOpts contains optional parameters for Scalar.WithinEpsilon.
type ScalarWithinEpsilonOpts struct {
	Epsilon *float64
}

// WithinEpsilon calls the WithinEpsilon method on the Scalar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#withinepsilon
func (s *Scalar) WithinEpsilon(a float64, b float64, opts *ScalarWithinEpsilonOpts) bool {
	if opts == nil {
		opts = &ScalarWithinEpsilonOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, a)
	args = append(args, b)

	if opts.Epsilon == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Epsilon)
	}

	retVal := s.p.Call("WithinEpsilon", args...)
	return retVal.Bool()
}

// TwoPi returns the TwoPi property of class Scalar.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#twopi
func (s *Scalar) TwoPi() float64 {
	retVal := s.p.Get("TwoPi")
	return retVal.Float()
}

// SetTwoPi sets the TwoPi property of class Scalar.
//
// https://doc.babylonjs.com/api/classes/babylon.scalar#twopi
func (s *Scalar) SetTwoPi(TwoPi float64) *Scalar {
	s.p.Set("TwoPi", TwoPi)
	return s
}
