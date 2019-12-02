// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RecastJSCrowd represents a babylon.js RecastJSCrowd.
// Recast detour crowd implementation
type RecastJSCrowd struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RecastJSCrowd) JSObject() js.Value { return r.p }

// RecastJSCrowd returns a RecastJSCrowd JavaScript class.
func (ba *Babylon) RecastJSCrowd() *RecastJSCrowd {
	p := ba.ctx.Get("RecastJSCrowd")
	return RecastJSCrowdFromJSObject(p, ba.ctx)
}

// RecastJSCrowdFromJSObject returns a wrapped RecastJSCrowd JavaScript class.
func RecastJSCrowdFromJSObject(p js.Value, ctx js.Value) *RecastJSCrowd {
	return &RecastJSCrowd{p: p, ctx: ctx}
}

// RecastJSCrowdArrayToJSArray returns a JavaScript Array for the wrapped array.
func RecastJSCrowdArrayToJSArray(array []*RecastJSCrowd) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRecastJSCrowd returns a new RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd
func (ba *Babylon) NewRecastJSCrowd(plugin *RecastJSPlugin, maxAgents float64, maxAgentRadius float64, scene *Scene) *RecastJSCrowd {

	args := make([]interface{}, 0, 4+0)

	args = append(args, plugin.JSObject())
	args = append(args, maxAgents)
	args = append(args, maxAgentRadius)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("RecastJSCrowd").New(args...)
	return RecastJSCrowdFromJSObject(p, ba.ctx)
}

// AddAgent calls the AddAgent method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#addagent
func (r *RecastJSCrowd) AddAgent(pos *Vector3, parameters *IAgentParameters, transform *TransformNode) float64 {

	args := make([]interface{}, 0, 3+0)

	args = append(args, pos.JSObject())
	args = append(args, parameters.JSObject())
	args = append(args, transform.JSObject())

	retVal := r.p.Call("addAgent", args...)
	return retVal.Float()
}

// AgentGoto calls the AgentGoto method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#agentgoto
func (r *RecastJSCrowd) AgentGoto(index float64, destination *Vector3) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, index)
	args = append(args, destination.JSObject())

	r.p.Call("agentGoto", args...)
}

// Dispose calls the Dispose method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#dispose
func (r *RecastJSCrowd) Dispose() {

	r.p.Call("dispose")
}

// GetAgentPosition calls the GetAgentPosition method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#getagentposition
func (r *RecastJSCrowd) GetAgentPosition(index float64) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := r.p.Call("getAgentPosition", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// GetAgentVelocity calls the GetAgentVelocity method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#getagentvelocity
func (r *RecastJSCrowd) GetAgentVelocity(index float64) *Vector3 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := r.p.Call("getAgentVelocity", args...)
	return Vector3FromJSObject(retVal, r.ctx)
}

// GetAgents calls the GetAgents method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#getagents
func (r *RecastJSCrowd) GetAgents() float64 {

	retVal := r.p.Call("getAgents")
	return retVal.Float()
}

// GetDefaultQueryExtent calls the GetDefaultQueryExtent method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#getdefaultqueryextent
func (r *RecastJSCrowd) GetDefaultQueryExtent() *Vector3 {

	retVal := r.p.Call("getDefaultQueryExtent")
	return Vector3FromJSObject(retVal, r.ctx)
}

// RemoveAgent calls the RemoveAgent method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#removeagent
func (r *RecastJSCrowd) RemoveAgent(index float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	r.p.Call("removeAgent", args...)
}

// SetDefaultQueryExtent calls the SetDefaultQueryExtent method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#setdefaultqueryextent
func (r *RecastJSCrowd) SetDefaultQueryExtent(extent *Vector3) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, extent.JSObject())

	r.p.Call("setDefaultQueryExtent", args...)
}

// Update calls the Update method on the RecastJSCrowd object.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#update
func (r *RecastJSCrowd) Update(deltaTime float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, deltaTime)

	r.p.Call("update", args...)
}

// Agents returns the Agents property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#agents
func (r *RecastJSCrowd) Agents() float64 {
	retVal := r.p.Get("agents")
	return retVal.Float()
}

// SetAgents sets the Agents property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#agents
func (r *RecastJSCrowd) SetAgents(agents float64) *RecastJSCrowd {
	r.p.Set("agents", agents)
	return r
}

// BjsRECASTPlugin returns the BjsRECASTPlugin property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#bjsrecastplugin
func (r *RecastJSCrowd) BjsRECASTPlugin() *RecastJSPlugin {
	retVal := r.p.Get("bjsRECASTPlugin")
	return RecastJSPluginFromJSObject(retVal, r.ctx)
}

// SetBjsRECASTPlugin sets the BjsRECASTPlugin property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#bjsrecastplugin
func (r *RecastJSCrowd) SetBjsRECASTPlugin(bjsRECASTPlugin *RecastJSPlugin) *RecastJSCrowd {
	r.p.Set("bjsRECASTPlugin", bjsRECASTPlugin.JSObject())
	return r
}

// RecastCrowd returns the RecastCrowd property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#recastcrowd
func (r *RecastJSCrowd) RecastCrowd() interface{} {
	retVal := r.p.Get("recastCrowd")
	return retVal
}

// SetRecastCrowd sets the RecastCrowd property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#recastcrowd
func (r *RecastJSCrowd) SetRecastCrowd(recastCrowd interface{}) *RecastJSCrowd {
	r.p.Set("recastCrowd", recastCrowd)
	return r
}

// Transforms returns the Transforms property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#transforms
func (r *RecastJSCrowd) Transforms() *TransformNode {
	retVal := r.p.Get("transforms")
	return TransformNodeFromJSObject(retVal, r.ctx)
}

// SetTransforms sets the Transforms property of class RecastJSCrowd.
//
// https://doc.babylonjs.com/api/classes/babylon.recastjscrowd#transforms
func (r *RecastJSCrowd) SetTransforms(transforms *TransformNode) *RecastJSCrowd {
	r.p.Set("transforms", transforms.JSObject())
	return r
}
