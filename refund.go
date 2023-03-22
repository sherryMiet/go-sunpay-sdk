package go_sunpay_sdk

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"strconv"
	"strings"
)

const (
	RefundURL     = "https://www.esafe.com.tw/Service/Hx_CardRefund.ashx"
	TestRefundURL = "https://test.esafe.com.tw/Service/Hx_CardRefund.ashx"
)

type RefundDataRequest struct {
	Web        string `json:"web"`
	MN         string `json:"MN"`
	BuySafeNo  string `json:"buysafeno"`
	Td         string `json:"Td"`
	RefundMemo string `json:"RefundMemo"`
	ChkValue   string `json:"ChkValue"`
}

func NewRefund() *RefundDataRequest {
	return &RefundDataRequest{}
}

func (r *RefundDataRequest) CreateRefund(MN int, BuySafeNo, Td, RefundMemo string) *RefundDataRequest {
	r.MN = strconv.Itoa(MN)
	r.BuySafeNo = BuySafeNo
	r.Td = Td
	r.RefundMemo = RefundMemo
	return r
}

func (c *Client) Refund(r *RefundDataRequest) (*string, error) {
	r.Web = c.MerchantID

	var StrBuild strings.Builder
	StrBuild.WriteString(c.MerchantID)
	StrBuild.WriteString(c.TransPassword)
	StrBuild.WriteString(r.BuySafeNo)
	StrBuild.WriteString(r.MN)
	StrBuild.WriteString(r.Td)

	fmt.Println(StrBuild.String())
	r.ChkValue = SHA256(StrBuild.String())
	vals, _ := query.Values(r)
	body, err := SendSunPayRequest(vals, TestRefundURL)
	if err != nil {
		return nil, err
	}
	res := string(*body)
	return &res, nil
}
