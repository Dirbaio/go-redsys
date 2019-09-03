package redsys

import (
	"encoding/xml"
)

type Request struct {
	XMLName            xml.Name      `xml:"DATOSENTRADA" json:"-"`
	OperationType      OperationType `xml:"DS_MERCHANT_TRANSACTIONTYPE" json:"DS_MERCHANT_TRANSACTIONTYPE"`
	MerchantCode       string        `xml:"DS_MERCHANT_MERCHANTCODE" json:"DS_MERCHANT_MERCHANTCODE"`
	Terminal           int           `xml:"DS_MERCHANT_TERMINAL" json:"DS_MERCHANT_TERMINAL"`
	OperationID        string        `xml:"DS_MERCHANT_ORDER,omitempty" json:"DS_MERCHANT_ORDER,omitempty"`
	Amount             int64         `xml:"DS_MERCHANT_AMOUNT" json:"DS_MERCHANT_AMOUNT,string"`
	Currency           Currency      `xml:"DS_MERCHANT_CURRENCY,omitempty" json:"DS_MERCHANT_CURRENCY,omitempty"`
	Language           Language      `xml:"DS_MERCHANT_CONSUMERLANGUAGE,omitempty" json:"Ds_Merchant_ConsumerLanguage,omitempty"`
	Descriptor         string        `xml:"DS_MERCHANT_MERCHANTDESCRIPTOR,omitempty" json:"DS_MERCHANT_MERCHANTDESCRIPTOR,omitempty"`
	MerchantGroup      string        `xml:"DS_MERCHANT_GROUP,omitempty" json:"DS_MERCHANT_GROUP,omitempty"`
	Reference          string        `xml:"DS_MERCHANT_IDENTIFIER,omitempty" json:"DS_MERCHANT_IDENTIFIER,omitempty"`
	PAN                string        `xml:"DS_MERCHANT_PAN,omitempty" json:"DS_MERCHANT_PAN,omitempty"`
	CVV                string        `xml:"DS_MERCHANT_CVV2,omitempty" json:"DS_MERCHANT_CVV2,omitempty"`
	ExpirationDate     string        `xml:"DS_MERCHANT_EXPIRYDATE,omitempty" json:"DS_MERCHANT_EXPIRYDATE,omitempty"`
	SecurePayment      string        `xml:"DS_MERCHANT_SECUREPAYMENT,omitempty" json:"DS_MERCHANT_SECUREPAYMENT,omitempty"`
	CallbackURL        string        `xml:"DS_MERCHANT_MERCHANTURL,omitempty" json:"DS_MERCHANT_MERCHANTURL,omitempty"`
	RedirectSuccessURL string        `xml:"DS_MERCHANT_URLOK,omitempty" json:"DS_MERCHANT_URLOK,omitempty"`
	RedirectFailureURL string        `xml:"DS_MERCHANT_URLKO,omitempty" json:"DS_MERCHANT_URLKO,omitempty"`
}

type Response struct {
	XMLName           xml.Name      `xml:"RETORNOXML" json:"-"`
	ErrorCode         string        `xml:"CODIGO" json:"Ds_ErrorCode"`
	Date              string        `xml:"OPERACION>Ds_Date" json:"Ds_Date"`
	Hour              string        `xml:"OPERACION>Ds_Hour" json:"Ds_Hour"`
	SecurePayment     string        `xml:"OPERACION>Ds_SecurePayment" json:"Ds_SecurePayment"`
	ExpirationDate    string        `xml:"OPERACION>Ds_ExpiryDate" json:"Ds_ExpiryDate"`
	Reference         string        `xml:"OPERACION>Ds_Merchant_Identifier" json:"Ds_Merchant_Identifier"`
	CardCountry       string        `xml:"OPERACION>Ds_Card_Country" json:"Ds_Card_Country"`
	Amount            int64         `xml:"OPERACION>Ds_Amount" json:"Ds_Amount"`
	Currency          Currency      `xml:"OPERACION>Ds_Currency" json:"Ds_Currency"`
	OperationID       string        `xml:"OPERACION>Ds_Order" json:"Ds_Order"`
	MerchantCode      string        `xml:"OPERACION>Ds_MerchantCode" json:"Ds_MerchantCode"`
	Terminal          int           `xml:"OPERACION>Ds_Terminal" json:"Ds_Terminal"`
	ResponseCode      string        `xml:"OPERACION>Ds_Response" json:"Ds_Response"`
	MerchantData      string        `xml:"OPERACION>Ds_MerchantData" json:"Ds_MerchantData"`
	OperationType     OperationType `xml:"OPERACION>Ds_TransactionType" json:"Ds_TransactionType"`
	Language          string        `xml:"OPERACION>Ds_ConsumerLanguage" json:"Ds_ConsumerLanguage"`
	AuthorizationCode string        `xml:"OPERACION>Ds_AuthorisationCode" json:"Ds_AuthorisationCode"`
	CardBrand         string        `xml:"OPERACION>Ds_Card_Brand" json:"Ds_Card_Brand"`
}
