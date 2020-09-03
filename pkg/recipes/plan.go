package recipes

import(
//    "fmt"
    "errors"
    "github.com/stripe/stripe-go/v71"
    "github.com/stripe/stripe-go/v71/plan"
)

func (c *Client)CreatePlan(amount int64, plan_name string, interval_num int, product_id string) string {
    stripe.Key = c.skey

    interval, interval_count, _ := getInterval(interval_num)
	/*if err != nil {
		return nil, errors.New("Invalid Private Key")
	}*/

    params := &stripe.PlanParams{
      Amount: stripe.Int64(amount),
      Currency: stripe.String(string(stripe.CurrencyJPY)),
      Interval: stripe.String(interval),
      IntervalCount: stripe.Int64(interval_count),
      ProductID: stripe.String(product_id),
    }
    p, _ := plan.New(params)

    return p.ID
}

func getInterval(interval_num int) (string, int64, error) {

    var interval string
    var interval_count int64
    switch(interval_num) {
    case 0 :
        interval = "month"
        interval_count = 1
    case 1 :
        interval = "month"
        interval_count = 3
    case 2 :
        interval = "month"
        interval_count = 6
    default:
        return "", 0, errors.New("Invalid interval_num")
    }
    return interval, interval_count, nil
}

