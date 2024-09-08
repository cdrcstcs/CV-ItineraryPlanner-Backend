package test

import (
	"context"
	"fmt"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/dal"
	"itineraryplanner/dal/db"
	"itineraryplanner/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
func TestCreateEvent(t *testing.T) {
    ctx := context.Background()
    type arg struct {
        event *models.Event
        ctx   context.Context
    }
    tests := []struct {
        name          string
        before        func(t *testing.T) string 
        arg           arg
        wantErr       error
        wantEvent     *models.Event
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
                attraction := &models.Attraction{
                    Name:       "Test Attraction",
                    Address:    "Test Address",
                    X:          1,
                    Y:          2,
                    TagIDs:     []string{"1", "2"},
                }
                createdAttraction, err := attractionDal.CreateAttraction(ctx, attraction)
                if err != nil {
                    t.Fatalf("failed to create attraction: %v", err)
                }
                return createdAttraction.Id
            },
            arg: arg{
                ctx: ctx,
                event: &models.Event{
                    StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                    EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                    AttractionId:  "", 
                    AttractionName: "test",
                    Description:   "TestDescription",
                },
            },
            wantEvent: &models.Event{
                StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                AttractionId:  "", 
                AttractionName: "test",
                Description:   "TestDescription",
            },
        },
        {
            name: "with id err",
            before: func(t *testing.T) string {
                attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
                attraction := &models.Attraction{
                    Name:       "Test Attraction",
                    Address:    "Test Address",
                    X:          1,
                    Y:          2,
                    TagIDs:     []string{"1", "2"},
                }
                createdAttraction, err := attractionDal.CreateAttraction(ctx, attraction)
                if err != nil {
                    t.Fatalf("failed to create attraction: %v", err)
                }
				fmt.Print(createdAttraction.Id)
                return createdAttraction.Id
            },
            arg: arg{
                ctx: ctx,
                event: &models.Event{
                    Id:          "1",
                    StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                    EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                    AttractionId:  "",
                    AttractionName: "test",
                    Description:   "TestDescription",
                },
            },
			wantEvent: nil,
            wantErr: custom_errs.DBErrCreateWithID,
        },
    }
    eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            attractionID := tt.before(t)
			if attractionID == "" {
                t.Fatalf("before function did not return a valid attraction ID")
				return 
            }
            tt.arg.event.AttractionId = attractionID
			if tt.wantEvent != nil {
				tt.wantEvent.AttractionId = attractionID
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
    ctx := context.Background()
    type arg struct {
        ctx      context.Context
        eventId  string
    }
    tests := []struct {
        name            string
        before          func(t *testing.T) string 
        arg
        wantEvent *models.Event
        wantErr   error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
                attraction := &models.Attraction{
                    Name:       "Test Attraction",
                    Address:    "Test Address",
                    X:          1,
                    Y:          2,
                    TagIDs:     []string{"1", "2"},
                }
                createdAttraction, err := attractionDal.CreateAttraction(ctx, attraction)
                if err != nil {
                    t.Fatalf("failed to create attraction: %v", err)
                }
                return createdAttraction.Id
            },
            arg: arg{
                ctx:     ctx,
                eventId: "",
            },
            wantEvent: &models.Event{
                StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                AttractionId:  "", 
                AttractionName: "test",
                Description:   "TestDescription",
            },
            wantErr: nil,
        },
        {
            name: "event not found",
            before: func(t *testing.T) string {
                attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
                attraction := &models.Attraction{
                    Name:       "Test Attraction",
                    Address:    "Test Address",
                    X:          1,
                    Y:          2,
                    TagIDs:     []string{"1", "2"},
                }
                createdAttraction, err := attractionDal.CreateAttraction(ctx, attraction)
                if err != nil {
                    t.Fatalf("failed to create attraction: %v", err)
                }
                return createdAttraction.Id
            },
            arg: arg{
                ctx:     ctx,
                eventId: "non-existent-id", 
            },
            wantEvent: nil,
            wantErr:   custom_errs.DBErrIDConversion, 
        },
    }
    eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            attractionID := tt.before(t)
            if tt.arg.eventId == "" {
                event := &models.Event{
                    StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                    EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                    AttractionId:  attractionID,
                    AttractionName: "test",
                    Description:   "TestDescription",
                }
                createdEvent, err := eventDal.CreateEvent(ctx, event)
                if err != nil {
                    t.Fatalf("failed to create event: %v", err)
                }
                tt.arg.eventId = createdEvent.Id
                tt.wantEvent.Id = createdEvent.Id
				tt.wantEvent.AttractionId = attractionID
            }
            gotEvent, err := eventDal.GetEventById(tt.arg.ctx, tt.arg.eventId)
            assert.Equal(t, tt.wantErr, err)
            if err != nil {
                return
            }
            if gotEvent == nil {
                t.Fatalf("GetEventById returned nil event")
            }
            assert.Equal(t, tt.wantEvent, gotEvent)
        })
    }
}
func TestGetEvent(t *testing.T) {
	ctx:= context.Background()
	type arg struct {
		ctx	context.Context
	}
	eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
	e, _ := eventDal.GetEvent(ctx)
	tests := []struct {
		name		string
		before 		func(t *testing.T)
		arg
		wantEvent 	[]*models.Event
		wantErr    	error
	}{
		{
			name: "success",
			arg: arg{
				ctx:	ctx,
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
    ctx := context.Background()
	type arg struct {
        event *models.Event
        ctx   context.Context
    }
    createEvent := func(t *testing.T) (string) {
        eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
        event := &models.Event{
            StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
            EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
            AttractionId:   "1",
            AttractionName: "test",
            Description:    "testdescription",
        }
        createdEvent, err := eventDal.CreateEvent(ctx, event) 
        if err != nil {
            t.Fatalf("Failed to create event: %v", err)
        }
        return createdEvent.Id
    }
    tests := []struct {
        name            string
        before          func(t *testing.T) (string)
        arg
        wantEvent *models.Event
        wantErr   error
    }{
        {
            name: "success",
            before: createEvent,
            arg: arg{
                ctx:   ctx,
                event: &models.Event{
                    StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                    EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                    AttractionId:   "1",
                    AttractionName: "test",
                    Description:    "testdescription2", 
                },
            },
            wantEvent: &models.Event{
                StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                AttractionId:   "1",
                AttractionName: "test",
                Description:    "testdescription2",
            },
            wantErr: nil,
        },
    }
    eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            createdEventID := tt.before(t)
            tt.arg.event.Id = createdEventID
			tt.wantEvent.Id = createdEventID
            gotEvent, err := eventDal.UpdateEvent(tt.arg.ctx, tt.arg.event)
            assert.Equal(t, tt.wantErr, err)
            if err != nil {
                return
            }
            assert.NotEmpty(t, gotEvent)
            assert.Equal(t, tt.wantEvent, gotEvent)
        })
    }
}
func TestDeleteEvent(t *testing.T) {
    ctx := context.Background()
	type arg struct {
        ctx      context.Context
        eventId  string
    }
    createEvent := func(t *testing.T) (string) {
        eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
        event := &models.Event{
            StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
            EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
            AttractionId:   "1",
            AttractionName: "test",
            Description:    "testdescription",
        }
        createdEvent, err := eventDal.CreateEvent(ctx, event) 
        if err != nil {
            t.Fatalf("Failed to create event: %v", err)
        }
        return createdEvent.Id
    }
    tests := []struct {
        name            string
        before          func(t *testing.T) (string) 
        arg
        wantEvent *models.Event
        wantErr   error
    }{
        {
            name: "success",
            before: createEvent,
            arg: arg{
                ctx:     ctx,
                eventId: "", 
            },
            wantEvent: &models.Event{ 
                StartTime:      time.Date(2024, 5, 13, 4, 33, 20, 430000000, time.UTC),
                EndTime:        time.Date(2024, 5, 13, 5, 33, 20, 430000000, time.UTC),
                AttractionId:   "1",
                AttractionName: "test",
                Description:    "testdescription",
            },
            wantErr: nil,
        },
        {
            name: "invalid id",
            before: func(t *testing.T) (string) {
                return "nonexistentID"
            },
            arg: arg{
                ctx:     ctx,
                eventId: "", 
            },
            wantEvent: nil,
            wantErr:   custom_errs.DBErrIDConversion,
        },
    }
    eventDal := &dal.EventDal{MainDB: db.GetMemoMongo()}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            eventId := tt.before(t)
            tt.arg.eventId = eventId
            gotEvent, err := eventDal.DeleteEvent(tt.arg.ctx, tt.arg.eventId)
            assert.Equal(t, tt.wantErr, err)
            if err != nil {
                assert.Equal(t, tt.wantEvent, gotEvent)
                return
            }
			tt.wantEvent.Id = eventId
            assert.Equal(t, tt.wantEvent, gotEvent)
        })
    }
}