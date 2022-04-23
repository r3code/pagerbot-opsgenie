package opsgenie

import (
	opsgenie "github.com/opsgenie/opsgenie-go-sdk-v2"
)

type Api struct {
	key      string
	org      string
	client   *opsgenie.Client
	timezone string
}

// ?? API doesn't provide a sane way of checking for auth
// so we just get the schedules at setup time
func New(key string, org string) (*Api, error) {
	a := Api{}
	a.key = key
	a.org = org
	a.timezone = "UTC"

	a.client = opsgenie.NewClient(key)

	_, err := a.client.ListSchedules(opsgenie.ListSchedulesOptions{})
	if err != nil {
		return &a, err
	}

	return &a, nil
}
