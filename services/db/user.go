package dbService

import (
	"fmt"
	"gRMS/modals"
	"strings"
)

// CreateUser function to create a new user entry
func (sr *dbs) CreateUser(firstName, lastName, username, email, password string) (*modals.User, error) {
	if len(firstName) < 2 {
		return nil, fmt.Errorf("invalid name provided")
	}
	if len(email) < 4 && !strings.Contains(email, "@") {
		return nil, fmt.Errorf("invalid email provided")
	}
	if len(username) < 4 {
		return nil, fmt.Errorf("length of username cannot be less than 4")
	}
	if len(password) < 5 {
		return nil, fmt.Errorf("length of password cannot be less than 5")
	}

	user := modals.User{}
	sr.db.First(&user, "email = ?", email)
	if user.GetEmail() != "" {
		return nil, fmt.Errorf("email already exists")
	}
	sr.db.First(&user, "username = ?", username)
	if user.GetUserName() != "" {
		return nil, fmt.Errorf("username already exists")
	}

	user = modals.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
		Chats:     make([]modals.Participant, 0),
	}

	sr.db.Create(&user)

	return &user, nil
}

// GetUser is used to find user by id
func (sr *dbs) GetUser(userID uint64) (*modals.User, error) {
	user := modals.User{}

	sr.db.Preload("Chats").First(&user, "id = ?", userID)
	if user.ID == 0 {
		return nil, fmt.Errorf("invalid user id: %v", userID)
	}

	return &user, nil
}

func (sr *dbs) FindUser(username string) (*modals.User, error) {
	user := modals.User{}

	sr.db.Preload("Chats").First(&user, "username = ?", username)
	if user.ID == 0 {
		return nil, fmt.Errorf("invalid username")
	}

	return &user, nil
}

func (sr *dbs) UpdateUser(user *modals.User) error {
	// TODO implement this function
	return nil
}

func (sr *dbs) DeleteUser(userID uint64) error {
	// TODO implement this function
	return nil
}