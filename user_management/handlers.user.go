package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Shows the main page with all the users
func showIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page", "payload": getAllUsers()})
}

// Shows the User creation page
func showUserCreationPage(c *gin.Context) {
	c.HTML(http.StatusOK, "create-user.html", gin.H{"title": "Create New User"})
}

// Shows the error page with the given message
func showErrorPage(c *gin.Context, message string) {
	c.HTML(http.StatusOK, "error.html", gin.H{"title": "Error Page", "message": message})
}

// Show the user edit page
func getUser(c *gin.Context) {
	// Check if the user ID exists in parameters
	if userID := c.Param("user_id"); userID != "" {
		// Check if the user exists
		if user, err := getUserByID(userID); err == nil {
			c.HTML(http.StatusOK, "user.html", gin.H{"title": "View or edit a user", "payload": user})
		} else {
			showErrorPage(c, err.Error())
		}
	} else {
		showErrorPage(c, "No ID not provided in request or invalid ID")
	}
}

// Edits the given user according to post form params
func editUser(c *gin.Context) {
	if userID := c.PostForm("userid"); userID != "" {
		if user, err := updateExistingUser(userID, c.PostForm("firstname"), c.PostForm("lastname")); err == nil {
			// If the user is edited successfully, show success page
			c.HTML(http.StatusOK, "edit-successful.html", gin.H{"title": "Edit User Successful", "payload": user})
		} else {
			showErrorPage(c, err.Error())
		}
	} else {
		showErrorPage(c, "Could not find user to edit")
	}
}

// Creates a new user with the given post form params
func createUser(c *gin.Context) {
	if user, err := createNewUser(c.PostForm("firstname"), c.PostForm("lastname")); err == nil {
		// If the user is created successfully, show success page
		c.HTML(http.StatusOK, "creation-successful.html", gin.H{"title": "Creation Successful", "payload": user})
	} else {
		showErrorPage(c, err.Error())
	}
}
