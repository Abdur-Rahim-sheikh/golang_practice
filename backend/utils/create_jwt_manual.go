package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwtManual(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteHeader, _ := json.Marshal(header)
	byteSecret, _ := json.Marshal(secret)
	headerbs64 := base64UrlEncode(byteHeader)
	dataArray, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	databs64 := base64UrlEncode(dataArray)
	message := headerbs64 + "." + databs64
	byteMsg := []byte(message)
	h := hmac.New(sha256.New, byteSecret)
	h.Write(byteMsg)
	signature := h.Sum(nil)
	signatureb64 := base64UrlEncode(signature)
	jwt := headerbs64 + "." + databs64 + "." + signatureb64
	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
