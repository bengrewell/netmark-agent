package users

import (
	"github.com/StackExchange/wmi"
	"log"
	"os/user"
)

type Win32_UserAccount struct {
	Name string
}

func (iu *InfoUsers) Run() (value map[string]interface{}) {
	key := "users"

	var users []Win32_UserAccount
	err := wmi.Query("Select Name from Win32_UserAccount where LocalAccount=True", &users)
	if err != nil {
		return map[string]interface{}{key: err}
	}
	for _, u := range users {
		usr, err := user.Lookup(u.Name)
		if err != nil {
			log.Printf("[WARN] Error looking up user %s: %w\n", u.Name, err)
		}
		iu.Users = append(iu.Users, usr)
	}

	return map[string]interface{}{key: iu.Users}
}
