// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IInspectorOptions represents a babylon.js IInspectorOptions.
// Interface used to define the options to use to create the Inspector
type IInspectorOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IInspectorOptions) JSObject() js.Value { return i.p }

// IInspectorOptions returns a IInspectorOptions JavaScript class.
func (ba *Babylon) IInspectorOptions() *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions")
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// IInspectorOptionsFromJSObject returns a wrapped IInspectorOptions JavaScript class.
func IInspectorOptionsFromJSObject(p js.Value, ctx js.Value) *IInspectorOptions {
	return &IInspectorOptions{p: p, ctx: ctx}
}

// IInspectorOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func IInspectorOptionsArrayToJSArray(array []*IInspectorOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// EmbedMode returns the EmbedMode property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#embedmode
func (i *IInspectorOptions) EmbedMode() bool {
	retVal := i.p.Get("embedMode")
	return retVal.Bool()
}

// SetEmbedMode sets the EmbedMode property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#embedmode
func (i *IInspectorOptions) SetEmbedMode(embedMode bool) *IInspectorOptions {
	i.p.Set("embedMode", embedMode)
	return i
}

// EnableClose returns the EnableClose property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enableclose
func (i *IInspectorOptions) EnableClose() bool {
	retVal := i.p.Get("enableClose")
	return retVal.Bool()
}

// SetEnableClose sets the EnableClose property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enableclose
func (i *IInspectorOptions) SetEnableClose(enableClose bool) *IInspectorOptions {
	i.p.Set("enableClose", enableClose)
	return i
}

// EnablePopup returns the EnablePopup property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enablepopup
func (i *IInspectorOptions) EnablePopup() bool {
	retVal := i.p.Get("enablePopup")
	return retVal.Bool()
}

// SetEnablePopup sets the EnablePopup property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enablepopup
func (i *IInspectorOptions) SetEnablePopup(enablePopup bool) *IInspectorOptions {
	i.p.Set("enablePopup", enablePopup)
	return i
}

// ExplorerExtensibility returns the ExplorerExtensibility property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#explorerextensibility
func (i *IInspectorOptions) ExplorerExtensibility() *IExplorerExtensibilityGroup {
	retVal := i.p.Get("explorerExtensibility")
	return IExplorerExtensibilityGroupFromJSObject(retVal, i.ctx)
}

// SetExplorerExtensibility sets the ExplorerExtensibility property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#explorerextensibility
func (i *IInspectorOptions) SetExplorerExtensibility(explorerExtensibility *IExplorerExtensibilityGroup) *IInspectorOptions {
	i.p.Set("explorerExtensibility", explorerExtensibility.JSObject())
	return i
}

// GlobalRoot returns the GlobalRoot property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#globalroot
func (i *IInspectorOptions) GlobalRoot() js.Value {
	retVal := i.p.Get("globalRoot")
	return retVal
}

// SetGlobalRoot sets the GlobalRoot property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#globalroot
func (i *IInspectorOptions) SetGlobalRoot(globalRoot js.Value) *IInspectorOptions {
	i.p.Set("globalRoot", globalRoot)
	return i
}

// HandleResize returns the HandleResize property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#handleresize
func (i *IInspectorOptions) HandleResize() bool {
	retVal := i.p.Get("handleResize")
	return retVal.Bool()
}

// SetHandleResize sets the HandleResize property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#handleresize
func (i *IInspectorOptions) SetHandleResize(handleResize bool) *IInspectorOptions {
	i.p.Set("handleResize", handleResize)
	return i
}

// InspectorURL returns the InspectorURL property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#inspectorurl
func (i *IInspectorOptions) InspectorURL() string {
	retVal := i.p.Get("inspectorURL")
	return retVal.String()
}

// SetInspectorURL sets the InspectorURL property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#inspectorurl
func (i *IInspectorOptions) SetInspectorURL(inspectorURL string) *IInspectorOptions {
	i.p.Set("inspectorURL", inspectorURL)
	return i
}

// Overlay returns the Overlay property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#overlay
func (i *IInspectorOptions) Overlay() bool {
	retVal := i.p.Get("overlay")
	return retVal.Bool()
}

// SetOverlay sets the Overlay property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#overlay
func (i *IInspectorOptions) SetOverlay(overlay bool) *IInspectorOptions {
	i.p.Set("overlay", overlay)
	return i
}

// ShowExplorer returns the ShowExplorer property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showexplorer
func (i *IInspectorOptions) ShowExplorer() bool {
	retVal := i.p.Get("showExplorer")
	return retVal.Bool()
}

// SetShowExplorer sets the ShowExplorer property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showexplorer
func (i *IInspectorOptions) SetShowExplorer(showExplorer bool) *IInspectorOptions {
	i.p.Set("showExplorer", showExplorer)
	return i
}

// ShowInspector returns the ShowInspector property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showinspector
func (i *IInspectorOptions) ShowInspector() bool {
	retVal := i.p.Get("showInspector")
	return retVal.Bool()
}

// SetShowInspector sets the ShowInspector property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showinspector
func (i *IInspectorOptions) SetShowInspector(showInspector bool) *IInspectorOptions {
	i.p.Set("showInspector", showInspector)
	return i
}
