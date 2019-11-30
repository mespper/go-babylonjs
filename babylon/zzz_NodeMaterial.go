// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NodeMaterial represents a babylon.js NodeMaterial.
// Class used to create a node based material built by assembling shader blocks
type NodeMaterial struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (n *NodeMaterial) JSObject() js.Value { return n.p }

// NodeMaterial returns a NodeMaterial JavaScript class.
func (ba *Babylon) NodeMaterial() *NodeMaterial {
	p := ba.ctx.Get("NodeMaterial")
	return NodeMaterialFromJSObject(p)
}

// NodeMaterialFromJSObject returns a wrapped NodeMaterial JavaScript class.
func NodeMaterialFromJSObject(p js.Value) *NodeMaterial {
	return &NodeMaterial{p: p}
}

// NewNodeMaterialOpts contains optional parameters for NewNodeMaterial.
type NewNodeMaterialOpts struct {
	Scene *Scene

	Options js.Value
}

// NewNodeMaterial returns a new NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial
func (ba *Babylon) NewNodeMaterial(name string, opts *NewNodeMaterialOpts) *NodeMaterial {
	if opts == nil {
		opts = &NewNodeMaterialOpts{}
	}

	p := ba.ctx.Get("NodeMaterial").New(name, opts.Scene.JSObject(), opts.Options)
	return NodeMaterialFromJSObject(p)
}

// TODO: methods
