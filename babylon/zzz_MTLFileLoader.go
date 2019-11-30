// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MTLFileLoader represents a babylon.js MTLFileLoader.
// Class reading and parsing the MTL file bundled with the obj file.
type MTLFileLoader struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (m *MTLFileLoader) JSObject() js.Value { return m.p }

// MTLFileLoader returns a MTLFileLoader JavaScript class.
func (ba *Babylon) MTLFileLoader() *MTLFileLoader {
	p := ba.ctx.Get("MTLFileLoader")
	return MTLFileLoaderFromJSObject(p)
}

// MTLFileLoaderFromJSObject returns a wrapped MTLFileLoader JavaScript class.
func MTLFileLoaderFromJSObject(p js.Value) *MTLFileLoader {
	return &MTLFileLoader{p: p}
}

// TODO: methods
