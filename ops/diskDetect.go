package ops

import (
	"fmt"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)
// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

//func DiskDetect(){
//	mounts ,_ := gofstab.ParseSystem()
//	if mounts == nil {
//		fmt.Printf("空的别看了\n")
//	}
//
//	for _,val := range mounts{
//		//fmt.Printf("%v\n",val.File)
//		if val.File == "swap"||val.File == "/dev/shm"||val.File == "/dev/pts"||val.File == "/proc"||val.File =="/sys"{
//			continue
//		}
//		disk := DiskUsage(val.File)
//		fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
//		fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
//		fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
//
//		diskAll:=float64(disk.All)/float64(GB)
//		diskFree:= float64(disk.Free)/float64(GB)
//
//		dfPercent:=float64(diskFree/diskAll)
//		fmt.Printf("%s %.2f%%\n",val.File, dfPercent*100)
//	}
//}

//windows下的解决方法
//type DiskStatus struct {
//	All  uint64
//	Used uint64
//	Free uint64
//}
//
//func DiskUsage(path string) (disk DiskStatus) {
//	h := windows.MustLoadDLL("kernel32.dll")
//	c := h.MustFindProc("GetDiskFreeSpaceExW")
//	lpFreeBytesAvailable := uint64(0)
//	lpTotalNumberOfBytes := uint64(0)
//	lpTotalNumberOfFreeBytes := uint64(0)
//	r1, r2, err := c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("C:"))),
//		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
//		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
//		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
//	disk.All = lpTotalNumberOfBytes
//	disk.Free = lpTotalNumberOfFreeBytes
//	disk.Used = lpFreeBytesAvailable
//	return
//}

func DirDetect()  {
	disk := DiskUsage("/")
	fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))

	diskAll:=float64(disk.All)/float64(GB)
	diskFree:= float64(disk.Free)/float64(GB)

	dfPercent:=float64(diskFree/diskAll)
	fmt.Printf("%s %.2f%%\n","根目录磁盘可用百分比：", dfPercent*100)
}
