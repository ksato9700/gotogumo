//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//

// Package gotogumo proivdes structs and functions for accessing soundcloud API.
package gotogumo

type Tracks []Track

type Track struct {
	ArtworkUrl          string `json:"artwork_url"`
	AttachmentsUri      string `json:"attachments_uri"`
	Bpm                 string `json:"bpm"`
	CommentCount        int    `json:"comment_count"`
	Commentable         bool   `json:"commentable"`
	CreatedAt           string `json:"created_at"`
	Description         string `json:"description"`
	DownloadCount       int    `json:"download_count"`
	Downloadable        bool   `json:"downloadable"`
	Duration            int    `json:"duration"`
	EmbeddableBy        string `json:"embeddable_by"`
	FavoritingsCount    int    `json:"favoritings_count"`
	Genre               string `json:"genre"`
	Id                  int    `json:"id"`
	Isrc                string `json:"isrc"`
	KeySignature        string `json:"key_signature"`
	Kind                string `json:"kind"`
	LabelId             string `json:"label_id"`
	LabelName           string `json:"label_name"`
	License             string `json:"license"`
	OriginalContentSize int    `json:"original_content_size"`
	OriginalFormat      string `json:"original_format"`
	Permalink           string `json:"permalink"`
	PermalinkUrl        string `json:"permalink_url"`
	PlaybackCount       int    `json:"playback_count"`
	PurchaseTitle       string `json:"purchase_title"`
	PurchaseUrl         string `json:"purchase_url"`
	Release             string `json:"release"`
	ReleaseDay          string `json:"release_day"`
	ReleaseMonth        string `json:"release_month"`
	ReleaseYear         string `json:"release_year"`
	Sharing             string `json:"sharing"`
	State               string `json:"state"`
	StreamUrl           string `json:"stream_url"`
	Streamable          bool   `json:"streamable"`
	TagList             string `json:"tag_list"`
	Title               string `json:"title"`
	TrackType           string `json:"track_type"`
	Uri                 string `json:"uri"`
	User                struct {
		AvatarUrl    string `json:"avatar_url"`
		Id           int    `json:"id"`
		Kind         string `json:"kind"`
		Permalink    string `json:"permalink"`
		PermalinkUrl string `json:"permalink_url"`
		Uri          string `json:"uri"`
		Username     string `json:"username"`
	} `json:"user"`
	UserFavorite      bool   `json:"user_favorite"`
	UserId            int    `json:"user_id"`
	UserPlaybackCount int    `json:"user_playback_count"`
	VideoUrl          string `json:"video_url"`
	WaveformUrl       string `json:"waveform_url"`
}
