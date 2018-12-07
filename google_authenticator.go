package common

/**
* 谷歌动态验证
 */

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	//	"math"
	"time"
)

/**
* 定义结构体
 */
type GoogleAuth struct {
	IntervalLength int64 //验证码的有效时间 30
	PinLength      int   //验证码长度 6
}

/**
*创建并实例化一个GoogleAuth结构体
 */
func SetGoogleAuth(intervalLength int64, pinLength int) *GoogleAuth {
	c := new(GoogleAuth)
	c.IntervalLength = intervalLength
	c.PinLength = pinLength
	return c
}

/**
* 比较用户输入的谷歌验证码是否正确
* @Secret_char   int64  数据库保存的Secret_char
* @codeValue    int64  用户输入的谷歌验证码
* @return int 状态值 200表示正确
   string   结果说明
*/
func (g *GoogleAuth) CheckGoogleCode(Secret_char string, codeValue int64) (int, string) {
	common_checkgooglecode_status := 100
	common_checkgooglecode_msg := "验证码错误"

	secretUnix := time.Now().Unix()
	common_checkgooglecode_value := IntToByte(secretUnix / g.IntervalLength)

	common_checkgooglecode_key, err := base32.StdEncoding.DecodeString(Secret_char)

	if err == nil {
		// sign the value using HMAC-SHA1
		hmacSha1 := hmac.New(sha1.New, common_checkgooglecode_key)
		hmacSha1.Write(common_checkgooglecode_value)
		hash := hmacSha1.Sum(nil)

		// We're going to use a subset of the generated hash.
		// Using the last nibble (half-byte) to choose the index to start from.
		// This number is always appropriate as it's maximum decimal 15, the hash will
		// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
		offset := hash[len(hash)-1] & 0x0F

		// get a 32-bit (4-byte) chunk from the hash starting at offset
		hashParts := hash[offset : offset+4]

		// ignore the most significant bit as per RFC 4226
		hashParts[0] = hashParts[0] & 0x7F

		number := ByteToUint32(hashParts)

		// size to 6 digits
		// one million is the first number with 7 digits so the remainder
		// of the division will always return < 7 digits

		common_checkgooglecode_pwd := number % 1000000
		if int64(common_checkgooglecode_pwd) == codeValue {
			common_checkgooglecode_status = 200
			common_checkgooglecode_msg = "验证码正确"
		}
	}

	return common_checkgooglecode_status, common_checkgooglecode_msg
}

/**
* 当前时间戳
* return int64
 */
func GetSecret() int64 {
	secret := time.Now().Unix()
	return secret
}

/**
* 获取谷歌验证二维码的url
* @identifier string 谷歌的账号(TBET)
* @key
* @width int 图片的宽
* @height int 图片的高
 */
func GetImageUrl(identifier string, key []byte, width, height int) string {
	keyString := base32.StdEncoding.EncodeToString(key)

	provisionUrl := fmt.Sprintf("otpauth://totp/%s?secret=%s", identifier, keyString)
	chartUrl := fmt.Sprintf("http://chart.apis.google.com/chart?cht=qr&chs=%dx%d&chld=L|1&chl=%s", width, height, provisionUrl)
	return chartUrl
}
