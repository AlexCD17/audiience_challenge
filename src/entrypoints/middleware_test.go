package entrypoints

import (
	"audiience_challenge/mocks"
	"audiience_challenge/resources"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestMiddleware_VerifyMiddleware(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantHeader http.Header
		wantBody   string
	}{
		{
			name: "failed estimation, missing state",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?type=premium&distance=23&base_amount=345",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.MissingState,
		},
		{
			name: "failed estimation, wrong state",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?state=bd&type=premium&distance=23&base_amount=345",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.WrongState,
		},
		{
			name: "failed estimation, unsupported state",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?state=DS&type=premium&distance=23&base_amount=345",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.UnsupportedState,
		},
		{
			name: "failed estimation, missing distance",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?state=NY&type=premium&base_amount=345",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.MissingDistance,
		},
		{
			name: "failed estimation, wrong distance type",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?state=NY&type=premium&distance=fsdf&base_amount=345",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.WrongDistance,
		},
		{
			name: "failed estimation, missing estimation type",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?state=NY&distance=23&base_amount=345",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.MissingType,
		},
		{
			name: "failed estimation, wrong estimation type",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?state=NY&type=quadratic&distance=23&base_amount=345",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.WrongType,
		},
		{
			name: "failed estimation, missing base amount",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"http://0.0.0.0/estimate?state=NY&type=premium&distance=23",
					strings.NewReader(""),
				),
			},
			wantStatus: http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   resources.MissingBaseAmount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := mocks.NewMockIService(mockCtrl)

			router := mux.NewRouter().StrictSlash(true)
			server := NewServer(mockService, router)
			server.SetupRouter()
			rec := tt.args.w.(*httptest.ResponseRecorder)
			tt.args.r.Header.Add("ip-client", "127.0.0.0")

			router.ServeHTTP(rec, tt.args.r)

			res := rec.Result()
			if !reflect.DeepEqual(res.StatusCode, tt.wantStatus) {
				t.Errorf("entrypoints.GetCredit() status code error,  got %v, wanted: %v", res.StatusCode, tt.wantStatus)
			}

			bodyBuffer := new(bytes.Buffer)
			_, _ = bodyBuffer.ReadFrom(res.Body)
			body := strings.TrimSpace(bodyBuffer.String())
			if body[1:len(body)-1] != tt.wantBody {
				t.Errorf("entrypoints.GetCredit() wrong response, got %v, wanted: %v", body, tt.wantBody)
			}

		})
	}

}
