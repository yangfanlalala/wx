package aes

import (
	"encoding/base64"
)

func base64decode(txt []byte) ([]byte, error) {
	l := len(txt) / 4 * 3
	dec := make([]byte, l)
	l, err := base64.StdEncoding.Decode(dec, txt)
	if err != nil {
		return nil, err
	}
	return dec[:l], nil
}

func WxDecrypt(cipherTxt, key []byte) ([]byte, error){
	//base64decode
	cipherTxt, err := base64decode(cipherTxt)
	if err != nil {
		return nil, err
	}
	iv := key[:16]
	key, err= base64decode(key)
	if err != nil {
		return nil, err
	}
	plainTxt, err := Decrypt(cipherTxt, key, iv, CBC, PKCS7Padding)
	if err != nil {
		return nil, err
	}
	return plainTxt, nil
}
