package userservice

import (
	"net/http"
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	userdto "github.com/ladmakhi81/golang-ecommerce-api/internal/user/dto"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
	userrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/user/repository"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo        userrepository.IUserRepository
	translation     translations.ITranslation
	eventsContainer *events.EventsContainer
}

func NewUserService(
	userRepo userrepository.IUserRepository,
	translation translations.ITranslation,
	eventsContainer *events.EventsContainer,
) UserService {
	return UserService{
		userRepo:        userRepo,
		translation:     translation,
		eventsContainer: eventsContainer,
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
		return types.NewClientError(
			translations.NewTranslation().Message("user.email_duplicate_error"),
			http.StatusConflict,
		)
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
		return nil, types.NewClientError(
			userService.translation.Message("user.not_found"),
			http.StatusNotFound,
		)
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
			userService.translation.Message("user.not_found_email"),
			http.StatusNotFound,
		)
	}
	return user, nil
}
func (userService UserService) FindBasicUserInfoById(id uint) (*userentity.User, error) {
	user, findUserErr := userService.userRepo.FindBasicUserInfoById(id)
	if findUserErr != nil {
		return nil, types.NewServerError(
			"error in finding user by email",
			"UserService.FindBasicUserInfoById",
			findUserErr,
		)
	}
	if user == nil {
		return nil, types.NewClientError(
			userService.translation.Message("user.not_found_id"),
			http.StatusNotFound,
		)
	}
	return user, nil
}
func (userService UserService) CompleteProfile(userId uint, data *userdto.CompleteProfileReqBody) (*userentity.User, error) {
	user, findUserErr := userService.FindBasicUserInfoById(userId)
	if findUserErr != nil {
		return nil, findUserErr
	}
	user.Address = data.Address
	user.NationalID = data.NationalID
	user.PostalCode = data.PostalCode
	user.FullName = data.FullName
	user.IsCompleteProfile = true
	user.CompleteProfileAt = time.Now()
	if updateErr := userService.userRepo.CompleteProfile(user); updateErr != nil {
		return nil, types.NewServerError(
			"error in update user",
			"UserService.CompleteProfile.UpdateUser",
			updateErr,
		)
	}
	userService.eventsContainer.PublishEvent(
		events.NewEvent(
			events.USER_COMPLETE_PROFILE_EVENT,
			events.NewUserCompleteProfileEventBody(user.Email),
		),
	)
	return user, nil
}
func (userService UserService) VerifyAccountByAdmin(adminId uint, vendorId uint) error {
	admin, findAdminErr := userService.FindBasicUserInfoById(adminId)
	if findAdminErr != nil {
		return findAdminErr
	}
	vendor, findVendorErr := userService.FindBasicUserInfoById(vendorId)
	if findVendorErr != nil {
		return findVendorErr
	}
	if !vendor.IsCompleteProfile {
		return types.NewClientError(
			translations.NewTranslation().Message("complete_profile_error"),
			http.StatusBadRequest,
		)
	}

	verificationErr := userService.userRepo.UpdateVerificationState(admin.ID, vendor.ID)

	if verificationErr != nil {
		return types.NewServerError(
			"error in verifying user",
			"UserService.VerifyAccountByAdmin",
			verificationErr,
		)
	}
	userService.eventsContainer.PublishEvent(
		events.NewEvent(
			events.USER_VERIFIED_EVENT,
			events.NewUserVerificationEventBody(
				admin.Email,
				vendor.Email,
				vendor.FullName,
				time.Now(),
			),
		),
	)
	return nil
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
func (userService UserService) SetActiveUserAddress(userId uint, addressId uint) error {
	err := userService.userRepo.SetActiveUserAddress(userId, addressId)
	if err != nil {
		return types.NewServerError(
			"error in setting active address for user",
			"UserService.SetActiveUserAddress",
			err,
		)
	}
	return nil
}
