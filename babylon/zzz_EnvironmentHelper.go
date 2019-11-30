// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EnvironmentHelper represents a babylon.js EnvironmentHelper.
// The Environment helper class can be used to add a fully featuread none expensive background to your scene.
// It includes by default a skybox and a ground relying on the BackgroundMaterial.
// It also helps with the default setup of your imageProcessing configuration.
type EnvironmentHelper struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (e *EnvironmentHelper) JSObject() js.Value { return e.p }

// EnvironmentHelper returns a EnvironmentHelper JavaScript class.
func (ba *Babylon) EnvironmentHelper() *EnvironmentHelper {
	p := ba.ctx.Get("EnvironmentHelper")
	return EnvironmentHelperFromJSObject(p)
}

// EnvironmentHelperFromJSObject returns a wrapped EnvironmentHelper JavaScript class.
func EnvironmentHelperFromJSObject(p js.Value) *EnvironmentHelper {
	return &EnvironmentHelper{p: p}
}

// NewEnvironmentHelper returns a new EnvironmentHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.environmenthelper
func (ba *Babylon) NewEnvironmentHelper(options js.Value, scene *Scene) *EnvironmentHelper {
	p := ba.ctx.Get("EnvironmentHelper").New(options, scene.JSObject())
	return EnvironmentHelperFromJSObject(p)
}

// TODO: methods
