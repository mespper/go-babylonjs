// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VideoRecorder represents a babylon.js VideoRecorder.
// This can help with recording videos from BabylonJS.
// This is based on the available WebRTC functionalities of the browser.
//
// See: http://doc.babylonjs.com/how_to/render_scene_on_a_video
type VideoRecorder struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (v *VideoRecorder) JSObject() js.Value { return v.p }

// VideoRecorder returns a VideoRecorder JavaScript class.
func (ba *Babylon) VideoRecorder() *VideoRecorder {
	p := ba.ctx.Get("VideoRecorder")
	return VideoRecorderFromJSObject(p)
}

// VideoRecorderFromJSObject returns a wrapped VideoRecorder JavaScript class.
func VideoRecorderFromJSObject(p js.Value) *VideoRecorder {
	return &VideoRecorder{p: p}
}

// NewVideoRecorderOpts contains optional parameters for NewVideoRecorder.
type NewVideoRecorderOpts struct {
	Options js.Value
}

// NewVideoRecorder returns a new VideoRecorder object.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorder
func (ba *Babylon) NewVideoRecorder(engine *Engine, opts *NewVideoRecorderOpts) *VideoRecorder {
	if opts == nil {
		opts = &NewVideoRecorderOpts{}
	}

	p := ba.ctx.Get("VideoRecorder").New(engine.JSObject(), opts.Options)
	return VideoRecorderFromJSObject(p)
}

// TODO: methods
