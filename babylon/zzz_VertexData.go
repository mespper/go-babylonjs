// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VertexData represents a babylon.js VertexData.
// This class contains the various kinds of data on every vertex of a mesh used in determining its shape and appearance
type VertexData struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VertexData) JSObject() js.Value { return v.p }

// VertexData returns a VertexData JavaScript class.
func (ba *Babylon) VertexData() *VertexData {
	p := ba.ctx.Get("VertexData")
	return VertexDataFromJSObject(p, ba.ctx)
}

// VertexDataFromJSObject returns a wrapped VertexData JavaScript class.
func VertexDataFromJSObject(p js.Value, ctx js.Value) *VertexData {
	return &VertexData{p: p, ctx: ctx}
}

// VertexDataArrayToJSArray returns a JavaScript Array for the wrapped array.
func VertexDataArrayToJSArray(array []*VertexData) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// VertexDataApplyToGeometryOpts contains optional parameters for VertexData.ApplyToGeometry.
type VertexDataApplyToGeometryOpts struct {
	Updatable *bool
}

// ApplyToGeometry calls the ApplyToGeometry method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#applytogeometry
func (v *VertexData) ApplyToGeometry(geometry *Geometry, opts *VertexDataApplyToGeometryOpts) *VertexData {
	if opts == nil {
		opts = &VertexDataApplyToGeometryOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, geometry.JSObject())

	if opts.Updatable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Updatable)
	}

	retVal := v.p.Call("applyToGeometry", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// VertexDataApplyToMeshOpts contains optional parameters for VertexData.ApplyToMesh.
type VertexDataApplyToMeshOpts struct {
	Updatable *bool
}

// ApplyToMesh calls the ApplyToMesh method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#applytomesh
func (v *VertexData) ApplyToMesh(mesh *Mesh, opts *VertexDataApplyToMeshOpts) *VertexData {
	if opts == nil {
		opts = &VertexDataApplyToMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, mesh.JSObject())

	if opts.Updatable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Updatable)
	}

	retVal := v.p.Call("applyToMesh", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// VertexDataComputeNormalsOpts contains optional parameters for VertexData.ComputeNormals.
type VertexDataComputeNormalsOpts struct {
	Options map[string]interface{}
}

// ComputeNormals calls the ComputeNormals method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#computenormals
func (v *VertexData) ComputeNormals(positions interface{}, indices interface{}, normals interface{}, opts *VertexDataComputeNormalsOpts) {
	if opts == nil {
		opts = &VertexDataComputeNormalsOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, positions)
	args = append(args, indices)
	args = append(args, normals)

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	v.p.Call("ComputeNormals", args...)
}

// CreateBox calls the CreateBox method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createbox
func (v *VertexData) CreateBox(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateBox", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateCylinder calls the CreateCylinder method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createcylinder
func (v *VertexData) CreateCylinder(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateCylinder", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateDashedLines calls the CreateDashedLines method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createdashedlines
func (v *VertexData) CreateDashedLines(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateDashedLines", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateDisc calls the CreateDisc method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createdisc
func (v *VertexData) CreateDisc(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateDisc", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateGround calls the CreateGround method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createground
func (v *VertexData) CreateGround(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateGround", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateGroundFromHeightMap calls the CreateGroundFromHeightMap method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#creategroundfromheightmap
func (v *VertexData) CreateGroundFromHeightMap(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateGroundFromHeightMap", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateIcoSphere calls the CreateIcoSphere method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createicosphere
func (v *VertexData) CreateIcoSphere(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateIcoSphere", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateLineSystem calls the CreateLineSystem method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createlinesystem
func (v *VertexData) CreateLineSystem(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateLineSystem", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreatePlane calls the CreatePlane method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createplane
func (v *VertexData) CreatePlane(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreatePlane", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// VertexDataCreatePolygonOpts contains optional parameters for VertexData.CreatePolygon.
type VertexDataCreatePolygonOpts struct {
	FUV      *Vector4
	FColors  *Color4
	FrontUVs *Vector4
	BackUVs  *Vector4
}

// CreatePolygon calls the CreatePolygon method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createpolygon
func (v *VertexData) CreatePolygon(polygon *Mesh, sideOrientation float64, opts *VertexDataCreatePolygonOpts) *VertexData {
	if opts == nil {
		opts = &VertexDataCreatePolygonOpts{}
	}

	args := make([]interface{}, 0, 2+4)

	args = append(args, polygon.JSObject())
	args = append(args, sideOrientation)

	if opts.FUV == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FUV.JSObject())
	}
	if opts.FColors == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FColors.JSObject())
	}
	if opts.FrontUVs == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FrontUVs.JSObject())
	}
	if opts.BackUVs == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.BackUVs.JSObject())
	}

	retVal := v.p.Call("CreatePolygon", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreatePolyhedron calls the CreatePolyhedron method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createpolyhedron
func (v *VertexData) CreatePolyhedron(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreatePolyhedron", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateRibbon calls the CreateRibbon method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createribbon
func (v *VertexData) CreateRibbon(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateRibbon", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateSphere calls the CreateSphere method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createsphere
func (v *VertexData) CreateSphere(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateSphere", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateTiledBox calls the CreateTiledBox method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createtiledbox
func (v *VertexData) CreateTiledBox(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateTiledBox", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateTiledGround calls the CreateTiledGround method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createtiledground
func (v *VertexData) CreateTiledGround(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateTiledGround", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateTiledPlane calls the CreateTiledPlane method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createtiledplane
func (v *VertexData) CreateTiledPlane(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateTiledPlane", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateTorus calls the CreateTorus method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createtorus
func (v *VertexData) CreateTorus(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateTorus", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// CreateTorusKnot calls the CreateTorusKnot method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#createtorusknot
func (v *VertexData) CreateTorusKnot(options js.Value) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, options)

	retVal := v.p.Call("CreateTorusKnot", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// VertexDataExtractFromGeometryOpts contains optional parameters for VertexData.ExtractFromGeometry.
type VertexDataExtractFromGeometryOpts struct {
	CopyWhenShared *bool
	ForceCopy      *bool
}

// ExtractFromGeometry calls the ExtractFromGeometry method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#extractfromgeometry
func (v *VertexData) ExtractFromGeometry(geometry *Geometry, opts *VertexDataExtractFromGeometryOpts) *VertexData {
	if opts == nil {
		opts = &VertexDataExtractFromGeometryOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, geometry.JSObject())

	if opts.CopyWhenShared == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CopyWhenShared)
	}
	if opts.ForceCopy == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceCopy)
	}

	retVal := v.p.Call("ExtractFromGeometry", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// VertexDataExtractFromMeshOpts contains optional parameters for VertexData.ExtractFromMesh.
type VertexDataExtractFromMeshOpts struct {
	CopyWhenShared *bool
	ForceCopy      *bool
}

// ExtractFromMesh calls the ExtractFromMesh method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#extractfrommesh
func (v *VertexData) ExtractFromMesh(mesh *Mesh, opts *VertexDataExtractFromMeshOpts) *VertexData {
	if opts == nil {
		opts = &VertexDataExtractFromMeshOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, mesh.JSObject())

	if opts.CopyWhenShared == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CopyWhenShared)
	}
	if opts.ForceCopy == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceCopy)
	}

	retVal := v.p.Call("ExtractFromMesh", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// ImportVertexData calls the ImportVertexData method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#importvertexdata
func (v *VertexData) ImportVertexData(parsedVertexData interface{}, geometry *Geometry) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, parsedVertexData)
	args = append(args, geometry.JSObject())

	v.p.Call("ImportVertexData", args...)
}

// VertexDataMergeOpts contains optional parameters for VertexData.Merge.
type VertexDataMergeOpts struct {
	Use32BitsIndices *bool
}

// Merge calls the Merge method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#merge
func (v *VertexData) Merge(other *VertexData, opts *VertexDataMergeOpts) *VertexData {
	if opts == nil {
		opts = &VertexDataMergeOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Use32BitsIndices == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Use32BitsIndices)
	}

	retVal := v.p.Call("merge", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// Serialize calls the Serialize method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#serialize
func (v *VertexData) Serialize() interface{} {

	retVal := v.p.Call("serialize")
	return retVal
}

// Set calls the Set method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#set
func (v *VertexData) Set(data js.Value, kind string) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, data)
	args = append(args, kind)

	v.p.Call("set", args...)
}

// Transform calls the Transform method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#transform
func (v *VertexData) Transform(matrix *Matrix) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, matrix.JSObject())

	retVal := v.p.Call("transform", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// UpdateGeometry calls the UpdateGeometry method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#updategeometry
func (v *VertexData) UpdateGeometry(geometry *Geometry) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, geometry.JSObject())

	retVal := v.p.Call("updateGeometry", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

// UpdateMesh calls the UpdateMesh method on the VertexData object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#updatemesh
func (v *VertexData) UpdateMesh(mesh *Mesh) *VertexData {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	retVal := v.p.Call("updateMesh", args...)
	return VertexDataFromJSObject(retVal, v.ctx)
}

/*

// BACKSIDE returns the BACKSIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#backside
func (v *VertexData) BACKSIDE(BACKSIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(BACKSIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetBACKSIDE sets the BACKSIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#backside
func (v *VertexData) SetBACKSIDE(BACKSIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(BACKSIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Colors returns the Colors property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#colors
func (v *VertexData) Colors(colors js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(colors)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetColors sets the Colors property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#colors
func (v *VertexData) SetColors(colors js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(colors)
	return VertexDataFromJSObject(p, ba.ctx)
}

// DEFAULTSIDE returns the DEFAULTSIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#defaultside
func (v *VertexData) DEFAULTSIDE(DEFAULTSIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(DEFAULTSIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetDEFAULTSIDE sets the DEFAULTSIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#defaultside
func (v *VertexData) SetDEFAULTSIDE(DEFAULTSIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(DEFAULTSIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// DOUBLESIDE returns the DOUBLESIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#doubleside
func (v *VertexData) DOUBLESIDE(DOUBLESIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(DOUBLESIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetDOUBLESIDE sets the DOUBLESIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#doubleside
func (v *VertexData) SetDOUBLESIDE(DOUBLESIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(DOUBLESIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// FRONTSIDE returns the FRONTSIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#frontside
func (v *VertexData) FRONTSIDE(FRONTSIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(FRONTSIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetFRONTSIDE sets the FRONTSIDE property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#frontside
func (v *VertexData) SetFRONTSIDE(FRONTSIDE float64) *VertexData {
	p := ba.ctx.Get("VertexData").New(FRONTSIDE)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Indices returns the Indices property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#indices
func (v *VertexData) Indices(indices js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(indices)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetIndices sets the Indices property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#indices
func (v *VertexData) SetIndices(indices js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(indices)
	return VertexDataFromJSObject(p, ba.ctx)
}

// MatricesIndices returns the MatricesIndices property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesindices
func (v *VertexData) MatricesIndices(matricesIndices js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesIndices)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetMatricesIndices sets the MatricesIndices property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesindices
func (v *VertexData) SetMatricesIndices(matricesIndices js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesIndices)
	return VertexDataFromJSObject(p, ba.ctx)
}

// MatricesIndicesExtra returns the MatricesIndicesExtra property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesindicesextra
func (v *VertexData) MatricesIndicesExtra(matricesIndicesExtra js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesIndicesExtra)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetMatricesIndicesExtra sets the MatricesIndicesExtra property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesindicesextra
func (v *VertexData) SetMatricesIndicesExtra(matricesIndicesExtra js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesIndicesExtra)
	return VertexDataFromJSObject(p, ba.ctx)
}

// MatricesWeights returns the MatricesWeights property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesweights
func (v *VertexData) MatricesWeights(matricesWeights js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesWeights)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetMatricesWeights sets the MatricesWeights property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesweights
func (v *VertexData) SetMatricesWeights(matricesWeights js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesWeights)
	return VertexDataFromJSObject(p, ba.ctx)
}

// MatricesWeightsExtra returns the MatricesWeightsExtra property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesweightsextra
func (v *VertexData) MatricesWeightsExtra(matricesWeightsExtra js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesWeightsExtra)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetMatricesWeightsExtra sets the MatricesWeightsExtra property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#matricesweightsextra
func (v *VertexData) SetMatricesWeightsExtra(matricesWeightsExtra js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(matricesWeightsExtra)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Normals returns the Normals property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#normals
func (v *VertexData) Normals(normals js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(normals)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetNormals sets the Normals property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#normals
func (v *VertexData) SetNormals(normals js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(normals)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Positions returns the Positions property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#positions
func (v *VertexData) Positions(positions js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(positions)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetPositions sets the Positions property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#positions
func (v *VertexData) SetPositions(positions js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(positions)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Tangents returns the Tangents property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#tangents
func (v *VertexData) Tangents(tangents js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(tangents)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetTangents sets the Tangents property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#tangents
func (v *VertexData) SetTangents(tangents js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(tangents)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Uvs returns the Uvs property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs
func (v *VertexData) Uvs(uvs js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetUvs sets the Uvs property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs
func (v *VertexData) SetUvs(uvs js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Uvs2 returns the Uvs2 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs2
func (v *VertexData) Uvs2(uvs2 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs2)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetUvs2 sets the Uvs2 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs2
func (v *VertexData) SetUvs2(uvs2 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs2)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Uvs3 returns the Uvs3 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs3
func (v *VertexData) Uvs3(uvs3 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs3)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetUvs3 sets the Uvs3 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs3
func (v *VertexData) SetUvs3(uvs3 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs3)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Uvs4 returns the Uvs4 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs4
func (v *VertexData) Uvs4(uvs4 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs4)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetUvs4 sets the Uvs4 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs4
func (v *VertexData) SetUvs4(uvs4 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs4)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Uvs5 returns the Uvs5 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs5
func (v *VertexData) Uvs5(uvs5 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs5)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetUvs5 sets the Uvs5 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs5
func (v *VertexData) SetUvs5(uvs5 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs5)
	return VertexDataFromJSObject(p, ba.ctx)
}

// Uvs6 returns the Uvs6 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs6
func (v *VertexData) Uvs6(uvs6 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs6)
	return VertexDataFromJSObject(p, ba.ctx)
}

// SetUvs6 sets the Uvs6 property of class VertexData.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexdata#uvs6
func (v *VertexData) SetUvs6(uvs6 js.Value) *VertexData {
	p := ba.ctx.Get("VertexData").New(uvs6)
	return VertexDataFromJSObject(p, ba.ctx)
}

*/
