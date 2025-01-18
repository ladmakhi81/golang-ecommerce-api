package orderdto

type CreateOrderReqBody struct {
	CartIds []uint `json:"carts" validate:"required"`
}

type CreateOrderResponse struct {
	OrderID uint   `json:"orderId"`
	PayLink string `json:"payLink"`
}

func NewCreateOrderResponse(orderID uint, payLink string) *CreateOrderResponse {
	return &CreateOrderResponse{
		OrderID: orderID,
		PayLink: payLink,
	}
}
