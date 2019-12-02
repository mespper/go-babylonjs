// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Button represents a babylon.js Button.
// Class used to create 2D buttons
type Button struct {
	*Rectangle
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *Button) JSObject() js.Value { return b.p }

// Button returns a Button JavaScript class.
func (ba *Babylon) Button() *Button {
	p := ba.ctx.Get("Button")
	return ButtonFromJSObject(p, ba.ctx)
}

// ButtonFromJSObject returns a wrapped Button JavaScript class.
func ButtonFromJSObject(p js.Value, ctx js.Value) *Button {
	return &Button{Rectangle: RectangleFromJSObject(p, ctx), ctx: ctx}
}

// ButtonArrayToJSArray returns a JavaScript Array for the wrapped array.
func ButtonArrayToJSArray(array []*Button) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewButtonOpts contains optional parameters for NewButton.
type NewButtonOpts struct {
	Name *string
}

// NewButton returns a new Button object.
//
// https://doc.babylonjs.com/api/classes/babylon.button
func (ba *Babylon) NewButton(opts *NewButtonOpts) *Button {
	if opts == nil {
		opts = &NewButtonOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("Button").New(args...)
	return ButtonFromJSObject(p, ba.ctx)
}

// CreateImageButton calls the CreateImageButton method on the Button object.
//
// https://doc.babylonjs.com/api/classes/babylon.button#createimagebutton
func (b *Button) CreateImageButton(name string, text string, imageUrl string) *Button {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, text)
	args = append(args, imageUrl)

	retVal := b.p.Call("CreateImageButton", args...)
	return ButtonFromJSObject(retVal, b.ctx)
}

// CreateImageOnlyButton calls the CreateImageOnlyButton method on the Button object.
//
// https://doc.babylonjs.com/api/classes/babylon.button#createimageonlybutton
func (b *Button) CreateImageOnlyButton(name string, imageUrl string) *Button {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, imageUrl)

	retVal := b.p.Call("CreateImageOnlyButton", args...)
	return ButtonFromJSObject(retVal, b.ctx)
}

// CreateImageWithCenterTextButton calls the CreateImageWithCenterTextButton method on the Button object.
//
// https://doc.babylonjs.com/api/classes/babylon.button#createimagewithcentertextbutton
func (b *Button) CreateImageWithCenterTextButton(name string, text string, imageUrl string) *Button {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, text)
	args = append(args, imageUrl)

	retVal := b.p.Call("CreateImageWithCenterTextButton", args...)
	return ButtonFromJSObject(retVal, b.ctx)
}

// CreateSimpleButton calls the CreateSimpleButton method on the Button object.
//
// https://doc.babylonjs.com/api/classes/babylon.button#createsimplebutton
func (b *Button) CreateSimpleButton(name string, text string) *Button {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, text)

	retVal := b.p.Call("CreateSimpleButton", args...)
	return ButtonFromJSObject(retVal, b.ctx)
}

// DelegatePickingToChildren returns the DelegatePickingToChildren property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#delegatepickingtochildren
func (b *Button) DelegatePickingToChildren() bool {
	retVal := b.p.Get("delegatePickingToChildren")
	return retVal.Bool()
}

// SetDelegatePickingToChildren sets the DelegatePickingToChildren property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#delegatepickingtochildren
func (b *Button) SetDelegatePickingToChildren(delegatePickingToChildren bool) *Button {
	b.p.Set("delegatePickingToChildren", delegatePickingToChildren)
	return b
}

// Image returns the Image property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#image
func (b *Button) Image() *Image {
	retVal := b.p.Get("image")
	return ImageFromJSObject(retVal, b.ctx)
}

// SetImage sets the Image property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#image
func (b *Button) SetImage(image *Image) *Button {
	b.p.Set("image", image.JSObject())
	return b
}

// Name returns the Name property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#name
func (b *Button) Name() string {
	retVal := b.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#name
func (b *Button) SetName(name string) *Button {
	b.p.Set("name", name)
	return b
}

// PointerDownAnimation returns the PointerDownAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointerdownanimation
func (b *Button) PointerDownAnimation() js.Value {
	retVal := b.p.Get("pointerDownAnimation")
	return retVal
}

// SetPointerDownAnimation sets the PointerDownAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointerdownanimation
func (b *Button) SetPointerDownAnimation(pointerDownAnimation func()) *Button {
	b.p.Set("pointerDownAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerDownAnimation(); return nil }))
	return b
}

// PointerEnterAnimation returns the PointerEnterAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointerenteranimation
func (b *Button) PointerEnterAnimation() js.Value {
	retVal := b.p.Get("pointerEnterAnimation")
	return retVal
}

// SetPointerEnterAnimation sets the PointerEnterAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointerenteranimation
func (b *Button) SetPointerEnterAnimation(pointerEnterAnimation func()) *Button {
	b.p.Set("pointerEnterAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerEnterAnimation(); return nil }))
	return b
}

// PointerOutAnimation returns the PointerOutAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointeroutanimation
func (b *Button) PointerOutAnimation() js.Value {
	retVal := b.p.Get("pointerOutAnimation")
	return retVal
}

// SetPointerOutAnimation sets the PointerOutAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointeroutanimation
func (b *Button) SetPointerOutAnimation(pointerOutAnimation func()) *Button {
	b.p.Set("pointerOutAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerOutAnimation(); return nil }))
	return b
}

// PointerUpAnimation returns the PointerUpAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointerupanimation
func (b *Button) PointerUpAnimation() js.Value {
	retVal := b.p.Get("pointerUpAnimation")
	return retVal
}

// SetPointerUpAnimation sets the PointerUpAnimation property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#pointerupanimation
func (b *Button) SetPointerUpAnimation(pointerUpAnimation func()) *Button {
	b.p.Set("pointerUpAnimation", js.FuncOf(func(this js.Value, args []js.Value) interface{} { pointerUpAnimation(); return nil }))
	return b
}

// TextBlock returns the TextBlock property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#textblock
func (b *Button) TextBlock() *TextBlock {
	retVal := b.p.Get("textBlock")
	return TextBlockFromJSObject(retVal, b.ctx)
}

// SetTextBlock sets the TextBlock property of class Button.
//
// https://doc.babylonjs.com/api/classes/babylon.button#textblock
func (b *Button) SetTextBlock(textBlock *TextBlock) *Button {
	b.p.Set("textBlock", textBlock.JSObject())
	return b
}
