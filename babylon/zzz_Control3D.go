// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Control3D represents a babylon.js Control3D.
// Class used as base class for controls
type Control3D struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *Control3D) JSObject() js.Value { return c.p }

// Control3D returns a Control3D JavaScript class.
func (ba *Babylon) Control3D() *Control3D {
	p := ba.ctx.Get("Control3D")
	return Control3DFromJSObject(p, ba.ctx)
}

// Control3DFromJSObject returns a wrapped Control3D JavaScript class.
func Control3DFromJSObject(p js.Value, ctx js.Value) *Control3D {
	return &Control3D{p: p, ctx: ctx}
}

// Control3DArrayToJSArray returns a JavaScript Array for the wrapped array.
func Control3DArrayToJSArray(array []*Control3D) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewControl3DOpts contains optional parameters for NewControl3D.
type NewControl3DOpts struct {
	Name *string
}

// NewControl3D returns a new Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d
func (ba *Babylon) NewControl3D(opts *NewControl3DOpts) *Control3D {
	if opts == nil {
		opts = &NewControl3DOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("Control3D").New(args...)
	return Control3DFromJSObject(p, ba.ctx)
}

// AddBehavior calls the AddBehavior method on the Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#addbehavior
func (c *Control3D) AddBehavior(behavior js.Value) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := c.p.Call("addBehavior", args...)
	return Control3DFromJSObject(retVal, c.ctx)
}

// Dispose calls the Dispose method on the Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#dispose
func (c *Control3D) Dispose() {

	c.p.Call("dispose")
}

// GetBehaviorByName calls the GetBehaviorByName method on the Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#getbehaviorbyname
func (c *Control3D) GetBehaviorByName(name string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := c.p.Call("getBehaviorByName", args...)
	return retVal
}

// GetClassName calls the GetClassName method on the Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#getclassname
func (c *Control3D) GetClassName() string {

	retVal := c.p.Call("getClassName")
	return retVal.String()
}

// LinkToTransformNode calls the LinkToTransformNode method on the Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#linktotransformnode
func (c *Control3D) LinkToTransformNode(node *TransformNode) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, node.JSObject())

	retVal := c.p.Call("linkToTransformNode", args...)
	return Control3DFromJSObject(retVal, c.ctx)
}

// RemoveBehavior calls the RemoveBehavior method on the Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#removebehavior
func (c *Control3D) RemoveBehavior(behavior js.Value) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := c.p.Call("removeBehavior", args...)
	return Control3DFromJSObject(retVal, c.ctx)
}

// Behaviors returns the Behaviors property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#behaviors
func (c *Control3D) Behaviors() js.Value {
	retVal := c.p.Get("behaviors")
	return retVal
}

// SetBehaviors sets the Behaviors property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#behaviors
func (c *Control3D) SetBehaviors(behaviors js.Value) *Control3D {
	c.p.Set("behaviors", behaviors)
	return c
}

// IsVisible returns the IsVisible property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#isvisible
func (c *Control3D) IsVisible() bool {
	retVal := c.p.Get("isVisible")
	return retVal.Bool()
}

// SetIsVisible sets the IsVisible property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#isvisible
func (c *Control3D) SetIsVisible(isVisible bool) *Control3D {
	c.p.Set("isVisible", isVisible)
	return c
}

// Mesh returns the Mesh property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#mesh
func (c *Control3D) Mesh() *AbstractMesh {
	retVal := c.p.Get("mesh")
	return AbstractMeshFromJSObject(retVal, c.ctx)
}

// SetMesh sets the Mesh property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#mesh
func (c *Control3D) SetMesh(mesh *AbstractMesh) *Control3D {
	c.p.Set("mesh", mesh.JSObject())
	return c
}

