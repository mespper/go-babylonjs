// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISceneLoaderPluginFactory represents a babylon.js ISceneLoaderPluginFactory.
// Interface used by SceneLoader plugin factory
type ISceneLoaderPluginFactory struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISceneLoaderPluginFactory) JSObject() js.Value { return i.p }

// ISceneLoaderPluginFactory returns a ISceneLoaderPluginFactory JavaScript class.
func (ba *Babylon) ISceneLoaderPluginFactory() *ISceneLoaderPluginFactory {
	p := ba.ctx.Get("ISceneLoaderPluginFactory")
	return ISceneLoaderPluginFactoryFromJSObject(p, ba.ctx)
}

// ISceneLoaderPluginFactoryFromJSObject returns a wrapped ISceneLoaderPluginFactory JavaScript class.
func ISceneLoaderPluginFactoryFromJSObject(p js.Value, ctx js.Value) *ISceneLoaderPluginFactory {
	return &ISceneLoaderPluginFactory{p: p, ctx: ctx}
}

// ISceneLoaderPluginFactoryArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISceneLoaderPluginFactoryArrayToJSArray(array []*ISceneLoaderPluginFactory) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CanDirectLoad calls the CanDirectLoad method on the ISceneLoaderPluginFactory object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginfactory#candirectload
func (i *ISceneLoaderPluginFactory) CanDirectLoad(data string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	retVal := i.p.Call("canDirectLoad", args...)
	return retVal.Bool()
}

// CreatePlugin calls the CreatePlugin method on the ISceneLoaderPluginFactory object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginfactory#createplugin
func (i *ISceneLoaderPluginFactory) CreatePlugin() *ISceneLoaderPlugin {

	retVal := i.p.Call("createPlugin")
	return ISceneLoaderPluginFromJSObject(retVal, i.ctx)
}

/*

// Name returns the Name property of class ISceneLoaderPluginFactory.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginfactory#name
func (i *ISceneLoaderPluginFactory) Name(name string) *ISceneLoaderPluginFactory {
	p := ba.ctx.Get("ISceneLoaderPluginFactory").New(name)
	return ISceneLoaderPluginFactoryFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class ISceneLoaderPluginFactory.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginfactory#name
func (i *ISceneLoaderPluginFactory) SetName(name string) *ISceneLoaderPluginFactory {
	p := ba.ctx.Get("ISceneLoaderPluginFactory").New(name)
	return ISceneLoaderPluginFactoryFromJSObject(p, ba.ctx)
}

*/