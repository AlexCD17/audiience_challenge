package services

import (
	"audiience_challenge/entities"
	"audiience_challenge/mocks"
	"github.com/golang/mock/gomock"

	"testing"
)

func TestEstimateService_Estimate(t *testing.T) {

	type args struct {
		input entities.InquiryModel
	}

	tests := []struct {
		name         string
		args         args
		repoResponse float32
		want         float32
	}{
		{
			name: "NY premium",
			args: args{
				input: entities.InquiryModel{
					State:      "NY",
					Type:       "premium",
					Distance:   22,
					BaseAmount: 342,
				},
			},
			repoResponse: 0.35,
			want:         558.657,
		},
		{
			name: "NY normal",
			args: args{
				input: entities.InquiryModel{
					State:      "NY",
					Type:       "normal",
					Distance:   34,
					BaseAmount: 342,
				},
			},
			repoResponse: 0.25,
			want:         517.275,
		},
		{
			name: "TX premium + 5% discount",
			args: args{
				input: entities.InquiryModel{
					State:      "TX",
					Type:       "premium",
					Distance:   34,
					BaseAmount: 400,
				},
			},
			repoResponse: 0.28,
			want:         486.4,
		},
		{
			name: "TX normal + 5% discount",
			args: args{
				input: entities.InquiryModel{
					State:      "TX",
					Type:       "normal",
					Distance:   34,
					BaseAmount: 400,
				},
			},
			repoResponse: 0.18,
			want:         448.4,
		},
		{
			name: "CA premium",
			args: args{
				input: entities.InquiryModel{
					State:      "CA",
					Type:       "premium",
					Distance:   34,
					BaseAmount: 530,
				},
			},
			repoResponse: 0.33,
			want:         669.655,
		},
		{
			name: "OH normal + 3% discount",
			args: args{
				input: entities.InquiryModel{
					State:      "OH",
					Type:       "normal",
					Distance:   24,
					BaseAmount: 333,
				},
			},
			repoResponse: 0.15,
			want:         372.96002,
		},
		{
			name: "AZ normal",
			args: args{
				input: entities.InquiryModel{
					State:      "AZ",
					Type:       "normal",
					Distance:   20,
					BaseAmount: 200,
				},
			},
			repoResponse: 0.20,
			want:         240,
		},
		{
			name: "AZ normal + 5% discount",
			args: args{
				input: entities.InquiryModel{
					State:      "AZ",
					Type:       "normal",
					Distance:   27,
					BaseAmount: 342,
				},
			},
			repoResponse: 0.20,
			want:         389.88,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//service mock setup
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mocks.NewMockIRepository(mockCtrl)
			mockRepo.EXPECT().GetRates(tt.args.input.State, tt.args.input.Type).Return(tt.repoResponse, nil)
			service := NewEstimateService(mockRepo)

			res, _ := service.Estimate(tt.args.input)
			if res != tt.want {
				t.Errorf("estimateService.Estimate() Result: %v, wanted: %v", res, tt.want)
			}

		})
	}
}
