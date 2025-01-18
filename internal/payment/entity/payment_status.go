package paymententity

type PaymentStatus string

const (
	PaymentStatusPending = "Pending"
	PaymentStatusSuccess = "Success"
	PaymentStatusFailed  = "Failed"
)

func IsValid(status PaymentStatus) bool {
	validStatuses := []PaymentStatus{
		PaymentStatusFailed,
		PaymentStatusPending,
		PaymentStatusSuccess,
	}

	for _, validStatus := range validStatuses {
		if validStatus == status {
			return true
		}
	}

	return false
}
