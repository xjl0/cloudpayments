package cp_go

import (
	"os"
	"testing"
)

func TestClient_FindPaymentByInvoiceID(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		request FindPaymentRequest
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSuccess bool
		wantStatus  string
		wantErr     bool
	}{
		{
			name: "success 123",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args:        args{FindPaymentRequest{InvoiceId: "123"}},
			wantSuccess: true,
			wantStatus:  string(PaymentStatusCompleted),
			wantErr:     false,
		},
		{
			name: "not found 1234",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args:        args{FindPaymentRequest{InvoiceId: "1234"}},
			wantSuccess: false,
			wantStatus:  "",
			wantErr:     false,
		},
		{
			name: "success 321",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args:        args{FindPaymentRequest{InvoiceId: "321"}},
			wantSuccess: true,
			wantStatus:  string(PaymentStatusCompleted),
			wantErr:     false,
		},
		{
			name: "success 322",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args:        args{FindPaymentRequest{InvoiceId: "322"}},
			wantSuccess: false,
			wantStatus:  string(PaymentStatusAwaitingAuthentication),
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config: tt.fields.config,
			}

			got, err := c.FindPaymentByInvoiceID(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindPaymentByInvoiceID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// pretty.Printf("response: %# v", got)

			if got.Success != tt.wantSuccess {
				t.Errorf("Client.FindPaymentByInvoiceID() success = %v, want %v", got.Success, tt.wantSuccess)
			}
			if got.Model.Status != tt.wantStatus {
				t.Errorf("Client.FindPaymentByInvoiceID() status = %v, want %v", got.Model.Status, tt.wantStatus)
			}
		})
	}
}

func TestClient_ChargeCryptogramPayment(t *testing.T) {
	ref := func(str string) *string {
		return &str
	}

	type fields struct {
		config Config
	}
	type args struct {
		cpr *CryptogramPaymentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PaymentResponse
		wantErr bool
	}{
		{
			name: "321 without 3DS",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args: args{
				&CryptogramPaymentRequest{
					PaymentRequest: PaymentRequest{
						BaseRequest: BaseRequest{
							RequestID: ref("321"), // for idempotency
						},
						Amount:    123,
						Currency:  RUB,
						InvoiceId: "321",
						IpAddress: "178.155.4.86",
					},
					CardCryptogramPacket: "014111111111240104DMVzx7FS9qS/iMfGb3MNkqwWGHoEOIdoxFJ+6I51V5iYjqUSYnElKL1VIU++1OImD9+Kay1FKe+fSSpMQ+wajiq4pXSl64ny1JSu7mEYNhECt9Myw2wF9U64IBOB3JM5oOYBm2FzWZLe3wpjBeXDVXMM6ZFXEamaBDYdoUMBc1QNF/icqaFXPBsSjo0UCXxXgA+vYoCKClmV0TZL3KiiyWqbZ5rhmULrReWJ8EDiM3AXedkKEZVV78vNP6MfYUEQUBsG6d39L3o0MLn43PblnFxYx+VZLZtNabT27sCevgGoH+6e+rZpbqgyI5Qf1K86uU7PPfboaLRho9TLSZv6BQ==",
				},
			},
			wantErr: false,
		},
		{
			name: "322 with 3DS",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args: args{
				&CryptogramPaymentRequest{
					PaymentRequest: PaymentRequest{
						BaseRequest: BaseRequest{
							RequestID: ref("322"), // for idempotency
						},
						Amount:    322,
						Currency:  RUB,
						InvoiceId: "322",
						IpAddress: "178.155.4.86",
					},
					CardCryptogramPacket: "015555554444240104SiI9OTlm6lUTJI7RG+MsLGKTvhuD+PzZZnb1BM1s+LcV7mHwjTQW3mzOwqscmDnCJKERQ1HECcDB4Y/gvWHzySVh5X3+/R7/Y+eZmpWZCd4u0NdmqVd2i1aO8yy4QReE+lK54xDRNz1I5DZdNFbCZS0iEwppsYFEMqOaDR7hQfEv83MmTo9Wm7c72YYwgu326tO3nv9jkknqWjlpj4Kzq1p+OhEs5CzXR5AuB9CB4pt274F3bvpCDsitTi4n6kZ7a/2M8U6a+gsKFy2wawRIkhzKpcsXg5rrEK7aDa2Ay9yZRvdFau1+wanOnR6zqL1rEanEH/4vVW12bcyufV3dDQ==",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config: tt.fields.config,
			}
			got, err := c.ChargeCryptogramPayment(tt.args.cpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ChargeCryptogramPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// pretty.Printf("response: %# v", got)

			if !got.Success || (got.Message == "" && got.Model.PaReq != "") {
				t.Errorf("payment failed")
			}
		})
	}
}

func TestClient_Payment(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		request GetPaymentRequest
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantStatus  string
		wantSuccess bool
		wantErr     bool
	}{
		{
			name: "success 1074219493",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args:        args{GetPaymentRequest{TransactionId: 1074219493}},
			wantErr:     false,
			wantStatus:  string(PaymentStatusCompleted),
			wantSuccess: true,
		},
		{
			name: "failure 1075282147",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args:        args{GetPaymentRequest{TransactionId: 1075282147}},
			wantErr:     false,
			wantStatus:  string(PaymentStatusAwaitingAuthentication),
			wantSuccess: false,
		},
		{
			name: "not found 1074219494",
			fields: fields{
				Config{
					ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
					PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
				},
			},
			args:        args{GetPaymentRequest{TransactionId: 1074219494}},
			wantErr:     false,
			wantStatus:  "",
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config: tt.fields.config,
			}
			got, err := c.Payment(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Payment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// pretty.Printf("response: %# v", got)

			if got.Model.Status != string(tt.wantStatus) {
				t.Errorf("Client.Payment() = %v, want %v", got.Model.Status, tt.wantStatus)
			}

			if got.Success != tt.wantSuccess {
				t.Errorf("Client.Payment() success = %v, want %v", got.Success, tt.wantSuccess)
			}
		})
	}
}
