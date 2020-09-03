package recipes

import(
//    "fmt"
 //   "errors"
    "github.com/stripe/stripe-go/v71"
    "github.com/stripe/stripe-go/v71/product"
)

func (c *Client)CreateProduct(product_name string) string {
    stripe.Key = c.skey
    params := &stripe.ProductParams{
      Name: stripe.String(product_name),
    }
    p, _ := product.New(params)

    return p.ID
}


