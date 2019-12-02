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
type AudioEngine struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AudioEngine) JSObject() js.Value { return a.p }

// AudioEngine returns a AudioEngine JavaScript class.
func (ba *Babylon) AudioEngine() *AudioEngine {
	p := ba.ctx.Get("AudioEngine")
	return AudioEngineFromJSObject(p, ba.ctx)
}

// AudioEngineFromJSObject returns a wrapped AudioEngine JavaScript class.
func AudioEngineFromJSObject(p js.Value, ctx js.Value) *AudioEngine {
	return &AudioEngine{p: p, ctx: ctx}
}

// AudioEngineArrayToJSArray returns a JavaScript Array for the wrapped array.
func AudioEngineArrayToJSArray(array []*AudioEngine) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAudioEngineOpts contains optional parameters for NewAudioEngine.
type NewAudioEngineOpts struct {
	HostElement js.Value
}

// NewAudioEngine returns a new AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine
func (ba *Babylon) NewAudioEngine(opts *NewAudioEngineOpts) *AudioEngine {
	if opts == nil {
		opts = &NewAudioEngineOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	args = append(args, opts.HostElement)

	p := ba.ctx.Get("AudioEngine").New(args...)
	return AudioEngineFromJSObject(p, ba.ctx)
}

// ConnectToAnalyser calls the ConnectToAnalyser method on the AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#connecttoanalyser
func (a *AudioEngine) ConnectToAnalyser(analyser *Analyser) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, analyser.JSObject())

	a.p.Call("connectToAnalyser", args...)
}

// Dispose calls the Dispose method on the AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#dispose
func (a *AudioEngine) Dispose() {

	a.p.Call("dispose")
}

// GetGlobalVolume calls the GetGlobalVolume method on the AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#getglobalvolume
func (a *AudioEngine) GetGlobalVolume() float64 {

	retVal := a.p.Call("getGlobalVolume")
	return retVal.Float()
}

// Lock calls the Lock method on the AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#lock
func (a *AudioEngine) Lock() {

	a.p.Call("lock")
}

// SetGlobalVolume calls the SetGlobalVolume method on the AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#setglobalvolume
func (a *AudioEngine) SetGlobalVolume(newVolume float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newVolume)

	a.p.Call("setGlobalVolume", args...)
}

// Unlock calls the Unlock method on the AudioEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#unlock
func (a *AudioEngine) Unlock() {

	a.p.Call("unlock")
}

// AudioContext returns the AudioContext property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#audiocontext
func (a *AudioEngine) AudioContext() js.Value {
	retVal := a.p.Get("audioContext")
	return retVal
}

// SetAudioContext sets the AudioContext property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#audiocontext
func (a *AudioEngine) SetAudioContext(audioContext js.Value) *AudioEngine {
	a.p.Set("audioContext", audioContext)
	return a
}

// CanUseWebAudio returns the CanUseWebAudio property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#canusewebaudio
func (a *AudioEngine) CanUseWebAudio() bool {
	retVal := a.p.Get("canUseWebAudio")
	return retVal.Bool()
}

// SetCanUseWebAudio sets the CanUseWebAudio property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#canusewebaudio
func (a *AudioEngine) SetCanUseWebAudio(canUseWebAudio bool) *AudioEngine {
	a.p.Set("canUseWebAudio", canUseWebAudio)
	return a
}

// IsMP3supported returns the IsMP3supported property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#ismp3supported
func (a *AudioEngine) IsMP3supported() bool {
	retVal := a.p.Get("isMP3supported")
	return retVal.Bool()
}

// SetIsMP3supported sets the IsMP3supported property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#ismp3supported
func (a *AudioEngine) SetIsMP3supported(isMP3supported bool) *AudioEngine {
	a.p.Set("isMP3supported", isMP3supported)
	return a
}

// IsOGGsupported returns the IsOGGsupported property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#isoggsupported
func (a *AudioEngine) IsOGGsupported() bool {
	retVal := a.p.Get("isOGGsupported")
	return retVal.Bool()
}

// SetIsOGGsupported sets the IsOGGsupported property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#isoggsupported
func (a *AudioEngine) SetIsOGGsupported(isOGGsupported bool) *AudioEngine {
	a.p.Set("isOGGsupported", isOGGsupported)
	return a
}

// MasterGain returns the MasterGain property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#mastergain
func (a *AudioEngine) MasterGain() js.Value {
	retVal := a.p.Get("masterGain")
	return retVal
}

// SetMasterGain sets the MasterGain property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#mastergain
func (a *AudioEngine) SetMasterGain(masterGain js.Value) *AudioEngine {
	a.p.Set("masterGain", masterGain)
	return a
}

// OnAudioLockedObservable returns the OnAudioLockedObservable property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#onaudiolockedobservable
func (a *AudioEngine) OnAudioLockedObservable() *Observable {
	retVal := a.p.Get("onAudioLockedObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnAudioLockedObservable sets the OnAudioLockedObservable property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#onaudiolockedobservable
func (a *AudioEngine) SetOnAudioLockedObservable(onAudioLockedObservable *Observable) *AudioEngine {
	a.p.Set("onAudioLockedObservable", onAudioLockedObservable.JSObject())
	return a
}

// OnAudioUnlockedObservable returns the OnAudioUnlockedObservable property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#onaudiounlockedobservable
func (a *AudioEngine) OnAudioUnlockedObservable() *Observable {
	retVal := a.p.Get("onAudioUnlockedObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnAudioUnlockedObservable sets the OnAudioUnlockedObservable property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#onaudiounlockedobservable
func (a *AudioEngine) SetOnAudioUnlockedObservable(onAudioUnlockedObservable *Observable) *AudioEngine {
	a.p.Set("onAudioUnlockedObservable", onAudioUnlockedObservable.JSObject())
	return a
}

// Unlocked returns the Unlocked property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#unlocked
func (a *AudioEngine) Unlocked() bool {
	retVal := a.p.Get("unlocked")
	return retVal.Bool()
}

// SetUnlocked sets the Unlocked property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#unlocked
func (a *AudioEngine) SetUnlocked(unlocked bool) *AudioEngine {
	a.p.Set("unlocked", unlocked)
	return a
}

// UseCustomUnlockedButton returns the UseCustomUnlockedButton property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#usecustomunlockedbutton
func (a *AudioEngine) UseCustomUnlockedButton() bool {
	retVal := a.p.Get("useCustomUnlockedButton")
	return retVal.Bool()
}

// SetUseCustomUnlockedButton sets the UseCustomUnlockedButton property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#usecustomunlockedbutton
func (a *AudioEngine) SetUseCustomUnlockedButton(useCustomUnlockedButton bool) *AudioEngine {
	a.p.Set("useCustomUnlockedButton", useCustomUnlockedButton)
	return a
}

// WarnedWebAudioUnsupported returns the WarnedWebAudioUnsupported property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#warnedwebaudiounsupported
func (a *AudioEngine) WarnedWebAudioUnsupported() bool {
	retVal := a.p.Get("WarnedWebAudioUnsupported")
	return retVal.Bool()
}

// SetWarnedWebAudioUnsupported sets the WarnedWebAudioUnsupported property of class AudioEngine.
//
// https://doc.babylonjs.com/api/classes/babylon.audioengine#warnedwebaudiounsupported
func (a *AudioEngine) SetWarnedWebAudioUnsupported(WarnedWebAudioUnsupported bool) *AudioEngine {
	a.p.Set("WarnedWebAudioUnsupported", WarnedWebAudioUnsupported)
	return a
}
