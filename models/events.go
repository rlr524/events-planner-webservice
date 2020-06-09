package models

import "go.mongodb.org/mongo-driver/bson"

type Event struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Date string `bson:"date" json:"date"`
	Time string `bson:"time" json:"time"`
	Location string `bson:"location" json:"location"`
	Description string `bson:"description" json:"description"`
	Organizer string `bson:"organizer" json:"organizer"`
	Category string `bson:"category" json:"category"`
}

type Attendee struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Fname string `bson:"fname" json:"fname"`
	Lname string `bson:"lname" json:"lname"`
	Email string `bson:"email" json:"email"`
}

var Events []Event
