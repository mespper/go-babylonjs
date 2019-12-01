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

/*

// EmbedMode returns the EmbedMode property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#embedmode
func (i *IInspectorOptions) EmbedMode(embedMode bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(embedMode)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetEmbedMode sets the EmbedMode property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#embedmode
func (i *IInspectorOptions) SetEmbedMode(embedMode bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(embedMode)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// EnableClose returns the EnableClose property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enableclose
func (i *IInspectorOptions) EnableClose(enableClose bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(enableClose)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetEnableClose sets the EnableClose property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enableclose
func (i *IInspectorOptions) SetEnableClose(enableClose bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(enableClose)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// EnablePopup returns the EnablePopup property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enablepopup
func (i *IInspectorOptions) EnablePopup(enablePopup bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(enablePopup)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetEnablePopup sets the EnablePopup property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#enablepopup
func (i *IInspectorOptions) SetEnablePopup(enablePopup bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(enablePopup)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// ExplorerExtensibility returns the ExplorerExtensibility property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#explorerextensibility
func (i *IInspectorOptions) ExplorerExtensibility(explorerExtensibility *IExplorerExtensibilityGroup) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(explorerExtensibility.JSObject())
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetExplorerExtensibility sets the ExplorerExtensibility property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#explorerextensibility
func (i *IInspectorOptions) SetExplorerExtensibility(explorerExtensibility *IExplorerExtensibilityGroup) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(explorerExtensibility.JSObject())
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// GlobalRoot returns the GlobalRoot property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#globalroot
func (i *IInspectorOptions) GlobalRoot(globalRoot js.Value) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(globalRoot)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetGlobalRoot sets the GlobalRoot property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#globalroot
func (i *IInspectorOptions) SetGlobalRoot(globalRoot js.Value) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(globalRoot)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// HandleResize returns the HandleResize property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#handleresize
func (i *IInspectorOptions) HandleResize(handleResize bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(handleResize)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetHandleResize sets the HandleResize property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#handleresize
func (i *IInspectorOptions) SetHandleResize(handleResize bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(handleResize)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// InspectorURL returns the InspectorURL property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#inspectorurl
func (i *IInspectorOptions) InspectorURL(inspectorURL string) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(inspectorURL)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetInspectorURL sets the InspectorURL property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#inspectorurl
func (i *IInspectorOptions) SetInspectorURL(inspectorURL string) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(inspectorURL)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// Overlay returns the Overlay property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#overlay
func (i *IInspectorOptions) Overlay(overlay bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(overlay)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetOverlay sets the Overlay property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#overlay
func (i *IInspectorOptions) SetOverlay(overlay bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(overlay)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// ShowExplorer returns the ShowExplorer property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showexplorer
func (i *IInspectorOptions) ShowExplorer(showExplorer bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(showExplorer)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetShowExplorer sets the ShowExplorer property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showexplorer
func (i *IInspectorOptions) SetShowExplorer(showExplorer bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(showExplorer)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// ShowInspector returns the ShowInspector property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showinspector
func (i *IInspectorOptions) ShowInspector(showInspector bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(showInspector)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

// SetShowInspector sets the ShowInspector property of class IInspectorOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.iinspectoroptions#showinspector
func (i *IInspectorOptions) SetShowInspector(showInspector bool) *IInspectorOptions {
	p := ba.ctx.Get("IInspectorOptions").New(showInspector)
	return IInspectorOptionsFromJSObject(p, ba.ctx)
}

*/