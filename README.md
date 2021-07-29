# openfiledialog
golang openfiledialog

一番シンプルな使い方
``` go
package main

import (
	"fmt"

	"github.com/Tobotobo/openfiledialog"
)

func main() {
	if ok, filePath := openfiledialog.Show(); ok {
		fmt.Println(filePath)
	}
}
```
タイトル、フィルター指定
``` go
openfiledialog
  .Title("Excel選択")
  .Filter("Excel(*.xlsx)|(*.xlsx)", 1)
  .Show()
```
複数選択可
``` go
openfiledialog.Mult().Show()
```
