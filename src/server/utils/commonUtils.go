package utils

import (
	"log"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func DelError(err error) {
	if (err != nil) {
		log.Println("error:", err)
	}
}

func Hmac(data ...string) string {
	h := hmac.New(md5.New, []byte(Key))
	var content string
	for i:=0;i<len(data) ;i++  {
		content += data[i]
	}
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}