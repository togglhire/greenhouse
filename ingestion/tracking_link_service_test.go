package ingestion

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trackingLinkService_Post(t *testing.T) {
	setup()
	defer teardown()

	type args struct {
		jobID int64
	}
	test := struct {
		args             args
		reqBody          string
		wantTrackingLink PostTrackingLinkResponse
		wantErr          bool
	}{
		reqBody: `{
			"job_id": 64
		}`,
		args: args{
			jobID: 64,
		},
		wantTrackingLink: PostTrackingLinkResponse{
			TrackingLink: "http://grnh.se/yvj0bj",
			Job:          "Auror",
			Source:       "Campus Recruiting",
			Referrer:     "Hermione Granger",
		},
	}

	mux.HandleFunc("/v1/partner/tracking_link", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		body := formatReadCloser(&r.Body)
		equal, err := areEqualJSON(test.reqBody, body)
		assert.NoError(t, err)
		if !equal {
			o1, err := jsonStringAsInterface(test.reqBody)
			assert.NoError(t, err)
			o2, err := jsonStringAsInterface(body)
			assert.NoError(t, err)
			assert.Equal(t, o1, o2) //just to get the diff
		}
		w.WriteHeader(200)
		io.WriteString(w, `
		{ 
			"tracking_link": "http://grnh.se/yvj0bj",
			"job": "Auror", 
			"source": "Campus Recruiting",
			"referrer": "Hermione Granger"
		}
		`)
	})

	gotTrackingLink, err := client.TrackingLinks.Post(test.args.jobID)

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantTrackingLink, gotTrackingLink)

}
