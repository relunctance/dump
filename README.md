![Dump](dump_logo.png "Dump")

# Dump

[![GoDev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/relunctance/dump?tab=doc)
[![codecov](https://codecov.io/gh/go-kit/kit/branch/master/graph/badge.svg)](https://codecov.io/gh/relunctance/dump)
[![Go Report Card](https://goreportcard.com/badge/relunctance/dump)](https://goreportcard.com/report/relunctance/dump)
[![Sourcegraph](https://sourcegraph.com/github.com/relunctance/dump/-/badge.svg)](https://sourcegraph.com/github.com/relunctance/dump?badge)

Dumps information about a variable  . The Same as [PHP:var_dump](https://www.php.net/var_dump)


Getting Started
===============

Installing
----------
```
go get -v github.com/relunctance/dump
```

Usage
===============

Example:Dump()
--------------

```go

import (
    "github.com/relunctance/dump"
)

func main(){
	dump.Dump(
		[]string{"a", "b", "c"},
		"abc",
		1,
		1.68,
		[]int{1, 2, 3, 4, 5},
		[]int32{1, 2, 3, 4, 5},
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
		map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		},
	)
}

/* returns:
[
   "a",
   "b",
   "c"
]
"abc"
1
1.68
[
   1,
   2,
   3,
   4,
   5
]
[
   1,
   2,
   3,
   4,
   5
]
{
   "a": "1",
   "b": "2",
   "c": "3"
}
{
   "a": 1,
   "b": 2,
   "c": 3
}
*/
```

Example:ColorDump()
-------------------

```go
import (
    "github.com/relunctance/dump"
    "github.com/fatih/color"
)

func main(){
	dump.ColorDump(
		color.New(color.FgRed).Add(color.BgHiWhite),    // with color
		[]string{"a", "b", "c"},
		"abc",
		1,
		1.68,
		[]int{1, 2, 3, 4, 5},
		[]int32{1, 2, 3, 4, 5},
		map[string]string{ "a": "1", "b": "2", "c": "3", },
		map[string]int{ "a": 1, "b": 2, "c": 3, },
	)
}
```

Example:P()
-----------


```go

import (
    "github.com/relunctance/dump"
)

func main(){
	dump.P(
		[]string{"a", "b", "c"},
		"abc",
		1,
		1.68,
		[]int{1, 2, 3, 4, 5},
		[]int32{1, 2, 3, 4, 5},
		map[string]string{ "a": "1", "b": "2", "c": "3", },
		map[string]int{ "a": 1, "b": 2, "c": 3, },
	)
}
```

## Thanks

* [github.com/fatih/color](https://github.com/fatih/color)
* [github.com/Jeffail/gabs](https://github.com/Jeffail/gabs) 

