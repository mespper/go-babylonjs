// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WaveBlock represents a babylon.js WaveBlock.
// Block used to apply wave operation to floats
type WaveBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (w *WaveBlock) JSObject() js.Value { return w.p }

// WaveBlock returns a WaveBlock JavaScript class.
func (b *Babylon) WaveBlock() *WaveBlock {
	p := b.ctx.Get("WaveBlock")
	return WaveBlockFromJSObject(p)
}

// WaveBlockFromJSObject returns a wrapped WaveBlock JavaScript class.
func WaveBlockFromJSObject(p js.Value) *WaveBlock {
	return &WaveBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewWaveBlock returns a new WaveBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.waveblock
func (b *Babylon) NewWaveBlock(todo parameters) *WaveBlock {
	p := b.ctx.Get("WaveBlock").New(todo)
	return WaveBlockFromJSObject(p)
}

// TODO: methods