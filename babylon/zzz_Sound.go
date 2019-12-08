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
type Sound struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *Sound) JSObject() js.Value { return s.p }

// Sound returns a Sound JavaScript class.
func (ba *Babylon) Sound() *Sound {
	p := ba.ctx.Get("Sound")
	return SoundFromJSObject(p, ba.ctx)
}

// SoundFromJSObject returns a wrapped Sound JavaScript class.
func SoundFromJSObject(p js.Value, ctx js.Value) *Sound {
	return &Sound{p: p, ctx: ctx}
}

// SoundArrayToJSArray returns a JavaScript Array for the wrapped array.
func SoundArrayToJSArray(array []*Sound) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSoundOpts contains optional parameters for NewSound.
type NewSoundOpts struct {
	ReadyToPlayCallback JSFunc
	Options             *ISoundOptions
}

// NewSound returns a new Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound
func (ba *Babylon) NewSound(name string, urlOrArrayBuffer JSObject, scene *Scene, opts *NewSoundOpts) *Sound {
	if opts == nil {
		opts = &NewSoundOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, name)
	args = append(args, urlOrArrayBuffer.JSObject())
	args = append(args, scene.JSObject())

	if opts.ReadyToPlayCallback == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.ReadyToPlayCallback) /* never freed! */)
	}
	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}

	p := ba.ctx.Get("Sound").New(args...)
	return SoundFromJSObject(p, ba.ctx)
}

// AttachToMesh calls the AttachToMesh method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#attachtomesh
func (s *Sound) AttachToMesh(transformNode *TransformNode) {

	args := make([]interface{}, 0, 1+0)

	if transformNode == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, transformNode.JSObject())
	}

	s.p.Call("attachToMesh", args...)
}

// Clone calls the Clone method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#clone
func (s *Sound) Clone() *Sound {

	retVal := s.p.Call("clone")
	return SoundFromJSObject(retVal, s.ctx)
}

// ConnectToSoundTrackAudioNode calls the ConnectToSoundTrackAudioNode method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#connecttosoundtrackaudionode
func (s *Sound) ConnectToSoundTrackAudioNode(soundTrackAudioNode js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, soundTrackAudioNode)

	s.p.Call("connectToSoundTrackAudioNode", args...)
}

// DetachFromMesh calls the DetachFromMesh method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#detachfrommesh
func (s *Sound) DetachFromMesh() {

	s.p.Call("detachFromMesh")
}

// Dispose calls the Dispose method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#dispose
func (s *Sound) Dispose() {

	s.p.Call("dispose")
}

// GetAudioBuffer calls the GetAudioBuffer method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#getaudiobuffer
func (s *Sound) GetAudioBuffer() js.Value {

	retVal := s.p.Call("getAudioBuffer")
	return retVal
}

// GetVolume calls the GetVolume method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#getvolume
func (s *Sound) GetVolume() float64 {

	retVal := s.p.Call("getVolume")
	return retVal.Float()
}

// IsReady calls the IsReady method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#isready
func (s *Sound) IsReady() bool {

	retVal := s.p.Call("isReady")
	return retVal.Bool()
}

// SoundParseOpts contains optional parameters for Sound.Parse.
type SoundParseOpts struct {
	SourceSound *Sound
}

// Parse calls the Parse method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#parse
func (s *Sound) Parse(parsedSound JSObject, scene *Scene, rootUrl string, opts *SoundParseOpts) *Sound {
	if opts == nil {
		opts = &SoundParseOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	if parsedSound == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedSound.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	if opts.SourceSound == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.SourceSound.JSObject())
	}

	retVal := s.p.Call("Parse", args...)
	return SoundFromJSObject(retVal, s.ctx)
}

// Pause calls the Pause method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#pause
func (s *Sound) Pause() {

	s.p.Call("pause")
}

// SoundPlayOpts contains optional parameters for Sound.Play.
type SoundPlayOpts struct {
	Time   *float64
	Offset *float64
	Length *float64
}

// Play calls the Play method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#play
func (s *Sound) Play(opts *SoundPlayOpts) {
	if opts == nil {
		opts = &SoundPlayOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Time == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Time)
	}
	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}
	if opts.Length == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Length)
	}

	s.p.Call("play", args...)
}

// Serialize calls the Serialize method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#serialize
func (s *Sound) Serialize() js.Value {

	retVal := s.p.Call("serialize")
	return retVal
}

// SetAttenuationFunction calls the SetAttenuationFunction method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#setattenuationfunction
func (s *Sound) SetAttenuationFunction(callback JSFunc) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(callback))

	s.p.Call("setAttenuationFunction", args...)
}

