// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoundingBoxRenderer represents a babylon.js BoundingBoxRenderer.
// Component responsible of rendering the bounding box of the meshes in a scene.
// This is usually used through the mesh.showBoundingBox or the scene.forceShowBoundingBoxes properties
type BoundingBoxRenderer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (b *BoundingBoxRenderer) JSObject() js.Value { return b.p }

// BoundingBoxRenderer returns a BoundingBoxRenderer JavaScript class.
func (ba *Babylon) BoundingBoxRenderer() *BoundingBoxRenderer {
	p := ba.ctx.Get("BoundingBoxRenderer")
	return BoundingBoxRendererFromJSObject(p)
}

// BoundingBoxRendererFromJSObject returns a wrapped BoundingBoxRenderer JavaScript class.
func BoundingBoxRendererFromJSObject(p js.Value) *BoundingBoxRenderer {
	return &BoundingBoxRenderer{p: p}
}

// NewBoundingBoxRenderer returns a new BoundingBoxRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxrenderer
func (ba *Babylon) NewBoundingBoxRenderer(scene *Scene) *BoundingBoxRenderer {
	p := ba.ctx.Get("BoundingBoxRenderer").New(scene.JSObject())
	return BoundingBoxRendererFromJSObject(p)
}

// TODO: methods
