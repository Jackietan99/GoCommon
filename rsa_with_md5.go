package common

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

/**
* 定义结构体
 */
type RSAWITHMD5 struct {
	privateKey string //加密的私钥
	publicKey  string //加密的私钥
}

/**
*创建并实例化一个RsaWithMd5结构体
 */
func SetRsaWithMd5(privateKey, publicKey string) *RSAWITHMD5 {
	c := new(RSAWITHMD5)
	c.privateKey = privateKey
	c.publicKey = publicKey
	return c
}

/**
* 数据加密
* 算法：Rsa私钥加密(md5(原文))
* @param	str	string	需要加密的原文
 */
func (c *RSAWITHMD5) RsaEncrypt(str string) ([]byte, error) {
	//步骤一，先对原文进行md5
	h := md5.New()
	h.Write([]byte(str))
	origData := h.Sum(nil)

	//设置私钥
	rsa_privateKey := []byte(c.privateKey)
	block, _ := pem.Decode(rsa_privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	prkI, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priv := prkI.(*rsa.PrivateKey)

	result, er := rsa.SignPKCS1v15(rand.Reader, priv, crypto.MD5, origData)

	return result, er
}

/**
* rsa数据加密
* @param	str	string	需要加密的字符串
* @return	string	返回加密后的字符串
 */
func (c *RSAWITHMD5) RsaEncryptString(str string) string {
	encode_byte, err := c.RsaEncrypt(str)
	if err != nil {
		fmt.Println("aes_encode_error->", err.Error())
		return str
	} else {
		return base64.StdEncoding.EncodeToString(encode_byte)
	}
}
