package golvm2

import (
	"fmt"
	"testing"
)

func Test_LvListLvsegHandler(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("vgtest", "w", 0)
	defer vgh.VgClose()
	checkError(err)
	lvh, err := vgh.LvFromName("lv1")
	checkError(err)
	l, err := lvh.LvListLvsegHandler()
	checkError(err)
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
}

func Test_LvGetProperty(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("vgtest", "w", 0)
	defer vgh.VgClose()
	checkError(err)
	lvh, err := vgh.LvFromName("lv1")
	checkError(err)
	prop, err := lvh.LvGetProperty("lv_size")
	checkError(err)
	if prop.isInteger {
		fmt.Println(prop.value.(uint64))
	}
	if prop.isString {
		fmt.Println(prop.value.(string))
	}

}
func Test_LvGetTags(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("vgtest", "w", 0)
	defer vgh.VgClose()
	checkError(err)
	lvh, err := vgh.LvFromName("lv1")
	checkError(err)

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
}
