package main

import (
	"crypto/md5"
	"fmt"
)

type User struct {
	ID             string
	Email          string
	HashedPassword string
	Username       string
}

func (user *User) AvatarURL() string {
	return fmt.Sprintf(
		"//www.gravatar.com/avatar/%x",
		md5.Sum([]byte(user.Email)),
	)
}

func (user *User) ImagesRoute() string {
	return "/user/" + user.ID
}

const (
	hashCost       = 10
	passwordLength = 6
	userIDLength   = 16
)

func NewUser(username, email, password string) (User, error) {
	user := User{
		Email:    email,
		Username: username,
	}
	if username == "" {
		return user, errNoUsername
	}

	if email == "" {
		return user, errNoEmail
	}

	if password == "" {
		return user, errNoPassword
	}

	if len(password) < passwordLength {
		return user, errPasswordTooShort
	}

	// Check if the username exists
	existingUser, err := globalUserStore.FindByUsername(username)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errUsernameExists
	}

	// Check if the email exists
	existingUser, err = globalUserStore.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errEmailExists
	}

	user.HashedPassword = password
	user.ID = GenerateID("usr", userIDLength)
	return user, err
}

func FindUser(username, password string) (*User, error) {

	existingUser := &User{
		Email:          "username",
		Username:       "username",
		HashedPassword: "hashedpassword",
		ID:             "1",
	}

	return existingUser, nil
}

func UpdateUser(user *User, email, currentPassword, newPassword string) (User, error) {
	out := *user
	out.Email = email

	// Check if the email exists
	existingUser, err := globalUserStore.FindByEmail(email)
	if err != nil {
		return out, err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return out, errEmailExists
	}

	// At this point, we can update the email address
	user.Email = email

	// No current password? Don't try update the password.
	if currentPassword == "" {
		return out, nil
	}

	if currentPassword == existingUser.HashedPassword {
		return out, errPasswordIncorrect
	}

	if newPassword == "" {
		return out, errNoPassword
	}

	if len(newPassword) < passwordLength {
		return out, errPasswordTooShort
	}

	user.HashedPassword = newPassword
	return out, err
}
