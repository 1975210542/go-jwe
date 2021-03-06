package hma

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"jwe/methodit"
)

type EncryptionMethodHMAC struct {
	Name string
}

var (
	EncryptionMethodHMAC256 *EncryptionMethodHMAC
)

func init() {
	// RS256
	fmt.Println("初始化HMAC")
	EncryptionMethodHMAC256 = &EncryptionMethodHMAC{"HS256"}
	methodit.RegisterSigningMethod(EncryptionMethodHMAC256.GetName(), func() methodit.EncryptionMethod {
		return EncryptionMethodHMAC256
	})

}

func (e *EncryptionMethodHMAC) GetName() string {
	return e.Name
}

func (e *EncryptionMethodHMAC) Encrypt(plantText []byte, key interface{}) ([]byte, error) {
	fmt.Println("HMAC,加密方法")
	return hmacEncrpt(plantText, key)
}

func (e *EncryptionMethodHMAC) Decrypt(cipherText []byte, key interface{}) ([]byte, error) {
	fmt.Println("HMAC,验证")
	return hmacDecrpt(cipherText, key)
}

func (e *EncryptionMethodHMAC) Verify(MessageHMAC, message string, key interface{}) bool {
	fmt.Println("HMAC验证！！！")
	messageHMAC, _ := hex.DecodeString(MessageHMAC)
	mac := hmac.New(sha256.New, key.([]byte))
	mac.Write([]byte(message))
	exmac := mac.Sum(nil)
	return hmac.Equal(messageHMAC, exmac)
}

/**********************************************************************/

func hmacEncrpt(plantText []byte, key interface{}) ([]byte, error) {

	if keyBytes, ok := key.([]byte); ok {

		mac := hmac.New(sha256.New, keyBytes)
		mac.Write(plantText)
		return []byte(hex.EncodeToString(mac.Sum(nil))), nil
	}
	return []byte{}, errors.New("key is invalid")
}

func hmacDecrpt(cipherText []byte, key interface{}) ([]byte, error) {

	return []byte{}, errors.New("AAAAA")
}