// Name returns the Name property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#name
func (c *Control3D) Name() string {
	retVal := c.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#name
func (c *Control3D) SetName(name string) *Control3D {
	c.p.Set("name", name)
	return c
}

// Node returns the Node property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#node
func (c *Control3D) Node() *TransformNode {
	retVal := c.p.Get("node")
	return TransformNodeFromJSObject(retVal, c.ctx)
}

// SetNode sets the Node property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#node
func (c *Control3D) SetNode(node *TransformNode) *Control3D {
	c.p.Set("node", node.JSObject())
	return c
}

// OnPointerClickObservable returns the OnPointerClickObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerclickobservable
func (c *Control3D) OnPointerClickObservable() *Observable {
	retVal := c.p.Get("onPointerClickObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnPointerClickObservable sets the OnPointerClickObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerclickobservable
func (c *Control3D) SetOnPointerClickObservable(onPointerClickObservable *Observable) *Control3D {
	c.p.Set("onPointerClickObservable", onPointerClickObservable.JSObject())
	return c
}

// OnPointerDownObservable returns the OnPointerDownObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerdownobservable
func (c *Control3D) OnPointerDownObservable() *Observable {
	retVal := c.p.Get("onPointerDownObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnPointerDownObservable sets the OnPointerDownObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerdownobservable
func (c *Control3D) SetOnPointerDownObservable(onPointerDownObservable *Observable) *Control3D {
	c.p.Set("onPointerDownObservable", onPointerDownObservable.JSObject())
	return c
}

// OnPointerEnterObservable returns the OnPointerEnterObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerenterobservable
func (c *Control3D) OnPointerEnterObservable() *Observable {
	retVal := c.p.Get("onPointerEnterObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnPointerEnterObservable sets the OnPointerEnterObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerenterobservable
func (c *Control3D) SetOnPointerEnterObservable(onPointerEnterObservable *Observable) *Control3D {
	c.p.Set("onPointerEnterObservable", onPointerEnterObservable.JSObject())
	return c
}

// OnPointerMoveObservable returns the OnPointerMoveObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointermoveobservable
func (c *Control3D) OnPointerMoveObservable() *Observable {
	retVal := c.p.Get("onPointerMoveObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnPointerMoveObservable sets the OnPointerMoveObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointermoveobservable
func (c *Control3D) SetOnPointerMoveObservable(onPointerMoveObservable *Observable) *Control3D {
	c.p.Set("onPointerMoveObservable", onPointerMoveObservable.JSObject())
	return c
}

// OnPointerOutObservable returns the OnPointerOutObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointeroutobservable
func (c *Control3D) OnPointerOutObservable() *Observable {
	retVal := c.p.Get("onPointerOutObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnPointerOutObservable sets the OnPointerOutObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointeroutobservable
func (c *Control3D) SetOnPointerOutObservable(onPointerOutObservable *Observable) *Control3D {
	c.p.Set("onPointerOutObservable", onPointerOutObservable.JSObject())
	return c
}

// OnPointerUpObservable returns the OnPointerUpObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerupobservable
func (c *Control3D) OnPointerUpObservable() *Observable {
	retVal := c.p.Get("onPointerUpObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnPointerUpObservable sets the OnPointerUpObservable property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#onpointerupobservable
func (c *Control3D) SetOnPointerUpObservable(onPointerUpObservable *Observable) *Control3D {
	c.p.Set("onPointerUpObservable", onPointerUpObservable.JSObject())
	return c
}

// Parent returns the Parent property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#parent
func (c *Control3D) Parent() *Container3D {
	retVal := c.p.Get("parent")
	return Container3DFromJSObject(retVal, c.ctx)
}

// SetParent sets the Parent property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#parent
func (c *Control3D) SetParent(parent *Container3D) *Control3D {
	c.p.Set("parent", parent.JSObject())
	return c
}

// PointerDownAnimation returns the PointerDownAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointerdownanimation
func (c *Control3D) PointerDownAnimation() js.Value {
	retVal := c.p.Get("pointerDownAnimation")
	return retVal
}

// SetPointerDownAnimation sets the PointerDownAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointerdownanimation
func (c *Control3D) SetPointerDownAnimation(pointerDownAnimation func()) *Control3D {
	c.p.Set("pointerDownAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerDownAnimation(); return nil }))
	return c
}

// PointerEnterAnimation returns the PointerEnterAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointerenteranimation
func (c *Control3D) PointerEnterAnimation() js.Value {
	retVal := c.p.Get("pointerEnterAnimation")
	return retVal
}

// SetPointerEnterAnimation sets the PointerEnterAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointerenteranimation
func (c *Control3D) SetPointerEnterAnimation(pointerEnterAnimation func()) *Control3D {
	c.p.Set("pointerEnterAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerEnterAnimation(); return nil }))
	return c
}

// PointerOutAnimation returns the PointerOutAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointeroutanimation
func (c *Control3D) PointerOutAnimation() js.Value {
	retVal := c.p.Get("pointerOutAnimation")
	return retVal
}

// SetPointerOutAnimation sets the PointerOutAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointeroutanimation
func (c *Control3D) SetPointerOutAnimation(pointerOutAnimation func()) *Control3D {
	c.p.Set("pointerOutAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerOutAnimation(); return nil }))
	return c
}

// PointerUpAnimation returns the PointerUpAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointerupanimation
func (c *Control3D) PointerUpAnimation() js.Value {
	retVal := c.p.Get("pointerUpAnimation")
	return retVal
}

// SetPointerUpAnimation sets the PointerUpAnimation property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#pointerupanimation
func (c *Control3D) SetPointerUpAnimation(pointerUpAnimation func()) *Control3D {
	c.p.Set("pointerUpAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerUpAnimation(); return nil }))
	return c
}

// Position returns the Position property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#position
func (c *Control3D) Position() *Vector3 {
	retVal := c.p.Get("position")
	return Vector3FromJSObject(retVal, c.ctx)
}

// SetPosition sets the Position property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#position
func (c *Control3D) SetPosition(position *Vector3) *Control3D {
	c.p.Set("position", position.JSObject())
	return c
}

// Scaling returns the Scaling property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#scaling
func (c *Control3D) Scaling() *Vector3 {
	retVal := c.p.Get("scaling")
	return Vector3FromJSObject(retVal, c.ctx)
}

// SetScaling sets the Scaling property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#scaling
func (c *Control3D) SetScaling(scaling *Vector3) *Control3D {
	c.p.Set("scaling", scaling.JSObject())
	return c
}

// TypeName returns the TypeName property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#typename
func (c *Control3D) TypeName() string {
	retVal := c.p.Get("typeName")
	return retVal.String()
}

// SetTypeName sets the TypeName property of class Control3D.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d#typename
func (c *Control3D) SetTypeName(typeName string) *Control3D {
	c.p.Set("typeName", typeName)
	return c
}
