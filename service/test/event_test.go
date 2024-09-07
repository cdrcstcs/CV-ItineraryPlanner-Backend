package test

import (
	"context"
	"testing"
	"time"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"itineraryplanner/dal/mock"
	"itineraryplanner/models"
	"itineraryplanner/service"
	ser_mock "itineraryplanner/service/mock"
)


func ConfigCreateEvent(t *testing.T) (context.Context, *mock.MockCreateEventDal, *ser_mock.MockEventDTOService, *service.EventService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateEventDal(ctrl)
	mockDTO := ser_mock.NewMockEventDTOService(ctrl)
	eventService := &service.EventService{CDal: mockC}
	return ctx, mockC, mockDTO, eventService
}

func ConfigGetEvent(t *testing.T) (context.Context, *mock.MockGetEventDal, *ser_mock.MockEventDTOService, *service.EventService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockG := mock.NewMockGetEventDal(ctrl)
	mockDTO := ser_mock.NewMockEventDTOService(ctrl)
	eventService := &service.EventService{GDal: mockG}
	return ctx, mockG, mockDTO, eventService
}

func ConfigGetEventById(t *testing.T) (context.Context, *mock.MockGetEventByIdDal, *ser_mock.MockEventDTOService, *service.EventService) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockB := mock.NewMockGetEventByIdDal(ctrl)
	mockDTO := ser_mock.NewMockEventDTOService(ctrl)
	eventService := &service.EventService{BDal: mockB}
	return ctx, mockB, mockDTO, eventService
}

func ConfigUpdateEvent(t *testing.T) (context.Context, *mock.MockUpdateEventDal, *ser_mock.MockEventDTOService, *service.EventService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockU := mock.NewMockUpdateEventDal(ctrl)
	mockDTO := ser_mock.NewMockEventDTOService(ctrl)
	eventService := &service.EventService{UDal: mockU}
	return ctx, mockU, mockDTO, eventService
}
func ConfigDeleteEvent(t *testing.T) (context.Context, *mock.MockDeleteEventDal, *ser_mock.MockEventDTOService, *service.EventService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockD := mock.NewMockDeleteEventDal(ctrl)
	mockDTO := ser_mock.NewMockEventDTOService(ctrl)
	eventService := &service.EventService{DDal: mockD}
	return ctx, mockD, mockDTO, eventService
}
func ConfigConvertToDTOEvent(t *testing.T) (context.Context, *ser_mock.MockFacadeDesignPatternService, *service.EventService){
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockC := mock.NewMockCreateEventDal(ctrl)
	mockG := mock.NewMockGetEventDal(ctrl)
	mockF := ser_mock.NewMockFacadeDesignPatternService(ctrl)
	eventService := &service.EventService{CDal: mockC, GDal: mockG}
	return ctx, mockF, eventService
}


