package users

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
)

func (iu *InfoUsers) Run() (value map[string]interface{}) {
	key := "users"
	iu.Users = make([]*user.User, 0)

	// Open the /etc/passwd file
	f, err := os.Open("/etc/passwd")
	if err != nil {
		return map[string]interface{}{key: err}
	}
	defer f.Close()

	// Build a list of the users
	users := []string{}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			return map[string]interface{}{key: err}
		}

		if eq := strings.HasPrefix(line, "#"); !eq {
			parts := strings.FieldsFunc(line, func(marker rune) bool {
				return marker == ':'
			})

			if len(parts) > 0 {
				users = append(users, parts[0])
			}
		}
	}

	// Get the users using the os.Users Lookup function
	for _, username := range users {
		usr, err := user.Lookup(username)
		if err != nil {
			log.Printf("[WARN] Error looking up user %s: %w\n", username, err)
		}

		iu.Users = append(iu.Users, usr)
	}

	return map[string]interface{}{key: iu.Users}
}
