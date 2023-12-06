package harvest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCandidateService_List(t *testing.T) {
	tests := []struct {
		name           string
		responseStatus int
		params         CandidateListParams
		want           []Candidate
		wantErr        bool
	}{
		{
			name:           "List candidates",
			responseStatus: http.StatusOK,
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
		},
		{
			name:           "List candidates with no params",
			responseStatus: http.StatusOK,
			params:         CandidateListParams{},
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
		{
			name:           "List candidates response error",
			responseStatus: http.StatusInternalServerError,
			params:         CandidateListParams{},
			want:           nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/candidates", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("CandidateService.List() request method = %v, want %v", r.Method, http.MethodGet)
			}
			w.WriteHeader(tt.responseStatus)
			if tt.responseStatus != http.StatusOK {
				return
			}
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
	candidateId := int64(17681532)
	tests := []struct {
		name           string
		responseStatus int
		id             int64
		want           *Candidate
		wantErr        bool
	}{
		{
			name:           "Retrieve candidate",
			responseStatus: http.StatusOK,
			id:             candidateId,
			want: &Candidate{
				Id:        candidateId,
				FirstName: "John",
				LastName:  "Doe",
				Company:   "IETF",
			},
			wantErr: false,
		},
		{
			name:           "Retrieve candidate response error",
			responseStatus: http.StatusInternalServerError,
			id:             candidateId,
			want:           &Candidate{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/candidates/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("CandidateService.Retrieve() request method = %v, want %v", r.Method, http.MethodGet)
			}
			w.WriteHeader(tt.responseStatus)
			if tt.responseStatus != http.StatusOK {
				return
			}
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
		name           string
		responseStatus int
		candidate      *Candidate
		wantErr        bool
	}{
		{
			name:           "Add candidate",
			responseStatus: http.StatusOK,
			candidate: &Candidate{
				FirstName: "John",
				LastName:  "Doe",
				Company:   "IETF",
			},
			wantErr: false,
		},
		{
			name:           "Add candidate with no first name",
			responseStatus: http.StatusUnprocessableEntity,
			candidate: &Candidate{
				LastName: "Doe",
				Company:  "IETF",
			},
			wantErr: true,
		},
		{
			name:           "Add candidate response error",
			responseStatus: http.StatusInternalServerError,
			candidate:      nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc("/v1/candidates", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("CandidateService.Add() request method = %v, want %v", r.Method, http.MethodPost)
			}
			w.WriteHeader(tt.responseStatus)
			if !tt.wantErr {
				payload, err := json.Marshal(tt.candidate)
				if err != nil {
					t.Fatalf("CandidateService.Add() error = %v", err)
				}
				w.Write(payload)
			} else {
				w.Write([]byte(`{}`))
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			err := client.Candidates.Add(tt.candidate)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCandidateService_Edit(t *testing.T) {
	candidateId := int64(17681532)
	candidateFirstName := "John"
	candidateLastName := "Doe"

	tests := []struct {
		name           string
		responseStatus int
		candidate      *Candidate
		wantErr        bool
	}{
		{
			name:           "Edit candidate",
			responseStatus: http.StatusOK,
			candidate: &Candidate{
				Title:   "Senior Software Engineer",
				Company: "NewCompany Co.",
				PhoneNumbers: []KeyValue[PhoneNumberType]{
					{
						Type:  PNHome,
						Value: "555-555-5555",
					},
				},
				Addresses: []KeyValue[AddressType]{
					{
						Type:  ATHome,
						Value: "123 Main St.",
					},
					{
						Type:  ATOther,
						Value: "321 Corner St.",
					},
				},
				SocialMediaAddresses: []KeyValue[string]{
					{
						Type:  "twitter",
						Value: "@johndoe-123",
					},
				},
			},
			wantErr: false,
		},
		{
			name:           "Edit candidate response error",
			responseStatus: http.StatusInternalServerError,
			candidate:      nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc(fmt.Sprintf("/v1/candidates/%d", candidateId), func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("CandidateService.Edit() request method = %v, want %v", r.Method, http.MethodPatch)
			}
			w.WriteHeader(tt.responseStatus)
			if tt.responseStatus != http.StatusOK {
				return
			}
			returnCandidate := tt.candidate
			returnCandidate.FirstName = candidateFirstName
			returnCandidate.LastName = candidateLastName

			payload, err := json.Marshal(tt.candidate)
			if err != nil {
				t.Fatalf("CandidateService.Edit() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			err := client.Candidates.Edit(candidateId, tt.candidate)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && tt.candidate.FirstName != candidateFirstName {
				t.Errorf("CandidateService.Edit() error = %v, expected FirstName to be John", err)
			}

			if !tt.wantErr && tt.candidate.LastName != candidateLastName {
				t.Errorf("CandidateService.Edit() error = %v, expected LastName to be Doe", err)
			}
		})
	}
}

func TestCandidateService_AddAttachment(t *testing.T) {
	candidateId := int64(17681532)

	tests := []struct {
		name           string
		responseStatus int
		attachment     *Attachment
		wantErr        bool
	}{
		{
			name:           "Add attachment",
			responseStatus: http.StatusOK,
			attachment: &Attachment{
				Type:     "test.pdf",
				Filename: string(ATResume),
				URL:      "https://www.johndoe-123.com/test.pdf",
			},
			wantErr: false,
		},
		{
			name:           "Add attachment response error",
			responseStatus: http.StatusInternalServerError,
			attachment:     nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc(fmt.Sprintf("/v1/candidates/%d/attachments", candidateId), func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("CandidateService.AddAttachment() request method = %v, want %v", r.Method, http.MethodPost)
			}
			w.WriteHeader(tt.responseStatus)
			if tt.responseStatus != http.StatusOK {
				return
			}
			payload, err := json.Marshal(tt.attachment)
			if err != nil {
				t.Fatalf("CandidateService.AddAttachment() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			err := client.Candidates.AddAttachment(candidateId, tt.attachment)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.AddAttachment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCandidateService_AddNote(t *testing.T) {
	candidateId := int64(17681532)

	tests := []struct {
		name           string
		responseStatus int
		note           *Note
		wantErr        bool
	}{
		{
			name:           "Add Note",
			responseStatus: http.StatusOK,
			note: &Note{
				UserId:     candidateId,
				Body:       "This is a test note",
				Visibility: NVPublic,
			},
			wantErr: false,
		},
		{
			name:           "Add Note response error",
			responseStatus: http.StatusInternalServerError,
			note:           nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		setup(TEST_API_KEY, TEST_ON_BEHALF_OF, "", t)
		defer teardown()

		mux.HandleFunc(fmt.Sprintf("/v1/candidates/%d/activity_feed/notes", candidateId), func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("CandidateService.AddNote() request method = %v, want %v", r.Method, http.MethodPost)
			}
			w.WriteHeader(tt.responseStatus)
			if tt.responseStatus != http.StatusOK {
				return
			}
			payload, err := json.Marshal(tt.note)
			if err != nil {
				t.Fatalf("CandidateService.AddNote() error = %v", err)
			}
			w.Write(payload)
		})

		t.Run(tt.name, func(t *testing.T) {
			err := client.Candidates.AddNote(candidateId, tt.note)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CandidateService.AddNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
