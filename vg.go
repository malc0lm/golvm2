package golvm2

// #cgo LDFLAGS: -llvm2app -L /usr/lib/
// #include <lvm2app.h>
// #include "lvm_property_extension.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type VolumeGroup struct {
	c_vgh   C.vg_t
	c_lvm2h C.lvm_t
}

/*
In lvm2app.h list is first field in struct lvm_lv_list,
so will don't use complex macro for get lvm_lv_list address, we can
use (*C.struct_lvm_lv_list)(unsafe.Pointer(list) in go.
typedef struct lvm_lv_list {
	struct dm_list list;
	lv_t lv;
} lv_list_t;
*/
func (vg *VolumeGroup) VgListLvs() ([]*LvDetail, error) {

	lvDetailList := make([]*LvDetail, 0, 10)
	var lvl *C.lv_list_t

	vgname := C.lvm_vg_get_name(vg.c_vgh)
	lvs := C.lvm_vg_list_lvs(vg.c_vgh)
	if lvs == nil {
		return nil, nil
	}

	for current := lvs.n; current != lvs; current = current.n {
		lvl = (*C.lv_list_t)(unsafe.Pointer(current))

		c_lvName := C.lvm_lv_get_name(lvl.lv)
		c_uuid := C.lvm_lv_get_uuid(lvl.lv)
		c_size := C.lvm_lv_get_size(lvl.lv)
		c_isActive := C.lvm_lv_is_active(lvl.lv)
		c_isSuspended := C.lvm_lv_is_suspended(lvl.lv)
		lvName := C.GoString(c_lvName)
		devNo, err := getDevNo(C.GoString(vgname), lvName)
		if err != nil {
			return nil, err
		}

		lvd := &LvDetail{
			name:        lvName,
			uuid:        C.GoString(c_uuid),
			devNo:       devNo,
			size:        uint64(c_size),
			isActive:    bool(uint64(c_isActive)&1 == 1),
			isSuspended: bool(uint64(c_isSuspended)&1 == 1),
		}
		lvDetailList = append(lvDetailList, lvd)
	}

	return lvDetailList, nil
}

/*
typedef struct lvm_pv_list {
        struct dm_list list;
        pv_t pv;
} pv_list_t;
*/
func (vg *VolumeGroup) VgListPvs() ([]*PvDetail, error) {
	pvDetailList := make([]*PvDetail, 0, 10)
	var pvl *C.pv_list_t

	pvs := C.lvm_vg_list_pvs(vg.c_vgh)
	if pvs == nil {
		return nil, nil
	}

	for current := pvs.n; current != pvs; current = current.n {
		pvl = (*C.pv_list_t)(unsafe.Pointer(current))

		c_pvName := C.lvm_pv_get_name(pvl.pv)
		c_uuid := C.lvm_pv_get_uuid(pvl.pv)
		c_size := C.lvm_pv_get_size(pvl.pv)
		c_freesize := C.lvm_pv_get_free(pvl.pv)
		c_devsize := C.lvm_pv_get_dev_size(pvl.pv)
		c_metadata_areas := C.lvm_pv_get_mda_count(pvl.pv)

		pvd := &PvDetail{
			name:               C.GoString(c_pvName),
			uuid:               C.GoString(c_uuid),
			size:               uint64(c_size),
			freeSize:           uint64(c_freesize),
			devSize:            uint64(c_devsize),
			metadataAreaNumber: uint64(c_metadata_areas),
		}
		pvDetailList = append(pvDetailList, pvd)
	}
	return pvDetailList, nil
}

func (vg *VolumeGroup) VgWrite() error {
	ret := C.lvm_vg_write(vg.c_vgh)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_write error!"}
	}
	return nil
}

func (vg *VolumeGroup) VgRemove() error {
	ret := C.lvm_vg_remove(vg.c_vgh)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_remove error!"}
	}
	return nil
}

func (vg *VolumeGroup) VgClose() error {
	ret := C.lvm_vg_close(vg.c_vgh)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_close error!"}
	}
	return nil
}

/*
 *	VgExtend()
 * \param   device
 * Absolute pathname of device to add to VG.
 *
 */
func (vg *VolumeGroup) VgExtend(device string) error {
	c_device := C.CString(device)
	defer C.free(unsafe.Pointer(c_device))

	ret := C.lvm_vg_extend(vg.c_vgh, c_device)

	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_extend error!"}
	}
	return nil

}

/*
 *	VgReduce()
 * \param   device
 * Absolute pathname of device to remove from VG.
 */
func (vg *VolumeGroup) VgReduce(device string) error {
	c_device := C.CString(device)
	defer C.free(unsafe.Pointer(c_device))

	ret := C.lvm_vg_reduce(vg.c_vgh, c_device)

	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_reduce error!"}
	}
	return nil
}

