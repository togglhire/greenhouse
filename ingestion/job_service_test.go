package ingestion

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jobService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	test := struct {
		wantJobs []Job
		wantErr  bool
	}{
		wantJobs: []Job{
			Job{
				ID:     146859,
				Name:   "Auror",
				Status: "open",
				Public: true,
			},
			Job{
				ID:     150050,
				Name:   "Professor",
				Status: "open",
				Public: true,
			},
			Job{
				ID:     147886,
				Name:   "Caretaker",
				Status: "open",
				Public: false,
			},
		},
	}

	mux.HandleFunc("/v1/partner/jobs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		io.WriteString(w, `
		[
			{
			  "id": 146859,
			  "name": "Auror",
			  "status": "open",
			  "public": true
			},
			{
			  "id": 150050,
			  "name": "Professor",
			  "status": "open",
			  "public": true
			},
			{
			  "id": 147886,
			  "name": "Caretaker",
			  "status": "open",
			  "public": false
			}
		  ]
		`)
	})

	gotJobs, err := client.Jobs.Retrieve()

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantJobs, gotJobs)
}
