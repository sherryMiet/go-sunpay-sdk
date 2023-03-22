package go_sunpay_sdk

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"strings"
)

func SHA1(s string) *string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	res := fmt.Sprintf("%x", bs)
	res = strings.ToUpper(res)
	return &res
}

func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	res := fmt.Sprintf("%x", bs)
	return strings.ToUpper(res)
}
