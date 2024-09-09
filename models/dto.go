package models

import (
	"time"
)

type CreateUserReq struct {
	Name          string    `json:"name" validate:"required"`
	Password      string    `json:"password" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	EmailPassword string    `json:"email_password" validate:"required"`
	Phone         string    `json:"phone" validate:"required"`
	Token         string    `json:"token"`
	UserType      string    `json:"user_type" validate:"required,oneof=ADMIN USER"`
	RefreshToken  string    `json:"refresh_token"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdateUserReq struct {
	Id            string    `json:"id" validate:"required"`
	Name          string    `json:"name" validate:"required"`
	Password      string    `json:"password" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	EmailPassword string    `json:"email_password" validate:"required"`
	Phone         string    `json:"phone" validate:"required"`
	Token         string    `json:"token"`
	UserType      string    `json:"user_type" validate:"required,oneof=ADMIN USER"`
	RefreshToken  string    `json:"refresh_token"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateAttractionReq struct {
	Name       string   `json:"name" validate:"required"`
	Address    string   `json:"address" validate:"required"`
	X          int      `json:"x" validate:"required"`
	Y          int      `json:"y" validate:"required"`
	TagIDs     []string `json:"tag_ids" validate:"required,dive,uuid"`
	RatingId   string   `json:"rating_id" validate:"required,uuid"`
	City       string   `json:"city" validate:"required"`
}

type CreateAttractionResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}

type GetAttractionByIdReq struct {
	Id string `json:"id" validate:"required,uuid"`
}

type GetAttractionByIdResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}

type GetAttractionReq struct {
}

type GetAttractionResp struct {
	Attractions []*AttractionDTO `json:"attractions"`
}

type UpdateAttractionReq struct {
	Id         string   `json:"id" validate:"required,uuid"`
	Name       string   `json:"name" validate:"required"`
	Address    string   `json:"address" validate:"required"`
	X          int      `json:"x" validate:"required"`
	Y          int      `json:"y" validate:"required"`
	TagIDs     []string `json:"tag_ids" validate:"required,dive,uuid"`
	RatingId   string   `json:"rating_id" validate:"required,uuid"`
	City       string   `json:"city" validate:"required"`
}

type UpdateAttractionResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}

type DeleteAttractionReq struct {
	Id string `json:"id" validate:"required,uuid"`
}

type DeleteAttractionResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}

