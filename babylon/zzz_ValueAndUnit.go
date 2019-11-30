// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ValueAndUnit represents a babylon.js ValueAndUnit.
// Class used to specific a value and its associated unit
type ValueAndUnit struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (v *ValueAndUnit) JSObject() js.Value { return v.p }

// ValueAndUnit returns a ValueAndUnit JavaScript class.
func (ba *Babylon) ValueAndUnit() *ValueAndUnit {
	p := ba.ctx.Get("ValueAndUnit")
	return ValueAndUnitFromJSObject(p)
}

// ValueAndUnitFromJSObject returns a wrapped ValueAndUnit JavaScript class.
func ValueAndUnitFromJSObject(p js.Value) *ValueAndUnit {
	return &ValueAndUnit{p: p}
}

// NewValueAndUnitOpts contains optional parameters for NewValueAndUnit.
type NewValueAndUnitOpts struct {
	Unit *JSFloat64

	NegativeValueAllowed *JSBool
}

// NewValueAndUnit returns a new ValueAndUnit object.
//
// https://doc.babylonjs.com/api/classes/babylon.valueandunit
func (ba *Babylon) NewValueAndUnit(value float64, opts *NewValueAndUnitOpts) *ValueAndUnit {
	if opts == nil {
		opts = &NewValueAndUnitOpts{}
	}

	p := ba.ctx.Get("ValueAndUnit").New(value, opts.Unit, opts.NegativeValueAllowed)
	return ValueAndUnitFromJSObject(p)
}

// TODO: methods