// SetAudioBuffer calls the SetAudioBuffer method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#setaudiobuffer
func (s *Sound) SetAudioBuffer(audioBuffer js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, audioBuffer)

	s.p.Call("setAudioBuffer", args...)
}

// SetDirectionalCone calls the SetDirectionalCone method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#setdirectionalcone
func (s *Sound) SetDirectionalCone(coneInnerAngle float64, coneOuterAngle float64, coneOuterGain float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, coneInnerAngle)

	args = append(args, coneOuterAngle)

	args = append(args, coneOuterGain)

	s.p.Call("setDirectionalCone", args...)
}

// SetLocalDirectionToMesh calls the SetLocalDirectionToMesh method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#setlocaldirectiontomesh
func (s *Sound) SetLocalDirectionToMesh(newLocalDirection *Vector3) {

	args := make([]interface{}, 0, 1+0)

	if newLocalDirection == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, newLocalDirection.JSObject())
	}

	s.p.Call("setLocalDirectionToMesh", args...)
}

// SetPlaybackRate calls the SetPlaybackRate method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#setplaybackrate
func (s *Sound) SetPlaybackRate(newPlaybackRate float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newPlaybackRate)

	s.p.Call("setPlaybackRate", args...)
}

// SetPosition calls the SetPosition method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#setposition
func (s *Sound) SetPosition(newPosition *Vector3) {

	args := make([]interface{}, 0, 1+0)

	if newPosition == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, newPosition.JSObject())
	}

	s.p.Call("setPosition", args...)
}

// SoundSetVolumeOpts contains optional parameters for Sound.SetVolume.
type SoundSetVolumeOpts struct {
	Time *float64
}

// SetVolume calls the SetVolume method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#setvolume
func (s *Sound) SetVolume(newVolume float64, opts *SoundSetVolumeOpts) {
	if opts == nil {
		opts = &SoundSetVolumeOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, newVolume)

	if opts.Time == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Time)
	}

	s.p.Call("setVolume", args...)
}

// SoundStopOpts contains optional parameters for Sound.Stop.
type SoundStopOpts struct {
	Time *float64
}

// Stop calls the Stop method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#stop
func (s *Sound) Stop(opts *SoundStopOpts) {
	if opts == nil {
		opts = &SoundStopOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Time == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Time)
	}

	s.p.Call("stop", args...)
}

// SwitchPanningModelToEqualPower calls the SwitchPanningModelToEqualPower method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#switchpanningmodeltoequalpower
func (s *Sound) SwitchPanningModelToEqualPower() {

	s.p.Call("switchPanningModelToEqualPower")
}

// SwitchPanningModelToHRTF calls the SwitchPanningModelToHRTF method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#switchpanningmodeltohrtf
func (s *Sound) SwitchPanningModelToHRTF() {

	s.p.Call("switchPanningModelToHRTF")
}

// UpdateOptions calls the UpdateOptions method on the Sound object.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#updateoptions
func (s *Sound) UpdateOptions(options *ISoundOptions) {

	args := make([]interface{}, 0, 1+0)

	if options == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, options.JSObject())
	}

	s.p.Call("updateOptions", args...)
}

// Autoplay returns the Autoplay property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#autoplay
func (s *Sound) Autoplay() bool {
	retVal := s.p.Get("autoplay")
	return retVal.Bool()
}

// SetAutoplay sets the Autoplay property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#autoplay
func (s *Sound) SetAutoplay(autoplay bool) *Sound {
	s.p.Set("autoplay", autoplay)
	return s
}

// DirectionalConeInnerAngle returns the DirectionalConeInnerAngle property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#directionalconeinnerangle
func (s *Sound) DirectionalConeInnerAngle() float64 {
	retVal := s.p.Get("directionalConeInnerAngle")
	return retVal.Float()
}

// SetDirectionalConeInnerAngle sets the DirectionalConeInnerAngle property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#directionalconeinnerangle
func (s *Sound) SetDirectionalConeInnerAngle(directionalConeInnerAngle float64) *Sound {
	s.p.Set("directionalConeInnerAngle", directionalConeInnerAngle)
	return s
}

// DirectionalConeOuterAngle returns the DirectionalConeOuterAngle property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#directionalconeouterangle
func (s *Sound) DirectionalConeOuterAngle() float64 {
	retVal := s.p.Get("directionalConeOuterAngle")
	return retVal.Float()
}

// SetDirectionalConeOuterAngle sets the DirectionalConeOuterAngle property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#directionalconeouterangle
func (s *Sound) SetDirectionalConeOuterAngle(directionalConeOuterAngle float64) *Sound {
	s.p.Set("directionalConeOuterAngle", directionalConeOuterAngle)
	return s
}

