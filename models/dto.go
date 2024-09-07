package models

import "time"

type AttractionDTO struct {
	Id         string     `bson:"_id" json:"id"`
	Name       string     `bson:"name" json:"name"`
	Address    string     `bson:"address" json:"address"`
	Coordinate *CoordinateDTO `bson:"coordinate" json:"coordinate"`
	Tags       []*TagDTO      `bson:"tags" json:"tags"`
	Rating    *RatingDTO    `bson:"rating" json:"rating"`
}
type ItineraryDTO struct {
	Id        string    `bson:"_id" json:"id"`
	CopiedId  string   `bson:"copied_id" json:"copied_id"`
	User    *UserDTO        `bson:"user" json:"user"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	Events    []*EventDTO   `bson:"events" json:"events"`
	Rating    *RatingDTO    `bson:"rating" json:"rating"`
}
type TagDTO struct {
	Id    string `bson:"_id,omitempty" json:"id"`
	Value string `bson:"value" json:"value"` // value should be unique, however tag value should not be passes around
}
type RatingDTO struct {
	Id       string `bson:"_id" json:"id"`
	Type     string   `bson:"type" json:"type"`
	User   *UserDTO `bson:"user" json:"user"`
	ObjectId string `bson:"object_id" json:"object_id"`
	Score    int    `bson:"score" json:"score"`
}
type EventDTO struct {
	Id          string     `bson:"_id" json:"id"`
	ItineraryId  string  `bson:"itinerary_id" json:"itinerary_id"`
	StartTime   time.Time  `bson:"start_time" json:"start_time"`
	EndTime     time.Time  `bson:"end_time" json:"end_time"`
	Attraction  *AttractionDTO `bson:"attraction" json:"attraction"`
	Description string     `bson:"description" json:"description"`
}
type UserDTO struct {
	Id            string `bson:"_id" json:"id"`
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	Email         string `bson:"email" json:"email"`
	EmailPassword string `bson:"email_password" json:"email_password"`
}
type CoordinateDTO struct {
	Id string `bson:"_id,omitempty" json:"id"`
	ItineraryId  string `bson:"itinerary_id"  json:"itinerary_id"`
	EventId    string    `bson:"event_id" json:"event_id"`
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}

type CreateAttractionReq struct {
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	CoordinateId string `json:"coordinate"`
	TagIDs     []string   `json:"tag_ids"`
	RatingId    string    `bson:"rating_id" json:"rating_id"`
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
	ItineraryId  string `bson:"itinerary_id"  json:"itinerary_id"`
	EventId    string    `bson:"event_id" json:"event_id"`
	CoordinateId string `json:"coordinate"`
	TagIDs     []string   `json:"tag_ids"`
	RatingId    string    `bson:"rating_id" json:"rating_id"`
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
	UserId    string        `bson:"user_id" json:"user_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	EventIds    []string   `bson:"event_ids" json:"event_ids"`
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
	UserId    string        `bson:"user_id" json:"user_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	EventIds    []string   `bson:"event_ids" json:"event_ids"`
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
	Type     Type   `bson:"type" json:"type"`
	UserId   string `bson:"user_id" json:"user_id"`
	ObjectId string `bson:"object_id" json:"object_id"`
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
	Type     Type   `bson:"type" json:"type"`
	UserId   string `bson:"user_id" json:"user_id"`
	ObjectId string `bson:"object_id" json:"object_id"`
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
	ItineraryId string         `bson:"itinerary_id" json:"itinerary_id"`
	StartTime   time.Time  `bson:"start_time" json:"start_time"`
	EndTime     time.Time  `bson:"end_time" json:"end_time"`
	AttractionId  string `bson:"attraction_id" json:"attraction_id"`
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
	ItineraryId string         `bson:"itinerary_id" json:"itinerary_id"`
	StartTime   time.Time  `bson:"start_time" json:"start_time"`
	EndTime     time.Time  `bson:"end_time" json:"end_time"`
	AttractionId  string `bson:"attraction_id" json:"attraction_id"`
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
}
type CreateUserResp struct {
	User *UserDTO `json:"user"`
}
type LoginUserReq struct {
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	Email         string `bson:"email" json:"email"`
	EmailPassword string `bson:"email_password" json:"email_password"`
}
type LoginUserResp struct {
	Check bool `json:"check"`
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
type CreateCoordinateReq struct{
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}
type CreateCoordinateResp struct {
	Coordinate *CoordinateDTO `json:"coordinate"`
}
type GetCoordinateByIdReq struct {
	Id string `bson:"_id" json:"_id"`
}
type GetCoordinateByIdResp struct {
	Coordinate *CoordinateDTO `json:"coordinate"`
}
type GetCoordinateReq struct {
}
type GetCoordinateResp struct {
	Coordinates []*CoordinateDTO `json:"coordinates"`
}
type UpdateCoordinateReq struct {
	Id string `bson:"_id" json:"id"`
	ItineraryId  string `bson:"itinerary_id"  json:"itinerary_id"`
	EventId    string    `bson:"event_id" json:"event_id"`
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}
type UpdateCoordinateResp struct {
	Coordinate *CoordinateDTO `json:"coordinate"`
}
type DeleteCoordinateReq struct {
	Id string `bson:"_id" json:"id"`
}
type DeleteCoordinateResp struct {
	Coordinate *CoordinateDTO `json:"coordinate"`
}
type CreateTagReq struct {
	Value string `bson:"value" json:"value"` // value should be unique, however tag value should not be passes around
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
	Value string `bson:"value" json:"value"` // value should be unique, however tag value should not be passes around
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

type RecommendItineraryReq struct {
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}
type RecommendItineraryResp struct {
	Itineraries []*ItineraryDTO `json:"itineraries"`
}
type BuildItineraryReq struct {
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}
type BuildItineraryResp struct {
	Itinerary *ItineraryDTO `json:"itinerary"`
}