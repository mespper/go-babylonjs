// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthOfFieldMergePostProcessOptions represents a babylon.js DepthOfFieldMergePostProcessOptions.
// Options to be set when merging outputs from the default pipeline.
type DepthOfFieldMergePostProcessOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DepthOfFieldMergePostProcessOptions) JSObject() js.Value { return d.p }

// DepthOfFieldMergePostProcessOptions returns a DepthOfFieldMergePostProcessOptions JavaScript class.
func (ba *Babylon) DepthOfFieldMergePostProcessOptions() *DepthOfFieldMergePostProcessOptions {
	p := ba.ctx.Get("DepthOfFieldMergePostProcessOptions")
	return DepthOfFieldMergePostProcessOptionsFromJSObject(p, ba.ctx)
}

// DepthOfFieldMergePostProcessOptionsFromJSObject returns a wrapped DepthOfFieldMergePostProcessOptions JavaScript class.
func DepthOfFieldMergePostProcessOptionsFromJSObject(p js.Value, ctx js.Value) *DepthOfFieldMergePostProcessOptions {
	return &DepthOfFieldMergePostProcessOptions{p: p, ctx: ctx}
}

// DepthOfFieldMergePostProcessOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func DepthOfFieldMergePostProcessOptionsArrayToJSArray(array []*DepthOfFieldMergePostProcessOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Bloom returns the Bloom property of class DepthOfFieldMergePostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocessoptions#bloom
func (d *DepthOfFieldMergePostProcessOptions) Bloom() js.Value {
	retVal := d.p.Get("bloom")
	return retVal
}

// SetBloom sets the Bloom property of class DepthOfFieldMergePostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocessoptions#bloom
func (d *DepthOfFieldMergePostProcessOptions) SetBloom(bloom js.Value) *DepthOfFieldMergePostProcessOptions {
	d.p.Set("bloom", bloom)
	return d
}

// DepthOfField returns the DepthOfField property of class DepthOfFieldMergePostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocessoptions#depthoffield
func (d *DepthOfFieldMergePostProcessOptions) DepthOfField() js.Value {
	retVal := d.p.Get("depthOfField")
	return retVal
}

// SetDepthOfField sets the DepthOfField property of class DepthOfFieldMergePostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocessoptions#depthoffield
func (d *DepthOfFieldMergePostProcessOptions) SetDepthOfField(depthOfField js.Value) *DepthOfFieldMergePostProcessOptions {
	d.p.Set("depthOfField", depthOfField)
	return d
}

// OriginalFromInput returns the OriginalFromInput property of class DepthOfFieldMergePostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocessoptions#originalfrominput
func (d *DepthOfFieldMergePostProcessOptions) OriginalFromInput() *PostProcess {
	retVal := d.p.Get("originalFromInput")
	return PostProcessFromJSObject(retVal, d.ctx)
}

// SetOriginalFromInput sets the OriginalFromInput property of class DepthOfFieldMergePostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldmergepostprocessoptions#originalfrominput
func (d *DepthOfFieldMergePostProcessOptions) SetOriginalFromInput(originalFromInput *PostProcess) *DepthOfFieldMergePostProcessOptions {
	d.p.Set("originalFromInput", originalFromInput.JSObject())
	return d
}
