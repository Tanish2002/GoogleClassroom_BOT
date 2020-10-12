package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/classroom/v1"
)

const inline = false

var (
	BEE_A             string
	YOGA_A            string
	SEMI_CONDUCTORS_A string
	ENGLISH_A         string
	PROFFESIONAL_A    string
	EG_A              string
	MATHS_A           string
	CONSTITUITION_A   string
	TEST_A            string
)

func AssignmentCheck(s *discordgo.Session, sr *classroom.Service) {
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
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}
						for _, a := range l.CourseWork {
							if len(YOGA_A) == 0 {
								YOGA_A = a.AlternateLink
							}
							if YOGA_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for YOGA CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								YOGA_A = a.AlternateLink
							}
						}
					case "167939178573":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}
						for _, a := range l.CourseWork {
							if len(CONSTITUITION_A) == 0 {
								CONSTITUITION_A = a.AlternateLink
							}
							if CONSTITUITION_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for CONSTITUITION OF INDIA CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								CONSTITUITION_A = a.AlternateLink
							}
						}
					case "169440007943":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}
						for _, a := range l.CourseWork {
							if len(BEE_A) == 0 {
								BEE_A = a.AlternateLink
							}
							if BEE_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for BEE CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								BEE_A = a.AlternateLink
							}
						}
					case "56523552967":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}
						for _, a := range l.CourseWork {
							if len(PROFFESIONAL_A) == 0 {
								PROFFESIONAL_A = a.AlternateLink
							}
							if PROFFESIONAL_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for PROFESSIONAL SKILLS CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								PROFFESIONAL_A = a.AlternateLink
							}
						}
					case "167609056736":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}

						for _, a := range l.CourseWork {
							if len(ENGLISH_A) == 0 {
								ENGLISH_A = a.AlternateLink
							}
							if ENGLISH_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for ENGLISH CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								ENGLISH_A = a.AlternateLink
							}
						}
					case "145573038389":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}

						for _, a := range l.CourseWork {
							if len(EG_A) == 0 {
								EG_A = a.AlternateLink
							}
							if EG_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for ENGINEERING GRAPHICS CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								EG_A = a.AlternateLink
							}
						}
					case "161369689917":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}

						for _, a := range l.CourseWork {
							if len(SEMI_CONDUCTORS_A) == 0 {
								SEMI_CONDUCTORS_A = a.AlternateLink
							}
							if SEMI_CONDUCTORS_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for SEMI CONDUCTORS PHYSICS CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								SEMI_CONDUCTORS_A = a.AlternateLink
							}
						}
					case "161968063971":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}

						for _, a := range l.CourseWork {
							if len(MATHS_A) == 0 {
								MATHS_A = a.AlternateLink
							}
							if MATHS_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for MATHS CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								MATHS_A = a.AlternateLink
							}
						}
					case "150436902293":
						l, err := sr.Courses.CourseWork.List(c.Id).PageSize(1).Do()
						if err != nil {
							log.Fatalf("Unable to retrieve Assignments for %s. %v", c.Name, err)
						}
						for _, a := range l.CourseWork {
							if len(TEST_A) == 0 {
								TEST_A = a.AlternateLink
							}
							if TEST_A != a.AlternateLink {
								embed := discordgo.MessageEmbed{
									Title:     "New Assignment for TEST CLASS",
									URL:       a.AlternateLink,
									Fields:    []*discordgo.MessageEmbedField{},
									Timestamp: a.CreationTime,
									Footer: &discordgo.MessageEmbedFooter{
										Text: teacher.Profile.Name.FullName,
									},
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
								TEST_A = a.AlternateLink
							}
						}
					}
				}
			}
		} else {
			log.Fatalf("No courses found when checking for assigments")
			return
		}
		time.Sleep(10 * time.Second)
	}
}
