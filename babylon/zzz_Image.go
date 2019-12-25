// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Image represents a babylon.js Image.
// Class used to create 2D images
type Image struct {
	*Control
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *Image) JSObject() js.Value { return i.p }

// Image returns a Image JavaScript class.
func (gui *GUI) Image() *Image {
	p := gui.ctx.Get("Image")
	return ImageFromJSObject(p, gui.ctx)
}

// ImageFromJSObject returns a wrapped Image JavaScript class.
func ImageFromJSObject(p js.Value, ctx js.Value) *Image {
	return &Image{Control: ControlFromJSObject(p, ctx), ctx: ctx}
}

// ImageArrayToJSArray returns a JavaScript Array for the wrapped array.
func ImageArrayToJSArray(array []*Image) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewImageOpts contains optional parameters for NewImage.
type NewImageOpts struct {
	Name *string
	Url  *string
}

// NewImage returns a new Image object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image
func (gui *GUI) NewImage(opts *NewImageOpts) *Image {
	if opts == nil {
		opts = &NewImageOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}
	if opts.Url == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Url)
	}

	p := gui.ctx.Get("Image").New(args...)
	return ImageFromJSObject(p, gui.ctx)
}

// Contains calls the Contains method on the Image object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#contains
func (i *Image) Contains(x float64, y float64) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, x)

	args = append(args, y)

	retVal := i.p.Call("contains", args...)
	return retVal.Bool()
}

// Dispose calls the Dispose method on the Image object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#dispose
func (i *Image) Dispose() {

	i.p.Call("dispose")
}

// SynchronizeSizeWithContent calls the SynchronizeSizeWithContent method on the Image object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#synchronizesizewithcontent
func (i *Image) SynchronizeSizeWithContent() {

	i.p.Call("synchronizeSizeWithContent")
}

// _draw calls the _draw method on the Image object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#_draw
func (i *Image) _draw(context js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, context)

	i.p.Call("_draw", args...)
}

// AutoScale returns the AutoScale property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#autoscale
func (i *Image) AutoScale() bool {
	retVal := i.p.Get("autoScale")
	return retVal.Bool()
}

// SetAutoScale sets the AutoScale property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#autoscale
func (i *Image) SetAutoScale(autoScale bool) *Image {
	i.p.Set("autoScale", autoScale)
	return i
}

// CellHeight returns the CellHeight property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#cellheight
func (i *Image) CellHeight() float64 {
	retVal := i.p.Get("cellHeight")
	return retVal.Float()
}

// SetCellHeight sets the CellHeight property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#cellheight
func (i *Image) SetCellHeight(cellHeight float64) *Image {
	i.p.Set("cellHeight", cellHeight)
	return i
}

// CellId returns the CellId property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#cellid
func (i *Image) CellId() float64 {
	retVal := i.p.Get("cellId")
	return retVal.Float()
}

// SetCellId sets the CellId property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#cellid
func (i *Image) SetCellId(cellId float64) *Image {
	i.p.Set("cellId", cellId)
	return i
}

// CellWidth returns the CellWidth property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#cellwidth
func (i *Image) CellWidth() float64 {
	retVal := i.p.Get("cellWidth")
	return retVal.Float()
}

// SetCellWidth sets the CellWidth property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#cellwidth
func (i *Image) SetCellWidth(cellWidth float64) *Image {
	i.p.Set("cellWidth", cellWidth)
	return i
}

// DetectPointerOnOpaqueOnly returns the DetectPointerOnOpaqueOnly property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#detectpointeronopaqueonly
func (i *Image) DetectPointerOnOpaqueOnly() bool {
	retVal := i.p.Get("detectPointerOnOpaqueOnly")
	return retVal.Bool()
}

// SetDetectPointerOnOpaqueOnly sets the DetectPointerOnOpaqueOnly property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#detectpointeronopaqueonly
func (i *Image) SetDetectPointerOnOpaqueOnly(detectPointerOnOpaqueOnly bool) *Image {
	i.p.Set("detectPointerOnOpaqueOnly", detectPointerOnOpaqueOnly)
	return i
}

// DomImage returns the DomImage property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#domimage
func (i *Image) DomImage() js.Value {
	retVal := i.p.Get("domImage")
	return retVal
}

// SetDomImage sets the DomImage property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#domimage
func (i *Image) SetDomImage(domImage js.Value) *Image {
	i.p.Set("domImage", domImage)
	return i
}