func (vg *VolumeGroup) VgAddTag(tag string) error {
	c_tag := C.CString(tag)
	defer C.free(unsafe.Pointer(c_tag))

	ret := C.lvm_vg_add_tag(vg.c_vgh, c_tag)

	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_add_tag error!"}
	}
	return nil
}

func (vg *VolumeGroup) VgRemoveTag(tag string) error {
	c_tag := C.CString(tag)
	defer C.free(unsafe.Pointer(c_tag))

	ret := C.lvm_vg_remove_tag(vg.c_vgh, c_tag)

	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_remove_tag error!"}
	}
	return nil
}

func (vg *VolumeGroup) VgSetExtentSize(newSize uint32) error {
	ret := C.lvm_vg_set_extent_size(vg.c_vgh, C.uint32_t(newSize))
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_vg_set_extent_size error!"}
	}
	return nil
}

func (vg *VolumeGroup) VgIsClustered() bool {
	ret := C.lvm_vg_is_clustered(vg.c_vgh)
	if uint64(ret) == 0 {
		return false
	}
	return true
}

func (vg *VolumeGroup) VgIsExported() bool {
	ret := C.lvm_vg_is_exported(vg.c_vgh)
	if uint64(ret) == 0 {
		return false
	}
	return true
}

func (vg *VolumeGroup) VgIsPartial() bool {
	ret := C.lvm_vg_is_partial(vg.c_vgh)
	if uint64(ret) == 0 {
		return false
	}
	return true
}

//TODO test err
func (vg *VolumeGroup) VgGetSeqno() uint64 {
	seqno := C.lvm_vg_get_seqno(vg.c_vgh)

	return uint64(seqno)
}

//TODO test err
func (vg *VolumeGroup) VgGetUUID() string {
	uuid := C.lvm_vg_get_uuid(vg.c_vgh)

	return C.GoString(uuid)
}

//TODO test err
func (vg *VolumeGroup) VgGetName() string {
	vgname := C.lvm_vg_get_name(vg.c_vgh)

	return C.GoString(vgname)
}

//TODO test err
func (vg *VolumeGroup) VgGetSize() uint64 {
	vgsize := C.lvm_vg_get_size(vg.c_vgh)

	return uint64(vgsize)
}

//TODO test err
func (vg *VolumeGroup) VgGetFreeSize() uint64 {
	vgFreeSize := C.lvm_vg_get_free_size(vg.c_vgh)

	return uint64(vgFreeSize)
}

//TODO test err
func (vg *VolumeGroup) VgGetExtentSize() uint64 {
	vgExtentSize := C.lvm_vg_get_free_size(vg.c_vgh)

	return uint64(vgExtentSize)
}

//TODO test err
func (vg *VolumeGroup) VgGetExtentCount() uint64 {
	vgExtentCount := C.lvm_vg_get_extent_count(vg.c_vgh)

	return uint64(vgExtentCount)
}

//TODO test err
func (vg *VolumeGroup) VgGetExtentFreeCount() uint64 {
	vgExtentFreeCount := C.lvm_vg_get_free_extent_count(vg.c_vgh)

	return uint64(vgExtentFreeCount)
}

//TODO test err
func (vg *VolumeGroup) VgGetPvCount() uint64 {
	vgPvCount := C.lvm_vg_get_pv_count(vg.c_vgh)

	return uint64(vgPvCount)
}

//TODO test err
func (vg *VolumeGroup) VgGetMaxPv() uint64 {
	vgMaxPv := C.lvm_vg_get_max_pv(vg.c_vgh)

	return uint64(vgMaxPv)
}

//TODO test err
func (vg *VolumeGroup) VgGetMaxLv() uint64 {
	vgMaxLv := C.lvm_vg_get_max_lv(vg.c_vgh)

	return uint64(vgMaxLv)
}

/*Return the list of volume group tags.
typedef struct lvm_str_list {
        struct dm_list list;
        const char *str;
} lvm_str_list_t;
*/
func (vg *VolumeGroup) VgGetTags() ([]string, error) {
	tagList := make([]string, 0, 10)
	var tagl *C.lvm_str_list_t

	tags := C.lvm_vg_get_tags(vg.c_vgh)
	if tags == nil {
		return nil, Lvm2Error{-1, "C function lvm_vg_get_tags failed!"}
	}

	for current := tags.n; current != tags; current = current.n {
		tagl = (*C.lvm_str_list_t)(unsafe.Pointer(current))

		tagList = append(tagList, C.GoString(tagl.str))
	}
	return tagList, nil
}

