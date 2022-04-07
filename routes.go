package main

func initializeRouts() {

	router.GET("/", indexHandler)
	router.GET("/register-successful", doRegisterHandler)
	router.POST("/", indexHandler)

	userRoutes := router.Group("/")
	{
		userRoutes.GET("/enter", enterHandler)

		userRoutes.POST("/enter", doEnterHandler)

		userRoutes.POST("/add_friend", doAddFriendHandler)

		userRoutes.GET("/register", registerHandler)

		userRoutes.POST("/register", doRegisterHandler)

		userRoutes.GET("/logout", doLogOutHandler)

		userRoutes.GET("/chat/:friend_id", chatWithFriendHandler)

		userRoutes.POST("/chat/:friend_id", sendMessageToFriendHandler)
	}
}
