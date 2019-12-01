// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Skeleton represents a babylon.js Skeleton.
// Class used to handle skinning animations
//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons
type Skeleton struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *Skeleton) JSObject() js.Value { return s.p }

// Skeleton returns a Skeleton JavaScript class.
func (ba *Babylon) Skeleton() *Skeleton {
	p := ba.ctx.Get("Skeleton")
	return SkeletonFromJSObject(p, ba.ctx)
}

// SkeletonFromJSObject returns a wrapped Skeleton JavaScript class.
func SkeletonFromJSObject(p js.Value, ctx js.Value) *Skeleton {
	return &Skeleton{p: p, ctx: ctx}
}

// SkeletonArrayToJSArray returns a JavaScript Array for the wrapped array.
func SkeletonArrayToJSArray(array []*Skeleton) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSkeleton returns a new Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton
func (ba *Babylon) NewSkeleton(name string, id string, scene *Scene) *Skeleton {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, id)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("Skeleton").New(args...)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SkeletonBeginAnimationOpts contains optional parameters for Skeleton.BeginAnimation.
type SkeletonBeginAnimationOpts struct {
	Loop           *bool
	SpeedRatio     *float64
	OnAnimationEnd func()
}

// BeginAnimation calls the BeginAnimation method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#beginanimation
func (s *Skeleton) BeginAnimation(name string, opts *SkeletonBeginAnimationOpts) *Animatable {
	if opts == nil {
		opts = &SkeletonBeginAnimationOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, name)

	if opts.Loop == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Loop)
	}
	if opts.SpeedRatio == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SpeedRatio)
	}
	if opts.OnAnimationEnd == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnAnimationEnd)
	}

	retVal := s.p.Call("beginAnimation", args...)
	return AnimatableFromJSObject(retVal, s.ctx)
}

// SkeletonCloneOpts contains optional parameters for Skeleton.Clone.
type SkeletonCloneOpts struct {
	Id *string
}

// Clone calls the Clone method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#clone
func (s *Skeleton) Clone(name string, opts *SkeletonCloneOpts) *Skeleton {
	if opts == nil {
		opts = &SkeletonCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, name)

	if opts.Id == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Id)
	}

	retVal := s.p.Call("clone", args...)
	return SkeletonFromJSObject(retVal, s.ctx)
}

// SkeletonComputeAbsoluteTransformsOpts contains optional parameters for Skeleton.ComputeAbsoluteTransforms.
type SkeletonComputeAbsoluteTransformsOpts struct {
	ForceUpdate *bool
}

// ComputeAbsoluteTransforms calls the ComputeAbsoluteTransforms method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#computeabsolutetransforms
func (s *Skeleton) ComputeAbsoluteTransforms(opts *SkeletonComputeAbsoluteTransformsOpts) {
	if opts == nil {
		opts = &SkeletonComputeAbsoluteTransformsOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceUpdate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceUpdate)
	}

	s.p.Call("computeAbsoluteTransforms", args...)
}

// SkeletonCopyAnimationRangeOpts contains optional parameters for Skeleton.CopyAnimationRange.
type SkeletonCopyAnimationRangeOpts struct {
	RescaleAsRequired *bool
}

// CopyAnimationRange calls the CopyAnimationRange method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#copyanimationrange
func (s *Skeleton) CopyAnimationRange(source *Skeleton, name string, opts *SkeletonCopyAnimationRangeOpts) bool {
	if opts == nil {
		opts = &SkeletonCopyAnimationRangeOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, source.JSObject())
	args = append(args, name)

	if opts.RescaleAsRequired == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RescaleAsRequired)
	}

	retVal := s.p.Call("copyAnimationRange", args...)
	return retVal.Bool()
}

// CreateAnimationRange calls the CreateAnimationRange method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#createanimationrange
func (s *Skeleton) CreateAnimationRange(name string, from float64, to float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, from)
	args = append(args, to)

	s.p.Call("createAnimationRange", args...)
}

// SkeletonDeleteAnimationRangeOpts contains optional parameters for Skeleton.DeleteAnimationRange.
type SkeletonDeleteAnimationRangeOpts struct {
	DeleteFrames *bool
}

// DeleteAnimationRange calls the DeleteAnimationRange method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#deleteanimationrange
func (s *Skeleton) DeleteAnimationRange(name string, opts *SkeletonDeleteAnimationRangeOpts) {
	if opts == nil {
		opts = &SkeletonDeleteAnimationRangeOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, name)

	if opts.DeleteFrames == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DeleteFrames)
	}

	s.p.Call("deleteAnimationRange", args...)
}

// Dispose calls the Dispose method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#dispose
func (s *Skeleton) Dispose() {

	s.p.Call("dispose")
}

// SkeletonEnableBlendingOpts contains optional parameters for Skeleton.EnableBlending.
type SkeletonEnableBlendingOpts struct {
	BlendingSpeed *float64
}