//TODO implement this
func (vg *VolumeGroup) VgGetProperty(prop string) (*LvmPropertyValue, error) {
	c_propname := C.CString(prop)
	defer C.free(unsafe.Pointer(c_propname))

	propValue := new(LvmPropertyValue)
	c_prop := C.lvm_vg_get_property_c(vg.c_vgh, c_propname)
	defer C.free(unsafe.Pointer(c_prop))
	if c_prop == nil {
		return nil, Lvm2Error{int(C.lvm_errno(vg.c_lvm2h)), C.GoString(C.lvm_errmsg(vg.c_lvm2h))}
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

//All of vg's property is unsettable so far, so VgSetProperty can't be tested well.
func (vg *VolumeGroup) VgSetProperty(prop string, setValue interface{}) error {
	lpv, err := vg.VgGetProperty(prop)
	c_propname := C.CString(prop)
	defer C.free(unsafe.Pointer(c_propname))

	var c_prop C.lvm_property_value_t
	if err != nil {
		return err
	}
	if !lpv.isSettable {
		return Lvm2Error{-1, fmt.Sprintf("Setting %v property is forbidden!", prop)}
	}
	if lpv.isInteger {
		setValueUint64, ok := setValue.(uint64)
		if !ok {
			return Lvm2Error{-1, fmt.Sprintf("The %v property's value must be uint64!", prop)}
		}
		c_prop.value = uint64ToUnion(setValueUint64)
	}
	if lpv.isString {
		setValueString, ok := setValue.(string)
		if !ok {
			return Lvm2Error{-1, fmt.Sprintf("The %v property's value must be string!", prop)}
		}
		c_string := C.CString(setValueString)
		defer C.free(unsafe.Pointer(c_string))
		c_prop.value = charptrToUnion(c_string)
	}
	ret := C.lvm_vg_set_property(vg.c_vgh, c_propname, &c_prop)
	if ret != 0 {
		return Lvm2Error{int(ret), fmt.Sprintf("C function lvm_vg_set_property failed!", prop)}
	}

	return nil
}

func (vg *VolumeGroup) VgCreateLvLinear(lvname string, lvsize uint64) (*LogicalVolume, error) {
	c_lvname := C.CString(lvname)
	defer C.free(unsafe.Pointer(c_lvname))
	c_lv := C.lvm_vg_create_lv_linear(vg.c_vgh, c_lvname, C.uint64_t(lvsize))
	if c_lv == nil {
		return nil, Lvm2Error{-1, "C Function lvm_vg_create_lv_linear error!"}
	}
	lv := new(LogicalVolume)
	lv.c_lvh = c_lv
	lv.c_lvm2h = vg.c_lvm2h
	return lv, nil
}

func (vg *VolumeGroup) LvFromName(lvname string) (*LogicalVolume, error) {
	c_lvname := C.CString(lvname)
	defer C.free(unsafe.Pointer(c_lvname))

	c_lv := C.lvm_lv_from_name(vg.c_vgh, c_lvname)
	if c_lv == nil {
		return nil, Lvm2Error{-1, "C function lvm_lv_from_name failed!"}
	}

	lv := new(LogicalVolume)
	lv.c_lvh = c_lv
	lv.c_lvm2h = vg.c_lvm2h
	return lv, nil
}

func (vg *VolumeGroup) LvFromUUID(uuid string) (*LogicalVolume, error) {
	c_uuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(c_uuid))

	c_lv := C.lvm_lv_from_uuid(vg.c_vgh, c_uuid)
	if c_lv == nil {
		return nil, Lvm2Error{-1, "C function lvm_lv_from_uuid failed!"}
	}

	lv := new(LogicalVolume)
	lv.c_lvh = c_lv
	lv.c_lvm2h = vg.c_lvm2h
	return lv, nil
}

func (vg *VolumeGroup) PvFromName(pvname string) (*PhysicalVolume, error) {
	c_pvname := C.CString(pvname)
	defer C.free(unsafe.Pointer(c_pvname))

	c_pv := C.lvm_pv_from_name(vg.c_vgh, c_pvname)
	if c_pv == nil {
		return nil, Lvm2Error{-1, "C function lvm_pv_from_name failed!"}
	}

	pv := new(PhysicalVolume)
	pv.c_pvh = c_pv
	pv.c_lvm2h = vg.c_lvm2h
	return pv, nil
}

func (vg *VolumeGroup) PvFromUUID(uuid string) (*PhysicalVolume, error) {
	c_uuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(c_uuid))

	c_pv := C.lvm_pv_from_uuid(vg.c_vgh, c_uuid)
	if c_pv == nil {
		return nil, Lvm2Error{-1, "Not Found error!"}
	}

	pv := new(PhysicalVolume)
	pv.c_pvh = c_pv
	pv.c_lvm2h = vg.c_lvm2h
	return pv, nil
}