func TestCreateEvent(t *testing.T) {
	ctx, mockC, mockDTO, eventService := ConfigCreateEvent(t)

	type arg struct {
		req *models.CreateEventReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.CreateEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.CreateEventReq{
					ItineraryId: "test_itinerary_id",
					StartTime:   time.Now(),
					EndTime:     time.Now().Add(time.Hour),
					AttractionId:  "test_attraction_id",
					Description: "test_description",
				},
			},
			before: func(t *testing.T) {
				mockC.EXPECT().CreateEvent(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOEvent(
					ctx,
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
				).Return(
					&models.EventDTO{
						Id:            "test_event_id",
						ItineraryId: "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						Attraction:    &models.AttractionDTO{
							Id:      "test_attraction_id",
							Name:    "test_name",
							Address: "test_address",
							Coordinate: &models.CoordinateDTO{
								Id: "1",
								X: 1,
								Y: 1,
							},
							Tags: []*models.TagDTO{
								{
									Id: "1", 
									Value: "test 1",
								},
								{
									Id: "2", 
									Value: "test 2",
								},
							},
						},
						Description:   "test_description",
					},
					nil,
				).Do(func(ctx context.Context, event *models.Event) {
					fmt.Println("ConvertDBOToDTOEventEvent called with", event)
				})
			},
			wantResp: &models.CreateEventResp{
				Event: &models.EventDTO{
					Id:            "test_event_id",
                    ItineraryId: "test_itinerary_id",
					StartTime:     time.Now(),
					EndTime:       time.Now().Add(time.Hour),
					Attraction:    &models.AttractionDTO{
						Id:      "test_attraction_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X: 1,
							Y: 1,
						},
						Tags: []*models.TagDTO{
							{
								Id: "1", 
								Value: "test 1",
							},
							{
								Id: "2", 
								Value: "test 2",
							},
						},
					},
					Description:   "test_description",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.CreateEvent(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetEventById(t *testing.T) {
	ctx, mockB, mockDTO, eventService := ConfigGetEventById(t)

	type arg struct {
		req *models.GetEventByIdReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetEventByIdResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetEventByIdReq{
					Id: "test_event_id",
				},
			},
			before: func(t *testing.T) {
				mockB.EXPECT().GetEventById(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOEvent(
					ctx,
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
				).Return(
					&models.EventDTO{
						Id:            "test_event_id",
						ItineraryId: "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						Attraction:    &models.AttractionDTO{
							Id:      "test_attraction_id",
							Name:    "test_name",
							Address: "test_address",
							Coordinate: &models.CoordinateDTO{
								Id: "1",
								X: 1,
								Y: 1,
							},
							Tags: []*models.TagDTO{
								{
									Id: "1", 
									Value: "test 1",
								},
								{
									Id: "2", 
									Value: "test 2",
								},
							},
						},
						Description:   "test_description",
					},
					nil,
				).Do(func(ctx context.Context, event *models.Event) {
					fmt.Println("ConvertDBOToDTOEventEvent called with", event)
				})
			},
			wantResp: &models.GetEventByIdResp{
				Event: &models.EventDTO{
					Id:            "test_event_id",
					ItineraryId: "test_itinerary_id",
					StartTime:     time.Now(),
					EndTime:       time.Now().Add(time.Hour),
					Attraction:    &models.AttractionDTO{
						Id:      "test_attraction_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X: 1,
							Y: 1,
						},
						Tags: []*models.TagDTO{
							{
								Id: "1", 
								Value: "test 1",
							},
							{
								Id: "2", 
								Value: "test 2",
							},
						},
					},
					Description:   "test_description",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.GetEventById(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetEvent(t *testing.T) {
	ctx, mockG, mockDTO, eventService := ConfigGetEvent(t)

	type arg struct {
		req *models.GetEventReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.GetEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.GetEventReq{
				},
			},
			before: func(t *testing.T) {
				mockG.EXPECT().GetEvent(ctx).Return(
					[]*models.Event{
						{
							Id:            "test_event_id",
							ItineraryId:   "test_itinerary_id",
							StartTime:     time.Now(),
							EndTime:       time.Now().Add(time.Hour),
							AttractionId:  "test_attraction_id",
							Description:   "test_description",
						},
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOEvent(
					ctx,
					gomock.Any(),
				).Return(
					&models.EventDTO{
						Id:            "test_event_id",
						ItineraryId: "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						Attraction:    &models.AttractionDTO{
							Id:      "test_attraction_id",
							Name:    "test_name",
							Address: "test_address",
							Coordinate: &models.CoordinateDTO{
								Id: "1",
								X: 1,
								Y: 1,
							},
							Tags: []*models.TagDTO{
								{
									Id: "1", 
									Value: "test 1",
								},
								{
									Id: "2", 
									Value: "test 2",
								},
							},
						},
						Description:   "test_description",
					},
					nil,
				).Do(func(ctx context.Context, event *models.Event) {
					fmt.Println("ConvertDBOToDTOEventEvent called with", event)
				})
			},
			wantResp: &models.GetEventResp{
				Events: []*models.EventDTO{
					{
						Id:            "test_event_id",
						ItineraryId: "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						Attraction:    &models.AttractionDTO{
							Id:      "test_attraction_id",
							Name:    "test_name",
							Address: "test_address",
							Coordinate: &models.CoordinateDTO{
								Id: "1",
								X: 1,
								Y: 1,
							},
							Tags: []*models.TagDTO{
								{
									Id: "1", 
									Value: "test 1",
								},
								{
									Id: "2", 
									Value: "test 2",
								},
							},
						},
						Description:   "test_description",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.GetEvent(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUpdateEvent(t *testing.T) {
	ctx, mockU, mockDTO, eventService := ConfigUpdateEvent(t)

	type arg struct {
		req *models.UpdateEventReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.UpdateEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.UpdateEventReq{
					Id:            "test_event_id",
					ItineraryId: "test_itinerary_id",
					StartTime:   time.Now(),
					EndTime:     time.Now().Add(time.Hour),
					AttractionId:  "test_attraction_id",
					Description: "test_description",
				},
			},
			before: func(t *testing.T) {
				mockU.EXPECT().UpdateEvent(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOEvent(
					ctx,
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
				).Return(
					&models.EventDTO{
						Id:            "test_event_id",
						ItineraryId: "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						Attraction:    &models.AttractionDTO{
							Id:      "test_attraction_id",
							Name:    "test_name",
							Address: "test_address",
							Coordinate: &models.CoordinateDTO{
								Id: "1",
								X: 1,
								Y: 1,
							},
							Tags: []*models.TagDTO{
								{
									Id: "1", 
									Value: "test 1",
								},
								{
									Id: "2", 
									Value: "test 2",
								},
							},
						},
						Description:   "test_description",
					},
					nil,
				).Do(func(ctx context.Context, event *models.Event) {
					fmt.Println("ConvertDBOToDTOEventEvent called with", event)
				})
			},
			wantResp: &models.UpdateEventResp{
				Event: &models.EventDTO{
					Id:            "test_event_id",
                    ItineraryId: "test_itinerary_id",
					StartTime:     time.Now(),
					EndTime:       time.Now().Add(time.Hour),
					Attraction:    &models.AttractionDTO{
						Id:      "test_attraction_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X: 1,
							Y: 1,
						},
						Tags: []*models.TagDTO{
							{
								Id: "1", 
								Value: "test 1",
							},
							{
								Id: "2", 
								Value: "test 2",
							},
						},
					},
					Description:   "test_description",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.UpdateEvent(ctx, tt.arg.req)
			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteEvent(t *testing.T) {
	ctx, mockD, mockDTO, eventService := ConfigDeleteEvent(t)

	type arg struct {
		req *models.DeleteEventReq
		ctx context.Context
	}

	tests := []struct {
		name     string
		arg      arg
		before   func(t *testing.T)
		wantResp *models.DeleteEventResp
		wantErr  error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				req: &models.DeleteEventReq{
					Id:    "test_event_id",
				},
			},
			before: func(t *testing.T) {
				mockD.EXPECT().DeleteEvent(
					ctx,
					gomock.Any(),
				).Return(
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
					nil,
				)
				mockDTO.EXPECT().ConvertDBOToDTOEvent(
					ctx,
					&models.Event{
						Id:            "test_event_id",
						ItineraryId:   "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						AttractionId:  "test_attraction_id",
						Description:   "test_description",
					},
				).Return(
					&models.EventDTO{
						Id:            "test_event_id",
						ItineraryId: "test_itinerary_id",
						StartTime:     time.Now(),
						EndTime:       time.Now().Add(time.Hour),
						Attraction:    &models.AttractionDTO{
							Id:      "test_attraction_id",
							Name:    "test_name",
							Address: "test_address",
							Coordinate: &models.CoordinateDTO{
								Id: "1",
								X: 1,
								Y: 1,
							},
							Tags: []*models.TagDTO{
								{
									Id: "1", 
									Value: "test 1",
								},
								{
									Id: "2", 
									Value: "test 2",
								},
							},
						},
						Description:   "test_description",
					},
					nil,
				).Do(func(ctx context.Context, event *models.Event) {
					fmt.Println("ConvertDBOToDTOEventEvent called with", event)
				})
			},
			wantResp: &models.DeleteEventResp{
				Event: &models.EventDTO{
					Id:            "test_event_id",
					ItineraryId: "test_itinerary_id",
					StartTime:     time.Now(),
					EndTime:       time.Now().Add(time.Hour),
					Attraction:    &models.AttractionDTO{
						Id:      "test_attraction_id",
						Name:    "test_name",
						Address: "test_address",
						Coordinate: &models.CoordinateDTO{
							Id: "1",
							X: 1,
							Y: 1,
						},
						Tags: []*models.TagDTO{
							{
								Id: "1", 
								Value: "test 1",
							},
							{
								Id: "2", 
								Value: "test 2",
							},
						},
					},
					Description:   "test_description",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(t)
			gotResp, err := eventService.DeleteEvent(ctx, tt.arg.req)

			assert.Equal(t, tt.wantResp, gotResp)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConvertDBOToDTOEvent(t *testing.T) {
    ctx, mockF, eventService := ConfigConvertToDTOEvent(t)

    type arg struct {
        req *models.Event
        ctx context.Context
    }

    tests := []struct {
        name     string
        before   func(t *testing.T)
        arg      arg
        wantResp *models.EventDTO
        wantErr  error
    }{
        {
            name: "success",
            arg: arg{
                req: &models.Event{
                    Id:          "test_event_id",
                    ItineraryId: "test_itinerary_id",
                    StartTime:   time.Now(),
                    EndTime:     time.Now().Add(time.Hour),
                    AttractionId:  "test_attraction_id",
                    Description: "test_description",
                },
                ctx: ctx,
            },
            before: func(t *testing.T) {
                mockF.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
                    &models.RespFacade{
                        GARB: &models.GetAttractionByIdResp{
                            Attraction: &models.AttractionDTO{
								Id:      "test_attraction_id",
								Name:    "test_attraction_name",
								Address: "test_attraction_address",
								Coordinate: &models.CoordinateDTO{
									Id: "test_coordinate_id",
									X:  1,
									Y:  2,
								},
								Tags: []*models.TagDTO{
									{
										Id:    "tag_1",
										Value: "tag_value_1",
									},
									{
										Id:    "tag_2",
										Value: "tag_value_2",
									},
								},
                            },
                        },
                    },
                    nil,
                )
            },
            wantResp: &models.EventDTO{
                Id:          "test_event_id",
				ItineraryId: "test_itinerary_id",
                StartTime:   time.Now(),
                EndTime:     time.Now().Add(time.Hour),
                Attraction: &models.AttractionDTO{
                    Id:      "test_attraction_id",
                    Name:    "test_attraction_name",
                    Address: "test_attraction_address",
                    Coordinate: &models.CoordinateDTO{
                        Id: "test_coordinate_id",
                        X:  1,
                        Y:  2,
                    },
                    Tags: []*models.TagDTO{
                        {
                            Id:    "tag_1",
                            Value: "tag_value_1",
                        },
                        {
                            Id:    "tag_2",
                            Value: "tag_value_2",
                        },
                    },
                },
                Description: "test_description",
            },
            wantErr: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.before(t)
            gotResp, err := eventService.ConvertDBOToDTOEvent(tt.arg.ctx, tt.arg.req)

            assert.Equal(t, tt.wantResp, gotResp)
            assert.Equal(t, tt.wantErr, err)
        })
    }
}


