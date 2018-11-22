package util

import "encoding/binary"

// Uint32ToBytes unit32转换为byte数组。
func Uint32ToBytes(n uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, n)

	return bytes
}

// BytesToUint32 byte数组转换为unit32。
func BytesToUint32(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}

// Uint64ToBytes unit64转换为byte数组。
func Uint64ToBytes(n uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, n)

	return bytes
}

// BytesToUint64 byte数组转换为unit64。
func BytesToUint64(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

// Int64ToBytes nit64转换为byte数组。
func Int64ToBytes(n int64) []byte {
	return Uint64ToBytes(uint64(n))
}

// BytesToInt64 byte数组转换为nit64。
func BytesToInt64(bytes []byte) int64 {
	return int64(BytesToUint64(bytes))
}
