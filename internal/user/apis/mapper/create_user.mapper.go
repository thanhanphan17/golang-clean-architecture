package mapper

import (
	"go-clean-architecture/internal/user/apis/req"
	"go-clean-architecture/internal/user/business/entity"
)

// TransformCreateUserReq transforms a CreateUserReq into a User entity.
// The function takes a CreateUserReq struct as a parameter and returns a User struct.
func TransformCreateUserReq(req req.CreateUserReq) entity.User {
	// Create a new User struct with the required fields from the CreateUserReq struct.
	user := entity.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Role:     req.Role,
		Status:   entity.UNVERIFIED.Value(),
	}

	// Return the transformed User entity.
	return user
}

// TransformCreateUserReq transforms a CreateUserReq into a User entity.
// The function takes a CreateUserReq struct as a parameter and returns a User struct.
func TransformLogineUserReq(req req.LoginUserReq) entity.User {
	// Create a new User struct with the required fields from the CreateUserReq struct.
	user := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	// Return the transformed User entity.
	return user
}
