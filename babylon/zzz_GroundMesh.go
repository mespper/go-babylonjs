// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GroundMesh represents a babylon.js GroundMesh.
// Mesh representing the gorund
type GroundMesh struct {
	*Mesh
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GroundMesh) JSObject() js.Value { return g.p }

// GroundMesh returns a GroundMesh JavaScript class.
func (ba *Babylon) GroundMesh() *GroundMesh {
	p := ba.ctx.Get("GroundMesh")
	return GroundMeshFromJSObject(p, ba.ctx)
}

// GroundMeshFromJSObject returns a wrapped GroundMesh JavaScript class.
func GroundMeshFromJSObject(p js.Value, ctx js.Value) *GroundMesh {
	return &GroundMesh{Mesh: MeshFromJSObject(p, ctx), ctx: ctx}
}

// GroundMeshArrayToJSArray returns a JavaScript Array for the wrapped array.
func GroundMeshArrayToJSArray(array []*GroundMesh) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGroundMesh returns a new GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh
func (ba *Babylon) NewGroundMesh(name string, scene *Scene) *GroundMesh {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("GroundMesh").New(args...)
	return GroundMeshFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#getclassname
func (g *GroundMesh) GetClassName() string {

	retVal := g.p.Call("getClassName")
	return retVal.String()
}

// GetHeightAtCoordinates calls the GetHeightAtCoordinates method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#getheightatcoordinates
func (g *GroundMesh) GetHeightAtCoordinates(x float64, z float64) float64 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)

	args = append(args, z)

	retVal := g.p.Call("getHeightAtCoordinates", args...)
	return retVal.Float()
}

// GetNormalAtCoordinates calls the GetNormalAtCoordinates method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#getnormalatcoordinates
func (g *GroundMesh) GetNormalAtCoordinates(x float64, z float64) *Vector3 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)

	args = append(args, z)

	retVal := g.p.Call("getNormalAtCoordinates", args...)
	return Vector3FromJSObject(retVal, g.ctx)
}

// GetNormalAtCoordinatesToRef calls the GetNormalAtCoordinatesToRef method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#getnormalatcoordinatestoref
func (g *GroundMesh) GetNormalAtCoordinatesToRef(x float64, z float64, ref *Vector3) *GroundMesh {

	args := make([]interface{}, 0, 3+0)

	args = append(args, x)

	args = append(args, z)

	if ref == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, ref.JSObject())
	}

	retVal := g.p.Call("getNormalAtCoordinatesToRef", args...)
	return GroundMeshFromJSObject(retVal, g.ctx)
}

// GroundMeshOptimizeOpts contains optional parameters for GroundMesh.Optimize.
type GroundMeshOptimizeOpts struct {
	OctreeBlocksSize *float64
}

// Optimize calls the Optimize method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#optimize
func (g *GroundMesh) Optimize(chunksCount float64, opts *GroundMeshOptimizeOpts) {
	if opts == nil {
		opts = &GroundMeshOptimizeOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, chunksCount)

	if opts.OctreeBlocksSize == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.OctreeBlocksSize)
	}

	g.p.Call("optimize", args...)
}

// Parse calls the Parse method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#parse
func (g *GroundMesh) Parse(parsedMesh JSObject, scene *Scene) *GroundMesh {

	args := make([]interface{}, 0, 2+0)

	if parsedMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedMesh.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := g.p.Call("Parse", args...)
	return GroundMeshFromJSObject(retVal, g.ctx)
}

// Serialize calls the Serialize method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#serialize
func (g *GroundMesh) Serialize(serializationObject JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	g.p.Call("serialize", args...)
}

// UpdateCoordinateHeights calls the UpdateCoordinateHeights method on the GroundMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#updatecoordinateheights
func (g *GroundMesh) UpdateCoordinateHeights() *GroundMesh {

	retVal := g.p.Call("updateCoordinateHeights")
	return GroundMeshFromJSObject(retVal, g.ctx)
}

// GenerateOctree returns the GenerateOctree property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#generateoctree
func (g *GroundMesh) GenerateOctree() bool {
	retVal := g.p.Get("generateOctree")
	return retVal.Bool()
}

// SetGenerateOctree sets the GenerateOctree property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#generateoctree
func (g *GroundMesh) SetGenerateOctree(generateOctree bool) *GroundMesh {
	g.p.Set("generateOctree", generateOctree)
	return g
}

// Subdivisions returns the Subdivisions property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#subdivisions
func (g *GroundMesh) Subdivisions() float64 {
	retVal := g.p.Get("subdivisions")
	return retVal.Float()
}

// SetSubdivisions sets the Subdivisions property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#subdivisions
func (g *GroundMesh) SetSubdivisions(subdivisions float64) *GroundMesh {
	g.p.Set("subdivisions", subdivisions)
	return g
}

// SubdivisionsX returns the SubdivisionsX property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#subdivisionsx
func (g *GroundMesh) SubdivisionsX() float64 {
	retVal := g.p.Get("subdivisionsX")
	return retVal.Float()
}

// SetSubdivisionsX sets the SubdivisionsX property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#subdivisionsx
func (g *GroundMesh) SetSubdivisionsX(subdivisionsX float64) *GroundMesh {
	g.p.Set("subdivisionsX", subdivisionsX)
	return g
}

// SubdivisionsY returns the SubdivisionsY property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#subdivisionsy
func (g *GroundMesh) SubdivisionsY() float64 {
	retVal := g.p.Get("subdivisionsY")
	return retVal.Float()
}

// SetSubdivisionsY sets the SubdivisionsY property of class GroundMesh.
//
// https://doc.babylonjs.com/api/classes/babylon.groundmesh#subdivisionsy
func (g *GroundMesh) SetSubdivisionsY(subdivisionsY float64) *GroundMesh {
	g.p.Set("subdivisionsY", subdivisionsY)
	return g
}
