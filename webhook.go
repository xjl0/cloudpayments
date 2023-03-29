package cloudpayments

import "time"

type PayWebhookRequest struct {
	TransactionId                         int                    `json:"TransactionId"`
	Amount                                float64                `json:"Amount"`
	Currency                              string                 `json:"Currency"`
	PaymentAmount                         string                 `json:"PaymentAmount"`
	PaymentCurrency                       string                 `json:"PaymentCurrency"`
	DateTime                              time.Time              `json:"DateTime"`
	CardFirstSix                          string                 `json:"CardFirstSix"`
	CardLastFour                          string                 `json:"CardLastFour"`
	CardExpDate                           string                 `json:"CardExpDate"`
	TestMode                              int                    `json:"TestMode"`
	Status                                string                 `json:"Status"`
	OperationType                         string                 `json:"OperationType"`
	GatewayName                           string                 `json:"GatewayName"`
	InvoiceId                             string                 `json:"InvoiceId,omitempty"`
	AccountId                             string                 `json:"AccountId,omitempty"`
	SubscriptionId                        string                 `json:"SubscriptionId,omitempty"`
	Name                                  string                 `json:"Name,omitempty"`
	Email                                 string                 `json:"Email,omitempty"`
	IpAddress                             string                 `json:"IpAddress,omitempty"`
	IpCountry                             string                 `json:"IpCountry,omitempty"`
	IpCity                                string                 `json:"IpCity,omitempty"`
	IpRegion                              string                 `json:"IpRegion,omitempty"`
	IpDistrict                            string                 `json:"IpDistrict,omitempty"`
	IpLatitude                            float64                `json:"IpLatitude,omitempty"`
	IpLongitude                           float64                `json:"IpLongitude,omitempty"`
	Issuer                                string                 `json:"Issuer,omitempty"`
	IssuerBankCountry                     string                 `json:"IssuerBankCountry,omitempty"`
	Description                           string                 `json:"Description,omitempty"`
	AuthCode                              string                 `json:"AuthCode,omitempty"`
	Data                                  map[string]interface{} `json:"Data,omitempty"`
	Token                                 string                 `json:"Token,omitempty"`
	CardProduct                           string                 `json:"CardProduct,omitempty"`
	PaymentMethod                         string                 `json:"PaymentMethod,omitempty"`
	FallBackScenarioDeclinedTransactionId int                    `json:"FallBackScenarioDeclinedTransactionId,omitempty"`
}

type FailWebhookRequest struct {
	TransactionId                         int                    `json:"TransactionId"`
	Amount                                float64                `json:"Amount"`
	Currency                              string                 `json:"Currency"`
	PaymentAmount                         string                 `json:"PaymentAmount"`
	PaymentCurrency                       string                 `json:"PaymentCurrency"`
	DateTime                              time.Time              `json:"DateTime"`
	CardFirstSix                          string                 `json:"CardFirstSix"`
	CardLastFour                          string                 `json:"CardLastFour"`
	CardExpDate                           string                 `json:"CardExpDate"`
	TestMode                              int                    `json:"TestMode"`
	Reason                                string                 `json:"Reason"`
	ReasonCode                            int                    `json:"ReasonCode"`
	OperationType                         string                 `json:"OperationType"`
	InvoiceId                             string                 `json:"InvoiceId,omitempty"`
	AccountId                             string                 `json:"AccountId,omitempty"`
	SubscriptionId                        string                 `json:"SubscriptionId,omitempty"`
	Name                                  string                 `json:"Name,omitempty"`
	Email                                 string                 `json:"Email,omitempty"`
	IpAddress                             string                 `json:"IpAddress,omitempty"`
	IpCountry                             string                 `json:"IpCountry,omitempty"`
	IpCity                                string                 `json:"IpCity,omitempty"`
	IpRegion                              string                 `json:"IpRegion,omitempty"`
	IpDistrict                            string                 `json:"IpDistrict,omitempty"`
	IpLatitude                            float64                `json:"IpLatitude,omitempty"`
	IpLongitude                           float64                `json:"IpLongitude,omitempty"`
	Issuer                                string                 `json:"Issuer,omitempty"`
	IssuerBankCountry                     string                 `json:"IssuerBankCountry,omitempty"`
	Description                           string                 `json:"Description,omitempty"`
	Data                                  map[string]interface{} `json:"Data,omitempty"`
	Token                                 string                 `json:"Token,omitempty"`
	PaymentMethod                         string                 `json:"PaymentMethod,omitempty"`
	FallBackScenarioDeclinedTransactionId int                    `json:"FallBackScenarioDeclinedTransactionId,omitempty"`
}

type RefundWebhookRequest struct {
	TransactionId        int                    `json:"TransactionId"`
	PaymentTransactionId int                    `json:"PaymentTransactionId"`
	Amount               float64                `json:"Amount"`
	DateTime             time.Time              `json:"DateTime"`
	OperationType        string                 `json:"OperationType"`
	InvoiceId            string                 `json:"InvoiceId,omitempty"`
	AccountId            string                 `json:"AccountId,omitempty"`
	Email                string                 `json:"Email,omitempty"`
	Data                 map[string]interface{} `json:"Data,omitempty"`
}

type CancelWebhookRequest struct {
	TransactionId int                    `json:"TransactionId"`
	Amount        float64                `json:"Amount"`
	DateTime      time.Time              `json:"DateTime"`
	InvoiceId     string                 `json:"InvoiceId,omitempty"`
	AccountId     string                 `json:"AccountId,omitempty"`
	Email         string                 `json:"Email,omitempty"`
	Data          map[string]interface{} `json:"Data,omitempty"`
}

type WebhookResponse struct {
	Code int `json:"code"`
}