// EnableBlending calls the EnableBlending method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#enableblending
func (s *Skeleton) EnableBlending(opts *SkeletonEnableBlendingOpts) {
	if opts == nil {
		opts = &SkeletonEnableBlendingOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.BlendingSpeed == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BlendingSpeed)
	}

	s.p.Call("enableBlending", args...)
}

// GetAnimatables calls the GetAnimatables method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getanimatables
func (s *Skeleton) GetAnimatables() *IAnimatable {

	retVal := s.p.Call("getAnimatables")
	return IAnimatableFromJSObject(retVal, s.ctx)
}

// GetAnimationRange calls the GetAnimationRange method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getanimationrange
func (s *Skeleton) GetAnimationRange(name string) *AnimationRange {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getAnimationRange", args...)
	return AnimationRangeFromJSObject(retVal, s.ctx)
}

// GetAnimationRanges calls the GetAnimationRanges method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getanimationranges
func (s *Skeleton) GetAnimationRanges() *AnimationRange {

	retVal := s.p.Call("getAnimationRanges")
	return AnimationRangeFromJSObject(retVal, s.ctx)
}

// GetBoneIndexByName calls the GetBoneIndexByName method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getboneindexbyname
func (s *Skeleton) GetBoneIndexByName(name string) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getBoneIndexByName", args...)
	return retVal.Float()
}

// GetChildren calls the GetChildren method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getchildren
func (s *Skeleton) GetChildren() []*Bone {

	retVal := s.p.Call("getChildren")
	result := []*Bone{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BoneFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getclassname
func (s *Skeleton) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// GetPoseMatrix calls the GetPoseMatrix method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getposematrix
func (s *Skeleton) GetPoseMatrix() *Matrix {

	retVal := s.p.Call("getPoseMatrix")
	return MatrixFromJSObject(retVal, s.ctx)
}

// GetScene calls the GetScene method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#getscene
func (s *Skeleton) GetScene() *Scene {

	retVal := s.p.Call("getScene")
	return SceneFromJSObject(retVal, s.ctx)
}

// GetTransformMatrices calls the GetTransformMatrices method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#gettransformmatrices
func (s *Skeleton) GetTransformMatrices(mesh *AbstractMesh) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	retVal := s.p.Call("getTransformMatrices", args...)
	return retVal
}

// GetTransformMatrixTexture calls the GetTransformMatrixTexture method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#gettransformmatrixtexture
func (s *Skeleton) GetTransformMatrixTexture(mesh *AbstractMesh) *RawTexture {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	retVal := s.p.Call("getTransformMatrixTexture", args...)
	return RawTextureFromJSObject(retVal, s.ctx)
}

// Parse calls the Parse method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#parse
func (s *Skeleton) Parse(parsedSkeleton interface{}, scene *Scene) *Skeleton {

	args := make([]interface{}, 0, 2+0)

	args = append(args, parsedSkeleton)
	args = append(args, scene.JSObject())

	retVal := s.p.Call("Parse", args...)
	return SkeletonFromJSObject(retVal, s.ctx)
}

// Prepare calls the Prepare method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#prepare
func (s *Skeleton) Prepare() {

	s.p.Call("prepare")
}

// ReturnToRest calls the ReturnToRest method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#returntorest
func (s *Skeleton) ReturnToRest() {

	s.p.Call("returnToRest")
}

// Serialize calls the Serialize method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#serialize
func (s *Skeleton) Serialize() interface{} {

	retVal := s.p.Call("serialize")
	return retVal
}

// SortBones calls the SortBones method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#sortbones
func (s *Skeleton) SortBones() {

	s.p.Call("sortBones")
}

// SkeletonToStringOpts contains optional parameters for Skeleton.ToString.
type SkeletonToStringOpts struct {
	FullDetails *bool
}

// ToString calls the ToString method on the Skeleton object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#tostring
func (s *Skeleton) ToString(opts *SkeletonToStringOpts) string {
	if opts == nil {
		opts = &SkeletonToStringOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.FullDetails == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FullDetails)
	}

	retVal := s.p.Call("toString", args...)
	return retVal.String()
}

