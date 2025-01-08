package user_entity

type UserRole string

const (
	CustomerRole = "Customer"
	AdminRole    = "Admin"
	VendorRole   = "Vendor"
)

func (UserRole) IsValid(role UserRole) bool {
	validRoles := []UserRole{
		AdminRole,
		CustomerRole,
		VendorRole,
	}

	for _, validRole := range validRoles {
		if validRole == role {
			return true
		}
	}

	return false
}
