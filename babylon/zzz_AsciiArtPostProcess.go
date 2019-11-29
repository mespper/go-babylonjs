// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AsciiArtPostProcess represents a babylon.js AsciiArtPostProcess.
// AsciiArtPostProcess helps rendering everithing in Ascii Art.

//
// Simmply add it to your scene and let the nerd that lives in you have fun.
// Example usage: var pp = new AsciiArtPostProcess(&amp;quot;myAscii&amp;quot;, &amp;quot;20px Monospace&amp;quot;, camera);
type AsciiArtPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (a *AsciiArtPostProcess) JSObject() js.Value { return a.p }

// AsciiArtPostProcess returns a AsciiArtPostProcess JavaScript class.
func (b *Babylon) AsciiArtPostProcess() *AsciiArtPostProcess {
	p := b.ctx.Get("AsciiArtPostProcess")
	return AsciiArtPostProcessFromJSObject(p)
}

// AsciiArtPostProcessFromJSObject returns a wrapped AsciiArtPostProcess JavaScript class.
func AsciiArtPostProcessFromJSObject(p js.Value) *AsciiArtPostProcess {
	return &AsciiArtPostProcess{PostProcessFromJSObject(p)}
}

// NewAsciiArtPostProcess returns a new AsciiArtPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartpostprocess
func (b *Babylon) NewAsciiArtPostProcess(todo parameters) *AsciiArtPostProcess {
	p := b.ctx.Get("AsciiArtPostProcess").New(todo)
	return AsciiArtPostProcessFromJSObject(p)
}

// TODO: methods