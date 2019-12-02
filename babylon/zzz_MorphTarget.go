// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MorphTarget represents a babylon.js MorphTarget.
// Defines a target to use with MorphTargetManager
//
// See: http://doc.babylonjs.com/how_to/how_to_use_morphtargets
type MorphTarget struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MorphTarget) JSObject() js.Value { return m.p }

// MorphTarget returns a MorphTarget JavaScript class.
func (ba *Babylon) MorphTarget() *MorphTarget {
	p := ba.ctx.Get("MorphTarget")
	return MorphTargetFromJSObject(p, ba.ctx)
}

// MorphTargetFromJSObject returns a wrapped MorphTarget JavaScript class.
func MorphTargetFromJSObject(p js.Value, ctx js.Value) *MorphTarget {
	return &MorphTarget{p: p, ctx: ctx}
}

// MorphTargetArrayToJSArray returns a JavaScript Array for the wrapped array.
func MorphTargetArrayToJSArray(array []*MorphTarget) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMorphTargetOpts contains optional parameters for NewMorphTarget.
type NewMorphTargetOpts struct {
	Influence *float64
	Scene     *Scene
}

// NewMorphTarget returns a new MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget
func (ba *Babylon) NewMorphTarget(name string, opts *NewMorphTargetOpts) *MorphTarget {
	if opts == nil {
		opts = &NewMorphTargetOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, name)

	if opts.Influence == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Influence)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	p := ba.ctx.Get("MorphTarget").New(args...)
	return MorphTargetFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#clone
func (m *MorphTarget) Clone() *MorphTarget {

	retVal := m.p.Call("clone")
	return MorphTargetFromJSObject(retVal, m.ctx)
}

// MorphTargetFromMeshOpts contains optional parameters for MorphTarget.FromMesh.
type MorphTargetFromMeshOpts struct {
	Name      *string
	Influence *float64
}

// FromMesh calls the FromMesh method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#frommesh
func (m *MorphTarget) FromMesh(mesh *AbstractMesh, opts *MorphTargetFromMeshOpts) *MorphTarget {
	if opts == nil {
		opts = &MorphTargetFromMeshOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, mesh.JSObject())

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}
	if opts.Influence == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Influence)
	}

	retVal := m.p.Call("FromMesh", args...)
	return MorphTargetFromJSObject(retVal, m.ctx)
}

// GetClassName calls the GetClassName method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#getclassname
func (m *MorphTarget) GetClassName() string {

	retVal := m.p.Call("getClassName")
	return retVal.String()
}

// GetNormals calls the GetNormals method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#getnormals
func (m *MorphTarget) GetNormals() js.Value {

	retVal := m.p.Call("getNormals")
	return retVal
}

// GetPositions calls the GetPositions method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#getpositions
func (m *MorphTarget) GetPositions() js.Value {

	retVal := m.p.Call("getPositions")
	return retVal
}

// GetTangents calls the GetTangents method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#gettangents
func (m *MorphTarget) GetTangents() js.Value {

	retVal := m.p.Call("getTangents")
	return retVal
}

// GetUVs calls the GetUVs method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#getuvs
func (m *MorphTarget) GetUVs() js.Value {

	retVal := m.p.Call("getUVs")
	return retVal
}

// Parse calls the Parse method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#parse
func (m *MorphTarget) Parse(serializationObject interface{}) *MorphTarget {

	args := make([]interface{}, 0, 1+0)

	args = append(args, serializationObject)

	retVal := m.p.Call("Parse", args...)
	return MorphTargetFromJSObject(retVal, m.ctx)
}

// Serialize calls the Serialize method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#serialize
func (m *MorphTarget) Serialize() interface{} {

	retVal := m.p.Call("serialize")
	return retVal
}

// SetNormals calls the SetNormals method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#setnormals
func (m *MorphTarget) SetNormals(data js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	m.p.Call("setNormals", args...)
}

// SetPositions calls the SetPositions method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#setpositions
func (m *MorphTarget) SetPositions(data js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	m.p.Call("setPositions", args...)
}

// SetTangents calls the SetTangents method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#settangents
func (m *MorphTarget) SetTangents(data js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	m.p.Call("setTangents", args...)
}

// SetUVs calls the SetUVs method on the MorphTarget object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#setuvs
func (m *MorphTarget) SetUVs(data js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	m.p.Call("setUVs", args...)
}

