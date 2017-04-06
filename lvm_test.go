package golvm2

import (
	"fmt"
	"testing"
)

func Test_ListVgNames(t *testing.T) {
	lvmh, _ := NewLvm2Handler()

	nameList, err := lvmh.ListVgNames()

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(nameList)
	lvmh.Quit()
}
func Test_ListVgUUIDs(t *testing.T) {
	lvmh, _ := NewLvm2Handler()

	nameList, err := lvmh.ListVgUUIDs()

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(nameList)
	lvmh.Quit()
}
