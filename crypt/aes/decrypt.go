package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

const (
	PKCS7Padding = iota
	PKCS5Padding
)

const (
	CBC = iota
	CTR
)

func Encrypt(plainTxt, key, iv []byte, mod, pad int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if pad == PKCS7Padding {
		plainTxt = pkcs7padding(plainTxt, block.BlockSize())
	} else {
		return nil, errors.New("padding mode not implement")
	}
	var blockMode cipher.BlockMode
	if mod == CBC {
		blockMode = cipher.NewCBCEncrypter(block, iv)
	} else if mod == CTR {
		blockMode = cipher.NewCBCEncrypter(block, iv)
	} else {
		return nil, errors.New("block mode not implement")
	}
	cipherTxt := make([]byte, len(plainTxt))
	blockMode.CryptBlocks(cipherTxt, plainTxt)
	return cipherTxt, nil
}

func Decrypt(cipherTxt, key, iv []byte, mod, pad int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockMode cipher.BlockMode
	if mod == CBC {
		blockMode = cipher.NewCBCDecrypter(block, iv)
	} else if mod == CTR {
		blockMode = cipher.NewCBCDecrypter(block, iv)
	} else {
		return nil, errors.New("block mode not implement")
	}
	plainTxt := make([]byte, len(cipherTxt))
	blockMode.CryptBlocks(plainTxt, cipherTxt)
	if pad == PKCS7Padding {
		plainTxt = pkcs7unpadding(plainTxt)
	} else {
		return nil, errors.New("padding mode not implement")
	}
	return plainTxt, nil
}

func pkcs7padding(cipherTxt []byte, blockSize int) []byte {
	padding := blockSize - len(cipherTxt)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherTxt, padText...)
}

func pkcs7unpadding(plainTxt []byte) []byte {
	length := len(plainTxt)
	padding := int(plainTxt[length-1])
	return plainTxt[:(length - padding)]
}