/*

// AnimationPropertiesOverride returns the AnimationPropertiesOverride property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#animationpropertiesoverride
func (s *Skeleton) AnimationPropertiesOverride(animationPropertiesOverride *AnimationPropertiesOverride) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(animationPropertiesOverride.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetAnimationPropertiesOverride sets the AnimationPropertiesOverride property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#animationpropertiesoverride
func (s *Skeleton) SetAnimationPropertiesOverride(animationPropertiesOverride *AnimationPropertiesOverride) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(animationPropertiesOverride.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#animations
func (s *Skeleton) Animations(animations []*Animation) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(animations)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#animations
func (s *Skeleton) SetAnimations(animations []*Animation) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(animations)
	return SkeletonFromJSObject(p, ba.ctx)
}

// Bones returns the Bones property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#bones
func (s *Skeleton) Bones(bones *Bone) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(bones.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetBones sets the Bones property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#bones
func (s *Skeleton) SetBones(bones *Bone) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(bones.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// DimensionsAtRest returns the DimensionsAtRest property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#dimensionsatrest
func (s *Skeleton) DimensionsAtRest(dimensionsAtRest *Vector3) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(dimensionsAtRest.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetDimensionsAtRest sets the DimensionsAtRest property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#dimensionsatrest
func (s *Skeleton) SetDimensionsAtRest(dimensionsAtRest *Vector3) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(dimensionsAtRest.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// DoNotSerialize returns the DoNotSerialize property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#donotserialize
func (s *Skeleton) DoNotSerialize(doNotSerialize bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(doNotSerialize)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetDoNotSerialize sets the DoNotSerialize property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#donotserialize
func (s *Skeleton) SetDoNotSerialize(doNotSerialize bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(doNotSerialize)
	return SkeletonFromJSObject(p, ba.ctx)
}

// Id returns the Id property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#id
func (s *Skeleton) Id(id string) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(id)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetId sets the Id property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#id
func (s *Skeleton) SetId(id string) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(id)
	return SkeletonFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#inspectablecustomproperties
func (s *Skeleton) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(inspectableCustomProperties.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#inspectablecustomproperties
func (s *Skeleton) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(inspectableCustomProperties.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// IsUsingTextureForMatrices returns the IsUsingTextureForMatrices property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#isusingtextureformatrices
func (s *Skeleton) IsUsingTextureForMatrices(isUsingTextureForMatrices bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(isUsingTextureForMatrices)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetIsUsingTextureForMatrices sets the IsUsingTextureForMatrices property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#isusingtextureformatrices
func (s *Skeleton) SetIsUsingTextureForMatrices(isUsingTextureForMatrices bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(isUsingTextureForMatrices)
	return SkeletonFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#name
func (s *Skeleton) Name(name string) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(name)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#name
func (s *Skeleton) SetName(name string) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(name)
	return SkeletonFromJSObject(p, ba.ctx)
}

// NeedInitialSkinMatrix returns the NeedInitialSkinMatrix property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#needinitialskinmatrix
func (s *Skeleton) NeedInitialSkinMatrix(needInitialSkinMatrix bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(needInitialSkinMatrix)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetNeedInitialSkinMatrix sets the NeedInitialSkinMatrix property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#needinitialskinmatrix
func (s *Skeleton) SetNeedInitialSkinMatrix(needInitialSkinMatrix bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(needInitialSkinMatrix)
	return SkeletonFromJSObject(p, ba.ctx)
}

// OnBeforeComputeObservable returns the OnBeforeComputeObservable property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#onbeforecomputeobservable
func (s *Skeleton) OnBeforeComputeObservable(onBeforeComputeObservable *Observable) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(onBeforeComputeObservable.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetOnBeforeComputeObservable sets the OnBeforeComputeObservable property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#onbeforecomputeobservable
func (s *Skeleton) SetOnBeforeComputeObservable(onBeforeComputeObservable *Observable) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(onBeforeComputeObservable.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// OverrideMesh returns the OverrideMesh property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#overridemesh
func (s *Skeleton) OverrideMesh(overrideMesh *AbstractMesh) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(overrideMesh.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetOverrideMesh sets the OverrideMesh property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#overridemesh
func (s *Skeleton) SetOverrideMesh(overrideMesh *AbstractMesh) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(overrideMesh.JSObject())
	return SkeletonFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#uniqueid
func (s *Skeleton) UniqueId(uniqueId float64) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(uniqueId)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#uniqueid
func (s *Skeleton) SetUniqueId(uniqueId float64) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(uniqueId)
	return SkeletonFromJSObject(p, ba.ctx)
}

// UseTextureToStoreBoneMatrices returns the UseTextureToStoreBoneMatrices property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#usetexturetostorebonematrices
func (s *Skeleton) UseTextureToStoreBoneMatrices(useTextureToStoreBoneMatrices bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(useTextureToStoreBoneMatrices)
	return SkeletonFromJSObject(p, ba.ctx)
}

// SetUseTextureToStoreBoneMatrices sets the UseTextureToStoreBoneMatrices property of class Skeleton.
//
// https://doc.babylonjs.com/api/classes/babylon.skeleton#usetexturetostorebonematrices
func (s *Skeleton) SetUseTextureToStoreBoneMatrices(useTextureToStoreBoneMatrices bool) *Skeleton {
	p := ba.ctx.Get("Skeleton").New(useTextureToStoreBoneMatrices)
	return SkeletonFromJSObject(p, ba.ctx)
}

*/
