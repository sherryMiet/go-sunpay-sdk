package go_sunpay_sdk

import (
	"reflect"
	"strconv"
)

const (
	EtopmURL     = "https://www.esafe.com.tw/Service/Etopm.aspx"
	TestEtopmURL = "https://test.esafe.com.tw/Service/Etopm.aspx"
)

var (
	Card_Type_Card = "0"
)

var (
	AgencyType_ATM  = "2"
	AgencyType_CVS  = "1"
	AgencyBank_LAND = "T"
	FixedQuantity   = "1"
)

type ETOPMRequestData struct {
	Web               *string `json:"web,omitempty"`
	MN                *string `json:"MN,omitempty"`
	OrderInfo         *string `json:"OrderInfo,omitempty"`
	Td                *string `json:"Td,omitempty"`
	SNA               *string `json:"sna,omitempty"`
	SDT               *string `json:"sdt,omitempty"`
	Email             *string `json:"email,omitempty"`
	Note1             *string `json:"note1,omitempty"`
	Note2             *string `json:"note2,omitempty"`
	Card_Type         *string `json:"Card_Type,omitempty"`
	Country_Type      *string `json:"Country_Type,omitempty"`
	Term              *string `json:"Term,omitempty"`
	DueDate           *string `json:"DueDate,omitempty"`
	UserNo            *string `json:"UserNo,omitempty"`
	BillDate          *string `json:"BillDate,omitempty"`
	ProductName1      string  `json:"ProductName1,omitempty"`
	ProductName2      string  `json:"ProductName2,omitempty"`
	ProductName3      string  `json:"ProductName3,omitempty"`
	ProductName4      string  `json:"ProductName4,omitempty"`
	ProductName5      string  `json:"ProductName5,omitempty"`
	ProductName6      string  `json:"ProductName6,omitempty"`
	ProductName7      string  `json:"ProductName7,omitempty"`
	ProductName8      string  `json:"ProductName8,omitempty"`
	ProductName9      string  `json:"ProductName9,omitempty"`
	ProductName10     string  `json:"ProductName10,omitempty"`
	ProductPrice1     string  `json:"ProductPrice1,omitempty"`
	ProductPrice2     string  `json:"ProductPrice2,omitempty"`
	ProductPrice3     string  `json:"ProductPrice3,omitempty"`
	ProductPrice4     string  `json:"ProductPrice4,omitempty"`
	ProductPrice5     string  `json:"ProductPrice5,omitempty"`
	ProductPrice6     string  `json:"ProductPrice6,omitempty"`
	ProductPrice7     string  `json:"ProductPrice7,omitempty"`
	ProductPrice8     string  `json:"ProductPrice8,omitempty"`
	ProductPrice9     string  `json:"ProductPrice9,omitempty"`
	ProductPrice10    string  `json:"ProductPrice10,omitempty"`
	ProductQuantity1  string  `json:"ProductQuantity1,omitempty"`
	ProductQuantity2  string  `json:"ProductQuantity2,omitempty"`
	ProductQuantity3  string  `json:"ProductQuantity3,omitempty"`
	ProductQuantity4  string  `json:"ProductQuantity4,omitempty"`
	ProductQuantity5  string  `json:"ProductQuantity5,omitempty"`
	ProductQuantity6  string  `json:"ProductQuantity6,omitempty"`
	ProductQuantity7  string  `json:"ProductQuantity7,omitempty"`
	ProductQuantity8  string  `json:"ProductQuantity8,omitempty"`
	ProductQuantity9  string  `json:"ProductQuantity9,omitempty"`
	ProductQuantity10 string  `json:"ProductQuantity10,omitempty"`
	AgencyType        *string `json:"AgencyType,omitempty"`
	AgencyBank        *string `json:"AgencyBank,omitempty"`
	CargoFlag         *string `json:"CargoFlag,omitempty"`
	StoreID           *string `json:"StoreID,omitempty"`
	StoreName         *string `json:"StoreName,omitempty"`
	BuyerCid          *string `json:"BuyerCid,omitempty"`
	DonationCode      *string `json:"DonationCode,omitempty"`
	Carrier_ID        *string `json:"Carrier_ID,omitempty"`
	EDI               *string `json:"Carrier_ID,omitempty"`
	ChkValue          *string `json:"ChkValue,omitempty"`
}

type Product struct {
	ProductName     string
	ProductPrice    int
	ProductQuantity int
}

type EDI struct {
	EDI_Name    string
	EDI_Tel     string
	EDI_Address string
	EDI_Size    string
	EDI_Type    string
}

