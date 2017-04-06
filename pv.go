package golvm2

// #cgo LDFLAGS: -llvm2app -L /usr/lib/
// #include <stdio.h>
// #include <lvm2app.h>
import "C"

type PhysicalVolume struct {
	c_pvh C.pv_t
}

type PvDetail struct {
	name               string
	uuid               string
	size               uint64
	freeSize           uint64
	devSize            uint64
	metadataAreaNumber uint64
}
