// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextureOptimization represents a babylon.js TextureOptimization.
// Defines an optimization used to reduce the size of render target textures
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type TextureOptimization struct {
	*SceneOptimization
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TextureOptimization) JSObject() js.Value { return t.p }

// TextureOptimization returns a TextureOptimization JavaScript class.
func (ba *Babylon) TextureOptimization() *TextureOptimization {
	p := ba.ctx.Get("TextureOptimization")
	return TextureOptimizationFromJSObject(p, ba.ctx)
}

// TextureOptimizationFromJSObject returns a wrapped TextureOptimization JavaScript class.
func TextureOptimizationFromJSObject(p js.Value, ctx js.Value) *TextureOptimization {
	return &TextureOptimization{SceneOptimization: SceneOptimizationFromJSObject(p, ctx), ctx: ctx}
}

// TextureOptimizationArrayToJSArray returns a JavaScript Array for the wrapped array.
func TextureOptimizationArrayToJSArray(array []*TextureOptimization) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewTextureOptimizationOpts contains optional parameters for NewTextureOptimization.
type NewTextureOptimizationOpts struct {
	Priority    *float64
	MaximumSize *float64
	Step        *float64
}

// NewTextureOptimization returns a new TextureOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization
func (ba *Babylon) NewTextureOptimization(opts *NewTextureOptimizationOpts) *TextureOptimization {
	if opts == nil {
		opts = &NewTextureOptimizationOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Priority == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Priority)
	}
	if opts.MaximumSize == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaximumSize)
	}
	if opts.Step == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Step)
	}

	p := ba.ctx.Get("TextureOptimization").New(args...)
	return TextureOptimizationFromJSObject(p, ba.ctx)
}

// Apply calls the Apply method on the TextureOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#apply
func (t *TextureOptimization) Apply(scene *Scene, optimizer *SceneOptimizer) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, optimizer.JSObject())

	retVal := t.p.Call("apply", args...)
	return retVal.Bool()
}

// GetDescription calls the GetDescription method on the TextureOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#getdescription
func (t *TextureOptimization) GetDescription() string {

	retVal := t.p.Call("getDescription")
	return retVal.String()
}

// MaximumSize returns the MaximumSize property of class TextureOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#maximumsize
func (t *TextureOptimization) MaximumSize() float64 {
	retVal := t.p.Get("maximumSize")
	return retVal.Float()
}

// SetMaximumSize sets the MaximumSize property of class TextureOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#maximumsize
func (t *TextureOptimization) SetMaximumSize(maximumSize float64) *TextureOptimization {
	t.p.Set("maximumSize", maximumSize)
	return t
}

// Priority returns the Priority property of class TextureOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#priority
func (t *TextureOptimization) Priority() float64 {
	retVal := t.p.Get("priority")
	return retVal.Float()
}

// SetPriority sets the Priority property of class TextureOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#priority
func (t *TextureOptimization) SetPriority(priority float64) *TextureOptimization {
	t.p.Set("priority", priority)
	return t
}

// Step returns the Step property of class TextureOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#step
func (t *TextureOptimization) Step() float64 {
	retVal := t.p.Get("step")
	return retVal.Float()
}

// SetStep sets the Step property of class TextureOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization#step
func (t *TextureOptimization) SetStep(step float64) *TextureOptimization {
	t.p.Set("step", step)
	return t
}
