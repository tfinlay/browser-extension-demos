package store

import (
	"errors"
	"fmt"
)

var ErrUnknownHostname = errors.New("no login entry exists for the given hostname")

type Login struct {
	Username string
	Password string
}

// logins is a map from host -> login entry
var logins = map[string]Login{
	"127.0.0.1:8000": {
		Username: "my.username",
		Password: "supersecretpassword",
	},
}

type LoginManager struct{}

func NewLoginManager() *LoginManager {
	return &LoginManager{}
}

func (p *LoginManager) GetLogin(host string) (Login, error) {
	if entry, ok := logins[host]; ok {
		return entry, nil
	}
	return Login{}, fmt.Errorf("failed to get password for %q: %w", host, ErrUnknownHostname)
}
