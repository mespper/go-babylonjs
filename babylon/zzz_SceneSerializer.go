// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneSerializer represents a babylon.js SceneSerializer.
// Class used to serialize a scene into a string
type SceneSerializer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *SceneSerializer) JSObject() js.Value { return s.p }

// SceneSerializer returns a SceneSerializer JavaScript class.
func (ba *Babylon) SceneSerializer() *SceneSerializer {
	p := ba.ctx.Get("SceneSerializer")
	return SceneSerializerFromJSObject(p)
}

// SceneSerializerFromJSObject returns a wrapped SceneSerializer JavaScript class.
func SceneSerializerFromJSObject(p js.Value) *SceneSerializer {
	return &SceneSerializer{p: p}
}

// TODO: methods
