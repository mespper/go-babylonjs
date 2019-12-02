// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IAccessorSparseValues represents a babylon.js IAccessorSparseValues.
// Array of size accessor.sparse.count times number of components storing the displaced accessor attributes pointed by accessor.sparse.indices
type IAccessorSparseValues struct {
	*IProperty
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IAccessorSparseValues) JSObject() js.Value { return i.p }

// IAccessorSparseValues returns a IAccessorSparseValues JavaScript class.
func (ba *Babylon) IAccessorSparseValues() *IAccessorSparseValues {
	p := ba.ctx.Get("IAccessorSparseValues")
	return IAccessorSparseValuesFromJSObject(p, ba.ctx)
}

// IAccessorSparseValuesFromJSObject returns a wrapped IAccessorSparseValues JavaScript class.
func IAccessorSparseValuesFromJSObject(p js.Value, ctx js.Value) *IAccessorSparseValues {
	return &IAccessorSparseValues{IProperty: IPropertyFromJSObject(p, ctx), ctx: ctx}
}

// IAccessorSparseValuesArrayToJSArray returns a JavaScript Array for the wrapped array.
func IAccessorSparseValuesArrayToJSArray(array []*IAccessorSparseValues) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// BufferView returns the BufferView property of class IAccessorSparseValues.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparsevalues#bufferview
func (i *IAccessorSparseValues) BufferView() float64 {
	retVal := i.p.Get("bufferView")
	return retVal.Float()
}

// SetBufferView sets the BufferView property of class IAccessorSparseValues.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparsevalues#bufferview
func (i *IAccessorSparseValues) SetBufferView(bufferView float64) *IAccessorSparseValues {
	i.p.Set("bufferView", bufferView)
	return i
}

// ByteOffset returns the ByteOffset property of class IAccessorSparseValues.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparsevalues#byteoffset
func (i *IAccessorSparseValues) ByteOffset() float64 {
	retVal := i.p.Get("byteOffset")
	return retVal.Float()
}

// SetByteOffset sets the ByteOffset property of class IAccessorSparseValues.
//
// https://doc.babylonjs.com/api/classes/babylon.iaccessorsparsevalues#byteoffset
func (i *IAccessorSparseValues) SetByteOffset(byteOffset float64) *IAccessorSparseValues {
	i.p.Set("byteOffset", byteOffset)
	return i
}
