package common

import (
	"encoding/base64"
	//	"fmt"
	"fmt"
)

/**
* 加密base64
 */
func Base64Encode(msg string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}

/**
* 解密base64
 */
func Base64Decode(msg string) string {
	res := ""
	decoded, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		fmt.Println("decode error:", err.Error())
		return ""
	}
	res = string(decoded)
	return res
}

func Base64DecodeNoPadding(msg string) string {
	res := ""
	out, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(msg)
	if err != nil {
		fmt.Println("decode error:", err.Error())
		return ""
	}
	res = string(out)
	return res
}
