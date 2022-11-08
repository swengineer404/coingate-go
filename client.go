package coingate

type Client struct {
	api *apiClient
}

func (c *Client) CreateOrder(params *CreateOrderParams) (res *CreateOrder, err error) {
	return res, c.api.send("/orders", "POST", params, res)
}
