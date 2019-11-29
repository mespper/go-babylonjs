// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScreenshotTools represents a babylon.js ScreenshotTools.
// Class containing a set of static utilities functions for screenshots
type ScreenshotTools struct{}

// JSObject returns the underlying js.Value.
func (s *ScreenshotTools) JSObject() js.Value { return s.p }

// ScreenshotTools returns a ScreenshotTools JavaScript class.
func (b *Babylon) ScreenshotTools() *ScreenshotTools {
	p := b.ctx.Get("ScreenshotTools")
	return ScreenshotToolsFromJSObject(p)
}

// ScreenshotToolsFromJSObject returns a wrapped ScreenshotTools JavaScript class.
func ScreenshotToolsFromJSObject(p js.Value) *ScreenshotTools {
	return &ScreenshotTools{p: p}
}

// NewScreenshotTools returns a new ScreenshotTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.screenshottools
func (b *Babylon) NewScreenshotTools(todo parameters) *ScreenshotTools {
	p := b.ctx.Get("ScreenshotTools").New(todo)
	return ScreenshotToolsFromJSObject(p)
}

// TODO: methods