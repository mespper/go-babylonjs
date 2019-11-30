// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EnvironmentTextureTools represents a babylon.js EnvironmentTextureTools.
// Sets of helpers addressing the serialization and deserialization of environment texture
// stored in a BabylonJS env file.
// Those files are usually stored as .env files.
type EnvironmentTextureTools struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (e *EnvironmentTextureTools) JSObject() js.Value { return e.p }

// EnvironmentTextureTools returns a EnvironmentTextureTools JavaScript class.
func (ba *Babylon) EnvironmentTextureTools() *EnvironmentTextureTools {
	p := ba.ctx.Get("EnvironmentTextureTools")
	return EnvironmentTextureToolsFromJSObject(p)
}

// EnvironmentTextureToolsFromJSObject returns a wrapped EnvironmentTextureTools JavaScript class.
func EnvironmentTextureToolsFromJSObject(p js.Value) *EnvironmentTextureTools {
	return &EnvironmentTextureTools{p: p}
}

// TODO: methods
