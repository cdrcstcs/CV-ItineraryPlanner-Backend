package test
import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"itineraryplanner/common/custom_errs"
	"itineraryplanner/dal/db"
	"itineraryplanner/models"
	"itineraryplanner/dal"
	"fmt"
)
func TestCreateAttraction(t *testing.T) {
	ctx := context.Background()
	type arg struct {
		attraction *models.Attraction
		ctx        context.Context
	}
	tests := []struct {
		name   			string
		before 			func(t *testing.T)
		arg
		wantErr        	error
		wantAttraction 	*models.Attraction
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
				attraction: &models.Attraction{
					Name: "test",
					X: 1,
					Y: 2,					
					Address: "123 Paya Lebar",
					TagIDs: []string{"1","2","3"},
				},
			},
			wantAttraction: &models.Attraction{
				Name: "test",
				X: 1,
				Y: 2,				
				Address: "123 Paya Lebar",
				TagIDs: []string{"1","2","3"},
			},
		},
		{
			name: "with id err",
			arg: arg{
				ctx: ctx,
				attraction: &models.Attraction{
					Id:   "1",
					Name: "test",
					X: 1,
					Y: 2,					
					Address: "123 Paya Lebar",
					TagIDs: []string{"1","2","3"},
				},
			},
			wantErr: custom_errs.DBErrCreateWithID,
		},
	}
	attractionDal := &dal.AttractionDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.CreateAttraction(tt.arg.ctx, tt.arg.attraction)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction.Id)
			gotAttraction.Id = ""
			assert.Equal(t, tt.wantAttraction, gotAttraction)
		})
	}
}
func TestGetAttractionById(t *testing.T) {
    ctx := context.Background()
    type arg struct {
        ctx           context.Context
        attractionId string
    }
    tests := []struct {
        name            string
        before          func(t *testing.T) string 
        arg
        wantAttraction *models.Attraction
        wantErr         error
    }{
        {
            name: "success",
            before: func(t *testing.T) string {
                attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo()}
                newAttraction := &models.Attraction{
                    Name:       "water slide",
                    Address:    "24 paya lebar",
                    X:          1,
                    Y:          2,
                    TagIDs:     []string{"1", "2", "3"},
                }
                createdAttraction, err := attractionDal.CreateAttraction(ctx, newAttraction)
                if err != nil {
                    t.Fatalf("failed to create attraction: %v", err)
                }
                fmt.Printf("Created Attraction ID: %s\n", createdAttraction.Id)
                return createdAttraction.Id
            },
            arg: arg{
                ctx:           ctx,
                attractionId:  "", 
            },
            wantAttraction: &models.Attraction{
                Id:         "", 
                Name:       "water slide",
                Address:    "24 paya lebar",
                X:          1,
                Y:          2,
                TagIDs:     []string{"1", "2", "3"},
            },
            wantErr: nil,
        },
        {
            name: "invalid id",
            arg: arg{
                ctx:           ctx,
                attractionId:  "123456789123456789123456", 
            },
            wantAttraction: nil,
            wantErr: custom_errs.DBErrGetWithID,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.before != nil {
                tt.arg.attractionId = tt.before(t)
            }
            fmt.Printf("Testing with Attraction ID: %s\n", tt.arg.attractionId)
            attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo()}
            gotAttraction, err := attractionDal.GetAttractionById(tt.arg.ctx, tt.arg.attractionId)
            fmt.Printf("Got Attraction: %+v\n", gotAttraction)
            fmt.Printf("Got Error: %v\n", err)
            assert.Equal(t, tt.wantErr, err)
            if err != nil {
                return
            }
            tt.wantAttraction.Id = tt.arg.attractionId
            assert.Equal(t, tt.wantAttraction, gotAttraction)
        })
    }
}
func TestGetAttraction(t *testing.T) {
	ctx := context.Background()
	type arg struct {
		ctx context.Context
	}
	attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo()}
	a, err := attractionDal.GetAttraction(ctx)
	if err != nil {
		t.Fatalf("Failed to get attractions: %v", err)
	}
	tests := []struct {
		name            string
		before          func(t *testing.T)
		arg
		wantAttraction []*models.Attraction
		wantErr         error
	}{
		{
			name: "success",
			arg: arg{
				ctx: ctx,
			},
			wantAttraction: a,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttraction, err := attractionDal.GetAttraction(tt.arg.ctx)
			fmt.Printf("Got Attraction: %+v\n", gotAttraction)
			fmt.Printf("Expected Attraction: %+v\n", tt.wantAttraction)
			fmt.Printf("Error: %v\n", err)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction)
			assert.ElementsMatch(t, tt.wantAttraction, gotAttraction)
		})
	}
}
func TestUpdateAttraction(t *testing.T) {
	ctx := context.Background()
	type arg struct {
		ctx           context.Context
		attraction    *models.Attraction
	}
	tests := []struct {
		name            string
		before          func(t *testing.T) string
		arg
		wantAttraction *models.Attraction
		wantErr         error
	}{
		{
			name: "success",
			before: func(t *testing.T) string {
				attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo()}
				newAttraction := &models.Attraction{
					Name:       "water slide",
					Address:    "24 paya lebar",
					X:          1,
					Y:          2,
					TagIDs:     []string{"1", "2", "3"},
				}
				createdAttraction, err := attractionDal.CreateAttraction(ctx, newAttraction)
				if err != nil {
					t.Fatalf("failed to create attraction: %v", err)
				}
				return createdAttraction.Id
			},
			arg: arg{
				ctx:           ctx,
				attraction: &models.Attraction{
					Id:         "", 
					Name:       "water slide 2",
					Address:    "24 paya lebar",
					X:          1,
					Y:          2,
					TagIDs:     []string{"1", "2", "3"},
				},
			},
			wantAttraction: &models.Attraction{
				Id:         "", 
				Name:       "water slide 2",
				Address:    "24 paya lebar",
				X:          1,
				Y:          2,
				TagIDs:     []string{"1", "2", "3"},
			},
			wantErr: nil,
		},
	}
	attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attractionID := tt.before(t)
			tt.arg.attraction.Id = attractionID
			tt.wantAttraction.Id = attractionID
			gotAttraction, err := attractionDal.UpdateAttraction(tt.arg.ctx, tt.arg.attraction)
			fmt.Printf("Got Updated Attraction: %+v\n", gotAttraction)
			fmt.Printf("Expected Updated Attraction: %+v\n", tt.wantAttraction)
			fmt.Printf("Error: %v\n", err)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.NotEmpty(t, gotAttraction)
			assert.Equal(t, tt.wantAttraction, gotAttraction)
		})
	}
}
func TestDeleteAttraction(t *testing.T) {
	ctx := context.Background()
	type arg struct {
		ctx           context.Context
		attractionId  string
	}
	tests := []struct {
		name            string
		before          func(t *testing.T) string
		arg
		wantAttraction *models.Attraction
		wantErr         error
	}{
		{
			name: "success",
			before: func(t *testing.T) string {
				attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo()}
				newAttraction := &models.Attraction{
					Name:       "water slide",
					Address:    "24 paya lebar",
					X:          1,
					Y:          2,
					TagIDs:     []string{"1", "2", "3"},
				}
				createdAttraction, err := attractionDal.CreateAttraction(ctx, newAttraction)
				if err != nil {
					t.Fatalf("failed to create attraction: %v", err)
				}
				return createdAttraction.Id
			},
			arg: arg{
				ctx:           ctx,
				attractionId:  "", 
			},
			wantAttraction: &models.Attraction{
				Id: 		"",		
				Name:       "water slide",
				Address:    "24 paya lebar",
				X:          1,
				Y:          2,
				TagIDs:     []string{"1", "2", "3"},
			}, 
			wantErr:         nil,
		},
	}
	attractionDal := dal.AttractionDal{MainDB: db.GetMemoMongo()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attractionID := tt.before(t)
			tt.arg.attractionId = attractionID
			gotAttraction, err := attractionDal.DeleteAttraction(tt.arg.ctx, tt.arg.attractionId)
			fmt.Printf("Got Deleted Attraction: %+v\n", gotAttraction)
			fmt.Printf("Error: %v\n", err)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			tt.wantAttraction.Id = tt.arg.attractionId
			assert.Equal(t, tt.wantAttraction, gotAttraction)
		})
	}
}