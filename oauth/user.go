package oauth

import (
	"github.com/sirupsen/logrus"

	"github.com/Shopify/goose/statsd"
)

// User is a retrieved and authenticated user.
// Corresponds to https://www.googleapis.com/oauth2/v3/userinfo, but it meant to be used as an abstract user
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Domain        string `json:"hd,omitempty"`
}

func (u *User) LogFields() logrus.Fields {
	return logrus.Fields{
		"email": u.Email,
	}
}

func (u *User) StatsTags() statsd.Tags {
	return statsd.Tags{}
}
