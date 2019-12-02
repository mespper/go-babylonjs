// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IExportOptions represents a babylon.js IExportOptions.
// Holds a collection of exporter options and parameters
type IExportOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IExportOptions) JSObject() js.Value { return i.p }

// IExportOptions returns a IExportOptions JavaScript class.
func (ba *Babylon) IExportOptions() *IExportOptions {
	p := ba.ctx.Get("IExportOptions")
	return IExportOptionsFromJSObject(p, ba.ctx)
}

// IExportOptionsFromJSObject returns a wrapped IExportOptions JavaScript class.
func IExportOptionsFromJSObject(p js.Value, ctx js.Value) *IExportOptions {
	return &IExportOptions{p: p, ctx: ctx}
}

// IExportOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func IExportOptionsArrayToJSArray(array []*IExportOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// MetadataSelector calls the MetadataSelector method on the IExportOptions object.
//
// https://doc.babylonjs.com/api/classes/babylon.iexportoptions#metadataselector
func (i *IExportOptions) MetadataSelector(metadata interface{}) interface{} {

	args := make([]interface{}, 0, 1+0)

	args = append(args, metadata)

	retVal := i.p.Call("metadataSelector", args...)
	return retVal
}

// ShouldExportNode calls the ShouldExportNode method on the IExportOptions object.
//
// https://doc.babylonjs.com/api/classes/babylon.iexportoptions#shouldexportnode
func (i *IExportOptions) ShouldExportNode(node *Node) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, node.JSObject())

	retVal := i.p.Call("shouldExportNode", args...)
	return retVal.Bool()
}

// AnimationSampleRate returns the AnimationSampleRate property of class IExportOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iexportoptions#animationsamplerate
func (i *IExportOptions) AnimationSampleRate() float64 {
	retVal := i.p.Get("animationSampleRate")
	return retVal.Float()
}

// SetAnimationSampleRate sets the AnimationSampleRate property of class IExportOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iexportoptions#animationsamplerate
func (i *IExportOptions) SetAnimationSampleRate(animationSampleRate float64) *IExportOptions {
	i.p.Set("animationSampleRate", animationSampleRate)
	return i
}

// ExportWithoutWaitingForScene returns the ExportWithoutWaitingForScene property of class IExportOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iexportoptions#exportwithoutwaitingforscene
func (i *IExportOptions) ExportWithoutWaitingForScene() bool {
	retVal := i.p.Get("exportWithoutWaitingForScene")
	return retVal.Bool()
}

// SetExportWithoutWaitingForScene sets the ExportWithoutWaitingForScene property of class IExportOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iexportoptions#exportwithoutwaitingforscene
func (i *IExportOptions) SetExportWithoutWaitingForScene(exportWithoutWaitingForScene bool) *IExportOptions {
	i.p.Set("exportWithoutWaitingForScene", exportWithoutWaitingForScene)
	return i
}
