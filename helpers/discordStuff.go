package helpers

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func DiscordSession() *discordgo.Session {
	var (
		token = os.Getenv("TOKEN")
	)

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Failed on discordgo.New(): %s\n", err)
	}
	return discord
}

func DiscordBot(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		game := discordgo.Game{Type: discordgo.GameTypeWatching, Name: "for GCR Updates"}
		update := discordgo.UpdateStatusData{Game: &game}
		s.UpdateStatusComplex(update)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Failed on s.Open(): %s\n", err)
	}
	fmt.Println("Bot is now running.  Press CTRL-c to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	s.Close()
}
