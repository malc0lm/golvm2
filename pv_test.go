package golvm2

import (
	"fmt"
	"log"
	"testing"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Test_PvGetUUID(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	fmt.Println(pvh.PvGetUUID())
}

func Test_PvGetName(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromUUID("pIjoJX-nPy9-oZ1Y-za70-kz3g-mwu3-VtD9aj")
	checkError(err)
	fmt.Println(pvh.PvGetName())
}

func Test_PvGetMdaCount(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	fmt.Println(pvh.PvGetMdaCount())
}

func Test_PvGetDevSize(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	fmt.Println(pvh.PvGetDevSize())
}

func Test_PvGetSize(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	fmt.Println(pvh.PvGetSize())
}

func Test_PvGetFree(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	fmt.Println(pvh.PvGetFree())
}

func Test_PvResize(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	fmt.Println(pvh.PvGetFree())
}

func Test_PvGetProperty(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	propString, err := pvh.PvGetProperty("pv_uuid")
	checkError(err)
	if propString.isInteger {
		fmt.Println(propString.value.(uint64))
	}
	if propString.isString {
		fmt.Println(propString.value.(string))
	}
	propUint64, err := pvh.PvGetProperty("dev_size")
	checkError(err)
	if propUint64.isInteger {
		fmt.Println(propUint64.value.(uint64))
	}
	if propUint64.isString {
		fmt.Println(propUint64.value.(string))
	}

}

func Test_PvsegGetProperty(t *testing.T) {
	lvmh, err := NewLvm2Handler()
	defer lvmh.Quit()
	checkError(err)
	vgh, err := lvmh.VgOpen("centos", "r", 0)
	defer vgh.VgClose()
	checkError(err)
	pvh, err := vgh.PvFromName("/dev/sda2")
	checkError(err)
	pvl, err := pvh.PvListPvsegHandler()
	checkError(err)
	if len(pvl) == 0 {
		fmt.Println("0 lvseg")
		return
	}
	for i := 0; i < len(pvl); i++ {
		prop, _ := pvl[i].PvsegGetProperty("pvseg_start")
		if prop.isInteger {
			fmt.Println(prop.value.(uint64))
		}
		if prop.isString {
			fmt.Println(prop.value.(string))
		}
	}
}
