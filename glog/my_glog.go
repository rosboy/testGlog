package glog

import (
	"flag"

	"github.com/golang/glog"
)

type Verbose int
type myLogging struct {
	infoVerbosity  int
	errorVerbosity int
	warnVerbosity  int
	fatalVerbosity int
}

var logging myLogging

func init() {
	// 导入info,error,warning,fatal的flag配置值
	flag.IntVar(&logging.infoVerbosity, "v_info", 1, "info log level for V logs")
	flag.IntVar(&logging.errorVerbosity, "v_error", 1, "error log level for V logs")
	flag.IntVar(&logging.warnVerbosity, "v_warn", 1, "warn log level for V logs")
	flag.IntVar(&logging.fatalVerbosity, "v_fatal", 1, "fatal log level for V logs")
}

func V(l int) *Verbose {
	v := Verbose(l)
	return &v
}

func (v *Verbose) Info(args ...interface{}) {
	if int(*v) < logging.infoVerbosity {
		glog.Info(args...)
	}
}
func (v *Verbose) Infof(format string, args ...interface{}) {
	if int(*v) < logging.infoVerbosity {
		glog.Infof(format, args...)
	}
}

func (v *Verbose) Error(args ...interface{}) {
	if int(*v) < logging.errorVerbosity {
		glog.Error(args...)
	}
}
func (v *Verbose) Errorf(format string, args ...interface{}) {
	if int(*v) < logging.errorVerbosity {
		glog.Errorf(format, args...)
	}
}

func (v *Verbose) Warning(args ...interface{}) {
	if int(*v) < logging.warnVerbosity {
		glog.Warning(args...)
	}
}
func (v *Verbose) Warningf(format string, args ...interface{}) {
	if int(*v) < logging.warnVerbosity {
		glog.Warningf(format, args...)
	}
}

func (v *Verbose) Fatal(args ...interface{}) {
	if int(*v) < logging.fatalVerbosity {
		glog.Fatal(args...)
	}
}
func (v *Verbose) Fatalf(format string, args ...interface{}) {
	if int(*v) < logging.fatalVerbosity {
		glog.Fatalf(format, args...)
	}
}

func Flush() {
	glog.Flush()
}
