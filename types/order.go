package types

import "fmt"

type SendOrderConfirmationRequest struct {
	Email string      `json:"email"`
	Order OrderResult `json:"order"`
}

type OrderResult struct {
	OrderID            string      `json:"order_id"`
	ShippingTrackingID string      `json:"shipping_tracking_id"`
	ShippingCost       Money       `json:"shipping_cost"`
	ShippingAddress    Address     `json:"shipping_address"`
	Items              []OrderItem `json:"items"`
}

type OrderItem struct {
	Item CartItem `json:"item"`
	Cost Money    `json:"cost"`
}

type Money struct {
	CurrencyCode string `json:"currency_code"`
	Units        int    `json:"units"`
	Nanos        int    `json:"nanos"`
}

func (m Money) Format(in int) string {
	out := in / 10000000
	return fmt.Sprintf("%02d", out) // in / 1000
}

type Address struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
}
type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
