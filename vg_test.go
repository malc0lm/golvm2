package golvm2

import (
	"fmt"
	"testing"
)

//var lvmh *Lvm2Handler
//var vgh *VolumeGroup

//func getVgHandler() (*VolumeGroup, error) {
//	lvmh, _ := NewLvm2Handler()
//	var err error
//	vgh, err = lvmh.VgOpen("vgtest", "r", 0)
//	if err != nil {
//		return nil, err
//	}
//	return vgh, nil
//}
//func cleanHandler() {
//	vgh.VgClose()
//	lvmh.Quit()
//}
func Test_VgListLvs(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "r", 0)

	vlist, err := vgh.VgListLvs()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(vlist); i++ {
		fmt.Println(*vlist[i])
	}
	vgh.VgClose()
	lvmh.Quit()

	//	var herr error
	//	vgh, herr = getVgHandler()
	//	if herr != nil {
	//		t.Error(herr)
	//	}
	//	cleanHandler()
}

func Test_VgListPvs(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "r", 0)

	vlist, err := vgh.VgListPvs()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(vlist); i++ {
		fmt.Println(*vlist[i])
	}
	vgh.VgClose()
	lvmh.Quit()
}

func Test_VgGetTags(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "w", 0)

	vgh.VgAddTag("tag1")
	vgh.VgAddTag("bar2")
	vgh.VgAddTag("foo3")

	vTagList, err := vgh.VgGetTags()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(vTagList); i++ {
		fmt.Println(vTagList[i])
	}
	vgh.VgClose()
	lvmh.Quit()
}

func Test_VgGetProperty(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "w", 0)

	prop, err := vgh.VgGetProperty("vg_extent_size")

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if prop.isInteger {
		fmt.Println(prop.value.(uint64))
	}
	if prop.isString {
		fmt.Println(prop.value.(string))
	}
	fmt.Println(prop.isSettable)
	vgh.VgClose()
	lvmh.Quit()
}
func Test_VgSetProperty(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "w", 0)

	err := vgh.VgSetProperty("vg_attr", "cvvv")

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	vgh.VgClose()
	lvmh.Quit()
}
