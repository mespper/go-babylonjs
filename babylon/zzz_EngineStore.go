// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EngineStore represents a babylon.js EngineStore.
// The engine store class is responsible to hold all the instances of Engine and Scene created
// during the life time of the application.
type EngineStore struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (e *EngineStore) JSObject() js.Value { return e.p }

// EngineStore returns a EngineStore JavaScript class.
func (ba *Babylon) EngineStore() *EngineStore {
	p := ba.ctx.Get("EngineStore")
	return EngineStoreFromJSObject(p)
}

// EngineStoreFromJSObject returns a wrapped EngineStore JavaScript class.
func EngineStoreFromJSObject(p js.Value) *EngineStore {
	return &EngineStore{p: p}
}

// TODO: methods
