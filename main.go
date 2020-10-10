package main

import (
	"classroom/helpers"
	"classroom/utils"
)

func main() {
	//Discord API stuff
	discord := helpers.DiscordSession()

	//Classroom API stuff
	srv := helpers.Server()
	utils.ClassroomInit(srv)
	go utils.ClassroomStuff(discord, srv)

	// init Bot
	helpers.DiscordBot(discord)
}
