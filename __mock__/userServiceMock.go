package __mock__

import (
	model "github.com/T2-1c2023/NotificationsService/app/model"
)

type UserServiceMock struct{}

func NewUserServiceMock() UserServiceMock {
	return UserServiceMock{}
}

func (service *UserServiceMock) GetUserById(id int, userInfo string) (model.User, error) {
	return model.User{
		Id:            id,
		Fullname:      "John Doe",
		Mail:          "user@example.com",
		PhoneNumber:   "1234567890",
		IsAthlete:     true,
		IsTrainer:     false,
		IsVerified:    true,
		PhotoId:       "",
		ExpoPushToken: "ExpoPushToken1234",
	}, nil
}
