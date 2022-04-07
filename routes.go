package main

func initializeRouts() {

	router.GET("/", indexHandler)
	router.GET("/register-successful", doRegisterHandler)
	router.POST("/", indexHandler)

	router.GET("/enter", enterHandler)

	router.POST("/enter", doEnterHandler)

	router.POST("/add_friend", doAddFriendHandler)

	router.GET("/register", registerHandler)

	router.POST("/register", doRegisterHandler)

	router.GET("/logout", doLogOutHandler)

	router.GET("/chat/:friend_id", chatWithFriendHandler)

	router.POST("/chat/:friend_id", sendMessageToFriendHandler)

}
