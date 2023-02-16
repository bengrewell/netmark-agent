package users

import "os/user"

type InfoUsers struct {
	Users []*user.User `json:"users" yaml:"users"`
}
