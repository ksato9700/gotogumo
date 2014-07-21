//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//

// Package gotogumo proivdes structs and functions for accessing soundcloud API.
package gotogumo

type User struct {
	Id                   int      `json:"id"`
	Kind                 string   `json:"kind"`
	Permalink            string   `json:"permalink"`
	Username             string   `json:"username"`
	Uri                  string   `json:"uri"`
	PermalinkUrl         string   `json:"permalink_url"`
	AvatarUrl            string   `json:"avatar_url"`
	Country              string   `json:"country"`
	FirstName            string   `json:"first_name"`
	LastName             string   `json:"last_name"`
	FullName             string   `json:"full_name"`
	Description          string   `json:"description"`
	City                 string   `json:"city"`
	DiscogsName          string   `json:"discogs_name"`
	MyspaceName          string   `json:"myspace_name"`
	Website              string   `json:"website"`
	WebsiteTitle         string   `json:"website_title"`
	Online               bool     `json:"online boolean"`
	TrackCount           int      `json:"track_count"`
	PlaylistCount        int      `json:"playlist_count"`
	Plan                 string   `json:"plan"`
	PublicFavoritesCount int      `json:"public_favorites_count"`
	FollowersCount       int      `json:"followers_count"`
	FollowingsCount      int      `json:"followings_count"`
	Subscriptions        []string `json:"subscriptions"`
	UploadSecondsLeft    int      `json:"upload_seconds_left"`
	Quota                struct {
		UnlimitedUploadQuota bool `json: "unlimited_upload_quota"`
		UploadSecondsUsed    int  `json: "upload_seconds_used"`
		UploadSecondsLeft    int  `json: "upload_seconds_left"`
	} `json:"quota"`
	PrivateTracksCount    int  `json:"private_tracks_count"`
	PrivatePlaylistsCount int  `json:"private_playlists_count"`
	PrimaryEmailConfirmed bool `json:"primary_email_confirmed`
}
