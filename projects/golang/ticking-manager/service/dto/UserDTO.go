package dto

import (
	"ticket-manager/model"
)

type UserDTO struct {
	Name  string  `json:"nickname,omitempty"`
	Password  string  `json:"password,omitempty"`
	Email     string  `json:"email,omitempty"`
}

func (dto *UserDTO) ToUser() (*model.User, error) {
	u := &model.User{
		Name: dto.Name,
		Email: dto.Email,
	}

	// if dto.Password != "" {
	// 	hashedPassword, err := passwd.HashPassword(dto.Password)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	u.PasswordHash = string(hashedPassword)
	// }

	return u, nil
}