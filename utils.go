package golvm2

import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unsafe"
)

// this method was tested in ubuntu14.04, centos6, centos7
func getDevNo(vgname, lvname string) (string, error) {
	devPath := fmt.Sprintf("/dev/%v/%v", vgname, lvname)
	realName, err := os.Readlink(devPath)
	if err != nil {
		return "", err
	}
	blockPath := fmt.Sprintf("/sys/devices/virtual/block/%v/dev", filepath.Base(realName))
	content, err := ioutil.ReadFile(blockPath)
	if err != nil {
		return "", err
	}
	devNo := strings.Trim(string(content), "\n")
	return devNo, nil
}

func unionToUint64(cbytes [8]byte) uint64 {
	return binary.LittleEndian.Uint64(cbytes[:])
}
func uint64ToUnion(num uint64) [8]byte {
	bSlice := make([]byte, 8)
	bArray := [8]byte{}
	binary.LittleEndian.PutUint64(bSlice, num)
	copy(bArray[:], bSlice)
	return bArray
}

func unionToCharptr(cbytes [8]byte) *C.char {
	buf := bytes.NewBuffer(cbytes[:])
	var ptr uint64
	if err := binary.Read(buf, binary.LittleEndian, &ptr); err == nil {
		uptr := uintptr(ptr)
		return (*C.char)(unsafe.Pointer(uptr))
	}
	return nil
}
func charptrToUnion(cString *C.char) [8]byte {
	addr := uintptr(unsafe.Pointer(cString))
	return uint64ToUnion(uint64(addr))
}
