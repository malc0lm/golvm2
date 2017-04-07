package golvm2

import (
	"fmt"
	"testing"
)

func Test_ListVgNames(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	nameList, err := lvmh.ListVgNames()
	checkError(err)
	fmt.Println(nameList)
}
func Test_ListVgUUIDs(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	nameList, err := lvmh.ListVgUUIDs()
	checkError(err)
	fmt.Println(nameList)
}
