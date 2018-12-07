package common

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

/**
* 整形转byte数组
 */
func IntToByte(data int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((data>>shift)&mask))
	}
	return result
}

/**
* byte数组转int64
 */
func ByteToInt(data []byte) int64 {
	b_buf := bytes.NewBuffer(data)
	var res int32
	binary.Read(b_buf, binary.BigEndian, &res)
	return int64(res)
}

func ByteToUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

//Int返回绝对值
func AbsInt(int1 int) int {

	int2 := int1
	if int1 < 0 {
		int2 = 0 - int1
	}
	return int2
}

//string to int
func Str2Int(s string) (int, error) {
	return strconv.Atoi(s)
}

//string to int64
func Str2Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
