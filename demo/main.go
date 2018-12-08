package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/akshaydeo/strip/demo/models"
	"github.com/satori/go.uuid"
	"time"
)

// Function to log the interface passed to it
// I am using logrus as an example, with different levels
func logIt(i interface{}) {
	logrus.Debug(i)
}

// Method to get student
// In real world, this might be either a db call or a network call
func getStudent() *models.Student {
	s := new(models.Student)
	s.Name = "akshay deo"
	s.ID = uuid.NewV4().String()
	s.Address = "askjdnalksjdbnaljsnbdkajsdjkasbljd  askjdbajlsdbjhlasbdjlhasblasjd"

	for i := 0; i < 10; i++ {
		s.Courses = append(s.Courses, models.Course{})
		s.Courses[i].ID = uuid.NewV4().String()
		s.Courses[i].Name = "Test"
		s.Courses[i].Credits = 4
		s.Courses[i].MinCredits = 3
		s.Courses[i].Rules = map[string]interface{}{}
	}

	for i := 0; i < 10; i++ {
		s := new(models.Student)
		s.Name = fmt.Sprintf("%d", i)
		s.ID = uuid.NewV4().String()
		s.Address = "askjdnalksjdbnaljsnbdkajsdjkasbljd  askjdbajlsdbjhlasbdjlhasblasjd"
		for i := 0; i < 10; i++ {
			s.Courses = append(s.Courses, models.Course{})
			s.Courses[i].ID = uuid.NewV4().String()
			s.Courses[i].Name = "Test"
			s.Courses[i].Credits = 4
			s.Courses[i].MinCredits = 3
			s.Courses[i].Rules = map[string]interface{}{}
		}
	}
	return s
}

func main() {
	// Setting up logrus level
	logrus.SetLevel(logrus.DebugLevel)
	t := time.Now()
	for i := 0; i < 10000; i++ {
		s := getStudent()
		// because we want to print the object we need
		logIt(*s)
		s.GetFriendsCount()
	}
	fmt.Printf("%d\n", time.Now().Sub(t).Nanoseconds()/(1000*1000))
}
