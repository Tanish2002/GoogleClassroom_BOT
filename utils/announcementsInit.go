package utils

import (
	"fmt"
	"log"

	"google.golang.org/api/classroom/v1"
)

var (
	BEE             string
	YOGA            string
	SEMI_CONDUCTORS string
	ENGLISH         string
	PROFFESIONAL    string
	EG              string
	MATHS           string
	CONSTITUITION   string
	TEST            string
	CHANNEL         string = "749951930062733412"
)

func ClassroomInit(s *classroom.Service) {
	r, err := s.Courses.List().PageSize(100).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}
	if len(r.Courses) > 0 {
		for _, c := range r.Courses {
			switch c.Id {
			case "150436902293":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					TEST = a.AlternateLink
				}
			case "161832138169":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					YOGA = a.AlternateLink
				}
			case "167939178573":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					CONSTITUITION = a.AlternateLink
				}
			case "169440007943":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					BEE = a.AlternateLink
				}

			case "56523552967":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					PROFFESIONAL = a.AlternateLink
				}

			case "167609056736":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					ENGLISH = a.AlternateLink
				}

			case "145573038389":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					EG = a.AlternateLink
				}

			case "161369689917":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					SEMI_CONDUCTORS = a.AlternateLink
				}

			case "161968063971":
				l, err := s.Courses.Announcements.List(c.Id).PageSize(1).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve Announcements for %s. %v", c.Name, err)
				}
				for _, a := range l.Announcements {
					MATHS = a.AlternateLink
				}
			}
		}
	} else {
		fmt.Print("No courses found.")
		return
	}
}
