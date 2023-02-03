package rates

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"log"
	"reflect"
	"testing"
)

func TestRepository_UpdateCustomer(t *testing.T) {

	type args struct {
		state        string
		estimateType string
	}

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "NY premium",
			args: args{
				state:        "NY",
				estimateType: "premium",
			},
			want: 0.35,
		},
		{
			name: "NY normal",
			args: args{
				state:        "NY",
				estimateType: "normal",
			},
			want: 0.25,
		},
		{
			name: "CA premium",
			args: args{
				state:        "CA",
				estimateType: "premium",
			},
			want: 0.33,
		},
		{
			name: "CA normal",
			args: args{
				state:        "CA",
				estimateType: "normal",
			},
			want: 0.23,
		},
		{
			name: "AZ premium",
			args: args{
				state:        "AZ",
				estimateType: "premium",
			},
			want: 0.3,
		},
		{
			name: "AZ normal",
			args: args{
				state:        "AZ",
				estimateType: "normal",
			},
			want: 0.2,
		},
		{
			name: "TX premium",
			args: args{
				state:        "TX",
				estimateType: "premium",
			},
			want: 0.28,
		},
		{
			name: "TX normal",
			args: args{
				state:        "TX",
				estimateType: "normal",
			},
			want: 0.18,
		},
		{
			name: "OH premium",
			args: args{
				state:        "OH",
				estimateType: "premium",
			},
			want: 0.25,
		},
		{
			name: "OH normal",
			args: args{
				state:        "OH",
				estimateType: "normal",
			},
			want: 0.15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := sql.Open("sqlite3", "../../resources/rates.db")

			if err != nil {
				log.Fatal(err)
			}

			defer func(db *sql.DB) {
				if err = db.Close(); err != nil {
					t.Errorf("ratesRepo.GetRates() error = %v", err)
				}
			}(db)

			repo := NewRepository(db)

			res, err := repo.GetRates(tt.args.state, tt.args.estimateType)
			if err != nil {
				t.Errorf("ratesRepo.GetRates() error = %v", err)
			}

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("GetRates() = %v, want %v", res, tt.want)
			}

		})

	}

}