// DistanceModel returns the DistanceModel property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#distancemodel
func (s *Sound) DistanceModel() string {
	retVal := s.p.Get("distanceModel")
	return retVal.String()
}

// SetDistanceModel sets the DistanceModel property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#distancemodel
func (s *Sound) SetDistanceModel(distanceModel string) *Sound {
	s.p.Set("distanceModel", distanceModel)
	return s
}

// IsPaused returns the IsPaused property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#ispaused
func (s *Sound) IsPaused() bool {
	retVal := s.p.Get("isPaused")
	return retVal.Bool()
}

// SetIsPaused sets the IsPaused property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#ispaused
func (s *Sound) SetIsPaused(isPaused bool) *Sound {
	s.p.Set("isPaused", isPaused)
	return s
}

// IsPlaying returns the IsPlaying property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#isplaying
func (s *Sound) IsPlaying() bool {
	retVal := s.p.Get("isPlaying")
	return retVal.Bool()
}

// SetIsPlaying sets the IsPlaying property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#isplaying
func (s *Sound) SetIsPlaying(isPlaying bool) *Sound {
	s.p.Set("isPlaying", isPlaying)
	return s
}

// Loop returns the Loop property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#loop
func (s *Sound) Loop() bool {
	retVal := s.p.Get("loop")
	return retVal.Bool()
}

// SetLoop sets the Loop property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#loop
func (s *Sound) SetLoop(loop bool) *Sound {
	s.p.Set("loop", loop)
	return s
}

// MaxDistance returns the MaxDistance property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#maxdistance
func (s *Sound) MaxDistance() float64 {
	retVal := s.p.Get("maxDistance")
	return retVal.Float()
}

// SetMaxDistance sets the MaxDistance property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#maxdistance
func (s *Sound) SetMaxDistance(maxDistance float64) *Sound {
	s.p.Set("maxDistance", maxDistance)
	return s
}

// Name returns the Name property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#name
func (s *Sound) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#name
func (s *Sound) SetName(name string) *Sound {
	s.p.Set("name", name)
	return s
}

// OnEndedObservable returns the OnEndedObservable property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#onendedobservable
func (s *Sound) OnEndedObservable() *Observable {
	retVal := s.p.Get("onEndedObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnEndedObservable sets the OnEndedObservable property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#onendedobservable
func (s *Sound) SetOnEndedObservable(onEndedObservable *Observable) *Sound {
	s.p.Set("onEndedObservable", onEndedObservable.JSObject())
	return s
}

// RefDistance returns the RefDistance property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#refdistance
func (s *Sound) RefDistance() float64 {
	retVal := s.p.Get("refDistance")
	return retVal.Float()
}

// SetRefDistance sets the RefDistance property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#refdistance
func (s *Sound) SetRefDistance(refDistance float64) *Sound {
	s.p.Set("refDistance", refDistance)
	return s
}

// RolloffFactor returns the RolloffFactor property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#rollofffactor
func (s *Sound) RolloffFactor() float64 {
	retVal := s.p.Get("rolloffFactor")
	return retVal.Float()
}

// SetRolloffFactor sets the RolloffFactor property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#rollofffactor
func (s *Sound) SetRolloffFactor(rolloffFactor float64) *Sound {
	s.p.Set("rolloffFactor", rolloffFactor)
	return s
}

// SoundTrackId returns the SoundTrackId property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#soundtrackid
func (s *Sound) SoundTrackId() float64 {
	retVal := s.p.Get("soundTrackId")
	return retVal.Float()
}

// SetSoundTrackId sets the SoundTrackId property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#soundtrackid
func (s *Sound) SetSoundTrackId(soundTrackId float64) *Sound {
	s.p.Set("soundTrackId", soundTrackId)
	return s
}

// SpatialSound returns the SpatialSound property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#spatialsound
func (s *Sound) SpatialSound() bool {
	retVal := s.p.Get("spatialSound")
	return retVal.Bool()
}

// SetSpatialSound sets the SpatialSound property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#spatialsound
func (s *Sound) SetSpatialSound(spatialSound bool) *Sound {
	s.p.Set("spatialSound", spatialSound)
	return s
}

// UseCustomAttenuation returns the UseCustomAttenuation property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#usecustomattenuation
func (s *Sound) UseCustomAttenuation() bool {
	retVal := s.p.Get("useCustomAttenuation")
	return retVal.Bool()
}

// SetUseCustomAttenuation sets the UseCustomAttenuation property of class Sound.
//
// https://doc.babylonjs.com/api/classes/babylon.sound#usecustomattenuation
func (s *Sound) SetUseCustomAttenuation(useCustomAttenuation bool) *Sound {
	s.p.Set("useCustomAttenuation", useCustomAttenuation)
	return s
}
