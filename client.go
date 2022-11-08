package coingate

type Client struct {
	api *apiClient
}

func New(key string, sandbox bool) *Client {
	url := "https://api.coingate.com/v2"
	if sandbox {
		url = "https://sandbox-api.coingate.com/v2"
	}
	return &Client{
		api: newAPIClient(url, key),
	}
}

func (c *Client) CreateOrder(params *CreateOrderParams) (res *CreateOrder, err error) {
	return res, c.api.send("/orders", "POST", params, res)
}
