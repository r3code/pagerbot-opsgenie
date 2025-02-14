package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Validate the configuration file for sanity
func (c *config) Validate() error {
	if c.ApiKeys.Slack == "" || c.ApiKeys.OpsGenie.Key == "" {
		return fmt.Errorf("You must provide API keys for Slack and OpsGenie")
	}

	if c.ApiKeys.OpsGenie.Org == "" {
		return fmt.Errorf("You must provide an org name for OpsGenie (<org>.opsgenie.com)")
	}

	if len(c.Groups) == 0 {
		return fmt.Errorf("You must specify at least one group")
	}

	for i, group := range c.Groups {
		if group.Name == "" {
			return fmt.Errorf("Must specify group name for group %d", i)
		}

		if len(group.Schedules) == 0 {
			return fmt.Errorf("Must specify at least one schedule for group %s", group.Name)
		}
	}

	log.WithFields(log.Fields{"groups": len(c.Groups)}).Debug("Loaded config")

	return nil
}
