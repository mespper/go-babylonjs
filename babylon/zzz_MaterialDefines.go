// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MaterialDefines represents a babylon.js MaterialDefines.
// Manages the defines for the Material
type MaterialDefines struct{}

// JSObject returns the underlying js.Value.
func (m *MaterialDefines) JSObject() js.Value { return m.p }

// MaterialDefines returns a MaterialDefines JavaScript class.
func (b *Babylon) MaterialDefines() *MaterialDefines {
	p := b.ctx.Get("MaterialDefines")
	return MaterialDefinesFromJSObject(p)
}

// MaterialDefinesFromJSObject returns a wrapped MaterialDefines JavaScript class.
func MaterialDefinesFromJSObject(p js.Value) *MaterialDefines {
	return &MaterialDefines{p: p}
}

// NewMaterialDefines returns a new MaterialDefines object.
//
// https://doc.babylonjs.com/api/classes/babylon.materialdefines
func (b *Babylon) NewMaterialDefines(todo parameters) *MaterialDefines {
	p := b.ctx.Get("MaterialDefines").New(todo)
	return MaterialDefinesFromJSObject(p)
}

// TODO: methods