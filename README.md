# golvm2 - An lvm2 Library for Go 

This library is only tested in Linux, such as(ubuntu14.04, centos7). Golvm2 depends liblvm api which is written by C.
You must install lvm2-devel(Centos) or liblvm2-dev(Ubuntu) before use this library which use cgo.

## Example Usage

```go
package main

import (
  "github.com/malc0lm/golvm2"
  "fmt"
  "log"
)

func main() {
    lvmh, err := NewLvm2Handler()
    if err != nil{
        log.Fatal(err)
    }
    vgh, err := lvmh.VgOpen("vgtest", "r", 0)
    if err != nil{
        log.Fatal(err)
    }
    vlist, err := vgh.VgListLvs()
    if err != nil {
    	log.Fatal(err)
    }
    for i := 0; i < len(vlist); i++ {
    	fmt.Println(*vlist[i])
    }
    vgh.VgClose()
    lvmh.Quit()
}
```





