// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScreenshotTools represents a babylon.js ScreenshotTools.
// Class containing a set of static utilities functions for screenshots
type ScreenshotTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ScreenshotTools) JSObject() js.Value { return s.p }

// ScreenshotTools returns a ScreenshotTools JavaScript class.
func (ba *Babylon) ScreenshotTools() *ScreenshotTools {
	p := ba.ctx.Get("ScreenshotTools")
	return ScreenshotToolsFromJSObject(p, ba.ctx)
}

// ScreenshotToolsFromJSObject returns a wrapped ScreenshotTools JavaScript class.
func ScreenshotToolsFromJSObject(p js.Value, ctx js.Value) *ScreenshotTools {
	return &ScreenshotTools{p: p, ctx: ctx}
}

// ScreenshotToolsArrayToJSArray returns a JavaScript Array for the wrapped array.
func ScreenshotToolsArrayToJSArray(array []*ScreenshotTools) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ScreenshotToolsCreateScreenshotOpts contains optional parameters for ScreenshotTools.CreateScreenshot.
type ScreenshotToolsCreateScreenshotOpts struct {
	SuccessCallback JSFunc
	MimeType        *string
}

// CreateScreenshot calls the CreateScreenshot method on the ScreenshotTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.screenshottools#createscreenshot
func (s *ScreenshotTools) CreateScreenshot(engine *Engine, camera *Camera, size *IScreenshotSize, opts *ScreenshotToolsCreateScreenshotOpts) {
	if opts == nil {
		opts = &ScreenshotToolsCreateScreenshotOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	if engine == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, engine.JSObject())
	}

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	if size == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, size.JSObject())
	}

	if opts.SuccessCallback == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.SuccessCallback) /* never freed! */)
	}
	if opts.MimeType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MimeType)
	}

	s.p.Call("CreateScreenshot", args...)
}

// ScreenshotToolsCreateScreenshotAsyncOpts contains optional parameters for ScreenshotTools.CreateScreenshotAsync.
type ScreenshotToolsCreateScreenshotAsyncOpts struct {
	MimeType *string
}

// CreateScreenshotAsync calls the CreateScreenshotAsync method on the ScreenshotTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.screenshottools#createscreenshotasync
func (s *ScreenshotTools) CreateScreenshotAsync(engine *Engine, camera *Camera, size JSObject, opts *ScreenshotToolsCreateScreenshotAsyncOpts) *Promise {
	if opts == nil {
		opts = &ScreenshotToolsCreateScreenshotAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	if engine == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, engine.JSObject())
	}

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	if size == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, size.JSObject())
	}

	if opts.MimeType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MimeType)
	}

	retVal := s.p.Call("CreateScreenshotAsync", args...)
	return PromiseFromJSObject(retVal, s.ctx)
}

// ScreenshotToolsCreateScreenshotUsingRenderTargetOpts contains optional parameters for ScreenshotTools.CreateScreenshotUsingRenderTarget.
type ScreenshotToolsCreateScreenshotUsingRenderTargetOpts struct {
	SuccessCallback JSFunc
	MimeType        *string
	Samples         *float64
	Antialiasing    *bool
	FileName        *string
}

// CreateScreenshotUsingRenderTarget calls the CreateScreenshotUsingRenderTarget method on the ScreenshotTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.screenshottools#createscreenshotusingrendertarget
func (s *ScreenshotTools) CreateScreenshotUsingRenderTarget(engine *Engine, camera *Camera, size *IScreenshotSize, opts *ScreenshotToolsCreateScreenshotUsingRenderTargetOpts) {
	if opts == nil {
		opts = &ScreenshotToolsCreateScreenshotUsingRenderTargetOpts{}
	}

	args := make([]interface{}, 0, 3+5)

	if engine == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, engine.JSObject())
	}

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	if size == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, size.JSObject())
	}

	if opts.SuccessCallback == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.SuccessCallback) /* never freed! */)
	}
	if opts.MimeType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MimeType)
	}
	if opts.Samples == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Samples)
	}
	if opts.Antialiasing == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Antialiasing)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	s.p.Call("CreateScreenshotUsingRenderTarget", args...)
}

// ScreenshotToolsCreateScreenshotUsingRenderTargetAsyncOpts contains optional parameters for ScreenshotTools.CreateScreenshotUsingRenderTargetAsync.
type ScreenshotToolsCreateScreenshotUsingRenderTargetAsyncOpts struct {
	MimeType     *string
	Samples      *float64
	Antialiasing *bool
	FileName     *string
}

// CreateScreenshotUsingRenderTargetAsync calls the CreateScreenshotUsingRenderTargetAsync method on the ScreenshotTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.screenshottools#createscreenshotusingrendertargetasync
func (s *ScreenshotTools) CreateScreenshotUsingRenderTargetAsync(engine *Engine, camera *Camera, size JSObject, opts *ScreenshotToolsCreateScreenshotUsingRenderTargetAsyncOpts) *Promise {
	if opts == nil {
		opts = &ScreenshotToolsCreateScreenshotUsingRenderTargetAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	if engine == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, engine.JSObject())
	}

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	if size == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, size.JSObject())
	}

	if opts.MimeType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MimeType)
	}
	if opts.Samples == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Samples)
	}
	if opts.Antialiasing == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Antialiasing)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := s.p.Call("CreateScreenshotUsingRenderTargetAsync", args...)
	return PromiseFromJSObject(retVal, s.ctx)
}
