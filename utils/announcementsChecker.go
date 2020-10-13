package utils

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/classroom/v1"
)

var (
	CHANNEL string = os.Getenv("CHANNEL")
)

func AnnouncementCheck(s *discordgo.Session, sr *classroom.Service) {
	ids := strings.Split(os.Getenv("COURSE_ID"), ",")
	names := strings.Split(os.Getenv("COURSE_NAME"), ",")
	if len(ids) != len(names) {
		log.Fatalf("no of Course_ID != no of Course_Name")
		return
	}
	courses := make([]*classroom.Course, len(ids))
	for i, id := range ids {
		course := &classroom.Course{Id: id, Name: names[i]}
		courses[i] = course
	}
	for {
		for _, cs := range courses {
			cs := cs
			l, err := sr.Courses.Announcements.List(cs.Id).PageSize(1).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve Announcements for %s. %v", cs.Name, err)
			}
			for _, a := range l.Announcements {
				if len(cs.AlternateLink) == 0 {
					cs.AlternateLink = a.AlternateLink
				}
				if cs.AlternateLink != a.AlternateLink {
					s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
						Title:       "New Announcement for " + cs.Name + " CLASS",
						URL:         a.AlternateLink,
						Description: a.Text,
						Timestamp:   a.CreationTime,
					})
					cs.AlternateLink = a.AlternateLink
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}
