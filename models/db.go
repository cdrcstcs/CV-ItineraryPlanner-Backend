package models
import (
	"time"
)
type Tag struct {
	Id    string `bson:"_id,omitempty" json:"id"`
	Value string `bson:"value" json:"value"` 
}
type Attraction struct {
	Id           string   `bson:"_id,omitempty" json:"id"`
	Name         string   `bson:"name" json:"name"`
	Address      string   `bson:"address" json:"address"`
	X  int    `bson:"x" json:"x"`
	Y  int    `bson:"y" json:"y"`	
	TagIDs       []string `bson:"tag_ids" json:"tag_ids"`
	RatingId     string   `bson:"rating_id" json:"rating_id"`
	City         string   `bson:"city" json:"city"`
}
type User struct {
	Id            string    `bson:"_id,omitempty" json:"id" validate:"required"`
	Name          string    `bson:"name" json:"name" validate:"required"`
	Password      string    `bson:"password" json:"password" validate:"required"`
	Email         string    `bson:"email" json:"email" validate:"required"`
	EmailPassword string    `bson:"email_password" json:"email_password" validate:"required"`
	Phone         string    `bson:"phone" json:"phone" validate:"required"`
	Token         string    `bson:"token" json:"token" validate:"required"`
	User_type     string    `bson:"user_type" json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token string    `bson:"refresh_token" json:"refresh_token" validate:"required"`
	Created_at    time.Time `bson:"created_at" json:"created_at" validate:"required"`
	Updated_at    time.Time `bson:"updated_at" json:"updated_at" validate:"required"`
	OTPToken  	  string    `bson:"otpToken" json:"otpToken" validate:"required"`
	SID			  string    `bson:"SID" json:"SID" validate:"required"`
}
type Rating struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Score    int    `bson:"score" json:"score"`
}
type Event struct {
	Id           string    `bson:"_id,omitempty" json:"id"`
	StartTime    time.Time `bson:"start_time" json:"start_time"`
	EndTime      time.Time `bson:"end_time" json:"end_time"`
	AttractionId string    `bson:"attraction_id" json:"attraction_id"`
	AttractionName string   `bson:"attraction_name" json:"attraction_name"`
	Description  string    `bson:"description" json:"description"`
}
type Itinerary struct {
	Id        string    `bson:"_id,omitempty" json:"id"`
	CopiedId  string    `bson:"copied_id" json:"copied_id"`
	CopiedName          string    `bson:"copied_name" json:"copied_name"`
	Name          string    `bson:"name" json:"name"`
	UserId    string    `bson:"user_id" json:"user_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	EventIds    []string   `bson:"event_ids" json:"event_ids"`
	EventCount  int        `bson:"event_count" json:"event_count"`	
	RatingId  string    `bson:"rating_id" json:"rating_id"`
}