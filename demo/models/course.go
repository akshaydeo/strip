package models

// Course model
type Course struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Credits    int                    `json:"grades"`
	MinCredits int                    `json:"min_grades"`
	Rules      map[string]interface{} `json:"rules"` // this is just to add a complex struct :-|
}
