package main

import "fmt"

type user struct {
	nickname string
	username string
}

type message struct {
	prefix     string
	command    string
	parameters []string
}

const (
	ERR_NONICKNAMEGIVEN  = ":No nickname given"          //431
	ERR_ERRONEUSNICKNAME = "<nick> :Erroneous nickname"  //432
	ERR_NICKNAMEINUSE    = ":Nickname is already in use" //433
	RPL_NICKCHANGE       = "Nickname change was successful"
	RPL_WELCOME          = "Welcome to the Internet Relay Network<nick>!<user>@<host>" //001
)

func processNickCommand(message message) string {

	users := []user{
		{nickname: "Jon", username: "Johnathan"},
		{nickname: "Ben", username: "Benjamin"},
	}

	nick := message.parameters[0]
	for _, u := range users {
		if nick == u.nickname {
			return fmt.Sprintf("%s %s", nick, ERR_NICKNAMEINUSE)
		} else {
			newUser := user{nickname: nick, username: ""}
			users = append(users, newUser)
		}
	}
	return fmt.Sprintf("%s %s", RPL_WELCOME, nick)
}
