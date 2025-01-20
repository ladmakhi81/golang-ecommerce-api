package userrepository

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type UserAddressRepository struct {
	storage *storage.Storage
}

func NewUserAddressRepository(
	storage *storage.Storage,
) UserAddressRepository {
	return UserAddressRepository{
		storage: storage,
	}
}

func (userAddressRepo UserAddressRepository) CreateUserAddress(userAddress *userentity.UserAddress) error {
	command := `
		INSERT INTO _user_addresses 
		(city, province, address, license_plate, description, user_id) 
		VALUES 
		($1, $2, $3, $4, $5, $6)
		RETURNING Id, created_at, updated_at;
	`
	row := userAddressRepo.storage.DB.QueryRow(
		command,
		userAddress.City,
		userAddress.Province,
		userAddress.Address,
		userAddress.LicensePlate,
		userAddress.Description,
		userAddress.User.ID,
	)
	scanErr := row.Scan(
		&userAddress.ID,
		&userAddress.CreatedAt,
		&userAddress.UpdatedAt,
	)
	return scanErr
}
func (userAddressRepo UserAddressRepository) GetUserAddresses(userId uint) ([]*userentity.UserAddress, error) {
	command := `
		SELECT 
		id, created_at, updated_at, city, province, address, license_plate, description
		FROM _user_addresses 
		WHERE user_id = $1
	`
	rows, rowsErr := userAddressRepo.storage.DB.Query(
		command,
		userId,
	)
	if rowsErr != nil {
		return nil, rowsErr
	}
	userAddresses := []*userentity.UserAddress{}
	defer rows.Close()
	for rows.Next() {
		userAddress := new(userentity.UserAddress)
		scanErr := rows.Scan(
			&userAddress.ID,
			&userAddress.CreatedAt,
			&userAddress.UpdatedAt,
			&userAddress.City,
			&userAddress.Province,
			&userAddress.Address,
			&userAddress.LicensePlate,
			&userAddress.Description,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		userAddresses = append(userAddresses, userAddress)
	}
	return userAddresses, nil
}
