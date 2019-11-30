// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SoundTrack represents a babylon.js SoundTrack.
// It could be useful to isolate your music &amp;amp; sounds on several tracks to better manage volume on a grouped instance of sounds.
// It will be also used in a future release to apply effects on a specific track.
//
// See: http://doc.babylonjs.com/how_to/playing_sounds_and_music#using-sound-tracks
type SoundTrack struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *SoundTrack) JSObject() js.Value { return s.p }

// SoundTrack returns a SoundTrack JavaScript class.
func (ba *Babylon) SoundTrack() *SoundTrack {
	p := ba.ctx.Get("SoundTrack")
	return SoundTrackFromJSObject(p)
}

// SoundTrackFromJSObject returns a wrapped SoundTrack JavaScript class.
func SoundTrackFromJSObject(p js.Value) *SoundTrack {
	return &SoundTrack{p: p}
}

// NewSoundTrackOpts contains optional parameters for NewSoundTrack.
type NewSoundTrackOpts struct {
	Options *JSValue
}

// NewSoundTrack returns a new SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack
func (ba *Babylon) NewSoundTrack(scene *Scene, opts *NewSoundTrackOpts) *SoundTrack {
	if opts == nil {
		opts = &NewSoundTrackOpts{}
	}

	p := ba.ctx.Get("SoundTrack").New(scene.JSObject(), opts.Options.JSObject())
	return SoundTrackFromJSObject(p)
}

// TODO: methods
