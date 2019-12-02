// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KhronosTextureContainer represents a babylon.js KhronosTextureContainer.
// for description see <a href="https://www.khronos.org/opengles/sdk/tools/KTX/">https://www.khronos.org/opengles/sdk/tools/KTX/</a>
// for file layout see <a href="https://www.khronos.org/opengles/sdk/tools/KTX/file_format_spec/">https://www.khronos.org/opengles/sdk/tools/KTX/file_format_spec/</a>
type KhronosTextureContainer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KhronosTextureContainer) JSObject() js.Value { return k.p }

// KhronosTextureContainer returns a KhronosTextureContainer JavaScript class.
func (ba *Babylon) KhronosTextureContainer() *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer")
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// KhronosTextureContainerFromJSObject returns a wrapped KhronosTextureContainer JavaScript class.
func KhronosTextureContainerFromJSObject(p js.Value, ctx js.Value) *KhronosTextureContainer {
	return &KhronosTextureContainer{p: p, ctx: ctx}
}

// KhronosTextureContainerArrayToJSArray returns a JavaScript Array for the wrapped array.
func KhronosTextureContainerArrayToJSArray(array []*KhronosTextureContainer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewKhronosTextureContainerOpts contains optional parameters for NewKhronosTextureContainer.
type NewKhronosTextureContainerOpts struct {
	ThreeDExpected       *bool
	TextureArrayExpected *bool
}

// NewKhronosTextureContainer returns a new KhronosTextureContainer object.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer
func (ba *Babylon) NewKhronosTextureContainer(arrayBuffer interface{}, facesExpected float64, opts *NewKhronosTextureContainerOpts) *KhronosTextureContainer {
	if opts == nil {
		opts = &NewKhronosTextureContainerOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, arrayBuffer)
	args = append(args, facesExpected)

	if opts.ThreeDExpected == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ThreeDExpected)
	}
	if opts.TextureArrayExpected == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TextureArrayExpected)
	}

	p := ba.ctx.Get("KhronosTextureContainer").New(args...)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// ArrayBuffer returns the ArrayBuffer property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#arraybuffer
func (k *KhronosTextureContainer) ArrayBuffer() interface{} {
	retVal := k.p.Get("arrayBuffer")
	return retVal
}

// SetArrayBuffer sets the ArrayBuffer property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#arraybuffer
func (k *KhronosTextureContainer) SetArrayBuffer(arrayBuffer interface{}) *KhronosTextureContainer {
	k.p.Set("arrayBuffer", arrayBuffer)
	return k
}

// BytesOfKeyValueData returns the BytesOfKeyValueData property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#bytesofkeyvaluedata
func (k *KhronosTextureContainer) BytesOfKeyValueData() float64 {
	retVal := k.p.Get("bytesOfKeyValueData")
	return retVal.Float()
}

// SetBytesOfKeyValueData sets the BytesOfKeyValueData property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#bytesofkeyvaluedata
func (k *KhronosTextureContainer) SetBytesOfKeyValueData(bytesOfKeyValueData float64) *KhronosTextureContainer {
	k.p.Set("bytesOfKeyValueData", bytesOfKeyValueData)
	return k
}

// GlBaseInternalFormat returns the GlBaseInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glbaseinternalformat
func (k *KhronosTextureContainer) GlBaseInternalFormat() float64 {
	retVal := k.p.Get("glBaseInternalFormat")
	return retVal.Float()
}

// SetGlBaseInternalFormat sets the GlBaseInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glbaseinternalformat
func (k *KhronosTextureContainer) SetGlBaseInternalFormat(glBaseInternalFormat float64) *KhronosTextureContainer {
	k.p.Set("glBaseInternalFormat", glBaseInternalFormat)
	return k
}

// GlFormat returns the GlFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glformat
func (k *KhronosTextureContainer) GlFormat() float64 {
	retVal := k.p.Get("glFormat")
	return retVal.Float()
}

// SetGlFormat sets the GlFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glformat
func (k *KhronosTextureContainer) SetGlFormat(glFormat float64) *KhronosTextureContainer {
	k.p.Set("glFormat", glFormat)
	return k
}

// GlInternalFormat returns the GlInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glinternalformat
func (k *KhronosTextureContainer) GlInternalFormat() float64 {
	retVal := k.p.Get("glInternalFormat")
	return retVal.Float()
}

// SetGlInternalFormat sets the GlInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glinternalformat
func (k *KhronosTextureContainer) SetGlInternalFormat(glInternalFormat float64) *KhronosTextureContainer {
	k.p.Set("glInternalFormat", glInternalFormat)
	return k
}

