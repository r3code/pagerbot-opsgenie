package opsgenie

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"
)

type Users []User

type User struct {
  Id string // uuid
	Email string // see username at https://docs.opsgenie.com/docs/user-api#list-user
	FullName  string
}

func (a *Api) Users() (Users, error) {
	var usr Users
	var opts user.ListRequest

	for {
		res, err := a.client.ListUsers(opts)
		if err != nil {
			return usr, err
		}

		for _, user := range res.Users {
			usr = append(usr, User{
				Id:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			})
		}

		if !res.More {
			break
		}

		opts.Offset = res.Offset + res.Limit
	}

	return usr, nil
}
