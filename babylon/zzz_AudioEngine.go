// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AudioEngine represents a babylon.js AudioEngine.
// This represents the default audio engine used in babylon.
// It is responsible to play, synchronize and analyse sounds throughout the  application.

//
// See: http://doc.babylonjs.com/how_to/playing_sounds_and_music
type AudioEngine struct{}

// JSObject returns the underlying js.Value.
func (a *AudioEngine) JSObject() js.Value { return a.p }

// AudioEngine returns a AudioEngine JavaScript class.
func (b *Babylon) AudioEngine() *AudioEngine {
	p := b.ctx.Get("AudioEngine")
	return AudioEngineFromJSObject(p)
}

// AudioEngineFromJSObject returns a wrapped AudioEngine JavaScript class.
func AudioEngineFromJSObject(p js.Value) *AudioEngine {
	return &AudioEngine{p: p}
}

// NewAudioEngine returns a new AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine
func (b *Babylon) NewAudioEngine(todo parameters) *AudioEngine {
	p := b.ctx.Get("AudioEngine").New(todo)
	return AudioEngineFromJSObject(p)
}

// TODO: methods