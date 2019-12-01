// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// INavMeshParameters represents a babylon.js INavMeshParameters.
// Configures the navigation mesh creation
type INavMeshParameters struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *INavMeshParameters) JSObject() js.Value { return i.p }

// INavMeshParameters returns a INavMeshParameters JavaScript class.
func (ba *Babylon) INavMeshParameters() *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters")
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// INavMeshParametersFromJSObject returns a wrapped INavMeshParameters JavaScript class.
func INavMeshParametersFromJSObject(p js.Value, ctx js.Value) *INavMeshParameters {
	return &INavMeshParameters{p: p, ctx: ctx}
}

// INavMeshParametersArrayToJSArray returns a JavaScript Array for the wrapped array.
func INavMeshParametersArrayToJSArray(array []*INavMeshParameters) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Ch returns the Ch property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#ch
func (i *INavMeshParameters) Ch(ch float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(ch)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetCh sets the Ch property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#ch
func (i *INavMeshParameters) SetCh(ch float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(ch)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// Cs returns the Cs property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#cs
func (i *INavMeshParameters) Cs(cs float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(cs)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetCs sets the Cs property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#cs
func (i *INavMeshParameters) SetCs(cs float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(cs)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// DetailSampleDist returns the DetailSampleDist property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#detailsampledist
func (i *INavMeshParameters) DetailSampleDist(detailSampleDist float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(detailSampleDist)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetDetailSampleDist sets the DetailSampleDist property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#detailsampledist
func (i *INavMeshParameters) SetDetailSampleDist(detailSampleDist float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(detailSampleDist)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// DetailSampleMaxError returns the DetailSampleMaxError property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#detailsamplemaxerror
func (i *INavMeshParameters) DetailSampleMaxError(detailSampleMaxError float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(detailSampleMaxError)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetDetailSampleMaxError sets the DetailSampleMaxError property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#detailsamplemaxerror
func (i *INavMeshParameters) SetDetailSampleMaxError(detailSampleMaxError float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(detailSampleMaxError)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// MaxEdgeLen returns the MaxEdgeLen property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#maxedgelen
func (i *INavMeshParameters) MaxEdgeLen(maxEdgeLen float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(maxEdgeLen)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetMaxEdgeLen sets the MaxEdgeLen property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#maxedgelen
func (i *INavMeshParameters) SetMaxEdgeLen(maxEdgeLen float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(maxEdgeLen)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// MaxSimplificationError returns the MaxSimplificationError property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#maxsimplificationerror
func (i *INavMeshParameters) MaxSimplificationError(maxSimplificationError float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(maxSimplificationError)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetMaxSimplificationError sets the MaxSimplificationError property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#maxsimplificationerror
func (i *INavMeshParameters) SetMaxSimplificationError(maxSimplificationError float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(maxSimplificationError)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// MaxVertsPerPoly returns the MaxVertsPerPoly property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#maxvertsperpoly
func (i *INavMeshParameters) MaxVertsPerPoly(maxVertsPerPoly float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(maxVertsPerPoly)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetMaxVertsPerPoly sets the MaxVertsPerPoly property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#maxvertsperpoly
func (i *INavMeshParameters) SetMaxVertsPerPoly(maxVertsPerPoly float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(maxVertsPerPoly)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// MergeRegionArea returns the MergeRegionArea property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#mergeregionarea
func (i *INavMeshParameters) MergeRegionArea(mergeRegionArea float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(mergeRegionArea)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetMergeRegionArea sets the MergeRegionArea property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#mergeregionarea
func (i *INavMeshParameters) SetMergeRegionArea(mergeRegionArea float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(mergeRegionArea)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// MinRegionArea returns the MinRegionArea property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#minregionarea
func (i *INavMeshParameters) MinRegionArea(minRegionArea float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(minRegionArea)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetMinRegionArea sets the MinRegionArea property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#minregionarea
func (i *INavMeshParameters) SetMinRegionArea(minRegionArea float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(minRegionArea)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// WalkableClimb returns the WalkableClimb property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableclimb
func (i *INavMeshParameters) WalkableClimb(walkableClimb float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableClimb)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetWalkableClimb sets the WalkableClimb property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableclimb
func (i *INavMeshParameters) SetWalkableClimb(walkableClimb float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableClimb)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// WalkableHeight returns the WalkableHeight property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableheight
func (i *INavMeshParameters) WalkableHeight(walkableHeight float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableHeight)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetWalkableHeight sets the WalkableHeight property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableheight
func (i *INavMeshParameters) SetWalkableHeight(walkableHeight float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableHeight)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// WalkableRadius returns the WalkableRadius property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableradius
func (i *INavMeshParameters) WalkableRadius(walkableRadius float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableRadius)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetWalkableRadius sets the WalkableRadius property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableradius
func (i *INavMeshParameters) SetWalkableRadius(walkableRadius float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableRadius)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// WalkableSlopeAngle returns the WalkableSlopeAngle property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableslopeangle
func (i *INavMeshParameters) WalkableSlopeAngle(walkableSlopeAngle float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableSlopeAngle)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

// SetWalkableSlopeAngle sets the WalkableSlopeAngle property of class INavMeshParameters.
//
// https://doc.babylonjs.com/api/classes/babylon.inavmeshparameters#walkableslopeangle
func (i *INavMeshParameters) SetWalkableSlopeAngle(walkableSlopeAngle float64) *INavMeshParameters {
	p := ba.ctx.Get("INavMeshParameters").New(walkableSlopeAngle)
	return INavMeshParametersFromJSObject(p, ba.ctx)
}

*/
