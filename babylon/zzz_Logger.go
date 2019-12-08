// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Logger represents a babylon.js Logger.
// Logger used througouht the application to allow configuration of
// the log level required for the messages.
type Logger struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *Logger) JSObject() js.Value { return l.p }

// Logger returns a Logger JavaScript class.
func (ba *Babylon) Logger() *Logger {
	p := ba.ctx.Get("Logger")
	return LoggerFromJSObject(p, ba.ctx)
}

// LoggerFromJSObject returns a wrapped Logger JavaScript class.
func LoggerFromJSObject(p js.Value, ctx js.Value) *Logger {
	return &Logger{p: p, ctx: ctx}
}

// LoggerArrayToJSArray returns a JavaScript Array for the wrapped array.
func LoggerArrayToJSArray(array []*Logger) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ClearLogCache calls the ClearLogCache method on the Logger object.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#clearlogcache
func (l *Logger) ClearLogCache() {

	l.p.Call("ClearLogCache")
}

// AllLogLevel returns the AllLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#allloglevel
func (l *Logger) AllLogLevel() float64 {
	retVal := l.p.Get("AllLogLevel")
	return retVal.Float()
}

// SetAllLogLevel sets the AllLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#allloglevel
func (l *Logger) SetAllLogLevel(AllLogLevel float64) *Logger {
	l.p.Set("AllLogLevel", AllLogLevel)
	return l
}

// Error returns the Error property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#error
func (l *Logger) Error() js.Value {
	retVal := l.p.Get("Error")
	return retVal
}

// SetError sets the Error property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#error
func (l *Logger) SetError(Error JSFunc) *Logger {
	l.p.Set("Error", js.FuncOf(Error))
	return l
}

// ErrorLogLevel returns the ErrorLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorloglevel
func (l *Logger) ErrorLogLevel() float64 {
	retVal := l.p.Get("ErrorLogLevel")
	return retVal.Float()
}

// SetErrorLogLevel sets the ErrorLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorloglevel
func (l *Logger) SetErrorLogLevel(ErrorLogLevel float64) *Logger {
	l.p.Set("ErrorLogLevel", ErrorLogLevel)
	return l
}

// ErrorsCount returns the ErrorsCount property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorscount
func (l *Logger) ErrorsCount() float64 {
	retVal := l.p.Get("errorsCount")
	return retVal.Float()
}

// SetErrorsCount sets the ErrorsCount property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorscount
func (l *Logger) SetErrorsCount(errorsCount float64) *Logger {
	l.p.Set("errorsCount", errorsCount)
	return l
}

// Log returns the Log property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#log
func (l *Logger) Log() js.Value {
	retVal := l.p.Get("Log")
	return retVal
}

// SetLog sets the Log property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#log
func (l *Logger) SetLog(Log JSFunc) *Logger {
	l.p.Set("Log", js.FuncOf(Log))
	return l
}

// LogCache returns the LogCache property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#logcache
func (l *Logger) LogCache() string {
	retVal := l.p.Get("LogCache")
	return retVal.String()
}

// SetLogCache sets the LogCache property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#logcache
func (l *Logger) SetLogCache(LogCache string) *Logger {
	l.p.Set("LogCache", LogCache)
	return l
}

// LogLevels returns the LogLevels property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#loglevels
func (l *Logger) LogLevels() float64 {
	retVal := l.p.Get("LogLevels")
	return retVal.Float()
}

// SetLogLevels sets the LogLevels property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#loglevels
func (l *Logger) SetLogLevels(LogLevels float64) *Logger {
	l.p.Set("LogLevels", LogLevels)
	return l
}

// MessageLogLevel returns the MessageLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#messageloglevel
func (l *Logger) MessageLogLevel() float64 {
	retVal := l.p.Get("MessageLogLevel")
	return retVal.Float()
}

// SetMessageLogLevel sets the MessageLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#messageloglevel
func (l *Logger) SetMessageLogLevel(MessageLogLevel float64) *Logger {
	l.p.Set("MessageLogLevel", MessageLogLevel)
	return l
}

// NoneLogLevel returns the NoneLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#noneloglevel
func (l *Logger) NoneLogLevel() float64 {
	retVal := l.p.Get("NoneLogLevel")
	return retVal.Float()
}

// SetNoneLogLevel sets the NoneLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#noneloglevel
func (l *Logger) SetNoneLogLevel(NoneLogLevel float64) *Logger {
	l.p.Set("NoneLogLevel", NoneLogLevel)
	return l
}

// OnNewCacheEntry returns the OnNewCacheEntry property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#onnewcacheentry
func (l *Logger) OnNewCacheEntry() js.Value {
	retVal := l.p.Get("OnNewCacheEntry")
	return retVal
}

// SetOnNewCacheEntry sets the OnNewCacheEntry property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#onnewcacheentry
func (l *Logger) SetOnNewCacheEntry(OnNewCacheEntry JSFunc) *Logger {
	l.p.Set("OnNewCacheEntry", js.FuncOf(OnNewCacheEntry))
	return l
}

// Warn returns the Warn property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warn
func (l *Logger) Warn() js.Value {
	retVal := l.p.Get("Warn")
	return retVal
}

// SetWarn sets the Warn property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warn
func (l *Logger) SetWarn(Warn JSFunc) *Logger {
	l.p.Set("Warn", js.FuncOf(Warn))
	return l
}

// WarningLogLevel returns the WarningLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warningloglevel
func (l *Logger) WarningLogLevel() float64 {
	retVal := l.p.Get("WarningLogLevel")
	return retVal.Float()
}

// SetWarningLogLevel sets the WarningLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warningloglevel
func (l *Logger) SetWarningLogLevel(WarningLogLevel float64) *Logger {
	l.p.Set("WarningLogLevel", WarningLogLevel)
	return l
}