// AnimationPropertiesOverride returns the AnimationPropertiesOverride property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#animationpropertiesoverride
func (m *MorphTarget) AnimationPropertiesOverride() *AnimationPropertiesOverride {
	retVal := m.p.Get("animationPropertiesOverride")
	return AnimationPropertiesOverrideFromJSObject(retVal, m.ctx)
}

// SetAnimationPropertiesOverride sets the AnimationPropertiesOverride property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#animationpropertiesoverride
func (m *MorphTarget) SetAnimationPropertiesOverride(animationPropertiesOverride *AnimationPropertiesOverride) *MorphTarget {
	m.p.Set("animationPropertiesOverride", animationPropertiesOverride.JSObject())
	return m
}

// Animations returns the Animations property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#animations
func (m *MorphTarget) Animations() *Animation {
	retVal := m.p.Get("animations")
	return AnimationFromJSObject(retVal, m.ctx)
}

// SetAnimations sets the Animations property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#animations
func (m *MorphTarget) SetAnimations(animations *Animation) *MorphTarget {
	m.p.Set("animations", animations.JSObject())
	return m
}

// HasNormals returns the HasNormals property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#hasnormals
func (m *MorphTarget) HasNormals() bool {
	retVal := m.p.Get("hasNormals")
	return retVal.Bool()
}

// SetHasNormals sets the HasNormals property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#hasnormals
func (m *MorphTarget) SetHasNormals(hasNormals bool) *MorphTarget {
	m.p.Set("hasNormals", hasNormals)
	return m
}

// HasPositions returns the HasPositions property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#haspositions
func (m *MorphTarget) HasPositions() bool {
	retVal := m.p.Get("hasPositions")
	return retVal.Bool()
}

// SetHasPositions sets the HasPositions property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#haspositions
func (m *MorphTarget) SetHasPositions(hasPositions bool) *MorphTarget {
	m.p.Set("hasPositions", hasPositions)
	return m
}

// HasTangents returns the HasTangents property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#hastangents
func (m *MorphTarget) HasTangents() bool {
	retVal := m.p.Get("hasTangents")
	return retVal.Bool()
}

// SetHasTangents sets the HasTangents property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#hastangents
func (m *MorphTarget) SetHasTangents(hasTangents bool) *MorphTarget {
	m.p.Set("hasTangents", hasTangents)
	return m
}

// HasUVs returns the HasUVs property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#hasuvs
func (m *MorphTarget) HasUVs() bool {
	retVal := m.p.Get("hasUVs")
	return retVal.Bool()
}

// SetHasUVs sets the HasUVs property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#hasuvs
func (m *MorphTarget) SetHasUVs(hasUVs bool) *MorphTarget {
	m.p.Set("hasUVs", hasUVs)
	return m
}

// Id returns the Id property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#id
func (m *MorphTarget) Id() string {
	retVal := m.p.Get("id")
	return retVal.String()
}

// SetId sets the Id property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#id
func (m *MorphTarget) SetId(id string) *MorphTarget {
	m.p.Set("id", id)
	return m
}

// Influence returns the Influence property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#influence
func (m *MorphTarget) Influence() float64 {
	retVal := m.p.Get("influence")
	return retVal.Float()
}

// SetInfluence sets the Influence property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#influence
func (m *MorphTarget) SetInfluence(influence float64) *MorphTarget {
	m.p.Set("influence", influence)
	return m
}

// Name returns the Name property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#name
func (m *MorphTarget) Name() string {
	retVal := m.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#name
func (m *MorphTarget) SetName(name string) *MorphTarget {
	m.p.Set("name", name)
	return m
}

// OnInfluenceChanged returns the OnInfluenceChanged property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#oninfluencechanged
func (m *MorphTarget) OnInfluenceChanged() *Observable {
	retVal := m.p.Get("onInfluenceChanged")
	return ObservableFromJSObject(retVal, m.ctx)
}

// SetOnInfluenceChanged sets the OnInfluenceChanged property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#oninfluencechanged
func (m *MorphTarget) SetOnInfluenceChanged(onInfluenceChanged *Observable) *MorphTarget {
	m.p.Set("onInfluenceChanged", onInfluenceChanged.JSObject())
	return m
}

// UniqueId returns the UniqueId property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#uniqueid
func (m *MorphTarget) UniqueId() float64 {
	retVal := m.p.Get("uniqueId")
	return retVal.Float()
}

// SetUniqueId sets the UniqueId property of class MorphTarget.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtarget#uniqueid
func (m *MorphTarget) SetUniqueId(uniqueId float64) *MorphTarget {
	m.p.Set("uniqueId", uniqueId)
	return m
}
