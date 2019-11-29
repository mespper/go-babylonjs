// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LayerSceneComponent represents a babylon.js LayerSceneComponent.
// Defines the layer scene component responsible to manage any layers
// in a given scene.
type LayerSceneComponent struct{}

// JSObject returns the underlying js.Value.
func (l *LayerSceneComponent) JSObject() js.Value { return l.p }

// LayerSceneComponent returns a LayerSceneComponent JavaScript class.
func (b *Babylon) LayerSceneComponent() *LayerSceneComponent {
	p := b.ctx.Get("LayerSceneComponent")
	return LayerSceneComponentFromJSObject(p)
}

// LayerSceneComponentFromJSObject returns a wrapped LayerSceneComponent JavaScript class.
func LayerSceneComponentFromJSObject(p js.Value) *LayerSceneComponent {
	return &LayerSceneComponent{p: p}
}

// NewLayerSceneComponent returns a new LayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.layerscenecomponent
func (b *Babylon) NewLayerSceneComponent(todo parameters) *LayerSceneComponent {
	p := b.ctx.Get("LayerSceneComponent").New(todo)
	return LayerSceneComponentFromJSObject(p)
}

// TODO: methods