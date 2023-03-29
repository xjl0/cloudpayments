package cloudpayments

type BaseResponse struct {
	Success bool   `json:"Success,omitempty"`
	Message string `json:"Message,omitempty"`
}
type PaymentModel struct {
	TransactionId     int     `json:"TransactionId,omitempty"`
	Amount            float64 `json:"Amount,omitempty"`
	Currency          string  `json:"Currency,omitempty"`
	CurrencyCode      int     `json:"CurrencyCode,omitempty"`
	InvoiceID         string  `json:"InvoiceID,omitempty"`
	AccountID         string  `json:"AccountID,omitempty"`
	Email             string  `json:"Email,omitempty"`
	Description       string  `json:"Description,omitempty"`
	Data              []byte  `json:"Data,omitempty"`
	TestMode          bool    `json:"TestMode,omitempty"`
	IpAddress         string  `json:"IpAddress,omitempty"`
	IpCountry         string  `json:"IpCountry,omitempty"`
	IpCity            string  `json:"IpCity,omitempty"`
	IpRegion          string  `json:"IpRegion,omitempty"`
	IpDistrict        string  `json:"IpDistrict,omitempty"`
	IpLatitude        float64 `json:"IpLatitude,omitempty"`
	IpLongitude       float64 `json:"IpLongitude,omitempty"`
	CardFirstSix      string  `json:"CardFirstSix,omitempty"`
	CardLastFour      string  `json:"CardLastFour,omitempty"`
	CardExpiredMonth  int     `json:"CardExpiredMonth,omitempty"`
	CardExpiredYear   int     `json:"CardExpiredYear,omitempty"`
	CardType          string  `json:"CardType,omitempty"`
	CardTypeCode      int     `json:"CardTypeCode,omitempty"`
	Issuer            string  `json:"Issuer,omitempty"`
	IssuerBankCountry string  `json:"IssuerBankCountry,omitempty"`
	Status            string  `json:"Status,omitempty"`
	StatusCode        int     `json:"StatusCode,omitempty"`
	Reason            string  `json:"Reason,omitempty"`
	ReasonCode        int     `json:"ReasonCode,omitempty"`
	CardHolderMessage string  `json:"CardHolderMessage,omitempty"`
	CardHolderName    string  `json:"CardHolderName,omitempty"`
	Token             string  `json:"Token,omitempty"`
	JsonData          string  `json:"JsonData,omitempty"`
	Name              string  `json:"Name,omitempty"`
}

type SubscriptionModel struct {
	Id                           string
	CurrencyCode                 int
	StartDateIso                 string
	IntervalCode                 int
	StatusCode                   int
	Status                       string
	SuccessfulTransactionsNumber int
	FailedTransactionsNumber     int
	LastTransactionDate          string
	LastTransactionDateIso       string
	NextTransactionDate          string
	NextTransactionDateIso       string
}

type Payment3DSModel struct {
	PaReq             string `json:"PaReq,omitempty"`
	AcsUrl            string `json:"AcsUrl,omitempty"`
	ThreeDsCallbackId string `jsno:"ThreeDsCallbackId,omitempty"`
	AuthDate          string `json:"AuthDate,omitempty"`
	AuthDateIso       string `json:"AuthDateIso,omitempty"`
	AuthCode          string `json:"AuthCode,omitempty"`
	ConfirmDate       string `json:"ConfirmDate,omitempty"`
	ConfirmDateIso    string `json:"ConfirmDateIso,omitempty"`
}

type PaymentResponse struct {
	BaseResponse
	Model struct {
		PaymentModel
		Payment3DSModel
	}
}

type Payment3DSResponse struct {
	BaseResponse
	Model struct {
		PaymentModel
		Payment3DSModel
	}
}

type SubscriptionResponse struct {
	BaseResponse
	Model SubscriptionModel
}

type SubscriptionsListResponse struct {
	BaseResponse
	Model []SubscriptionModel
}
