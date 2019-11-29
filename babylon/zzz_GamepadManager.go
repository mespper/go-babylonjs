// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GamepadManager represents a babylon.js GamepadManager.
// Manager for handling gamepads
type GamepadManager struct{}

// JSObject returns the underlying js.Value.
func (g *GamepadManager) JSObject() js.Value { return g.p }

// GamepadManager returns a GamepadManager JavaScript class.
func (b *Babylon) GamepadManager() *GamepadManager {
	p := b.ctx.Get("GamepadManager")
	return GamepadManagerFromJSObject(p)
}

// GamepadManagerFromJSObject returns a wrapped GamepadManager JavaScript class.
func GamepadManagerFromJSObject(p js.Value) *GamepadManager {
	return &GamepadManager{p: p}
}

// NewGamepadManager returns a new GamepadManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager
func (b *Babylon) NewGamepadManager(todo parameters) *GamepadManager {
	p := b.ctx.Get("GamepadManager").New(todo)
	return GamepadManagerFromJSObject(p)
}

// TODO: methods