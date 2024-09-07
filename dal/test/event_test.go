package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"itineraryplanner/common/custom_errs"
	"itineraryplanner/dal/db"
	"itineraryplanner/models"
	"itineraryplanner/dal"

)

func TestCreateEvent(t *testing.T) {
	ctx := context.Background()

	type arg struct {
		event *models.Event
		ctx   context.Context
	}

	tests := []struct {
		name          string
		before        func(t *testing.T)
		arg           arg
		wantErr       error
		wantEvent     *models.Event
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				event: &models.Event{
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "66402ddd27a2989c866cac35",
					AttractionName: "test",
					Description: "TestDescription",
				},
			},
			wantEvent: &models.Event{
				StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
				EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
				AttractionId:  "66402ddd27a2989c866cac35",
				AttractionName: "test",
				Description: "TestDescription",
			},
		},
		{
			name: "with id err",
			arg: arg{
				ctx: ctx,
				event: &models.Event{
					Id:          "1",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId:  "66402ddd27a2989c866cac35",
					AttractionName: "test",
					Description: "TestDescription",
				},
			},
			wantErr: custom_errs.DBErrCreateWithID,
		},
	}
	eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t)
			}
			gotEvent, err := eventDal.CreateEvent(tt.arg.ctx, tt.arg.event)

			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotEvent.Id)
			gotEvent.Id = "" 
			assert.Equal(t, tt.wantEvent, gotEvent)

		})
	}
}

func TestGetEventById(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		eventId 	string
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantEvent *models.Event
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				eventId:  "66420cf5a9df66a0a0737cb5",
			},
			wantEvent: &models.Event{
				Id:           "66420cf5a9df66a0a0737cb5",
				StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
				EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
				AttractionId: "66402ddd27a2989c866cac35",
				AttractionName: "test",
				Description:  "TestDescription",
			},
			wantErr: nil,
		},
	}
	
	eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvent, err := eventDal.GetEventById(tt.arg.ctx, tt.arg.eventId)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotEvent)
			assert.Equal(t, tt.wantEvent, gotEvent)
		})	
	}
}
func TestGetEvent(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
	}
	eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
	e, _ := eventDal.GetEvent(ctx)
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantEvent []*models.Event
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
			},
			wantEvent: e,
			wantErr: nil,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvent, err := eventDal.GetEvent(tt.arg.ctx)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotEvent)
			assert.Equal(t, tt.wantEvent, gotEvent)
		})	
	}
}

func TestUpdateEvent(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		event *models.Event
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantEvent *models.Event
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				event: &models.Event{
					Id:"66419810c1c4d8e352e17786",
					StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
					EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
					AttractionId :"1",
					AttractionName: "test",
					Description :"testdescription2",
				},
			},
			wantEvent: &models.Event{
				Id:"66419810c1c4d8e352e17786",
				StartTime:    time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
				EndTime:      time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
				AttractionId :"1",
				AttractionName: "test",
				Description :"testdescription2",
			},
			wantErr: nil,
		},
	}
	
	eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvent, err := eventDal.UpdateEvent(tt.arg.ctx, tt.arg.event)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotEvent)
			assert.ElementsMatch(t, tt.wantEvent, gotEvent)
		})	
	}
}

func TestDeleteEvent(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx           context.Context
		eventId string
	}
	
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantEvent *models.Event
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx:           ctx,
				eventId:  "1234",
			},
			wantEvent: &models.Event{
					Id:"1234",
					StartTime:   time.Now(),
					EndTime:     time.Now().Add(1 * time.Hour),
					AttractionId :"1",
					AttractionName: "test",
					Description :"testdescription",
			},
			wantErr: nil,
		},
		{
			name: "invalid id",
			arg: arg{
				ctx:           ctx,
				eventId:  "12345",
			},
			wantEvent: nil,
			wantErr: custom_errs.DBErrDeleteWithID,
		},
	}
	
	eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvent, err := eventDal.DeleteEvent(tt.arg.ctx, tt.arg.eventId)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotEvent)
			assert.ElementsMatch(t, tt.wantEvent, gotEvent)
		})	
	}
}