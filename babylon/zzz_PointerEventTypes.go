// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointerEventTypes represents a babylon.js PointerEventTypes.
// Gather the list of pointer event types as constants.
type PointerEventTypes struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PointerEventTypes) JSObject() js.Value { return p.p }

// PointerEventTypes returns a PointerEventTypes JavaScript class.
func (ba *Babylon) PointerEventTypes() *PointerEventTypes {
	p := ba.ctx.Get("PointerEventTypes")
	return PointerEventTypesFromJSObject(p, ba.ctx)
}

// PointerEventTypesFromJSObject returns a wrapped PointerEventTypes JavaScript class.
func PointerEventTypesFromJSObject(p js.Value, ctx js.Value) *PointerEventTypes {
	return &PointerEventTypes{p: p, ctx: ctx}
}

// PointerEventTypesArrayToJSArray returns a JavaScript Array for the wrapped array.
func PointerEventTypesArrayToJSArray(array []*PointerEventTypes) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// POINTERDOUBLETAP returns the POINTERDOUBLETAP property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerdoubletap
func (p *PointerEventTypes) POINTERDOUBLETAP() float64 {
	retVal := p.p.Get("POINTERDOUBLETAP")
	return retVal.Float()
}

// SetPOINTERDOUBLETAP sets the POINTERDOUBLETAP property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerdoubletap
func (p *PointerEventTypes) SetPOINTERDOUBLETAP(POINTERDOUBLETAP float64) *PointerEventTypes {
	p.p.Set("POINTERDOUBLETAP", POINTERDOUBLETAP)
	return p
}

// POINTERDOWN returns the POINTERDOWN property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerdown
func (p *PointerEventTypes) POINTERDOWN() float64 {
	retVal := p.p.Get("POINTERDOWN")
	return retVal.Float()
}

// SetPOINTERDOWN sets the POINTERDOWN property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerdown
func (p *PointerEventTypes) SetPOINTERDOWN(POINTERDOWN float64) *PointerEventTypes {
	p.p.Set("POINTERDOWN", POINTERDOWN)
	return p
}

// POINTERMOVE returns the POINTERMOVE property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointermove
func (p *PointerEventTypes) POINTERMOVE() float64 {
	retVal := p.p.Get("POINTERMOVE")
	return retVal.Float()
}

// SetPOINTERMOVE sets the POINTERMOVE property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointermove
func (p *PointerEventTypes) SetPOINTERMOVE(POINTERMOVE float64) *PointerEventTypes {
	p.p.Set("POINTERMOVE", POINTERMOVE)
	return p
}

// POINTERPICK returns the POINTERPICK property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerpick
func (p *PointerEventTypes) POINTERPICK() float64 {
	retVal := p.p.Get("POINTERPICK")
	return retVal.Float()
}

// SetPOINTERPICK sets the POINTERPICK property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerpick
func (p *PointerEventTypes) SetPOINTERPICK(POINTERPICK float64) *PointerEventTypes {
	p.p.Set("POINTERPICK", POINTERPICK)
	return p
}

// POINTERTAP returns the POINTERTAP property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointertap
func (p *PointerEventTypes) POINTERTAP() float64 {
	retVal := p.p.Get("POINTERTAP")
	return retVal.Float()
}

// SetPOINTERTAP sets the POINTERTAP property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointertap
func (p *PointerEventTypes) SetPOINTERTAP(POINTERTAP float64) *PointerEventTypes {
	p.p.Set("POINTERTAP", POINTERTAP)
	return p
}

// POINTERUP returns the POINTERUP property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerup
func (p *PointerEventTypes) POINTERUP() float64 {
	retVal := p.p.Get("POINTERUP")
	return retVal.Float()
}

// SetPOINTERUP sets the POINTERUP property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerup
func (p *PointerEventTypes) SetPOINTERUP(POINTERUP float64) *PointerEventTypes {
	p.p.Set("POINTERUP", POINTERUP)
	return p
}

// POINTERWHEEL returns the POINTERWHEEL property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerwheel
func (p *PointerEventTypes) POINTERWHEEL() float64 {
	retVal := p.p.Get("POINTERWHEEL")
	return retVal.Float()
}

// SetPOINTERWHEEL sets the POINTERWHEEL property of class PointerEventTypes.
//
// https://doc.babylonjs.com/api/classes/babylon.pointereventtypes#pointerwheel
func (p *PointerEventTypes) SetPOINTERWHEEL(POINTERWHEEL float64) *PointerEventTypes {
	p.p.Set("POINTERWHEEL", POINTERWHEEL)
	return p
}