type CreateItineraryReq struct {
	CopiedId    string    `json:"copied_id" validate:"required,uuid"`
	CopiedName  string    `json:"copied_name" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	UserId      string    `json:"user_id" validate:"required,uuid"`
	StartTime   time.Time `json:"start_time" validate:"required"`
	EndTime     time.Time `json:"end_time" validate:"required"`
	EventIds    []string  `json:"event_ids" validate:"required,dive,uuid"`
	EventCount  int       `json:"event_count" validate:"required"`
	RatingId    string    `json:"rating_id" validate:"required,uuid"`
}
type CreateItineraryResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}
type GetItineraryByIdReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type GetItineraryByIdResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}
type GetItineraryReq struct {
}
type GetItineraryResp struct {
	Itineraries []*ItineraryDTO `json:"itineraries"`
}
type UpdateItineraryReq struct {
	Id         string    `json:"id" validate:"required,uuid"`
	CopiedId   string    `json:"copied_id" validate:"required,uuid"`
	CopiedName string    `json:"copied_name" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	UserId     string    `json:"user_id" validate:"required,uuid"`
	StartTime  time.Time `json:"start_time" validate:"required"`
	EndTime    time.Time `json:"end_time" validate:"required"`
	EventIds   []string  `json:"event_ids" validate:"required,dive,uuid"`
	EventCount int       `json:"event_count" validate:"required"`
	RatingId   string    `json:"rating_id" validate:"required,uuid"`
}
type UpdateItineraryResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}
type DeleteItineraryReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type DeleteItineraryResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}
type CreateRatingReq struct {
	Score int `json:"score" validate:"required,min=1,max=5"`
}
type CreateRatingResp struct {
	Rating *RatingDTO `json:"rating"`
}
type GetRatingByIdReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type GetRatingByIdResp struct {
	Rating *RatingDTO `json:"rating"`
}
type GetRatingReq struct {
}
type GetRatingResp struct {
	Ratings []*RatingDTO `json:"ratings"`
}
type UpdateRatingReq struct {
	Id    string `json:"id" validate:"required,uuid"`
	Score int    `json:"score" validate:"required,min=1,max=5"`
}
type UpdateRatingResp struct {
	Rating *RatingDTO `json:"rating"`
}
type DeleteRatingReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type DeleteRatingResp struct {
	Rating *RatingDTO `json:"rating"`
}
type CreateEventReq struct {
	StartTime      time.Time `json:"start_time" validate:"required"`
	EndTime        time.Time `json:"end_time" validate:"required"`
	AttractionId   string    `json:"attraction_id" validate:"required,uuid"`
	AttractionName string    `json:"attraction_name" validate:"required"`
	Description    string    `json:"description" validate:"required"`
}
type CreateEventResp struct {
	Event *EventDTO `json:"event"`
}
type GetEventByIdReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type GetEventByIdResp struct {
	Event *EventDTO `json:"event"`
}
type GetEventReq struct {
}
type GetEventResp struct {
	Events []*EventDTO `json:"events"`
}
type UpdateEventReq struct {
	Id             string    `json:"id" validate:"required,uuid"`
	StartTime      time.Time `json:"start_time" validate:"required"`
	EndTime        time.Time `json:"end_time" validate:"required"`
	AttractionId   string    `json:"attraction_id" validate:"required,uuid"`
	AttractionName string    `json:"attraction_name" validate:"required"`
	Description    string    `json:"description" validate:"required"`
}
type UpdateEventResp struct {
	Event *EventDTO `json:"event"`
}
type DeleteEventReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type DeleteEventResp struct {
	Event *EventDTO `json:"event"`
}
type CreateTagReq struct {
	Value string `json:"value" validate:"required"`
}
type CreateTagResp struct {
	Tag *TagDTO `json:"tag"`
}
type GetTagByIdReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type GetTagByIdResp struct {
	Tag *TagDTO `json:"tag"`
}
type GetTagReq struct {
}
type GetTagResp struct {
	Tags []*TagDTO `json:"tags"`
}
type UpdateTagReq struct {
	Id    string `json:"id" validate:"required,uuid"`
	Value string `json:"value" validate:"required"`
}
type UpdateTagResp struct {
	Tag *TagDTO `json:"tag"`
}
type DeleteTagReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type DeleteTagResp struct {
	Tag *TagDTO `json:"tag"`
}
type GetUserByIdReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type GetUserByIdResp struct {
	User *UserDTO `json:"user"`
}
type GetUserReq struct {
}
type GetUserResp struct {
	Users []*UserDTO `json:"users"`
}
type DeleteUserReq struct {
	Id string `json:"id" validate:"required,uuid"`
}
type DeleteUserResp struct {
	User *UserDTO `json:"user"`
}
type UpdateUserResp struct {
	User *UserDTO `json:"user"`
}
type CreateUserResp struct {
	User *UserDTO `json:"user"`
}
type AttractionDTO struct {
	Id         string     `bson:"_id" json:"id"`
	Name       string     `bson:"name" json:"name"`
	Address    string     `bson:"address" json:"address"`
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`	
	Tags       []*TagDTO      `bson:"tags" json:"tags"`
	Rating    *RatingDTO    `bson:"rating" json:"rating"`
	City     string `bson:"city" json:"city"`

}
type ItineraryDTO struct {
	Id        string    `bson:"_id" json:"id"`
	CopiedId  string        `bson:"copied_id" json:"copied_id"`
	CopiedName          string    `bson:"copied_name" json:"copied_name"`
	Name          string    `bson:"name" json:"name"`
	UserId    string        `bson:"user_id" json:"user_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	EventIds    []string   `bson:"event_ids" json:"event_ids"`
	EventCount  int        `bson:"event_count" json:"event_count"`	
	Rating    *RatingDTO    `bson:"rating" json:"rating"`
}
type TagDTO struct {
	Id    string `bson:"_id,omitempty" json:"id"`
	Value string `bson:"value" json:"value"` 
}
type RatingDTO struct {
	Id       string `bson:"_id" json:"id"`
	Score    int    `bson:"score" json:"score"`
}
type EventDTO struct {
	Id          string     `bson:"_id" json:"id"`
	StartTime   time.Time  `bson:"start_time" json:"start_time"`
	EndTime     time.Time  `bson:"end_time" json:"end_time"`
	AttractionId  string `bson:"attraction_id" json:"attraction_id"`
	AttractionName string `bson:"attraction_name" json:"attraction_name"`
	Description string     `bson:"description" json:"description"`
}
type UserDTO struct {
	Id            string `bson:"_id" json:"id"`
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	Email         string `bson:"email" json:"email"`
	EmailPassword string `bson:"email_password" json:"email_password"`
	OTPToken  	  string `bson:"otpToken" json:"otpToken"`
	SID			  string `bson:"SID" json:"SID"`
}