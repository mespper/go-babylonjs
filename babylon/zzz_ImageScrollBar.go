// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageScrollBar represents a babylon.js ImageScrollBar.
// Class used to create slider controls
type ImageScrollBar struct {
	*BaseSlider
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ImageScrollBar) JSObject() js.Value { return i.p }

// ImageScrollBar returns a ImageScrollBar JavaScript class.
func (ba *Babylon) ImageScrollBar() *ImageScrollBar {
	p := ba.ctx.Get("ImageScrollBar")
	return ImageScrollBarFromJSObject(p, ba.ctx)
}

// ImageScrollBarFromJSObject returns a wrapped ImageScrollBar JavaScript class.
func ImageScrollBarFromJSObject(p js.Value, ctx js.Value) *ImageScrollBar {
	return &ImageScrollBar{BaseSlider: BaseSliderFromJSObject(p, ctx), ctx: ctx}
}

// ImageScrollBarArrayToJSArray returns a JavaScript Array for the wrapped array.
func ImageScrollBarArrayToJSArray(array []*ImageScrollBar) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewImageScrollBarOpts contains optional parameters for NewImageScrollBar.
type NewImageScrollBarOpts struct {
	Name *string
}

// NewImageScrollBar returns a new ImageScrollBar object.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar
func (ba *Babylon) NewImageScrollBar(opts *NewImageScrollBarOpts) *ImageScrollBar {
	if opts == nil {
		opts = &NewImageScrollBarOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("ImageScrollBar").New(args...)
	return ImageScrollBarFromJSObject(p, ba.ctx)
}

// _draw calls the _draw method on the ImageScrollBar object.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#_draw
func (i *ImageScrollBar) _draw(context js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, context)

	i.p.Call("_draw", args...)
}

// _onPointerDown calls the _onPointerDown method on the ImageScrollBar object.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#_onpointerdown
func (i *ImageScrollBar) _onPointerDown(target *Control, coordinates *Vector2, pointerId float64, buttonIndex float64) bool {

	args := make([]interface{}, 0, 4+0)

	args = append(args, target.JSObject())
	args = append(args, coordinates.JSObject())
	args = append(args, pointerId)
	args = append(args, buttonIndex)

	retVal := i.p.Call("_onPointerDown", args...)
	return retVal.Bool()
}

// BackgroundImage returns the BackgroundImage property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#backgroundimage
func (i *ImageScrollBar) BackgroundImage() *Image {
	retVal := i.p.Get("backgroundImage")
	return ImageFromJSObject(retVal, i.ctx)
}

// SetBackgroundImage sets the BackgroundImage property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#backgroundimage
func (i *ImageScrollBar) SetBackgroundImage(backgroundImage *Image) *ImageScrollBar {
	i.p.Set("backgroundImage", backgroundImage.JSObject())
	return i
}

// BarImageHeight returns the BarImageHeight property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#barimageheight
func (i *ImageScrollBar) BarImageHeight() float64 {
	retVal := i.p.Get("barImageHeight")
	return retVal.Float()
}

// SetBarImageHeight sets the BarImageHeight property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#barimageheight
func (i *ImageScrollBar) SetBarImageHeight(barImageHeight float64) *ImageScrollBar {
	i.p.Set("barImageHeight", barImageHeight)
	return i
}

// Name returns the Name property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#name
func (i *ImageScrollBar) Name() string {
	retVal := i.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#name
func (i *ImageScrollBar) SetName(name string) *ImageScrollBar {
	i.p.Set("name", name)
	return i
}

// ThumbHeight returns the ThumbHeight property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#thumbheight
func (i *ImageScrollBar) ThumbHeight() float64 {
	retVal := i.p.Get("thumbHeight")
	return retVal.Float()
}

// SetThumbHeight sets the ThumbHeight property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#thumbheight
func (i *ImageScrollBar) SetThumbHeight(thumbHeight float64) *ImageScrollBar {
	i.p.Set("thumbHeight", thumbHeight)
	return i
}

// ThumbImage returns the ThumbImage property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#thumbimage
func (i *ImageScrollBar) ThumbImage() *Image {
	retVal := i.p.Get("thumbImage")
	return ImageFromJSObject(retVal, i.ctx)
}

// SetThumbImage sets the ThumbImage property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#thumbimage
func (i *ImageScrollBar) SetThumbImage(thumbImage *Image) *ImageScrollBar {
	i.p.Set("thumbImage", thumbImage.JSObject())
	return i
}

// ThumbLength returns the ThumbLength property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#thumblength
func (i *ImageScrollBar) ThumbLength() float64 {
	retVal := i.p.Get("thumbLength")
	return retVal.Float()
}

// SetThumbLength sets the ThumbLength property of class ImageScrollBar.
//
// https://doc.babylonjs.com/api/classes/babylon.imagescrollbar#thumblength
func (i *ImageScrollBar) SetThumbLength(thumbLength float64) *ImageScrollBar {
	i.p.Set("thumbLength", thumbLength)
	return i
}
