// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HolographicButton represents a babylon.js HolographicButton.
// Class used to create a holographic button in 3D
type HolographicButton struct {
	*Button3D
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HolographicButton) JSObject() js.Value { return h.p }

// HolographicButton returns a HolographicButton JavaScript class.
func (ba *Babylon) HolographicButton() *HolographicButton {
	p := ba.ctx.Get("HolographicButton")
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// HolographicButtonFromJSObject returns a wrapped HolographicButton JavaScript class.
func HolographicButtonFromJSObject(p js.Value, ctx js.Value) *HolographicButton {
	return &HolographicButton{Button3D: Button3DFromJSObject(p, ctx), ctx: ctx}
}

// HolographicButtonArrayToJSArray returns a JavaScript Array for the wrapped array.
func HolographicButtonArrayToJSArray(array []*HolographicButton) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHolographicButtonOpts contains optional parameters for NewHolographicButton.
type NewHolographicButtonOpts struct {
	Name           *string
	ShareMaterials *bool
}

// NewHolographicButton returns a new HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton
func (ba *Babylon) NewHolographicButton(opts *NewHolographicButtonOpts) *HolographicButton {
	if opts == nil {
		opts = &NewHolographicButtonOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}
	if opts.ShareMaterials == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ShareMaterials)
	}

	p := ba.ctx.Get("HolographicButton").New(args...)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#dispose
func (h *HolographicButton) Dispose() {

	h.p.Call("dispose")
}

// BackMaterial returns the BackMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#backmaterial
func (h *HolographicButton) BackMaterial() *FluentMaterial {
	retVal := h.p.Get("backMaterial")
	return FluentMaterialFromJSObject(retVal, h.ctx)
}

// SetBackMaterial sets the BackMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#backmaterial
func (h *HolographicButton) SetBackMaterial(backMaterial *FluentMaterial) *HolographicButton {
	h.p.Set("backMaterial", backMaterial.JSObject())
	return h
}

// FrontMaterial returns the FrontMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#frontmaterial
func (h *HolographicButton) FrontMaterial() *FluentMaterial {
	retVal := h.p.Get("frontMaterial")
	return FluentMaterialFromJSObject(retVal, h.ctx)
}

// SetFrontMaterial sets the FrontMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#frontmaterial
func (h *HolographicButton) SetFrontMaterial(frontMaterial *FluentMaterial) *HolographicButton {
	h.p.Set("frontMaterial", frontMaterial.JSObject())
	return h
}

// ImageUrl returns the ImageUrl property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#imageurl
func (h *HolographicButton) ImageUrl() string {
	retVal := h.p.Get("imageUrl")
	return retVal.String()
}

// SetImageUrl sets the ImageUrl property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#imageurl
func (h *HolographicButton) SetImageUrl(imageUrl string) *HolographicButton {
	h.p.Set("imageUrl", imageUrl)
	return h
}

// PlateMaterial returns the PlateMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#platematerial
func (h *HolographicButton) PlateMaterial() *StandardMaterial {
	retVal := h.p.Get("plateMaterial")
	return StandardMaterialFromJSObject(retVal, h.ctx)
}

// SetPlateMaterial sets the PlateMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#platematerial
func (h *HolographicButton) SetPlateMaterial(plateMaterial *StandardMaterial) *HolographicButton {
	h.p.Set("plateMaterial", plateMaterial.JSObject())
	return h
}

// RenderingGroupId returns the RenderingGroupId property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#renderinggroupid
func (h *HolographicButton) RenderingGroupId() float64 {
	retVal := h.p.Get("renderingGroupId")
	return retVal.Float()
}

// SetRenderingGroupId sets the RenderingGroupId property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#renderinggroupid
func (h *HolographicButton) SetRenderingGroupId(renderingGroupId float64) *HolographicButton {
	h.p.Set("renderingGroupId", renderingGroupId)
	return h
}

// ShareMaterials returns the ShareMaterials property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#sharematerials
func (h *HolographicButton) ShareMaterials() bool {
	retVal := h.p.Get("shareMaterials")
	return retVal.Bool()
}

// SetShareMaterials sets the ShareMaterials property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#sharematerials
func (h *HolographicButton) SetShareMaterials(shareMaterials bool) *HolographicButton {
	h.p.Set("shareMaterials", shareMaterials)
	return h
}

// Text returns the Text property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#text
func (h *HolographicButton) Text() string {
	retVal := h.p.Get("text")
	return retVal.String()
}

// SetText sets the Text property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#text
func (h *HolographicButton) SetText(text string) *HolographicButton {
	h.p.Set("text", text)
	return h
}

// TooltipText returns the TooltipText property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#tooltiptext
func (h *HolographicButton) TooltipText() string {
	retVal := h.p.Get("tooltipText")
	return retVal.String()
}

// SetTooltipText sets the TooltipText property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#tooltiptext
func (h *HolographicButton) SetTooltipText(tooltipText string) *HolographicButton {
	h.p.Set("tooltipText", tooltipText)
	return h
}