type ETOPMResponseData struct {
	BuySafeNo    string `form:"buysafeno"`
	Web          string `form:"web"`
	Td           string `form:"Td"`
	MN           string `form:"MN"`
	WebName      string `form:"webname"`
	Name         string `form:"Name"`
	Note1        string `form:"note1"`
	Note2        string `form:"note2"`
	ApproveCode  string `form:"ApproveCode"`
	CardNo       string `form:"Card_NO"`
	SendType     string `form:"SendType"`
	ErrCode      string `form:"errcode"`
	ErrMsg       string `form:"errmsg"`
	CardType     string `form:"Card_Type"`
	UserNo       string `form:"UserNo"`
	BarcodeA     string `form:"BarcodeA"`
	BarcodeB     string `form:"BarcodeB"`
	BarcodeC     string `form:"BarcodeC"`
	PostBarcodeA string `form:"PostBarcodeA"`
	PostBarcodeB string `form:"PostBarcodeB"`
	PostBarcodeC string `form:"PostBarcodeC"`
	BankCode     string `form:"BankCode"`
	EntityATM    string `form:"EntityATM"`
	BankName     string `form:"BankName"`
	PayCode      string `form:"paycode"`
	PayType      string `form:"PayType"`
	CargoNo      string `form:"CargoNo"`
	StoreID      string `form:"StoreID"`
	StoreName    string `form:"StoreName"`
	InvoiceNo    string `form:"InvoiceNo"`
	ChkValue     string `form:"ChkValue"`
}

func (e *ETOPMRequestData) CreateOrder(MN string, Td string) *ETOPMRequestData {
	e.MN = &MN
	e.Td = &Td
	return e
}

func (e *ETOPMRequestData) SetUserInfo(Name, Phone, Email string) *ETOPMRequestData {
	e.SNA = &Name
	e.SDT = &Phone
	e.Email = &Email
	return e
}

func (e *ETOPMRequestData) SetCard(Term string) *ETOPMRequestData {
	e.Card_Type = &Card_Type_Card
	e.Term = &Term
	return e
}

func (e *ETOPMRequestData) SetATM(AgencyBank, DueDate string) *ETOPMRequestData {
	e.AgencyType = &AgencyType_ATM
	e.DueDate = &DueDate
	if AgencyBank != "" {
		e.AgencyBank = &AgencyBank
	} else {
		e.AgencyBank = &AgencyBank_LAND
	}
	return e
}

func (e *ETOPMRequestData) SetCVS(DueDate string) *ETOPMRequestData {
	e.AgencyType = &AgencyType_CVS
	e.DueDate = &DueDate
	return e
}

func (e *ETOPMRequestData) SetProduct(Data []Product) *ETOPMRequestData {
	for i, p := range Data {
		if i >= 10 {
			return e
		}
		reflect.ValueOf(e).Elem().FieldByName("ProductName" + strconv.Itoa(i+1)).SetString(p.ProductName)
		reflect.ValueOf(e).Elem().FieldByName("ProductPrice" + strconv.Itoa(i+1)).SetString(strconv.Itoa(p.ProductPrice))
		reflect.ValueOf(e).Elem().FieldByName("ProductQuantity" + strconv.Itoa(i+1)).SetString(strconv.Itoa(p.ProductQuantity))
	}
	return e
}

func (c *Client) ETOPM(Data *ETOPMRequestData) string {
	Data.Web = &c.MerchantID
	if Data.Term != nil {
		Data.ChkValue = SHA1(c.MerchantID + c.TransPassword + *Data.MN + *Data.Term)
	} else {
		Data.ChkValue = SHA1(c.MerchantID + c.TransPassword + *Data.MN)
	}

	params := StructToParamsMap(Data)
	html := GenerateAutoSubmitHtmlForm(params, EtopmURL)
	return html
}

func (c *Client) ETOPMTest(Data *ETOPMRequestData) string {
	Data.Web = &c.MerchantID
	if Data.Term != nil {
		Data.ChkValue = SHA1(c.MerchantID + c.TransPassword + *Data.MN + *Data.Term)
	} else {
		Data.ChkValue = SHA1(c.MerchantID + c.TransPassword + *Data.MN)
	}

	params := StructToParamsMap(Data)
	html := GenerateAutoSubmitHtmlForm(params, TestEtopmURL)
	return html
}

func NewRequestData() *ETOPMRequestData {
	return &ETOPMRequestData{}
}

func NewResponseData() *ETOPMResponseData {
	return &ETOPMResponseData{}
}
