package harvest

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestCandidateService_List(t *testing.T) {
	tests := []struct {
		name    string
		params  CandidateListParams
		want    []Candidate
		wantErr bool
	}{
		{
			name: "List candidates",
			params: CandidateListParams{
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
			params: CandidateListParams{},
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
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/candidates", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("CandidateService.List() request method = %v, want %v", r.Method, http.MethodGet)
			}
			w.WriteHeader(http.StatusOK)
			payload, err := json.Marshal(tt.want)
			if err != nil {
				t.Fatalf("CandidateService.List() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Candidates.List(tt.params)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.List() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CandidateService.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCandidateService_Retrieve(t *testing.T) {
	tests := []struct {
		name    string
		id      int64
		want    *Candidate
		wantErr bool
	}{
		{
			name: "Retrieve candidate",
			id:   17681532,
			want: &Candidate{
				Id:        17681532,
				FirstName: "John",
				LastName:  "Doe",
				Company:   "IETF",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/candidates/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("CandidateService.Retrieve() request method = %v, want %v", r.Method, http.MethodGet)
			}
			w.WriteHeader(http.StatusOK)
			payload, err := json.Marshal(tt.want)
			if err != nil {
				t.Fatalf("CandidateService.Retrieve() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Candidates.Retrieve(tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CandidateService.Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCandidateService_Add(t *testing.T) {
	tests := []struct {
		name        string
		candidate   *Candidate
		wantErr     bool
		expectedErr error
	}{
		// {
		// 	name: "Add candidate",
		// 	candidate: &Candidate{
		// 		FirstName: "John",
		// 		LastName:  "Doe",
		// 		Company:   "IETF",
		// 	},
		// 	wantErr:     false,
		// 	expectedErr: nil,
		// },
		{
			name: "Add candidate with no first name",
			candidate: &Candidate{
				LastName: "Doe",
				Company:  "IETF",
			},
			wantErr:     true,
			expectedErr: &ValidationError{}, // Couldn't use this;
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/candidates", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("CandidateService.Add() request method = %v, want %v", r.Method, http.MethodPost)
			}
			if !tt.wantErr {
				w.WriteHeader(http.StatusOK)
				payload, err := json.Marshal(tt.candidate)
				if err != nil {
					t.Fatalf("CandidateService.Add() error = %v", err)
				}
				w.Write(payload)
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(`{}`))
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			err := client.Candidates.Add(tt.candidate)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				var validationErr *ValidationError
				if !errors.As(err, &validationErr) {
					t.Errorf("CandidateService.Add() error = %v, expectedErr %v", err, tt.expectedErr)
				}
			}
		})
	}
}