// GlType returns the GlType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltype
func (k *KhronosTextureContainer) GlType() float64 {
	retVal := k.p.Get("glType")
	return retVal.Float()
}

// SetGlType sets the GlType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltype
func (k *KhronosTextureContainer) SetGlType(glType float64) *KhronosTextureContainer {
	k.p.Set("glType", glType)
	return k
}

// GlTypeSize returns the GlTypeSize property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltypesize
func (k *KhronosTextureContainer) GlTypeSize() float64 {
	retVal := k.p.Get("glTypeSize")
	return retVal.Float()
}

// SetGlTypeSize sets the GlTypeSize property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltypesize
func (k *KhronosTextureContainer) SetGlTypeSize(glTypeSize float64) *KhronosTextureContainer {
	k.p.Set("glTypeSize", glTypeSize)
	return k
}

// IsInvalid returns the IsInvalid property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#isinvalid
func (k *KhronosTextureContainer) IsInvalid() bool {
	retVal := k.p.Get("isInvalid")
	return retVal.Bool()
}

// SetIsInvalid sets the IsInvalid property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#isinvalid
func (k *KhronosTextureContainer) SetIsInvalid(isInvalid bool) *KhronosTextureContainer {
	k.p.Set("isInvalid", isInvalid)
	return k
}

// LoadType returns the LoadType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#loadtype
func (k *KhronosTextureContainer) LoadType() float64 {
	retVal := k.p.Get("loadType")
	return retVal.Float()
}

// SetLoadType sets the LoadType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#loadtype
func (k *KhronosTextureContainer) SetLoadType(loadType float64) *KhronosTextureContainer {
	k.p.Set("loadType", loadType)
	return k
}

// NumberOfArrayElements returns the NumberOfArrayElements property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofarrayelements
func (k *KhronosTextureContainer) NumberOfArrayElements() float64 {
	retVal := k.p.Get("numberOfArrayElements")
	return retVal.Float()
}

// SetNumberOfArrayElements sets the NumberOfArrayElements property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofarrayelements
func (k *KhronosTextureContainer) SetNumberOfArrayElements(numberOfArrayElements float64) *KhronosTextureContainer {
	k.p.Set("numberOfArrayElements", numberOfArrayElements)
	return k
}

// NumberOfFaces returns the NumberOfFaces property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberoffaces
func (k *KhronosTextureContainer) NumberOfFaces() float64 {
	retVal := k.p.Get("numberOfFaces")
	return retVal.Float()
}

// SetNumberOfFaces sets the NumberOfFaces property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberoffaces
func (k *KhronosTextureContainer) SetNumberOfFaces(numberOfFaces float64) *KhronosTextureContainer {
	k.p.Set("numberOfFaces", numberOfFaces)
	return k
}

// NumberOfMipmapLevels returns the NumberOfMipmapLevels property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofmipmaplevels
func (k *KhronosTextureContainer) NumberOfMipmapLevels() float64 {
	retVal := k.p.Get("numberOfMipmapLevels")
	return retVal.Float()
}

// SetNumberOfMipmapLevels sets the NumberOfMipmapLevels property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofmipmaplevels
func (k *KhronosTextureContainer) SetNumberOfMipmapLevels(numberOfMipmapLevels float64) *KhronosTextureContainer {
	k.p.Set("numberOfMipmapLevels", numberOfMipmapLevels)
	return k
}

// PixelDepth returns the PixelDepth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixeldepth
func (k *KhronosTextureContainer) PixelDepth() float64 {
	retVal := k.p.Get("pixelDepth")
	return retVal.Float()
}

// SetPixelDepth sets the PixelDepth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixeldepth
func (k *KhronosTextureContainer) SetPixelDepth(pixelDepth float64) *KhronosTextureContainer {
	k.p.Set("pixelDepth", pixelDepth)
	return k
}

// PixelHeight returns the PixelHeight property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelheight
func (k *KhronosTextureContainer) PixelHeight() float64 {
	retVal := k.p.Get("pixelHeight")
	return retVal.Float()
}

// SetPixelHeight sets the PixelHeight property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelheight
func (k *KhronosTextureContainer) SetPixelHeight(pixelHeight float64) *KhronosTextureContainer {
	k.p.Set("pixelHeight", pixelHeight)
	return k
}

// PixelWidth returns the PixelWidth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelwidth
func (k *KhronosTextureContainer) PixelWidth() float64 {
	retVal := k.p.Get("pixelWidth")
	return retVal.Float()
}

// SetPixelWidth sets the PixelWidth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelwidth
func (k *KhronosTextureContainer) SetPixelWidth(pixelWidth float64) *KhronosTextureContainer {
	k.p.Set("pixelWidth", pixelWidth)
	return k
}
