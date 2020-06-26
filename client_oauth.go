package api

// ClientOAuth ...
type ClientOAuth struct {
	ResponseType        string
	ClientID            string
	Scope               string
	State               string
	CodeChallenge       string
	CodeChallengeMethod string
}
