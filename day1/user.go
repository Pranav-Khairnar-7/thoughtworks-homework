package user

import (
	customError "pranav/error"
	util "pranav/util"
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base
	Name  string
	Age   int
	Email string
}

func NewUser(name string, age int, email string) (*User, error) {

	var user = User{
		Base: Base{
			Id:        uuid.NewV4(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  name,
		Age:   age,
		Email: email,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Validate() error {

	if !util.IsValidEmailString(u.Email) {
		return customError.NewValidationError("email doesn't contain @")
	}

	if util.IsStringEmpty(u.Name) {
		return customError.NewValidationError("Name is empty.")
	}

	if u.Age < 0 || u.Age > 150 {
		return customError.NewValidationError("Age is less than 0 or greater than 150.")
	}
	return nil
}

func (u *User) IsAdult() bool {
	return u.Age >= 18
}
