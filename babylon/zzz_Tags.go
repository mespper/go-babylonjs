// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Tags represents a babylon.js Tags.
// Class used to store custom tags
type Tags struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *Tags) JSObject() js.Value { return t.p }

// Tags returns a Tags JavaScript class.
func (ba *Babylon) Tags() *Tags {
	p := ba.ctx.Get("Tags")
	return TagsFromJSObject(p, ba.ctx)
}

// TagsFromJSObject returns a wrapped Tags JavaScript class.
func TagsFromJSObject(p js.Value, ctx js.Value) *Tags {
	return &Tags{p: p, ctx: ctx}
}

// TagsArrayToJSArray returns a JavaScript Array for the wrapped array.
func TagsArrayToJSArray(array []*Tags) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddTagsTo calls the AddTagsTo method on the Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags#addtagsto
func (t *Tags) AddTagsTo(obj JSObject, tagsString string) {

	args := make([]interface{}, 0, 2+0)

	if obj == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, obj.JSObject())
	}

	args = append(args, tagsString)

	t.p.Call("AddTagsTo", args...)
}

// DisableFor calls the DisableFor method on the Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags#disablefor
func (t *Tags) DisableFor(obj JSObject) {

	args := make([]interface{}, 0, 1+0)

	if obj == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, obj.JSObject())
	}

	t.p.Call("DisableFor", args...)
}

// EnableFor calls the EnableFor method on the Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags#enablefor
func (t *Tags) EnableFor(obj JSObject) {

	args := make([]interface{}, 0, 1+0)

	if obj == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, obj.JSObject())
	}

	t.p.Call("EnableFor", args...)
}

// TagsGetTagsOpts contains optional parameters for Tags.GetTags.
type TagsGetTagsOpts struct {
	AsString *bool
}

// GetTags calls the GetTags method on the Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags#gettags
func (t *Tags) GetTags(obj JSObject, opts *TagsGetTagsOpts) js.Value {
	if opts == nil {
		opts = &TagsGetTagsOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if obj == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, obj.JSObject())
	}

	if opts.AsString == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AsString)
	}

	retVal := t.p.Call("GetTags", args...)
	return retVal
}

// HasTags calls the HasTags method on the Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags#hastags
func (t *Tags) HasTags(obj JSObject) bool {

	args := make([]interface{}, 0, 1+0)

	if obj == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, obj.JSObject())
	}

	retVal := t.p.Call("HasTags", args...)
	return retVal.Bool()
}

// MatchesQuery calls the MatchesQuery method on the Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags#matchesquery
func (t *Tags) MatchesQuery(obj JSObject, tagsQuery string) bool {

	args := make([]interface{}, 0, 2+0)

	if obj == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, obj.JSObject())
	}

	args = append(args, tagsQuery)

	retVal := t.p.Call("MatchesQuery", args...)
	return retVal.Bool()
}

// RemoveTagsFrom calls the RemoveTagsFrom method on the Tags object.
//
// https://doc.babylonjs.com/api/classes/babylon.tags#removetagsfrom
func (t *Tags) RemoveTagsFrom(obj JSObject, tagsString string) {

	args := make([]interface{}, 0, 2+0)

	if obj == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, obj.JSObject())
	}

	args = append(args, tagsString)

	t.p.Call("RemoveTagsFrom", args...)
}
