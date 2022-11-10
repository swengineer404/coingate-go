package coingate

type Client struct {
	api *apiClient
}

func New(key string, sandbox bool) *Client {
	url := "https://api.coingate.com/v2"
	if sandbox {
		url = "https://api-sandbox.coingate.com/v2"
	}
	return &Client{
		api: newAPIClient(url, key),
	}
}

func (c *Client) CreateOrder(params *CreateOrderParams) (*CreateOrder, error) {
	var res CreateOrder
	return &res, c.api.send("/orders", "POST", params, &res)
}
