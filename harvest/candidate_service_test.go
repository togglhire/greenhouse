package harvest

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestCandidateService_List(t *testing.T) {
	setup(TEST_API_KEY, TEST_ON_BEHALF_OF, t)
	defer teardown()

	tests := []struct {
		name    string
		params  *CandidateListParams
		want    []Candidate
		wantErr bool
	}{
		{
			name: "List candidates",
			params: &CandidateListParams{
				UpdatedAfter: "2019-01-01T00:00:00Z",
				PerPage:      100,
			},
			want: []Candidate{
				{
					Id:        17681532,
					FirstName: "John",
					LastName:  "Doe",
					Company:   "IETF",
				},
				{
					Id:        23881535,
					FirstName: "Jane",
					LastName:  "Doe",
					Company:   "IETF",
				},
			},
			wantErr: false,
		},
		{
			name:   "List candidates with no params",
			params: &CandidateListParams{},
			want: []Candidate{
				{
					Id:        17681532,
					FirstName: "John",
					LastName:  "Doe",
					Company:   "IETF",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/v1/candidates", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			payload, err := json.Marshal(tt.want)
			if err != nil {
				t.Fatalf("CandidateService.List() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Candidates.List(*tt.params)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.List() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CandidateService.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
