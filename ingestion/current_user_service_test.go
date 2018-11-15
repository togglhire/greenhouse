package ingestion

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_currentUserService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	test := struct {
		wantUser User
		wantErr  bool
	}{
		wantUser: User{
			FirstName: "Ron",
			LastName:  "Weasley",
			Email:     "rweasley@hogwarts.edu",
		},
	}

	mux.HandleFunc("/v1/partner/current_user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(200)
		io.WriteString(w, `
		{ 
			"first_name": "Ron", 
			"last_name": "Weasley", 
			"email": "rweasley@hogwarts.edu"
		}
		`)
	})

	gotUser, err := client.CurrentUser.Retrieve()

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantUser, gotUser)
}
