package main

func initializeRouts() {

	router.GET("/", indexHandler)
	router.GET("/register-successful", doRegisterHandler)
	router.POST("/", indexHandler)

	userRoutes := router.Group("/")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/enter", enterHandler)

		//// Handle POST requests at /u/login
		//// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/enter", doEnterHandler)

		userRoutes.POST("/add_friend", doAddFriendHandler)
		//
		//// Handle GET requests at /u/logout
		//// Ensure that the user is logged in by using the middleware
		//userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", registerHandler)

		//
		//// Handle POST requests at /u/register
		//// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", doRegisterHandler)

		userRoutes.GET("/logout", doLogOutHandler)

		userRoutes.GET("/chat/:friend_id", chatWithFriendHandler)

		userRoutes.POST("/chat/:friend_id", sendMessageToFriendHandler)
	}
}
