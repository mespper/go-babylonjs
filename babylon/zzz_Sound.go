// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Sound represents a babylon.js Sound.
// Defines a sound that can be played in the application.
// The sound can either be an ambient track or a simple sound played in reaction to a user action.
//
// See: http://doc.babylonjs.com/how_to/playing_sounds_and_music
type Sound struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *Sound) JSObject() js.Value { return s.p }

// Sound returns a Sound JavaScript class.
func (ba *Babylon) Sound() *Sound {
	p := ba.ctx.Get("Sound")
	return SoundFromJSObject(p)
}

// SoundFromJSObject returns a wrapped Sound JavaScript class.
func SoundFromJSObject(p js.Value) *Sound {
	return &Sound{p: p}
}

// NewSoundOpts contains optional parameters for NewSound.
type NewSoundOpts struct {
	ReadyToPlayCallback *func()

	Options *JSValue
}

// NewSound returns a new Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound
func (ba *Babylon) NewSound(name string, urlOrArrayBuffer interface{}, scene *Scene, opts *NewSoundOpts) *Sound {
	if opts == nil {
		opts = &NewSoundOpts{}
	}

	p := ba.ctx.Get("Sound").New(name, urlOrArrayBuffer, scene.JSObject(), opts.ReadyToPlayCallback, opts.Options.JSObject())
	return SoundFromJSObject(p)
}

// TODO: methods
