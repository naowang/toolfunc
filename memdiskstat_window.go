package toolfunc

import (
	"C"
	"log"
	"runtime"
	"syscall"
	"unsafe"
)

/*
#if defined(_WIN32) || defined(_WIN64)
#include <windows.h>
unsigned long long int GetFreeMem()  {
    MEMORYSTATUS ms;     //记录内容空间信息的结构体变量
    GlobalMemoryStatus(&ms);//调用GlobalMemoryStatus()函数获取内存信息
	return (unsigned long long int)ms.dwAvailPhys;
}
unsigned long long int GetTotalMem()  {
    MEMORYSTATUS ms;     //记录内容空间信息的结构体变量
    GlobalMemoryStatus(&ms);//调用GlobalMemoryStatus()函数获取内存信息
	return (unsigned long long int)ms.dwTotalPhys;
}
#endif
*/
import "C"

type MemStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Self uint64 `json:"self"`
}

func MemStat() MemStatus {
	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := MemStatus{}
	mem.Self = memStat.Alloc
	mem.All = uint64(C.GetTotalMem())
	mem.Free = uint64(C.GetFreeMem())
	mem.Used = mem.All - mem.Free

	return mem
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		log.Panic(err)
	}
	defer syscall.FreeLibrary(kernel32)
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	if err != nil {
		log.Panic(err)
	}

	lpFreeBytesAvailable := int64(0)
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)
	_, _, _ = syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)

	disk.All = uint64(lpTotalNumberOfBytes)
	disk.Free = uint64(lpTotalNumberOfFreeBytes)
	disk.Used = uint64(lpTotalNumberOfBytes - lpTotalNumberOfFreeBytes)
	return
}
