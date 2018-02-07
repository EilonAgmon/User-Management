package main

import (
	"github.com/gin-gonic/gin"
)

// Inits the app's routes
func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	// User Routes init:
	userRoutes := router.Group("/user")
	{
		// Handle GET requests at /user/view/some_user_id to show a specific user
		userRoutes.GET("/view/:user_id", getUser)
		
		// Handle POST requests at /user/edit/ to edit a user
		userRoutes.POST("/edit", doBeforeRoute(), editUser)
		
		// Handle the GET requests at /user/create and show the user creation page
		userRoutes.GET("/create", doBeforeRoute(), showUserCreationPage)

		// Handle POST requests at /user/create and actually create the user
		userRoutes.POST("/create", doBeforeRoute(), createUser)
	}
}

// Gets called before routing
func doBeforeRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Can be used to improve later on (validation, tempplates, etc')
	}
}