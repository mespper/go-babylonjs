// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IMotorEnabledJoint represents a babylon.js IMotorEnabledJoint.
// Interface for a motor enabled joint
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type IMotorEnabledJoint struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IMotorEnabledJoint) JSObject() js.Value { return i.p }

// IMotorEnabledJoint returns a IMotorEnabledJoint JavaScript class.
func (ba *Babylon) IMotorEnabledJoint() *IMotorEnabledJoint {
	p := ba.ctx.Get("IMotorEnabledJoint")
	return IMotorEnabledJointFromJSObject(p, ba.ctx)
}

// IMotorEnabledJointFromJSObject returns a wrapped IMotorEnabledJoint JavaScript class.
func IMotorEnabledJointFromJSObject(p js.Value, ctx js.Value) *IMotorEnabledJoint {
	return &IMotorEnabledJoint{p: p, ctx: ctx}
}

// IMotorEnabledJointArrayToJSArray returns a JavaScript Array for the wrapped array.
func IMotorEnabledJointArrayToJSArray(array []*IMotorEnabledJoint) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// IMotorEnabledJointSetLimitOpts contains optional parameters for IMotorEnabledJoint.SetLimit.
type IMotorEnabledJointSetLimitOpts struct {
	LowerLimit *float64
	MotorIndex *float64
}

// SetLimit calls the SetLimit method on the IMotorEnabledJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.imotorenabledjoint#setlimit
func (i *IMotorEnabledJoint) SetLimit(upperLimit float64, opts *IMotorEnabledJointSetLimitOpts) {
	if opts == nil {
		opts = &IMotorEnabledJointSetLimitOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, upperLimit)

	if opts.LowerLimit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LowerLimit)
	}
	if opts.MotorIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MotorIndex)
	}

	i.p.Call("setLimit", args...)
}

// IMotorEnabledJointSetMotorOpts contains optional parameters for IMotorEnabledJoint.SetMotor.
type IMotorEnabledJointSetMotorOpts struct {
	Force      *float64
	MaxForce   *float64
	MotorIndex *float64
}

// SetMotor calls the SetMotor method on the IMotorEnabledJoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.imotorenabledjoint#setmotor
func (i *IMotorEnabledJoint) SetMotor(opts *IMotorEnabledJointSetMotorOpts) {
	if opts == nil {
		opts = &IMotorEnabledJointSetMotorOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Force == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Force)
	}
	if opts.MaxForce == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxForce)
	}
	if opts.MotorIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MotorIndex)
	}

	i.p.Call("setMotor", args...)
}

/*

// PhysicsJoint returns the PhysicsJoint property of class IMotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.imotorenabledjoint#physicsjoint
func (i *IMotorEnabledJoint) PhysicsJoint(physicsJoint interface{}) *IMotorEnabledJoint {
	p := ba.ctx.Get("IMotorEnabledJoint").New(physicsJoint)
	return IMotorEnabledJointFromJSObject(p, ba.ctx)
}

// SetPhysicsJoint sets the PhysicsJoint property of class IMotorEnabledJoint.
//
// https://doc.babylonjs.com/api/classes/babylon.imotorenabledjoint#physicsjoint
func (i *IMotorEnabledJoint) SetPhysicsJoint(physicsJoint interface{}) *IMotorEnabledJoint {
	p := ba.ctx.Get("IMotorEnabledJoint").New(physicsJoint)
	return IMotorEnabledJointFromJSObject(p, ba.ctx)
}

*/