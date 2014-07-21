package gotogumo

import (
	"os"
	"testing"
)

var cid = os.Getenv("SC_CLIENT_ID")
var csec = os.Getenv("SC_CLIENT_SECRET")
var redurl = os.Getenv("SC_REDIRECT_URI")

func TestClient(t *testing.T) {
	client := NewClient(cid, csec, redurl)
	if client == nil {
		t.Error("client is nil")
	}
	me, err := client.GetMe()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", me)

	// var tracks string
	var tracks Tracks
	err = client.Get("/tracks", &tracks)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", tracks)
}
