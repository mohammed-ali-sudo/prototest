package model

import "github.com/kamva/mgm/v3"

type Teacher struct {
	mgm.DefaultModel `bson:",inline"` // ✅ adds ID, created_at, updated_at
	FirstName        string           `json:"first_name" bson:"first_name"`
	LastName         string           `json:"last_name" bson:"last_name"`
	Email            string           `json:"email" bson:"email"`
	Class            string           `json:"class" bson:"class"`
	Subject          string           `json:"subject" bson:"subject"`
}
