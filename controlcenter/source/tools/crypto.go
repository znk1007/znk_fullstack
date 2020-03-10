package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const (
	key = "fullstack-znk$#!"
)

//CBCEncrypt CBC模式加密
func CBCEncrypt(org string) (string, error) {
	//转字节数组
	orgData := []byte(org)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	//获取密钥块长度
	blockSize := block.BlockSize()
	//补全码
	orgData = pkcs7Padding(orgData, blockSize)
	//加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	//创建数组
	crypted := make([]byte, len(orgData))
	//加密
	blockMode.CryptBlocks(crypted, orgData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func pkcs7UnPadding(orgData []byte) []byte {
	l := len(orgData)
	unpadding := int(orgData[l-1])
	return orgData[:(l - unpadding)]
}

//CBCDecrypt CBC解密
func CBCDecrypt(crypted string) (string, error) {
	cryptedByte, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	k := []byte(key)
	//分组密钥
	block, err := aes.NewCipher(k)
	//获取密钥块长度
	blockSize := block.BlockSize()
	//解密方式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	//创建数组
	org := make([]byte, len(cryptedByte))
	//解密
	blockMode.CryptBlocks(org, cryptedByte)
	//去补全码
	org = pkcs7UnPadding(org)
	return string(org), nil
}

//ECBEncrypt ECB加密
func ECBEncrypt(org string) (string, error) {
	c, err := aes.NewCipher(generateKey([]byte(key)))
	if err != nil {
		return "", err
	}
	src := []byte(org)
	l := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, l*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	//分组分块加密
	for bs, be := 0, c.BlockSize(); bs <= len(src); bs, be = bs+c.BlockSize(), be+c.BlockSize() {
		c.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return string(encrypted), nil
}

//ECBDecrypt ECB解密
func ECBDecrypt(encrypted string) (string, error) {
	c, err := aes.NewCipher(generateKey([]byte(key)))
	if err != nil {
		return "", nil
	}
	encryptedByte := []byte(encrypted)
	decrypted := make([]byte, len(encryptedByte))
	for bs, be := 0, c.BlockSize(); bs < len(encryptedByte); bs, be = bs+c.BlockSize(), be+c.BlockSize() {
		c.Decrypt(decrypted[bs:be], encryptedByte[bs:be])
	}
	trim := 0
	decryptedLen := len(decrypted)
	if decryptedLen > 0 {
		trim = decryptedLen - int(decrypted[decryptedLen-1])
	}
	return string(decrypted[:trim]), nil
}

//generateKey 生成加解密key
func generateKey(keyByte []byte) (genKey []byte) {
	keyLen := len(keyByte)
	genKey = make([]byte, 16)
	copy(genKey, keyByte)
	for i := 16; i < keyLen; {
		for j := 0; j < 16 && i < keyLen; j, i = j+1, i+1 {
			genKey[j] ^= keyByte[i]
		}
	}
	return genKey
}

//GetSecurityKey 获取加密密钥
func GetSecurityKey() string {
	return key
}
