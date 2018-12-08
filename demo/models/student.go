package models

// Student model
type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`

	Courses []Course  `json:"courses"`
	Friends []Student `json:"friends"`
}

func (s *Student) GetFriendsCount() int {
	return len(s.Friends)
}
