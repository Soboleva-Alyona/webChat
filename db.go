package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func connectAndGet() (e error, db *sql.DB) {
	db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", "5432", "postgres", "123456", "users"))
	if e != nil {
		return e, nil
	}
	return nil, db
}

func getUsersList() []UserData {
	rows, e := db.Query(`SELECT * FROM "Users"`)

	if e != nil {
		fmt.Println("can't get users")
		return nil
	} else {
		var users []UserData
		for rows.Next() {
			var id int
			var login string
			var password string
			err := rows.Scan(&id, &login, &password)
			if err != nil {
				return nil
			} else {
				if id != userSessionId {
					curUser := UserData{id, login, password, 0, 0}
					users = append(users, curUser)
				}
			}

		}
		friends := getUserFriendsList()
		for i := 0; i < len(users); i++ {
			for j := 0; j < len(friends); j++ {
				if users[i].Id == friends[j].Id {
					users[i].IsMyFriend = 1
				}
			}
		}
		return users
	}
}

type UserData struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	IsMyFriend   int    `json:"isMyFriend"`
	ChatWithMeId int    `json:"chatWithId"`
}

type UserFriendData struct {
	Id       int `json:"id"`
	UserId   int `json:"friend_one"`
	FriendId int `json:"friend_two"`
	Status   int `json:"status"`
}

func getUserFriendsList() []UserData {
	rows, e := db.Query(`SELECT id, friend_two, status FROM "Friends" WHERE "friend_one"=$1`, userSessionId)

	if e != nil {
		fmt.Println("can't get users")
		return nil
	} else {
		var userFriends []UserFriendData
		for rows.Next() {
			var id int
			var friendTwo int
			var status int
			err := rows.Scan(&id, &friendTwo, &status)
			if err != nil {
				return nil
			} else {
				curFriend := UserFriendData{id, userSessionId, friendTwo, status}
				if status == 2 {
					userFriends = append(userFriends, curFriend)
				}
			}

		}
		var userFriendsAsUsers []UserData
		for i := 0; i < len(userFriends); i++ {
			rows, e := db.Query(`SELECT * FROM "Users" WHERE "id"=$1`, userFriends[i].FriendId)
			if e != nil {
				fmt.Println(e.Error())
			} else {
				for rows.Next() {
					var id int
					var login string
					var password string
					err := rows.Scan(&id, &login, &password)
					if err != nil {
						fmt.Println(e.Error())
					}
					userFriendsAsUsers = append(userFriendsAsUsers, UserData{id, login, password, 0, userFriends[i].Id})
				}
			}
		}
		return userFriendsAsUsers
	}
}
