package userrepository

import (
	"database/sql"
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type UserRepository struct {
	storage *storage.Storage
}

func NewUserRepository(storage *storage.Storage) UserRepository {
	return UserRepository{storage: storage}
}

func (userRepo UserRepository) CreateUser(user *userentity.User) error {
	command := `
		INSERT INTO _users (email, password, user_role) VALUES ($1, $2, $3)
		RETURNING id;
	`
	row := userRepo.storage.DB.QueryRow(command, user.Email, user.Password, user.Role)
	if scanErr := row.Scan(&user.ID); scanErr != nil {
		return scanErr
	}
	return nil
}
func (userRepo UserRepository) IsEmailExist(email string) (bool, error) {
	command := `
		SELECT id FROM _users WHERE email = $1
	`
	row := userRepo.storage.DB.QueryRow(command, email)
	var id uint
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (userRepo UserRepository) FindBasicInfoByEmail(email string) (*userentity.User, error) {
	command := `
		SELECT id, password, user_role FROM _users WHERE email = $1
	`
	user := new(userentity.User)
	user.Email = email
	row := userRepo.storage.DB.QueryRow(command, email)
	scanErr := row.Scan(
		&user.ID,
		&user.Password,
		&user.Role,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	return user, nil
}
func (userRepo UserRepository) CompleteProfile(user *userentity.User) error {
	command := `
		UPDATE _users SET
		is_complete_profile = TRUE,
		address = $1,
		postal_code = $2,
		national_id = $3,
		full_name = $4,
		complete_profile_at = $5
		WHERE id = $6;
	`
	row := userRepo.storage.DB.QueryRow(
		command,
		user.Address,
		user.PostalCode,
		user.NationalID,
		user.FullName,
		time.Now(),
		user.ID,
	)
	return row.Err()
}
func (userRepo UserRepository) FindBasicUserInfoById(id uint) (*userentity.User, error) {
	command := `
		SELECT email, user_role, is_complete_profile, is_verified FROM _users WHERE id = $1
	`
	row := userRepo.storage.DB.QueryRow(command, id)
	user := new(userentity.User)
	scanErr := row.Scan(
		&user.Email,
		&user.Role,
		&user.IsCompleteProfile,
		&user.IsVerified,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	user.ID = id
	return user, nil
}
func (userRepo UserRepository) UpdateVerificationState(adminId uint, vendorId uint) error {
	command := `
		UPDATE _users SET
		is_verified = TRUE,
		verified_by_id = $1,
		verified_date = $2
		WHERE id = $3
	`
	row := userRepo.storage.DB.QueryRow(
		command,
		adminId,
		time.Now(),
		vendorId,
	)
	return row.Err()
}
