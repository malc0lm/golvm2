package golvm2

// #cgo LDFLAGS: -llvm2app -L /usr/lib/
// #include <stdio.h>
// #include <lvm2app.h>
// #include "lvm_property_extension.h"
import "C"

import (
	"unsafe"
)

type LogicalVolume struct {
	c_lvh   C.lv_t
	c_lvm2h C.lvm_t
}

type LogicalVolumeSegment struct {
	c_lvseg C.lvseg_t
	c_lvm2h C.lvm_t
}

type LvDetail struct {
	name        string
	uuid        string
	devNo       string
	size        uint64
	isActive    bool
	isSuspended bool
}

//TODO implement
func (lv *LogicalVolume) LvListLvsegHandler() ([]*LogicalVolumeSegment, error) {
	lvsegList := make([]*LogicalVolumeSegment, 0, 10)
	var lvsegl *C.lvseg_list_t

	lvsegs := C.lvm_lv_list_lvsegs(lv.c_lvh)
	if lvsegs == nil {
		return nil, Lvm2Error{-1, "C Function lvm_lv_list_lvsegs error!"}
	}

	for current := lvsegs.n; current != lvsegs; current = current.n {
		lvsegl = (*C.lvseg_list_t)(unsafe.Pointer(current))

		lvsegh := &LogicalVolumeSegment{
			c_lvseg: lvsegl.lvseg,
			c_lvm2h: lv.c_lvm2h,
		}
		lvsegList = append(lvsegList, lvsegh)
	}
	return lvsegList, nil
}

func (lv *LogicalVolume) LvActivate() error {
	ret := C.lvm_lv_activate(lv.c_lvh)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_lv_activate error!"}
	}
	return nil
}

func (lv *LogicalVolume) LvDeactivate() error {
	ret := C.lvm_lv_deactivate(lv.c_lvh)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_lv_deactivate error!"}
	}
	return nil
}

func (lv *LogicalVolume) VgRemoveLv() error {
	ret := C.lvm_vg_remove_lv(lv.c_lvh)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_remove_lv error!"}
	}
	return nil
}

func (lv *LogicalVolume) LvGetUUID() string {
	uuid := C.lvm_lv_get_uuid(lv.c_lvh)

	return C.GoString(uuid)
}

func (lv *LogicalVolume) LvGetName() string {
	lvname := C.lvm_lv_get_name(lv.c_lvh)

	return C.GoString(lvname)
}

func (lv *LogicalVolume) LvGetSize() uint64 {
	lvsize := C.lvm_lv_get_size(lv.c_lvh)

	return uint64(lvsize)
}

func (lv *LogicalVolume) LvGetProperty(prop string) (*LvmPropertyValue, error) {
	c_propname := C.CString(prop)
	defer C.free(unsafe.Pointer(c_propname))

	propValue := new(LvmPropertyValue)
	c_prop := C.lvm_lv_get_property_c(lv.c_lvh, c_propname)
	//release memory which malloc in lvm_property_extension.c
	defer C.free(unsafe.Pointer(c_prop))

	if c_prop == nil {
		return nil, Lvm2Error{int(C.lvm_errno(lv.c_lvm2h)), C.GoString(C.lvm_errmsg(lv.c_lvm2h))}
	}
	if uint32(c_prop.is_settable) == 1 {
		propValue.isSettable = true
	} else {
		propValue.isSettable = false
	}
	if uint32(c_prop.is_string) == 1 {
		propValue.isString = true
		propValue.value = C.GoString(unionToCharptr(c_prop.value))
	}
	if uint32(c_prop.is_integer) == 1 {
		propValue.isInteger = true
		propValue.value = unionToUint64(c_prop.value)
	}
	return propValue, nil

}

func (lv *LogicalVolume) LvIsActive() bool {
	ret := C.lvm_lv_is_active(lv.c_lvh)
	if uint64(ret) == 0 {
		return false
	}
	return true
}

func (lv *LogicalVolume) LvIsSuspended() bool {
	ret := C.lvm_lv_is_suspended(lv.c_lvh)
	if uint64(ret) == 0 {
		return false
	}
	return true
}

func (lv *LogicalVolume) LvAddTag(tag string) error {
	c_tag := C.CString(tag)
	defer C.free(unsafe.Pointer(c_tag))
	ret := C.lvm_lv_add_tag(lv.c_lvh, c_tag)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_lv_add_tag error!"}
	}
	return nil
}

func (lv *LogicalVolume) LvRemoveTag(tag string) error {
	c_tag := C.CString(tag)
	defer C.free(unsafe.Pointer(c_tag))
	ret := C.lvm_lv_remove_tag(lv.c_lvh, c_tag)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_lv_remove_tag error!"}
	}
	return nil
}

func (lv *LogicalVolume) LvGetTags() ([]string, error) {
	tagList := make([]string, 0, 10)
	var tagl *C.lvm_str_list_t

	tags := C.lvm_lv_get_tags(lv.c_lvh)
	if tags == nil {
		return nil, Lvm2Error{-1, "C function lvm_lv_get_tags failed!"}
	}

	for current := tags.n; current != tags; current = current.n {
		tagl = (*C.lvm_str_list_t)(unsafe.Pointer(current))

		tagList = append(tagList, C.GoString(tagl.str))
	}
	return tagList, nil

}

func (lv *LogicalVolume) LvRename(newName string) error {
	c_newname := C.CString(newName)
	defer C.free(unsafe.Pointer(c_newname))
	ret := C.lvm_lv_rename(lv.c_lvh, c_newname)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_lv_rename error!"}
	}
	return nil
}

func (lv *LogicalVolume) LvResize(newSize uint64) error {
	ret := C.lvm_lv_resize(lv.c_lvh, C.uint64_t(newSize))
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_lv_resize error!"}
	}
	return nil
}

func (lvseg *LogicalVolumeSegment) LvsegGetProperty(prop string) (*LvmPropertyValue, error) {
	c_propname := C.CString(prop)
	defer C.free(unsafe.Pointer(c_propname))

	propValue := new(LvmPropertyValue)
	c_prop := C.lvm_lvseg_get_property_c(lvseg.c_lvseg, c_propname)
	//release memory which malloc in lvm_property_extension.c
	defer C.free(unsafe.Pointer(c_prop))

	if c_prop == nil {
		return nil, Lvm2Error{int(C.lvm_errno(lvseg.c_lvm2h)), C.GoString(C.lvm_errmsg(lvseg.c_lvm2h))}
	}
	if uint32(c_prop.is_settable) == 1 {
		propValue.isSettable = true
	} else {
		propValue.isSettable = false
	}
	if uint32(c_prop.is_string) == 1 {
		propValue.isString = true
		propValue.value = C.GoString(unionToCharptr(c_prop.value))
	}
	if uint32(c_prop.is_integer) == 1 {
		propValue.isInteger = true
		propValue.value = unionToUint64(c_prop.value)
	}
	return propValue, nil

}
