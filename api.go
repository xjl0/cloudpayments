package cloudpayments

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "https://api.cloudpayments.ru/"

func (c *Client) sendRequest(endpoint string, params []byte, requestID *string) ([]byte, error) {
	req, err := http.NewRequest("POST", url+endpoint, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.config.PublicId, c.config.ApiSecret)
	req.Header.Set("Content-Type", "application/json")

	if requestID != nil {
		req.Header.Set("X-Request-ID", *requestID)
	}

	client := &http.Client{Timeout: c.config.Timeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resp.Close = true
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) Ping() (map[string]interface{}, error) {
	var data map[string]interface{}
	response, err := c.sendRequest("test", nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &data)
	return data, err
}

// ChargeCryptogramPayment Payment by cryptogram
func (c *Client) ChargeCryptogramPayment(cpr *CryptogramPaymentRequest) (*PaymentResponse, error) {
	paymentReponse := &PaymentResponse{}

	params, err := json.Marshal(cpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/cards/charge", params, cpr.RequestID)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &paymentReponse)
	if err != nil {
		return nil, err
	}

	return paymentReponse, nil
}

// AuthorizeCryptogramPayment two-stage payment
func (c *Client) AuthorizeCryptogramPayment(cpr CryptogramPaymentRequest) (*Payment3DSResponse, error) {
	paymentReponse := &Payment3DSResponse{}

	params, _ := json.Marshal(cpr)

	response, err := c.sendRequest("payments/cards/auth", params, cpr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &paymentReponse)
	if err != nil {
		return nil, err
	}

	return paymentReponse, nil
}

// ChargeTokenPayment payment by token
func (c *Client) ChargeTokenPayment(tpr TokenPaymentRequest) (*PaymentResponse, error) {
	paymentReponse := &PaymentResponse{}

	params, err := json.Marshal(tpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/tokens/charge", params, tpr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &paymentReponse)
	if err != nil {
		return nil, err
	}

	return paymentReponse, nil
}

// AuthorizeTokenPayment authorize token payment
func (c *Client) AuthorizeTokenPayment(tpr TokenPaymentRequest) (*PaymentResponse, error) {
	paymentReponse := &PaymentResponse{}

	params, err := json.Marshal(tpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/tokens/auth", params, tpr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &paymentReponse)
	if err != nil {
		return nil, nil
	}

	return paymentReponse, nil
}

// Confirm3DSPayment Confirm a 3DS payment
func (c *Client) Confirm3DSPayment(confirm3DS Confirm3DSRequest) (*PaymentResponse, error) {
	paymentReponse := &PaymentResponse{}

	params, err := json.Marshal(confirm3DS)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/cards/post3ds", params, confirm3DS.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &paymentReponse)
	if err != nil {
		return nil, nil
	}

	return paymentReponse, nil
}

// ConfirmPayment confirm an authorized payment
func (c *Client) ConfirmPayment(confirm ConfirmPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(confirm)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/confirm", params, confirm.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// RefundPayment refund
func (c *Client) RefundPayment(rpp RefundPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, _ := json.Marshal(rpp)
	response, err := c.sendRequest("payments/refund", params, rpp.RequestID)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)

	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// VoidPayment cancellation of payment
func (c *Client) VoidPayment(vpr VoidPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, _ := json.Marshal(vpr)
	response, err := c.sendRequest("payments/void", params, vpr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// Payment getting pyment by transaction id
func (c *Client) Payment(gpr GetPaymentRequest) (*PaymentResponse, error) {

	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(gpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/get", params, gpr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &paymentResponse)
	if err != nil {
		return nil, err
	}

	return paymentResponse, nil
}

// FindPaymentByInvoiceID
func (c *Client) FindPaymentByInvoiceID(fpr FindPaymentRequest) (*PaymentResponse, error) {

	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(fpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("v2/payments/find", params, fpr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &paymentResponse)
	if err != nil {
		return nil, err
	}

	return paymentResponse, nil
}

// PaymentsList
func (c *Client) PaymentsList() {
	// TODO implement
}

// CreateOrder to send by email
func (c *Client) CreateOrder(lpr LinkPaymentRequest) (*BaseResponse, error) {

	baseResponse := &BaseResponse{}

	params, err := json.Marshal(lpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("orders/create", params, lpr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// CreateSubscription create subscription
func (c *Client) CreateSubscription(scr SubscriptionCreateRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, err := json.Marshal(scr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/create", params, scr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &subscriptionResponse)

	if err != nil {
		return nil, err
	}

	return subscriptionResponse, nil
}

// UpdateSubscription update subscription
func (c *Client) UpdateSubscription(sur SubscriptionUpdateRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, err := json.Marshal(sur)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/update", params, sur.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &subscriptionResponse)
	if err != nil {
		return nil, err
	}

	return subscriptionResponse, nil
}

// CancelSubscription cancel subscription
func (c *Client) CancelSubscription(sur SubscriptionUpdateRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(sur)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/cancel", params, sur.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)

	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// GetSubscription get subscription by ID of transaction
func (c *Client) GetSubscription(sgr SubscriptionGetRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, err := json.Marshal(sgr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/get", params, sgr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &subscriptionResponse)
	if err != nil {
		return nil, err
	}

	return subscriptionResponse, nil
}

// GetSubscriptionsList get subscription list by AccountId
func (c *Client) GetSubscriptionsList(slr SubscriptionListRequest) (*SubscriptionsListResponse, error) {
	list := &SubscriptionsListResponse{}

	params, err := json.Marshal(slr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/find", params, slr.RequestID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c *Client) ChargeCryptogramPayout() {
	// TODO implement
}

func (c *Client) ChargeTokenPayout() {
	// TODO implement
}
