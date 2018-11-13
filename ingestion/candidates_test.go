package ingestion

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_candidateService_Retrieve(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/partner/candidates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		io.WriteString(w, `
		[
			{
				"id": 17681532,
				"name": "Harry Potter",
				"external_id": "24680",
				"applications": [
					{
						"id": 59724,
						"job": "Auror",
						"status": "Active",
						"stage": "Application Review",
						"profile_url": "https://app.greenhouse.io/people/17681532?application_id=26234709"
					}
				]
			}
		]
		`)
	})

	type fields struct {
		client *Client
	}
	type args struct {
		ids []int64
	}
	tests := []struct {
		name           string
		args           args
		wantCandidates []Candidate
		wantErr        bool
	}{
		{
			name: "Test parse",
			args: args{
				ids: []int64{12},
			},
			wantCandidates: []Candidate{
				Candidate{
					ID:         17681532,
					Name:       "Harry Potter",
					ExternalID: "24680",
					Applications: []Application{
						Application{
							ID:         59724,
							Job:        "Auror",
							Status:     "Active",
							Stage:      "Application Review",
							ProfileURL: "https://app.greenhouse.io/people/17681532?application_id=26234709",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCandidates, err := client.Candidates.Retrieve(tt.args.ids)

			switch tt.wantErr {
			case true:
				assert.Error(t, err)
			case false:
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantCandidates, gotCandidates)
		})
	}
}
