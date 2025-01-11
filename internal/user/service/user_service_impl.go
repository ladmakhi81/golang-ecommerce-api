package userservice

import (
	"net/http"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
	userrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo userrepository.IUserRepository
}

func NewUserService(userRepo userrepository.IUserRepository) UserService {
	return UserService{
		userRepo,
	}
}

func (userService UserService) CheckDuplicatedUserEmail(email string) error {
	isExist, existErr := userService.userRepo.IsEmailExist(email)
	if existErr != nil {
		return types.NewServerError(
			"error in checking email exist",
			"UserService.CheckDuplicatedUserEmail.UserRepo.IsEmailExist",
			existErr,
		)
	}
	if isExist {
		return types.NewClientError("email is already exist", http.StatusConflict)
	}
	return nil
}

func (userService UserService) CreateBasicUser(email, password string, role userentity.UserRole) (*userentity.User, error) {
	userDuplicatedErr := userService.CheckDuplicatedUserEmail(email)
	if userDuplicatedErr != nil {
		return nil, userDuplicatedErr
	}

	hashedPassword, hashedPasswordErr := userService.hashUserPassword(password)
	if hashedPasswordErr != nil {
		return nil, hashedPasswordErr
	}

	user := userentity.NewUser(email, hashedPassword, role)
	if createUserErr := userService.userRepo.CreateUser(user); createUserErr != nil {
		return nil, types.NewServerError(
			"error in creating user",
			"UserService.CreateBasicUser.CreateUser",
			createUserErr,
		)
	}

	return user, nil
}
func (userService UserService) FindUserByEmailAndPassword(email, password string) (*userentity.User, error) {
	user, findUserErr := userService.FindUserByEmail(email)
	if findUserErr != nil {
		return nil, findUserErr
	}
	if !userService.isValidPassword(user.Password, password) {
		return nil, types.NewClientError("user not found with this information", http.StatusNotFound)
	}
	// TODO: move this logic into separated middleware
	// if user.Role != userentity.CustomerRole && !user.IsVerified {
	// 	return nil, types.NewClientError("you must verify your account first", http.StatusForbidden)
	// }
	return user, nil
}
func (userService UserService) FindUserByEmail(email string) (*userentity.User, error) {
	user, userErr := userService.userRepo.FindBasicInfoByEmail(email)
	if userErr != nil {
		return nil, types.NewServerError(
			"error in finding user by email",
			"UserService.FindUserByEmailAddress",
			userErr,
		)
	}
	if user == nil {
		return nil, types.NewClientError(
			"user not found with this email address",
			http.StatusNotFound,
		)
	}
	return user, nil
}

func (userService UserService) isValidPassword(hashedPassword, password string) bool {
	if passwordErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); passwordErr != nil {
		return false
	}
	return true
}

func (userService UserService) hashUserPassword(password string) (string, error) {
	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		return "", types.NewServerError(
			"error in hashing password",
			"UserService.HashUserPassword.GenerateFromPassword",
			hashedPasswordErr,
		)
	}
	return string(hashedPassword), nil
}
