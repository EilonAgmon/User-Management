package main

import (
	"errors"
	"github.com/rs/xid"
	"time"
	"strconv"
)

// A user object
type user struct {
	ID      string
	FirstName   string
	LastName string
	CreatedTime string
}

// Map of user id -> user index in list of quick retrieve of user
var userMap = make(map[string]int)

// Set of user names to prevent doubles
var existingNames = make(map[string]bool)

// User list
var userList = []user {}

// Creates 3 example users for the page not to be loaded empty
func createExampleUsers() {
	for i := 0; i < 3; i++ {
		createSingleUser("User" + strconv.Itoa(i + 1) + "_First", "User" + strconv.Itoa(i + 1) + "_Last")
	}
}

// Creates a single user and updates the global variables (maps and list)
func createSingleUser(firstName, lastName string) user {
	id := genXid();
	result := user{ID: id, FirstName: firstName, LastName: lastName, CreatedTime: time.Now().String()}
	userList = append(userList, result) // keeps order
	userMap[id] = len(userList) - 1 // add the index of this user to the map with the user id as key for quick retrieve
	existingNames[firstName] = true // add to existing names DB
	return result
}

// Returns a list of all the users
func getAllUsers() []user {

	if len(userList) == 0 {
		createExampleUsers()
	}

	return userList
}

// Fetches a user based on the ID given
func getUserByID(id string) (*user, error) {
	if value, ok := userMap[id]; ok {
		return &userList[value], nil
	} else {
		return nil, errors.New("User not found")
	}
}

// Create a new user with the first name and last name given
func createNewUser(firstName, lastName string) (*user, error) {
	// Check if the supplied user name is available
	if ok := existingNames[firstName]; ok {
		return nil, errors.New("Create cancelled: A user with the name < " + firstName + " > already exists. (In our world... first names must be unique!)")
	}
	newUser := createSingleUser(firstName, lastName)
	return &newUser, nil
}

// Updates an existing user with the firstName and lastName provided
func updateExistingUser(userID, firstName, lastName string) (*user, error) {
	// Check if the edited user name is available
	if ok := existingNames[firstName]; ok {
		return nil, errors.New("Update cancelled: A user with the name < " + firstName + " > already exists. (In our world... first names must be unique!)")
	}
	if user, err := getUserByID(userID); err == nil {
	
		delete(existingNames, user.FirstName) // Remove the previous name from the existin names map
		existingNames[firstName] = true // And add the new one
		
		// Update the user
		user.FirstName = firstName
		user.LastName = lastName
		
		return user, nil
	} else {
		return nil, errors.New("User to update not found")
	}
}

// Generates a unique ID for users
func genXid() string {
	return xid.New().String()
}