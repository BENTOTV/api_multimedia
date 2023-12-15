package lib

import (
	"implement_middleware/config"
	middlewares "implement_middleware/middleware"
	"implement_middleware/model"
)

// Lib func login user
func LoginUsers(user *model.Users) (interface{}, error) {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	Token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	userResponse := model.UsersResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: Token}

	return userResponse, nil
}
