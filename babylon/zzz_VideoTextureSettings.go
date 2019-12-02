// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VideoTextureSettings represents a babylon.js VideoTextureSettings.
// Settings for finer control over video usage
type VideoTextureSettings struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VideoTextureSettings) JSObject() js.Value { return v.p }

// VideoTextureSettings returns a VideoTextureSettings JavaScript class.
func (ba *Babylon) VideoTextureSettings() *VideoTextureSettings {
	p := ba.ctx.Get("VideoTextureSettings")
	return VideoTextureSettingsFromJSObject(p, ba.ctx)
}

// VideoTextureSettingsFromJSObject returns a wrapped VideoTextureSettings JavaScript class.
func VideoTextureSettingsFromJSObject(p js.Value, ctx js.Value) *VideoTextureSettings {
	return &VideoTextureSettings{p: p, ctx: ctx}
}

// VideoTextureSettingsArrayToJSArray returns a JavaScript Array for the wrapped array.
func VideoTextureSettingsArrayToJSArray(array []*VideoTextureSettings) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AutoPlay returns the AutoPlay property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#autoplay
func (v *VideoTextureSettings) AutoPlay() bool {
	retVal := v.p.Get("autoPlay")
	return retVal.Bool()
}

// SetAutoPlay sets the AutoPlay property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#autoplay
func (v *VideoTextureSettings) SetAutoPlay(autoPlay bool) *VideoTextureSettings {
	v.p.Set("autoPlay", autoPlay)
	return v
}

// AutoUpdateTexture returns the AutoUpdateTexture property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#autoupdatetexture
func (v *VideoTextureSettings) AutoUpdateTexture() bool {
	retVal := v.p.Get("autoUpdateTexture")
	return retVal.Bool()
}

// SetAutoUpdateTexture sets the AutoUpdateTexture property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#autoupdatetexture
func (v *VideoTextureSettings) SetAutoUpdateTexture(autoUpdateTexture bool) *VideoTextureSettings {
	v.p.Set("autoUpdateTexture", autoUpdateTexture)
	return v
}

// Loop returns the Loop property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#loop
func (v *VideoTextureSettings) Loop() bool {
	retVal := v.p.Get("loop")
	return retVal.Bool()
}

// SetLoop sets the Loop property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#loop
func (v *VideoTextureSettings) SetLoop(loop bool) *VideoTextureSettings {
	v.p.Set("loop", loop)
	return v
}

// Poster returns the Poster property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#poster
func (v *VideoTextureSettings) Poster() string {
	retVal := v.p.Get("poster")
	return retVal.String()
}

// SetPoster sets the Poster property of class VideoTextureSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.videotexturesettings#poster
func (v *VideoTextureSettings) SetPoster(poster string) *VideoTextureSettings {
	v.p.Set("poster", poster)
	return v
}
