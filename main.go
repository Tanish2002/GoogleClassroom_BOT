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
	go utils.AnnouncementCheck(discord, srv)
	go utils.AssignmentCheck(discord, srv)

	// Start Bot
	helpers.DiscordBot(discord)
}
