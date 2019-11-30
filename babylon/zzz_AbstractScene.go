// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AbstractScene represents a babylon.js AbstractScene.
// Base class of the scene acting as a container for the different elements composing a scene.
// This class is dynamically extended by the different components of the scene increasing
// flexibility and reducing coupling
type AbstractScene struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (a *AbstractScene) JSObject() js.Value { return a.p }

// AbstractScene returns a AbstractScene JavaScript class.
func (ba *Babylon) AbstractScene() *AbstractScene {
	p := ba.ctx.Get("AbstractScene")
	return AbstractSceneFromJSObject(p)
}

// AbstractSceneFromJSObject returns a wrapped AbstractScene JavaScript class.
func AbstractSceneFromJSObject(p js.Value) *AbstractScene {
	return &AbstractScene{p: p}
}

// TODO: methods
