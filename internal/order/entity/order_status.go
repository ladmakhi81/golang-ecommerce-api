package orderentity

type OrderStatus string

const (
	OrderStatusPending     = "Pending"
	OrderStatusPayed       = "Payed"
	OrderStatusPreparation = "Preparation"
	OrderStatusDelivery    = "Delivery"
	OrderStatusDone        = "Done"
)

func IsValid(status OrderStatus) bool {
	validStatuses := []OrderStatus{
		OrderStatusDelivery,
		OrderStatusDone,
		OrderStatusPayed,
		OrderStatusPending,
		OrderStatusPreparation,
	}

	for _, validStatus := range validStatuses {
		if validStatus == status {
			return true
		}
	}

	return false
}
