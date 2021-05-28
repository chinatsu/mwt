package actions

import (
	"encoding/json"
	"fmt"

	"github.com/chinatsu/mwt/client"
	"github.com/chinatsu/mwt/config"

	log "github.com/sirupsen/logrus"
)

const (
	FlexiTimeCode = "Fleksitid"
)

type Status struct {
	Data Data
}

type Data struct {
	Registrations    []Registration
	RemainingBalance []Balance
}

type Registration struct {
	Time string
	Code string
}

type Balance struct {
	Category string
	Amount   string
}

func GetStatus(cfg config.Config) (string, error) {
	resp, err := client.Request(cfg, StatusURL)
	if err != nil {
		return "", err
	}
	log.Debugf("Successfully retreived %v", StatusURL)

	var status Status
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return "", err
	}

	log.Debug("Successfully decoded body")

	var checkInStatus string
	if len(status.Data.Registrations) > 0 {
		switch status.Data.Registrations[len(status.Data.Registrations)-1].Code {
		case "INN":
			checkInStatus = "Currently checked in."
		case "UT":
			checkInStatus = "Currently checked out."
		default:
			checkInStatus = "Unknown status."
		}
	}

	var flexiTimeStatus string
	for _, el := range status.Data.RemainingBalance {
		if el.Category == FlexiTimeCode {
			flexiTimeStatus = el.Amount
		}
	}

	log.Debug("Returning status")

	return fmt.Sprintf("%v Flexi time according to MinWinTid: %v", checkInStatus, flexiTimeStatus), nil
}
