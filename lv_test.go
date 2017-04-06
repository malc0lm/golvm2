package golvm2

import (
	"fmt"
	"testing"
)

func Test_LvListLvsegHandler(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "w", 0)
	lvh, _ := vgh.LvFromName("lv1")
	l, err := lvh.LvListLvsegHandler()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(l) == 0 {
		fmt.Println("0 lvseg")
		return
	}
	for i := 0; i < len(l); i++ {
		prop, _ := l[i].LvsegGetProperty("segtype")
		if prop.isInteger {
			fmt.Println(prop.value.(uint64))
		}
		if prop.isString {
			fmt.Println(prop.value.(string))
		}
	}
	vgh.VgClose()
	lvmh.Quit()
}

func Test_LvGetProperty(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "w", 0)
	lvh, _ := vgh.LvFromName("lv1")
	prop, err := lvh.LvGetProperty("lv_size")
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

	vgh.VgClose()
	lvmh.Quit()
}
func Test_LvGetTags(t *testing.T) {
	lvmh, _ := NewLvm2Handler()
	vgh, _ := lvmh.VgOpen("vgtest", "w", 0)
	lvh, _ := vgh.LvFromName("lv1")

	lvh.LvAddTag("tag1")
	lvh.LvAddTag("tag2")
	lvh.LvAddTag("tag3")

	vTagList, err := lvh.LvGetTags()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(vTagList); i++ {
		fmt.Println(vTagList[i])
	}
	vgh.VgClose()
	lvmh.Quit()
}
