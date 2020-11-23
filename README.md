# dump
Color print package for golang


## Install

```
go get -v github.com/relunctance/dump
```


## Use Example:Dump()

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
```

## Use Example:ColorDump()

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

## Use Example:P()


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

