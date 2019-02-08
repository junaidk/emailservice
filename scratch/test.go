package scratch

import (
	"emailservice/types"
	"fmt"
	"html/template"
	"net/http"
)

func mainX() {
	tmpl := template.Must(template.ParseFiles("template/conf-go.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		shipment := types.SendOrderConfirmationRequest{
			Email: "email@gmail.com",
			Order: types.OrderResult{
				OrderID:            "900000021",
				ShippingTrackingID: "7999888891",
				ShippingCost: types.Money{
					CurrencyCode: "USD",
					Units:        99,
					Nanos:        100,
				},
				ShippingAddress: types.Address{
					StreetAddress: "248 GII",
					City:          "Lahore",
					State:         "Punjab",
					Country:       "Pakistan",
					ZipCode:       "54000",
				},
				Items: []types.OrderItem{
					{
						Item: types.CartItem{
							ProductID: "ABS3894723894",
							Quantity:  100,
						},
						Cost: types.Money{
							CurrencyCode: "USD",
							Units:        99,
							Nanos:        6600000000,
						},
					},
					{
						Item: types.CartItem{
							ProductID: "LPO3894723894",
							Quantity:  10,
						},
						Cost: types.Money{
							CurrencyCode: "USD",
							Units:        99,
							Nanos:        8700000000,
						},
					},
					{
						Item: types.CartItem{
							ProductID: "KKK3894723894",
							Quantity:  90,
						},
						Cost: types.Money{
							CurrencyCode: "USD",
							Units:        99,
							Nanos:        7500000000,
						},
					},
				},
			},
		}

		err := tmpl.Execute(w, shipment)
		fmt.Println(err)
	})

	http.ListenAndServe(":8089", nil)
}
