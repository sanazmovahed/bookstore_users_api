package services

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/utils/crypto_utils"
	"bookstore_users_api/utils/date_utils"
	"bookstore_users_api/utils/errors"
	"fmt"
)

var (
	UsersService userServiceInterface = &usersService{}
)

type usersService struct {
}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (u *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {

	result := &users.User{Id: userId}
	fmt.Println()
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
	// if userId <= 0 {
	// 	return nil, errors.NewBadRequestError("invalid user id")
	// }
}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := u.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil

}

func (u *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (u *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}
