package recipes

import(
    "fmt"
    "github.com/stripe/stripe-go/v71"
    "github.com/stripe/stripe-go/v71/customer"
)

func (c *Client)Test() {
    stripe.Key = c.skey
    fmt.Println(c.skey)
    params := &stripe.CustomerParams{
        Description: stripe.String("My First Test Customer (created for API docs)"),
    }
    customer, _ := customer.New(params)

    fmt.Println(customer)
}

