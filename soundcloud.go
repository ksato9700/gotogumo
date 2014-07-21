//
// Copyright 2014 Kenichi Sato <ksato9700@gmail.com>
//

// Package gotogumo proivdes structs and functions for accessing soundcloud API.
package gotogumo

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	AuthUrl    = "https://soundcloud.com/connect"
	ApiUrlBase = "https://api.soundcloud.com"
	TokenUrl   = ApiUrlBase + "/oauth2/token"
	CacheFile  = "cache.json"
)

type Client struct {
	HttpClient *http.Client
	Token      *oauth.Token
}

//var code string = "63b508a0c5d5968a69571d17633d542f"

var code string = ""

// Create a new instance of Client
func NewClient(cid string, csec string, redurl string) *Client {
	config := &oauth.Config{
		ClientId:     cid,
		ClientSecret: csec,
		Scope:        "non-expiring",
		AuthURL:      AuthUrl,
		TokenURL:     TokenUrl,
		RedirectURL:  redurl,
		TokenCache:   oauth.CacheFile(CacheFile),
	}
	//fmt.Printf("%+v\n", config)

	transport := &oauth.Transport{Config: config}
	token, err := config.TokenCache.Token()
	if err != nil {
		if code == "" {
			url := config.AuthCodeURL("")
			fmt.Println(url)
			return nil
		}
		token, err = transport.Exchange(code)
	}
	transport.Token = token

	return &Client{
		HttpClient: transport.Client(),
		Token:      token,
	}
}

//
func (c *Client) _ApiGet(path string, params url.Values, out interface{}) error {
	if params == nil {
		params = url.Values{}
	}
	params.Add("oauth_token", c.Token.AccessToken)
	urlstring := ApiUrlBase + path + "?" + params.Encode()
	// fmt.Println("urlstring", urlstring)
	res, err := c.HttpClient.Get(urlstring)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if strout, ok := out.(*string); ok {
		*strout = string(data)
	} else {
		err = json.Unmarshal(data, out)
	}
	return err
}

//
func (c *Client) Get(resource string, out interface{}) error {
	return c._ApiGet(resource+".json", nil, out)
}

//
func (c *Client) GetMe() (User, error) {
	var me User
	err := c.Get("/me", &me)
	return me, err
}
