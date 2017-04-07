package golvm2

// #cgo LDFLAGS: -llvm2app -L /usr/lib/
// #include <lvm2app.h>
// #include "lvm_property_extension.h"
import "C"
import "unsafe"

type PhysicalVolume struct {
	c_pvh   C.pv_t
	c_lvm2h C.lvm_t
}

type PhysicalVolumeSegment struct {
	c_pvsegh C.pvseg_t
	c_lvm2h  C.lvm_t
}

type PvDetail struct {
	name               string
	uuid               string
	size               uint64
	freeSize           uint64
	devSize            uint64
	metadataAreaNumber uint64
}

func (pv *PhysicalVolume) PvGetUUID() string {
	uuid := C.lvm_pv_get_uuid(pv.c_pvh)
	return C.GoString(uuid)
}

func (pv *PhysicalVolume) PvGetName() string {
	pvname := C.lvm_pv_get_name(pv.c_pvh)
	return C.GoString(pvname)
}

func (pv *PhysicalVolume) PvGetMdaCount() uint64 {
	mdaCount := C.lvm_pv_get_mda_count(pv.c_pvh)
	return uint64(mdaCount)
}

func (pv *PhysicalVolume) PvGetDevSize() uint64 {
	devSize := C.lvm_pv_get_dev_size(pv.c_pvh)
	return uint64(devSize)
}

func (pv *PhysicalVolume) PvGetSize() uint64 {
	size := C.lvm_pv_get_size(pv.c_pvh)
	return uint64(size)
}

func (pv *PhysicalVolume) PvGetFree() uint64 {
	freeSize := C.lvm_pv_get_free(pv.c_pvh)
	return uint64(freeSize)
}

func (pv *PhysicalVolume) PvGetProperty(prop string) (*LvmPropertyValue, error) {
	c_propname := C.CString(prop)
	defer C.free(unsafe.Pointer(c_propname))

	propValue := new(LvmPropertyValue)
	c_prop := C.lvm_pv_get_property_c(pv.c_pvh, c_propname)
	//Release memory which malloc in lvm_property_extension.c
	defer C.free(unsafe.Pointer(c_prop))

	if c_prop == nil {
		return nil, Lvm2Error{int(C.lvm_errno(pv.c_lvm2h)), C.GoString(C.lvm_errmsg(pv.c_lvm2h))}
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

func (pvseg *PhysicalVolumeSegment) PvsegGetProperty(prop string) (*LvmPropertyValue, error) {
	c_propname := C.CString(prop)
	defer C.free(unsafe.Pointer(c_propname))

	propValue := new(LvmPropertyValue)
	c_prop := C.lvm_pvseg_get_property_c(pvseg.c_pvsegh, c_propname)
	//Release memory which malloc in lvm_property_extension.c
	defer C.free(unsafe.Pointer(c_prop))

	if c_prop == nil {
		return nil, Lvm2Error{int(C.lvm_errno(pvseg.c_lvm2h)), C.GoString(C.lvm_errmsg(pvseg.c_lvm2h))}
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

func (pv *PhysicalVolume) PvListPvsegHandler() ([]*PhysicalVolumeSegment, error) {
	pvsegList := make([]*PhysicalVolumeSegment, 0, 10)
	var pvsegl *C.pvseg_list_t

	pvsegs := C.lvm_pv_list_pvsegs(pv.c_pvh)
	if pvsegs == nil {
		return nil, Lvm2Error{-1, "C Function lvm_pv_list_pvsegs error!"}
	}

	for current := pvsegs.n; current != pvsegs; current = current.n {
		pvsegl = (*C.pvseg_list_t)(unsafe.Pointer(current))

		pvsegh := &PhysicalVolumeSegment{
			c_pvsegh: pvsegl.pvseg,
			c_lvm2h:  pv.c_lvm2h,
		}
		pvsegList = append(pvsegList, pvsegh)
	}
	return pvsegList, nil
}

func (pv *PhysicalVolume) PvResize(newSize uint64) error {
	ret := C.lvm_pv_resize(pv.c_pvh, C.uint64_t(newSize))
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_pv_resize error!"}
	}
	return nil
}
