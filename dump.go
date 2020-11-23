package dump

import (
	"fmt"

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
	for _, v := range vals {
		c.Println(JsonDump(v))
	}
}

// Dump
// The same as PHP:var_dump()
// https://www.php.net/var_dump
func Dump(vals ...interface{}) {
	for _, v := range vals {
		fmt.Println(JsonDump(v))
	}
}

// JsonDump
func JsonDump(v interface{}) string {
	return gabs.Wrap(v).StringIndent("", "   ")
}
