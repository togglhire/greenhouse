package harvest

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestJobService_List(t *testing.T) {
	tests := []struct {
		name           string
		responseStatus int
		params         JobListParams
		want           []Job
		wantErr        bool
	}{
		{
			name:           "List jobs",
			responseStatus: http.StatusOK,
			params: JobListParams{
				UpdatedAfter: "2019-01-01T00:00:00Z",
				PerPage:      100,
			},
			want: []Job{
				{
					Id:   17681532,
					Name: "Software Engineer",
				},
				{
					Id:   23881535,
					Name: "UX Engineer",
				},
			},
		},
		{
			name:           "List jobs with no params",
			responseStatus: http.StatusOK,
			params:         JobListParams{},
			want: []Job{
				{
					Id:   17681532,
					Name: "Software Engineer",
				},
			},
			wantErr: false,
		},
		{
			name:           "List jobs response error",
			responseStatus: http.StatusInternalServerError,
			params:         JobListParams{},
			want:           nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/jobs", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("JobService.List() request method = %v, want %v", r.Method, http.MethodGet)
			}
			w.WriteHeader(tt.responseStatus)
			if tt.responseStatus != http.StatusOK {
				return
			}
			payload, err := json.Marshal(tt.want)
			if err != nil {
				t.Fatalf("JobService.List() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Jobs.List(tt.params)
			if (err != nil) != tt.wantErr {
				t.Fatalf("JobService.List() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JobService.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobService_Retrieve(t *testing.T) {
	jobId := int64(17681532)
	tests := []struct {
		name           string
		responseStatus int
		id             int64
		want           *Job
		wantErr        bool
	}{
		{
			name:           "Retrieve job",
			responseStatus: http.StatusOK,
			id:             jobId,
			want: &Job{
				Id:   jobId,
				Name: "Software Engineer",
			},
			wantErr: false,
		},
		{
			name:           "Retrieve job response error",
			responseStatus: http.StatusInternalServerError,
			id:             jobId,
			want:           nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/jobs/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("JobService.Retrieve() request method = %v, want %v", r.Method, http.MethodGet)
			}
			w.WriteHeader(tt.responseStatus)
			if tt.responseStatus != http.StatusOK {
				return
			}
			payload, err := json.Marshal(tt.want)
			if err != nil {
				t.Fatalf("JobService.Retrieve() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Jobs.Retrieve(tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("JobService.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JobService.Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}
