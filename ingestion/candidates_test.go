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
	test := struct {
		name           string
		args           args
		wantCandidates []Candidate
		wantErr        bool
	}{

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
	}

	gotCandidates, err := client.Candidates.Retrieve(test.args.ids)

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantCandidates, gotCandidates)

}

func Test_candidateService_Retrieve_client_error(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/partner/candidates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(400)
		io.WriteString(w, `
		{
			"errors": [
				{
					"message": "Your request included invalid JSON.",
					"field": "email"
				}
			]
		}
		`)
	})

	type fields struct {
		client *Client
	}
	type args struct {
		ids []int64
	}
	test := struct {
		name           string
		args           args
		wantCandidates []Candidate
		wantErr        bool
		wantErrorType  error
	}{
		name: "Error",
		args: args{
			ids: []int64{1, 2},
		},
		wantErr:       true,
		wantErrorType: ClientError{},
	}

	gotCandidates, err := client.Candidates.Retrieve(test.args.ids)
	switch test.wantErr {
	case true:
		assert.Error(t, err)
		assert.IsType(t, test.wantErrorType, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantCandidates, gotCandidates)

	// extra assertion
	clientError, _ := IsClientError(err)
	assert.Equal(t, ClientError{
		Errors: []Error{
			Error{
				Message: "Your request included invalid JSON.",
				Field:   "email",
			},
		},
	}, clientError)
}

func Test_candidateService_Retrieve_server_error(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/partner/candidates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(500)
		io.WriteString(w, `
		{
			"errors": [
				{
					"message": "There was a server error.",
					"field": ""
				}
			]
		}
		`)
	})

	type fields struct {
		client *Client
	}
	type args struct {
		ids []int64
	}
	test := struct {
		name           string
		args           args
		wantCandidates []Candidate
		wantErr        bool
		wantErrorType  error
	}{
		name: "Error",
		args: args{
			ids: []int64{1, 2},
		},
		wantErr:       true,
		wantErrorType: ServerError{},
	}

	gotCandidates, err := client.Candidates.Retrieve(test.args.ids)
	switch test.wantErr {
	case true:
		assert.Error(t, err)
		assert.IsType(t, test.wantErrorType, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantCandidates, gotCandidates)

	// extra assertion
	serverError, _ := IsServerError(err)
	assert.Equal(t, ServerError{
		Errors: []Error{
			Error{
				Message: "There was a server error.",
				Field:   "",
			},
		},
	}, serverError)
}
