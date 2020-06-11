package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `bson:"title,omitempty" json:"title,omitempty"`
	Date string `bson:"date,omitempty" json:"date,omitempty"`
	Time string `bson:"time,omitempty" json:"time,omitempty"`
	Location string `bson:"location,omitempty" json:"location,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Organizer string `bson:"organizer,omitempty" json:"organizer,omitempty"`
	Category string `bson:"category,omitempty" json:"category,omitempty"`
	Tags []string `bson:"tags,omitempty" json:"tags,omitempty"`
	Attendees []Attendee `bson:"attendees,omitempty" json:"attendees,omitempty"`
}

type Attendee struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fname string `bson:"fname,omitempty" json:"fname,omitempty"`
	Lname string `bson:"lname,omitempty" json:"lname,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}

var Events []Event
