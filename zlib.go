package common

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"net/url"
	"strconv"
)

//进行zlib压缩
func EnZlib(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DeZlib(compressSrc []byte) []byte {
	r, err := zlib.NewReader(bytes.NewReader(compressSrc))
	if err != nil {
		panic(err)
	}
	enflated, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return enflated
}

//投注數據解壓三次
func DeCodeBetData(sStr string, iCount int) string {
	sStr, _ = url.QueryUnescape(sStr)
	for i := 0; i < iCount; i++ {
		if i >= 1 {
			sStr, _ = strconv.Unquote(sStr)
		}
		b := Base64Decode(sStr)
		sStr = string(DeZlib([]byte(b)))

	}
	return sStr
}
