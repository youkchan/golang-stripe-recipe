package recipes

import(
    //"github.com/stripe/stripe-go/v71"
)

type Client struct{
    skey string
}

func NewClient(skey string)(*Client) {
    client := Client{
        skey: skey,
    }
    return &client
}

