package cli

import (
	"fmt"

	"skadi/backend/internal/app/user"
)

// UserController represents a controller for all auth routes.
type UserController struct {
	userUCManager user.UsecaseManager
}

// NewUserController returns a new instance of [UserController].
func NewUserController(userUCManager user.UsecaseManager) *UserController {
	return &UserController{
		userUCManager: userUCManager,
	}
}

// CreateAdmin creates a new admin user.
func (c *UserController) CreateAdmin() error {
	var (
		inputBody          = &userBody{}
		username, password string
	)
	// ask for username
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	if err := inputBody.ParseUsername(username); err != nil {
		return err
	}
	// ask for password
	fmt.Println(`Password rules:
	1) At least 8 letters.
	2) At least 1 number.
	3) At least 1 upper case.
	4) At least 1 special character.`)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)
	if err := inputBody.ParsePassword(password); err != nil {
		return err
	}

	// create a new admin
	userObj, err := c.userUCManager.CreateAdmin(inputBody.Username, inputBody.Password)
	if err != nil {
		return err
	}
	fmt.Printf("Admin user %q was created successfully!\n", userObj.Username)
	return nil
}

// DeleteAdmin deletes an admin user.
func (c *UserController) DeleteAdmin() error {
	var (
		inputBody = &userID{}
		id        string
	)
	// ask for ID
	fmt.Print("Enter admin ID: ")
	fmt.Scan(&id)
	if err := inputBody.ParseID(id); err != nil {
		return err
	}
	// delete admin user
	if err := c.userUCManager.DeleteAdminByID(inputBody.ID); err != nil {
		return err
	}
	fmt.Printf("Admin user %d was deleted successfully!\n", inputBody.ID)
	return nil
}

// DeleteAdmin deletes an admin user.
func (c *UserController) GetAdmins() error {
	// get all admins
	admins, err := c.userUCManager.GetAdmins()
	if err != nil {
		return err
	}
	// no one admin
	adminAmount := len(admins)
	if adminAmount == 0 {
		fmt.Println("No one admin user was found")
		return nil
	}
	// print out all admins
	fmt.Printf("Found: %d admins\n", adminAmount)
	for idx, admin := range admins {
		fmt.Printf("%d: ID - %d | Username - %s\n", idx+1, admin.ID, admin.Username)
	}
	return nil
}
