package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

/**
* 定义结构体
 */
type RSA struct {
	privateKey string //加密的私钥
	publicKey  string //加密的私钥
}

/**
*创建并实例化一个aes结构体
 */
func SetRSA(privateKey, publicKey string) *RSA {
	c := new(RSA)
	c.privateKey = privateKey
	c.publicKey = publicKey
	return c
}

/**
* rsa数据加密
* @param	str	string	需要加密的字符串
* @return	string	返回加密后的字符串
 */
func (c *RSA) RsaEncryptString(str string) string {
	encode_before := []byte(str)
	encode_end, err := c.RsaEncrypt(encode_before)
	if err != nil {
		//加密失败
		return str
	} else {
		return base64.StdEncoding.EncodeToString(encode_end)
	}
}

/**
* rsa数据解密
* @param	str	string	需要解密的字符串
* @param 	keyType	string	私钥的类型(PKCS1,PKCS8)
* @return	string	返回解密后的字符串
 */
func (c *RSA) RsaDecryptString(str, keyType string) string {
	if len(str) < 1 {
		return ""
	}
	decode_before, er := base64.StdEncoding.DecodeString(str)
	if er != nil {
		//解密失败
		return ""
	}
	decode_end, err := c.RsaDecrypt(decode_before, keyType)
	if err != nil {
		//解密失败
		return ""
	}
	return string(decode_end)
}

// 加密获得byte
func (c *RSA) RsaEncrypt(origData []byte) ([]byte, error) {
	rsa_publicKey := []byte(c.publicKey)
	block, _ := pem.Decode(rsa_publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密得到byte
func (c *RSA) RsaDecrypt(ciphertext []byte, keyType string) ([]byte, error) {
	rsa_privateKey := []byte(c.privateKey)

	block, _ := pem.Decode(rsa_privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	var err error
	var priv *rsa.PrivateKey
	if keyType == "PKCS1" {
		priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	if keyType == "PKCS8" {
		prkI, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		priv = prkI.(*rsa.PrivateKey)
	}

	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
