package golvm2

// #cgo LDFLAGS: -llvm2app -L /usr/lib/
// #include <stdio.h>
// #include <lvm2app.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Lvm2Error struct {
	errCode int
	errMsg  string
}

func (e Lvm2Error) Error() string {
	return fmt.Sprintf("golvm2: %s ret=%d", e.errMsg, e.errCode)
}

type Lvm2Handler struct {
	c_lvm2h C.lvm_t
}

type LvmPropertyValue struct {
	isSettable bool
	isString   bool
	isInteger  bool
	isValid    bool
	value      interface{}
}

func NewLvm2Handler() (*Lvm2Handler, error) {
	handle := C.lvm_init(nil)
	if handle == nil {
		return nil, Lvm2Error{-1, "C Function lvm_init error!"}
	}
	lvmh := new(Lvm2Handler)
	lvmh.c_lvm2h = handle
	return lvmh, nil
}

func (lvmh *Lvm2Handler) Quit() {
	C.lvm_quit(lvmh.c_lvm2h)
}

func (lvmh *Lvm2Handler) Scan() error {
	ret := C.lvm_scan(lvmh.c_lvm2h)
	if ret != 0 {
		return Lvm2Error{int(ret), "C Function lvm_scan error!"}
	}
	return nil
}

/*Return the list of vg names.
typedef struct lvm_str_list {
        struct dm_list list;
        const char *str;
} lvm_str_list_t;
*/
func (lvmh *Lvm2Handler) ListVgNames() ([]string, error) {
	nameList := make([]string, 0, 10)
	var namel *C.lvm_str_list_t

	names := C.lvm_list_vg_names(lvmh.c_lvm2h)
	if names == nil {
		return nil, Lvm2Error{-1, "C function lvm_list_vg_names failed!"}
	}

	for current := names.n; current != names; current = current.n {
		namel = (*C.lvm_str_list_t)(unsafe.Pointer(current))
		nameList = append(nameList, C.GoString(namel.str))
	}
	return nameList, nil

}

/*Return the list of vg uuids.
typedef struct lvm_str_list {
        struct dm_list list;
        const char *str;
} lvm_str_list_t;
*/
func (lvmh *Lvm2Handler) ListVgUUIDs() ([]string, error) {
	uuidList := make([]string, 0, 10)
	var uuidl *C.lvm_str_list_t

	uuids := C.lvm_list_vg_uuids(lvmh.c_lvm2h)
	if uuids == nil {
		return nil, Lvm2Error{-1, "C function lvm_list_vg_uuids failed!"}
	}

	for current := uuids.n; current != uuids; current = current.n {
		uuidl = (*C.lvm_str_list_t)(unsafe.Pointer(current))
		uuidList = append(uuidList, C.GoString(uuidl.str))
	}
	return uuidList, nil

}

func (lvmh *Lvm2Handler) GetVgNameFromPvId(pvid string) (string, error) {
	c_pvid := C.CString(pvid)
	defer C.free(unsafe.Pointer(c_pvid))
	vgname := C.lvm_vgname_from_pvid(lvmh.c_lvm2h, c_pvid)
	if vgname == nil {
		return "", Lvm2Error{-1, "Get VgName From PvId error!"}
	}
	return C.GoString(vgname), nil
}

func (lvmh *Lvm2Handler) GetVgNameFromDevice(deviceid string) (string, error) {

	c_deviceid := C.CString(deviceid)
	defer C.free(unsafe.Pointer(c_deviceid))
	vgname := C.lvm_vgname_from_device(lvmh.c_lvm2h, c_deviceid)
	if vgname == nil {
		return "", Lvm2Error{-1, "Get VgName From Device error!"}
	}
	return C.GoString(vgname), nil
}

func (lvmh *Lvm2Handler) VgOpen(vgname, mode string, flags int) (*VolumeGroup, error) {
	if mode != "r" && mode != "w" {
		return nil, Lvm2Error{-1, "Paramemter is not valid!"}
	}
	c_vgname := C.CString(vgname)
	defer C.free(unsafe.Pointer(c_vgname))
	c_mode := C.CString(mode)
	defer C.free(unsafe.Pointer(c_mode))
	c_vg := C.lvm_vg_open(lvmh.c_lvm2h, c_vgname, c_mode, C.uint32_t(flags))
	if c_vg == nil {
		return nil, Lvm2Error{-1, "C Function lvm_vg_open error!"}
	}
	vg := new(VolumeGroup)
	vg.c_vgh = c_vg
	vg.c_lvm2h = lvmh.c_lvm2h
	return vg, nil
}

func (lvmh *Lvm2Handler) VgCreate(vgname string) (*VolumeGroup, error) {
	c_vgname := C.CString(vgname)
	defer C.free(unsafe.Pointer(c_vgname))
	c_vg := C.lvm_vg_create(lvmh.c_lvm2h, c_vgname)
	if c_vg == nil {
		return nil, Lvm2Error{-1, "C Function lvm_vg_create error!"}
	}
	vg := new(VolumeGroup)
	vg.c_vgh = c_vg
	vg.c_lvm2h = lvmh.c_lvm2h
	return vg, nil
}
