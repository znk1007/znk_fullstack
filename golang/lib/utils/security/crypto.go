package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

const (
	// Key 加解密key
	Key = "aes-znk1007man20"
)

// AESDecode aes指定key解密
func AESDecode(src string) (string, error) {
	return AesDecodeStr(src, Key)
}

// MD5Str md5验证
func MD5Str(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha256 hmacsha256验证
func HmacSha256(src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// HmacSha512 hmacsha512验证
func HmacSha512(src, key string) string {
	m := hmac.New(sha512.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// HmacSha1 hmacsha1验证
func HmacSha1(src, key string) string {
	m := hmac.New(sha1.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// Sha256Str sha256验证
func Sha256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512Str sha512验证
func Sha512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

// Base64EncodeStr base编码
func Base64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

// Base64DecodeStr base解码
func Base64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(a)
}

var ivspec = []byte("0000000000000000")

// AesEncodeStr aes加密
func AesEncodeStr(src, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, ivspec)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return hex.EncodeToString(crypted)
}

// AesDecodeStr aes解密
func AesDecodeStr(crypt, key string) (string, error) {
	crypted, err := hex.DecodeString(strings.ToLower(crypt))
	if err != nil {
		fmt.Println("plain content empty")
		return "", err
	}
	if len(crypted) == 0 {
		return "", errors.New("crypted empty")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
		return "", err
	}
	ecb := cipher.NewCBCDecrypter(block, ivspec)
	decrypted := make([]byte, len(crypted))
	ecb.CryptBlocks(decrypted, crypted)

	return string(PKCS5Trimming(decrypted)), nil
}

// PKCS5Padding PKCS5Padding边距
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Trimming 去头尾
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// RsaEncrypt rsa加密
func RsaEncrypt(src, key string) string {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return ""
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		// logger.SysLogger.Err(err.Error())
		return ""
	}

	pub := pubInterface.(*rsa.PublicKey)

	crypted, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(src))
	if err != nil {
		// logger.SysLogger.Err(err.Error())
		return ""
	}

	return hex.EncodeToString(crypted)
}
