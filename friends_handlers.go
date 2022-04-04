package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func doAddFriendHandler(c *gin.Context) {
	if userSessionId > 0 {
		userFriend := c.PostForm("user_id")
		_, e := db.Exec(`INSERT INTO "Friends"(
			friend_one, friend_two, status)
		VALUES ($1, $2, $3)`, userSessionId, userFriend, 2)

		_, err := db.Exec(`INSERT INTO "Friends"(
			friend_one, friend_two, status)
		VALUES ($2, $1, $3)`, userSessionId, userFriend, 2)

		if e != nil && err != nil {
			fmt.Println(e.Error())
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func chatWithFriendHandler(c *gin.Context) {
	// find messages
	friendId, _ := c.Get("friend_id")
	chatId := c.PostForm("chat_id")
	c.Set("friend_id", friendId)
	c.Set("chat_id", chatId)

	render(c, gin.H{}, "chat.html")
}

func sendMessageToFriendHandler(c *gin.Context) {
	messageContent := c.PostForm("message_content")
	authorId := userSessionId
	chatId, _ := c.Get("chat_id")
	friendId, _ := c.Get("friend_id")

	db.Query(`INSERT INTO "Messages" (chat_id, content, author) VALUES ($1, $2, $3)`,
		chatId, messageContent, authorId)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/chat/%d", friendId))
}
