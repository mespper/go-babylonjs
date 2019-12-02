// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PerfCounter represents a babylon.js PerfCounter.
// This class is used to track a performance counter which is number based.
// The user has access to many properties which give statistics of different nature.
//
// The implementer can track two kinds of Performance Counter: time and count.
// For time you can optionally call fetchNewFrame() to notify the start of a new frame to monitor, then call beginMonitoring() to start and endMonitoring() to record the lapsed time. endMonitoring takes a newFrame parameter for you to specify if the monitored time should be set for a new frame or accumulated to the current frame being monitored.
// For count you first have to call fetchNewFrame() to notify the start of a new frame to monitor, then call addCount() how many time required to increment the count value you monitor.
type PerfCounter struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PerfCounter) JSObject() js.Value { return p.p }

// PerfCounter returns a PerfCounter JavaScript class.
func (ba *Babylon) PerfCounter() *PerfCounter {
	p := ba.ctx.Get("PerfCounter")
	return PerfCounterFromJSObject(p, ba.ctx)
}

// PerfCounterFromJSObject returns a wrapped PerfCounter JavaScript class.
func PerfCounterFromJSObject(p js.Value, ctx js.Value) *PerfCounter {
	return &PerfCounter{p: p, ctx: ctx}
}

// PerfCounterArrayToJSArray returns a JavaScript Array for the wrapped array.
func PerfCounterArrayToJSArray(array []*PerfCounter) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPerfCounter returns a new PerfCounter object.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter
func (ba *Babylon) NewPerfCounter() *PerfCounter {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("PerfCounter").New(args...)
	return PerfCounterFromJSObject(p, ba.ctx)
}

// AddCount calls the AddCount method on the PerfCounter object.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#addcount
func (p *PerfCounter) AddCount(newCount float64, fetchResult bool) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, newCount)
	args = append(args, fetchResult)

	p.p.Call("addCount", args...)
}

// BeginMonitoring calls the BeginMonitoring method on the PerfCounter object.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#beginmonitoring
func (p *PerfCounter) BeginMonitoring() {

	p.p.Call("beginMonitoring")
}

// PerfCounterEndMonitoringOpts contains optional parameters for PerfCounter.EndMonitoring.
type PerfCounterEndMonitoringOpts struct {
	NewFrame *bool
}

// EndMonitoring calls the EndMonitoring method on the PerfCounter object.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#endmonitoring
func (p *PerfCounter) EndMonitoring(opts *PerfCounterEndMonitoringOpts) {
	if opts == nil {
		opts = &PerfCounterEndMonitoringOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.NewFrame == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NewFrame)
	}

	p.p.Call("endMonitoring", args...)
}

// FetchNewFrame calls the FetchNewFrame method on the PerfCounter object.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#fetchnewframe
func (p *PerfCounter) FetchNewFrame() {

	p.p.Call("fetchNewFrame")
}

// Average returns the Average property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#average
func (p *PerfCounter) Average() float64 {
	retVal := p.p.Get("average")
	return retVal.Float()
}

// SetAverage sets the Average property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#average
func (p *PerfCounter) SetAverage(average float64) *PerfCounter {
	p.p.Set("average", average)
	return p
}

// Count returns the Count property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#count
func (p *PerfCounter) Count() float64 {
	retVal := p.p.Get("count")
	return retVal.Float()
}

// SetCount sets the Count property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#count
func (p *PerfCounter) SetCount(count float64) *PerfCounter {
	p.p.Set("count", count)
	return p
}

// Current returns the Current property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#current
func (p *PerfCounter) Current() float64 {
	retVal := p.p.Get("current")
	return retVal.Float()
}

// SetCurrent sets the Current property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#current
func (p *PerfCounter) SetCurrent(current float64) *PerfCounter {
	p.p.Set("current", current)
	return p
}

// Enabled returns the Enabled property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#enabled
func (p *PerfCounter) Enabled() bool {
	retVal := p.p.Get("Enabled")
	return retVal.Bool()
}

// SetEnabled sets the Enabled property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#enabled
func (p *PerfCounter) SetEnabled(Enabled bool) *PerfCounter {
	p.p.Set("Enabled", Enabled)
	return p
}

// LastSecAverage returns the LastSecAverage property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#lastsecaverage
func (p *PerfCounter) LastSecAverage() float64 {
	retVal := p.p.Get("lastSecAverage")
	return retVal.Float()
}

// SetLastSecAverage sets the LastSecAverage property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#lastsecaverage
func (p *PerfCounter) SetLastSecAverage(lastSecAverage float64) *PerfCounter {
	p.p.Set("lastSecAverage", lastSecAverage)
	return p
}

// Max returns the Max property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#max
func (p *PerfCounter) Max() float64 {
	retVal := p.p.Get("max")
	return retVal.Float()
}

// SetMax sets the Max property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#max
func (p *PerfCounter) SetMax(max float64) *PerfCounter {
	p.p.Set("max", max)
	return p
}

// Min returns the Min property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#min
func (p *PerfCounter) Min() float64 {
	retVal := p.p.Get("min")
	return retVal.Float()
}

// SetMin sets the Min property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#min
func (p *PerfCounter) SetMin(min float64) *PerfCounter {
	p.p.Set("min", min)
	return p
}

// Total returns the Total property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#total
func (p *PerfCounter) Total() float64 {
	retVal := p.p.Get("total")
	return retVal.Float()
}

// SetTotal sets the Total property of class PerfCounter.
//
// https://doc.babylonjs.com/api/classes/babylon.perfcounter#total
func (p *PerfCounter) SetTotal(total float64) *PerfCounter {
	p.p.Set("total", total)
	return p
}
