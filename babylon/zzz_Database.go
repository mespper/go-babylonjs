// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Database represents a babylon.js Database.
// Class used to enable access to IndexedDB
//
// See: http://doc.babylonjs.com/how_to/caching_resources_in_indexeddb
type Database struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *Database) JSObject() js.Value { return d.p }

// Database returns a Database JavaScript class.
func (ba *Babylon) Database() *Database {
	p := ba.ctx.Get("Database")
	return DatabaseFromJSObject(p, ba.ctx)
}

// DatabaseFromJSObject returns a wrapped Database JavaScript class.
func DatabaseFromJSObject(p js.Value, ctx js.Value) *Database {
	return &Database{p: p, ctx: ctx}
}

// DatabaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func DatabaseArrayToJSArray(array []*Database) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDatabaseOpts contains optional parameters for NewDatabase.
type NewDatabaseOpts struct {
	DisableManifestCheck *bool
}

// NewDatabase returns a new Database object.
//
// https://doc.babylonjs.com/api/classes/babylon.database
func (ba *Babylon) NewDatabase(urlToScene string, callbackManifestChecked JSFunc, opts *NewDatabaseOpts) *Database {
	if opts == nil {
		opts = &NewDatabaseOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, urlToScene)
	args = append(args, js.FuncOf(callbackManifestChecked))

	if opts.DisableManifestCheck == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisableManifestCheck)
	}

	p := ba.ctx.Get("Database").New(args...)
	return DatabaseFromJSObject(p, ba.ctx)
}

// DatabaseLoadFileOpts contains optional parameters for Database.LoadFile.
type DatabaseLoadFileOpts struct {
	ProgressCallBack JSFunc
	ErrorCallback    JSFunc
	UseArrayBuffer   *bool
}

// LoadFile calls the LoadFile method on the Database object.
//
// https://doc.babylonjs.com/api/classes/babylon.database#loadfile
func (d *Database) LoadFile(url string, sceneLoaded JSFunc, opts *DatabaseLoadFileOpts) {
	if opts == nil {
		opts = &DatabaseLoadFileOpts{}
	}

	args := make([]interface{}, 0, 2+3)

	args = append(args, url)

	args = append(args, js.FuncOf(sceneLoaded))

	if opts.ProgressCallBack == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.ProgressCallBack) /* never freed! */)
	}
	if opts.ErrorCallback == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.ErrorCallback) /* never freed! */)
	}
	if opts.UseArrayBuffer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseArrayBuffer)
	}

	d.p.Call("loadFile", args...)
}

// LoadImage calls the LoadImage method on the Database object.
//
// https://doc.babylonjs.com/api/classes/babylon.database#loadimage
func (d *Database) LoadImage(url string, image js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, url)

	args = append(args, image)

	d.p.Call("loadImage", args...)
}

// Open calls the Open method on the Database object.
//
// https://doc.babylonjs.com/api/classes/babylon.database#open
func (d *Database) Open(successCallback JSFunc, errorCallback JSFunc) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, js.FuncOf(successCallback))

	args = append(args, js.FuncOf(errorCallback))

	d.p.Call("open", args...)
}

// EnableSceneOffline returns the EnableSceneOffline property of class Database.
//
// https://doc.babylonjs.com/api/classes/babylon.database#enablesceneoffline
func (d *Database) EnableSceneOffline() bool {
	retVal := d.p.Get("enableSceneOffline")
	return retVal.Bool()
}

// SetEnableSceneOffline sets the EnableSceneOffline property of class Database.
//
// https://doc.babylonjs.com/api/classes/babylon.database#enablesceneoffline
func (d *Database) SetEnableSceneOffline(enableSceneOffline bool) *Database {
	d.p.Set("enableSceneOffline", enableSceneOffline)
	return d
}

// EnableTexturesOffline returns the EnableTexturesOffline property of class Database.
//
// https://doc.babylonjs.com/api/classes/babylon.database#enabletexturesoffline
func (d *Database) EnableTexturesOffline() bool {
	retVal := d.p.Get("enableTexturesOffline")
	return retVal.Bool()
}

// SetEnableTexturesOffline sets the EnableTexturesOffline property of class Database.
//
// https://doc.babylonjs.com/api/classes/babylon.database#enabletexturesoffline
func (d *Database) SetEnableTexturesOffline(enableTexturesOffline bool) *Database {
	d.p.Set("enableTexturesOffline", enableTexturesOffline)
	return d
}

// IDBStorageEnabled returns the IDBStorageEnabled property of class Database.
//
// https://doc.babylonjs.com/api/classes/babylon.database#idbstorageenabled
func (d *Database) IDBStorageEnabled() bool {
	retVal := d.p.Get("IDBStorageEnabled")
	return retVal.Bool()
}

// SetIDBStorageEnabled sets the IDBStorageEnabled property of class Database.
//
// https://doc.babylonjs.com/api/classes/babylon.database#idbstorageenabled
func (d *Database) SetIDBStorageEnabled(IDBStorageEnabled bool) *Database {
	d.p.Set("IDBStorageEnabled", IDBStorageEnabled)
	return d
}
