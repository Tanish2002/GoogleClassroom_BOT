package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/classroom/v1"
)

const inline = false

func AssignmentCheck(s *discordgo.Session, sr *classroom.Service) {
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
			l, err := sr.Courses.CourseWork.List(cs.Id).PageSize(1).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve Announcements for %s. %v", cs.Name, err)
			}
			for _, a := range l.CourseWork {
				if len(cs.AlternateLink) == 0 {
					cs.AlternateLink = a.AlternateLink
				}
				if cs.AlternateLink != a.AlternateLink {
					embed := discordgo.MessageEmbed{
						Title:     "New Assignment for " + cs.Name + " CLASS",
						URL:       a.AlternateLink,
						Fields:    []*discordgo.MessageEmbedField{},
						Timestamp: a.CreationTime,
					}
					if len(a.Title) != 0 {
						embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
							Name:   "Topic",
							Value:  a.Title,
							Inline: inline,
						})
					}
					if len(a.Description) != 0 {
						embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
							Name:   "Description",
							Value:  a.Description,
							Inline: inline,
						})
					}
					if a.DueDate != nil {
						embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
							Name:   "Due Date",
							Value:  fmt.Sprintf("%d/%d/%d", a.DueDate.Day, a.DueDate.Month, a.DueDate.Year),
							Inline: inline,
						})
					}
					s.ChannelMessageSendEmbed(CHANNEL, &embed)
					cs.AlternateLink = a.AlternateLink
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}
