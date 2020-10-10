package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/classroom/v1"
)

func ClassroomStuff(s *discordgo.Session, sr *classroom.Service) {
	r, err := sr.Courses.List().PageSize(100).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}
	for {
		if len(r.Courses) > 0 {
			for _, c := range r.Courses {
				t, err := sr.Courses.Teachers.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Teacher Details. %s", err)
				}
				for _, teacher := range t.Teachers {
					switch c.Id {
					case "161832138169":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if YOGA != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for YOGA CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								YOGA = a.AlternateLink
							}
						}
					case "167939178573":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if CONSTITUITION != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for CONSTITUITION OF INDIA CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								CONSTITUITION = a.AlternateLink
							}
						}

					case "169440007943":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if BEE != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for BEE CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								BEE = a.AlternateLink
							}
						}

					case "56523552967":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if PROFFESIONAL != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for PROFESSIONAL CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								PROFFESIONAL = a.AlternateLink
							}
						}

					case "167609056736":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if ENGLISH != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for ENGLISH CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								ENGLISH = a.AlternateLink
							}
						}

					case "145573038389":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if EG != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for ENGINEERING GRAPHICS CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								EG = a.AlternateLink
							}
						}

					case "161369689917":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if SEMI_CONDUCTORS != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for SEMICONDUCTORS CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								SEMI_CONDUCTORS = a.AlternateLink
							}
						}

					case "161968063971":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if MATHS != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for MATHS CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								MATHS = a.AlternateLink
							}
						}

					case "150436902293":
						l, err := sr.Courses.Announcements.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
						}
						for _, a := range l.Announcements {
							if TEST != a.AlternateLink {
								s.ChannelMessageSendEmbed(CHANNEL, &discordgo.MessageEmbed{
									Title:       "New Announcement for TEST CLASS",
									URL:         a.AlternateLink,
									Description: a.Text,
									Timestamp:   a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
								})
								TEST = a.AlternateLink
							}
						}
					}
				}
			}
		} else {
			fmt.Print("No courses found.")
			return
		}
		time.Sleep(10 * time.Second)
	}
}