// IsLoaded returns the IsLoaded property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#isloaded
func (i *Image) IsLoaded() bool {
	retVal := i.p.Get("isLoaded")
	return retVal.Bool()
}

// SetIsLoaded sets the IsLoaded property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#isloaded
func (i *Image) SetIsLoaded(isLoaded bool) *Image {
	i.p.Set("isLoaded", isLoaded)
	return i
}

// Name returns the Name property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#name
func (i *Image) Name() string {
	retVal := i.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#name
func (i *Image) SetName(name string) *Image {
	i.p.Set("name", name)
	return i
}

// OnImageLoadedObservable returns the OnImageLoadedObservable property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#onimageloadedobservable
func (i *Image) OnImageLoadedObservable() *Observable {
	retVal := i.p.Get("onImageLoadedObservable")
	return ObservableFromJSObject(retVal, i.ctx)
}

// SetOnImageLoadedObservable sets the OnImageLoadedObservable property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#onimageloadedobservable
func (i *Image) SetOnImageLoadedObservable(onImageLoadedObservable *Observable) *Image {
	i.p.Set("onImageLoadedObservable", onImageLoadedObservable.JSObject())
	return i
}

// OnSVGAttributesComputedObservable returns the OnSVGAttributesComputedObservable property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#onsvgattributescomputedobservable
func (i *Image) OnSVGAttributesComputedObservable() *Observable {
	retVal := i.p.Get("onSVGAttributesComputedObservable")
	return ObservableFromJSObject(retVal, i.ctx)
}

// SetOnSVGAttributesComputedObservable sets the OnSVGAttributesComputedObservable property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#onsvgattributescomputedobservable
func (i *Image) SetOnSVGAttributesComputedObservable(onSVGAttributesComputedObservable *Observable) *Image {
	i.p.Set("onSVGAttributesComputedObservable", onSVGAttributesComputedObservable.JSObject())
	return i
}

// PopulateNinePatchSlicesFromImage returns the PopulateNinePatchSlicesFromImage property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#populateninepatchslicesfromimage
func (i *Image) PopulateNinePatchSlicesFromImage() bool {
	retVal := i.p.Get("populateNinePatchSlicesFromImage")
	return retVal.Bool()
}

// SetPopulateNinePatchSlicesFromImage sets the PopulateNinePatchSlicesFromImage property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#populateninepatchslicesfromimage
func (i *Image) SetPopulateNinePatchSlicesFromImage(populateNinePatchSlicesFromImage bool) *Image {
	i.p.Set("populateNinePatchSlicesFromImage", populateNinePatchSlicesFromImage)
	return i
}

// STRETCH_EXTEND returns the STRETCH_EXTEND property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_extend
func (i *Image) STRETCH_EXTEND() float64 {
	retVal := i.p.Get("STRETCH_EXTEND")
	return retVal.Float()
}

// SetSTRETCH_EXTEND sets the STRETCH_EXTEND property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_extend
func (i *Image) SetSTRETCH_EXTEND(STRETCH_EXTEND float64) *Image {
	i.p.Set("STRETCH_EXTEND", STRETCH_EXTEND)
	return i
}

// STRETCH_FILL returns the STRETCH_FILL property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_fill
func (i *Image) STRETCH_FILL() float64 {
	retVal := i.p.Get("STRETCH_FILL")
	return retVal.Float()
}

// SetSTRETCH_FILL sets the STRETCH_FILL property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_fill
func (i *Image) SetSTRETCH_FILL(STRETCH_FILL float64) *Image {
	i.p.Set("STRETCH_FILL", STRETCH_FILL)
	return i
}

// STRETCH_NINE_PATCH returns the STRETCH_NINE_PATCH property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_nine_patch
func (i *Image) STRETCH_NINE_PATCH() float64 {
	retVal := i.p.Get("STRETCH_NINE_PATCH")
	return retVal.Float()
}

// SetSTRETCH_NINE_PATCH sets the STRETCH_NINE_PATCH property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_nine_patch
func (i *Image) SetSTRETCH_NINE_PATCH(STRETCH_NINE_PATCH float64) *Image {
	i.p.Set("STRETCH_NINE_PATCH", STRETCH_NINE_PATCH)
	return i
}

// STRETCH_NONE returns the STRETCH_NONE property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_none
func (i *Image) STRETCH_NONE() float64 {
	retVal := i.p.Get("STRETCH_NONE")
	return retVal.Float()
}

// SetSTRETCH_NONE sets the STRETCH_NONE property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_none
func (i *Image) SetSTRETCH_NONE(STRETCH_NONE float64) *Image {
	i.p.Set("STRETCH_NONE", STRETCH_NONE)
	return i
}

// STRETCH_UNIFORM returns the STRETCH_UNIFORM property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_uniform
func (i *Image) STRETCH_UNIFORM() float64 {
	retVal := i.p.Get("STRETCH_UNIFORM")
	return retVal.Float()
}

// SetSTRETCH_UNIFORM sets the STRETCH_UNIFORM property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch_uniform
func (i *Image) SetSTRETCH_UNIFORM(STRETCH_UNIFORM float64) *Image {
	i.p.Set("STRETCH_UNIFORM", STRETCH_UNIFORM)
	return i
}

// SliceBottom returns the SliceBottom property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#slicebottom
func (i *Image) SliceBottom() float64 {
	retVal := i.p.Get("sliceBottom")
	return retVal.Float()
}

// SetSliceBottom sets the SliceBottom property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#slicebottom
func (i *Image) SetSliceBottom(sliceBottom float64) *Image {
	i.p.Set("sliceBottom", sliceBottom)
	return i
}

// SliceLeft returns the SliceLeft property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sliceleft
func (i *Image) SliceLeft() float64 {
	retVal := i.p.Get("sliceLeft")
	return retVal.Float()
}

// SetSliceLeft sets the SliceLeft property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sliceleft
func (i *Image) SetSliceLeft(sliceLeft float64) *Image {
	i.p.Set("sliceLeft", sliceLeft)
	return i
}

// SliceRight returns the SliceRight property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sliceright
func (i *Image) SliceRight() float64 {
	retVal := i.p.Get("sliceRight")
	return retVal.Float()
}

// SetSliceRight sets the SliceRight property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sliceright
func (i *Image) SetSliceRight(sliceRight float64) *Image {
	i.p.Set("sliceRight", sliceRight)
	return i
}

// SliceTop returns the SliceTop property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#slicetop
func (i *Image) SliceTop() float64 {
	retVal := i.p.Get("sliceTop")
	return retVal.Float()
}

// SetSliceTop sets the SliceTop property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#slicetop
func (i *Image) SetSliceTop(sliceTop float64) *Image {
	i.p.Set("sliceTop", sliceTop)
	return i
}

// Source returns the Source property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#source
func (i *Image) Source() string {
	retVal := i.p.Get("source")
	return retVal.String()
}

// SetSource sets the Source property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#source
func (i *Image) SetSource(source string) *Image {
	i.p.Set("source", source)
	return i
}

// SourceHeight returns the SourceHeight property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourceheight
func (i *Image) SourceHeight() float64 {
	retVal := i.p.Get("sourceHeight")
	return retVal.Float()
}

// SetSourceHeight sets the SourceHeight property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourceheight
func (i *Image) SetSourceHeight(sourceHeight float64) *Image {
	i.p.Set("sourceHeight", sourceHeight)
	return i
}

// SourceLeft returns the SourceLeft property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourceleft
func (i *Image) SourceLeft() float64 {
	retVal := i.p.Get("sourceLeft")
	return retVal.Float()
}

// SetSourceLeft sets the SourceLeft property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourceleft
func (i *Image) SetSourceLeft(sourceLeft float64) *Image {
	i.p.Set("sourceLeft", sourceLeft)
	return i
}

// SourceTop returns the SourceTop property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourcetop
func (i *Image) SourceTop() float64 {
	retVal := i.p.Get("sourceTop")
	return retVal.Float()
}

// SetSourceTop sets the SourceTop property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourcetop
func (i *Image) SetSourceTop(sourceTop float64) *Image {
	i.p.Set("sourceTop", sourceTop)
	return i
}

// SourceWidth returns the SourceWidth property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourcewidth
func (i *Image) SourceWidth() float64 {
	retVal := i.p.Get("sourceWidth")
	return retVal.Float()
}

// SetSourceWidth sets the SourceWidth property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#sourcewidth
func (i *Image) SetSourceWidth(sourceWidth float64) *Image {
	i.p.Set("sourceWidth", sourceWidth)
	return i
}

// Stretch returns the Stretch property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch
func (i *Image) Stretch() float64 {
	retVal := i.p.Get("stretch")
	return retVal.Float()
}

// SetStretch sets the Stretch property of class Image.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.image#stretch
func (i *Image) SetStretch(stretch float64) *Image {
	i.p.Set("stretch", stretch)
	return i
}
