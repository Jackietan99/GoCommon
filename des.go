package common

import (
	"bytes"
	//	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	//"fmt"
	"fmt"
	"strings"
	"time"
)

/**
* des加密
 */
func DesEncryptString(str, keyStr string) string {
	//key := make([]byte, 8)    //设置加密数组
	//copy(key, []byte(keyStr)) //合并数组补位
	result, err := DDesEncrypt([]byte(str), []byte(keyStr))
	res := ""
	if err == nil {
		res = base64.StdEncoding.EncodeToString(result)
	}
	return res
}

func DDesEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	//src = DZeroPadding(src, bs)
	src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

/**
* des解密
 */
func DesDecryptString(str, key string) string {
	result, _ := base64.StdEncoding.DecodeString(str)
	res := ""
	if len(result) > 0 {
		origData, err := DDesDecrypt(result, []byte(key))
		if err == nil {
			res = string(origData)
		}
	}
	return res
}

func DDesDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	//out = ZeroUnPadding(out)
	out = PKCS5UnPadding(out)
	return out, nil
}

/**
* Zero补位算法
 */
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
* PKCS5补位算法
 */
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Encrypt(sString, sOperation, sKey string, iExpiry int64) string {
	iChekeyLength := 4
	now := time.Now()
	secs := now.Unix()
	//nanos := now.UnixNano()
	//millis := nanos / 1000000

	if len(sKey) > 0 {
		sKey = GetMd5(sKey)
	} else {
		sKey = GetMd5("US_KEY")

	}

	sKeyA := GetMd5(Substr(sKey, 0, 16))
	sKeyB := GetMd5(Substr(sKey, 16, 16))
	sKeyC := ""

	if sOperation == "DECODE" {
		sKeyC = Substr(sString, 0, iChekeyLength)
	} else {
		sKeyC = Substr(GetMd5("0.16561000 "+fmt.Sprintf("%d",secs)), len(GetMd5("0.16561000 "+fmt.Sprintf("%d",secs)))-4, 4)
	}

	sCryptkey := sKeyA + GetMd5(sKeyA+sKeyC)

	ikeyLength := len(sCryptkey)

	if sOperation == "DECODE" {
		//fmt.Println(common.Substr(sString, iChekeyLength, len(sString)))
		sString = Base64DecodeNoPadding(Substr(sString, iChekeyLength, len(sString)))
	} else {
		if iExpiry > 0 {
			iExpiryAndTime := iExpiry + secs
			sString = fmt.Sprintf("%010d", iExpiryAndTime) + Substr(GetMd5(sString+sKeyB), 0, 16) + sString
		} else {
			sString = fmt.Sprintf("%010d", 0) + Substr(GetMd5(sString+sKeyB), 0, 16) + sString
		}
	}

	sStringLength := len(sString)

	sResult := ""

	var ABox [256]int
	for i := 0; i <= 255; i++ {
		ABox[i] = i
	}

	var ARandKey [256]int

	for i := 0; i <= 255; i++ {
		var c rune = int32(sCryptkey[i%ikeyLength])

		ordC := int(c)

		ARandKey[i] = ordC
	}

	j := 0
	for i := 0; i < 256; i++ {
		j = (j + ABox[i] + ARandKey[i]) % 256
		ATmp := ABox[i]
		ABox[i] = ABox[j]
		ABox[j] = ATmp
	}

	a  := 0
	j1 := 0

	var sByte []byte
	for i1 := 0; i1 < sStringLength; i1++{
		a = (a + 1) % 256
		j1 = (j1 + ABox[a]) % 256
		ATmp := ABox[a]
		ABox[a] = ABox[j1]
		ABox[j1] = ATmp
		var m rune = int32(sString[i1]) ^ int32((ABox[(ABox[a]+ABox[j1])%256]))
		b1 := byte(int(m))
		sByte = append(sByte,b1)
	}
	sResult = string(sByte)

	if sOperation == "DECODE" {

		fmt.Println("sresutl2->"+sResult)
		iResult10, _ := Str2Int64(Substr(sResult, 0, 10))
		iResult16 := Substr(sResult, 10, 16)
		if (iResult10 == 0 || iResult10-secs > 0) && iResult16 == Substr(GetMd5(Substr(sResult, 26, len(sResult))+sKeyB), 0, 16) {
			return Substr(sResult, 26, len(sResult))
		} else {
			return ""
		}
	} else {
		return sKeyC+strings.TrimRight(Base64Encode(sResult), "=")
	}

}
