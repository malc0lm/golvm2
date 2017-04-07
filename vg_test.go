package golvm2

import (
	"fmt"
	"testing"
)

func Test_VgListLvs(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("vgtest", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	vlist, err := vgh.VgListLvs()
	checkError(err)
	for i := 0; i < len(vlist); i++ {
		fmt.Println(*vlist[i])
	}
}

func Test_VgListPvs(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("vgtest", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	vlist, err := vgh.VgListPvs()
	checkError(err)
	for i := 0; i < len(vlist); i++ {
		fmt.Println(*vlist[i])
	}

}

func Test_VgGetTags(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "w", 0)
	defer vgh.VgClose()
	checkError(err)

	vgh.VgAddTag("tag1")
	vgh.VgAddTag("bar2")
	vgh.VgAddTag("foo3")

	vTagList, err := vgh.VgGetTags()
	checkError(err)
	for i := 0; i < len(vTagList); i++ {
		fmt.Println(vTagList[i])
	}
}

func Test_VgGetProperty(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("vgtest", "w", 0)
	defer vgh.VgClose()
	checkError(err)
	prop, err := vgh.VgGetProperty("vg_extent_size")
	checkError(err)
	if prop.isInteger {
		fmt.Println(prop.value.(uint64))
	}
	if prop.isString {
		fmt.Println(prop.value.(string))
	}
	fmt.Println(prop.isSettable)
}

// func Test_VgSetProperty(t *testing.T) {
// 	lvmh, err := NewLvm2Handler()
// 	defer lvmh.Quit()
// 	checkError(err)
// 	vgh, err := lvmh.VgOpen("vgtest", "w", 0)
// 	defer vgh.VgClose()
// 	checkError(err)
// 	err := vgh.VgSetProperty("vg_attr", "cvvv")
// 	checkError(err)
// }
