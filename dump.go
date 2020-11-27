package dump

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/fatih/color"
)

// P
// alias VarDump()
func P(vals ...interface{}) {
	VarDump(vals...)

}
func VarDump(vals ...interface{}) {
	ColorDump(color.New(color.FgRed).Add(color.BgHiWhite), vals...)
}

// ColorDump
// c := color.New(color.FgCyan).Add(color.Underline)
// c := color.New(color.FgCyan, color.Bold)
// https://github.com/fatih/color
func ColorDump(c *color.Color, vals ...interface{}) {
	if c == nil {
		Dump(vals...)
		return
	}
	commonPrintln(c.Println, vals...)
}

// Println
// alias Dump()
func Println(vals ...interface{}) {
	Dump(vals...)
}

// Dump
// The same as PHP:var_dump()
// https://www.php.net/var_dump
func Dump(vals ...interface{}) {
	commonPrintln(fmt.Println, vals...)
}

// JsonDump
func JsonDump(v interface{}) string {
	return gabs.Wrap(v).StringIndent("", "   ")
}

type F func(...interface{}) (int, error)

func commonPrintln(f F, vals ...interface{}) {
	tnow := timeNow()
	file := debugTrace()
	for _, v := range vals {
		f(tnow, file, JsonDump(v))
	}
}

func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Traces() []string {
	i := 0
	ret := make([]string, 0, 10)
	funcName := ""
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fc := runtime.FuncForPC(pc)
		if fc != nil {
			funcName = fc.Name()
			if pos := strings.LastIndex(funcName, "/"); pos > 0 {
				funcName = funcName[pos+1:]
			}
		}
		ret = append(ret, fmt.Sprintf("%s:%d:%s()\n", file, line, funcName))
		i++
	}
	return ret
}

func debugTrace() string {
	traces := Traces()
	for _, line := range traces {
		if pos := strings.Index(line, "dump.go:"); pos != -1 {
			continue
		}
		return line
	}
	return ""
}
