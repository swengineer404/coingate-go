package coingate

type CreateOrderParams struct {
	OrderID         string  `json:"order_id,omitempty"`
	PriceAmount     float64 `json:"price_amount,omitempty"`
	PriceCurrency   string  `json:"price_currency,omitempty"`
	ReceiveCurrency string  `json:"receive_currency,omitempty"`
	Title           string  `json:"title,omitempty"`
	Description     string  `json:"description,omitempty"`
	CallbackURL     string  `json:"callback_url,omitempty"`
	CancelURL       string  `json:"cancel_url,omitempty"`
	SuccessURL      string  `json:"success_url,omitempty"`
	Token           string  `json:"token,omitempty"`
	PurchaserEmail  string  `json:"purchaser_email,omitempty"`
}

type CreateOrder struct {
	ID         int    `json:"id"`
	PaymentURL string `json:"payment_url"`
}

type PaymentEvent struct {
	ID              int     `form:"id"`
	OrderID         string  `form:"order_id"`
	Status          string  `form:"status"`
	UnderpaidAmount float64 `form:"underpaid_amount"`
}
