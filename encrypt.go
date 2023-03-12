package go_sunpay_sdk

import (
	"crypto/sha1"
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
