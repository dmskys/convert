# convert
golang convert type utils | golang类型转换工具

# init package
```go
go get github.com/dmskys/convert
```

## How To Use
```go

import (
    "github.com/dmskys/convert"
)
func main(){
	convert.ToStr(1)
	convert.BinToDec([]byte{1, 2})
	convert.DecToBin(1)
	convert.StrToInt("1")
	convert.StrToInt8("2")
	convert.Float32ToStr(1.234, 2)
	convert.Float64ToStr(1.3445, 3)
	convert.AddSingleQuotes(1)
}

```