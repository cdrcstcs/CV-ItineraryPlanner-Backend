package models
import (
	"time"
)
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
type CreateAttractionReq struct {
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`	
	TagIDs     []string   `json:"tag_ids"`
	RatingId    string    `bson:"rating_id" json:"rating_id"`
	City     string `bson:"city" json:"city"`
}
type CreateAttractionResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}
type GetAttractionByIdReq struct {
	Id string `bson:"_id" json:"_id"`
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
	Id string `bson:"_id" json:"id"`
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`	
	TagIDs     []string   `json:"tag_ids"`
	RatingId    string    `bson:"rating_id" json:"rating_id"`
	City     string `bson:"city" json:"city"`
}
type UpdateAttractionResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}
type DeleteAttractionReq struct {
	Id string `bson:"_id" json:"id"`
}
type DeleteAttractionResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}
type CreateItineraryReq struct {
	CopiedId  string        `bson:"copied_id" json:"copied_id"`
	CopiedName          string    `bson:"copied_name" json:"copied_name"`
	Name          string    `bson:"name" json:"name"`
	UserId    string        `bson:"user_id" json:"user_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	EventIds    []string   `bson:"event_ids" json:"event_ids"`
	EventCount  int        `bson:"event_count" json:"event_count"`
	RatingId    string    `bson:"rating_id" json:"rating_id"`
}
type CreateItineraryResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}
type GetItineraryByIdReq struct {
	Id string `bson:"_id" json:"_id"`
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
	Id string `bson:"_id" json:"id"`
	CopiedId  string        `bson:"copied_id" json:"copied_id"`
	CopiedName          string    `bson:"copied_name" json:"copied_name"`
	Name          string    `bson:"name" json:"name"`
	UserId    string        `bson:"user_id" json:"user_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	EventIds    []string   `bson:"event_ids" json:"event_ids"`
	EventCount  int        `bson:"event_count" json:"event_count"`
	RatingId    string    `bson:"rating_id" json:"rating_id"`
}
type UpdateItineraryResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}
type DeleteItineraryReq struct {
	Id string `bson:"_id" json:"id"`
}
type DeleteItineraryResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}
type CreateRatingReq struct {
	Score    int    `bson:"score" json:"score"`
}
type CreateRatingResp struct {
	Rating *RatingDTO `json:"rating"`
}
type GetRatingByIdReq struct {
	Id string `bson:"_id" json:"_id"`
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
	Id string `bson:"_id" json:"id"`
	Score    int    `bson:"score" json:"score"`
}
type UpdateRatingResp struct {
	Rating *RatingDTO `json:"rating"`
}
type DeleteRatingReq struct {
	Id string `bson:"_id" json:"id"`
}
type DeleteRatingResp struct {
	Rating *RatingDTO `json:"rating"`
}
type CreateEventReq struct {
	StartTime   time.Time  `bson:"start_time" json:"start_time"`
	EndTime     time.Time  `bson:"end_time" json:"end_time"`
	AttractionId  string `bson:"attraction_id" json:"attraction_id"`
	AttractionName string `bson:"attraction_name" json:"attraction_name"`
	Description string     `bson:"description" json:"description"`
}
type CreateEventResp struct {
	Event *EventDTO `json:"event"`
}
type GetEventByIdReq struct {
	Id string `bson:"_id" json:"_id"`
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
	Id string `bson:"_id" json:"id"`
	StartTime   time.Time  `bson:"start_time" json:"start_time"`
	EndTime     time.Time  `bson:"end_time" json:"end_time"`
	AttractionId  string `bson:"attraction_id" json:"attraction_id"`
	AttractionName string `bson:"attraction_name" json:"attraction_name"`
	Description string     `bson:"description" json:"description"`
}
type UpdateEventResp struct {
	Event *EventDTO `json:"event"`
}
type DeleteEventReq struct {
	Id string `bson:"_id" json:"id"`
}
type DeleteEventResp struct {
	Event *EventDTO `json:"event"`
}
type CreateUserReq struct {
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	Email         string `bson:"email" json:"email"`
	EmailPassword string `bson:"email_password" json:"email_password"`
	Phone			string		`json:"phone" validate:"required"`
	Token			string		`json:"token"`
	User_type		string		`json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token	string		`json:"refresh_token"`
	Created_at		time.Time	`json:"created_at"`
	Updated_at		time.Time	`json:"updated_at"`
}
type CreateUserResp struct {
	User *UserDTO `json:"user"`
}
type GetUserByIdReq struct {
	Id string `bson:"_id" json:"_id"`
}
type GetUserByIdResp struct {
	User *UserDTO `json:"user"`
}
type GetUserReq struct {
}
type GetUserResp struct {
	Users []*UserDTO `json:"users"`
}
type UpdateUserReq struct {
	Id string `bson:"_id" json:"id"`
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	Email         string `bson:"email" json:"email"`
	EmailPassword string `bson:"email_password" json:"email_password"`
	Phone			string		`json:"phone" validate:"required"`
	Token			string		`json:"token"`
	User_type		string		`json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token	string		`json:"refresh_token"`
	Created_at		time.Time	`json:"created_at"`
	Updated_at		time.Time	`json:"updated_at"`
}
type UpdateUserResp struct {
	User *UserDTO `json:"user"`
}
type DeleteUserReq struct {
	Id string `bson:"_id" json:"id"`
}
type DeleteUserResp struct {
	User *UserDTO `json:"user"`
}
type CreateTagReq struct {
	Value string `bson:"value" json:"value"` 
}
type CreateTagResp struct {
	Tag *TagDTO `json:"tag"`
}
type GetTagByIdReq struct {
	Id string `bson:"_id" json:"_id"`
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
	Id string `bson:"_id" json:"id"`
	Value string `bson:"value" json:"value"` 
}
type UpdateTagResp struct {
	Tag *TagDTO `json:"tag"`
}
type DeleteTagReq struct {
	Id string `bson:"_id" json:"id"`
}
type DeleteTagResp struct {
	Tag *TagDTO `json:"tag"`
}