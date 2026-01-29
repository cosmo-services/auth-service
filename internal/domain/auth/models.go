package auth

import "time"

type TokenPair struct {
	Access  JwtToken `json:"access_token"`
	Refresh JwtToken `json:"refresh_token"`
}

type JwtToken struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
