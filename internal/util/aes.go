package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"github.com/spf13/viper"
)

var AES = &Aes{}

type Aes struct {
	key   []byte
	block cipher.Block
}

func init() {
	var err error
	var configViperConfig = viper.New()
	configViperConfig.SetConfigName("config")
	configViperConfig.SetConfigType("yaml")
	configViperConfig.AddConfigPath("./configs/")

	//读取配置文件内容
	if err = configViperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	//读取配置文件内容
	if err = configViperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	AES.key = []byte(configViperConfig.GetString("server.ase_key"))

	AES.block, err = aes.NewCipher(AES.key)
	if err != nil {
		panic(err)
	}
}

func (u *Aes) pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func (u *Aes) pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

func (u *Aes) AesEncrypt(in string) string {
	data := []byte(in)
	blockSize := u.block.BlockSize()
	encryptBytes := u.pkcs7Padding(data, blockSize)
	crypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(u.block, u.key[:blockSize])
	blockMode.CryptBlocks(crypted, encryptBytes)
	return hex.EncodeToString(crypted)
}

func (u *Aes) AesDecrypt(in string) string {
	data, _ := hex.DecodeString(in)
	blockSize := u.block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(u.block, u.key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted = u.pkcs7UnPadding(crypted)
	return string(crypted)
}

// EncryptByAes Aes加密 后 base64 再加
func (u *Aes) EncryptByAes(in string) string {
	res := u.AesEncrypt(in)
	return base64.StdEncoding.EncodeToString([]byte(res))
}

// DecryptByAes Aes 解密
func (u *Aes) DecryptByAes(in string) (out string, err error) {
	var dataByte []byte
	dataByte, err = base64.StdEncoding.DecodeString(in)
	if err != nil {
		return
	}
	out = u.AesDecrypt(string(dataByte))
	return
}
