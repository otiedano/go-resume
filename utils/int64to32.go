package utils

import "unsafe"

//Int64to32 转换int格式
func Int64to32(id64 int64) int {
	return *(*int)(unsafe.Pointer(&id64))
}
