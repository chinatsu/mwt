package actions

import (
	"fmt"

	"github.com/chinatsu/mwt/client"
	"github.com/chinatsu/mwt/config"
	log "github.com/sirupsen/logrus"
)

func CheckOut(cfg config.Config) (string, error) {
	resp, err := client.Request(cfg, OutURL)
	if err != nil {
		return "", err
	}
	log.Debugf("Successfully sent a sign-out request at %v", OutURL)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Http Status: %v", resp.StatusCode)
	}
	return "You have checked out", nil
}

func CheckIn(cfg config.Config) (string, error) {
	resp, err := client.Request(cfg, InURL)
	if err != nil {
		return "", err
	}

	log.Debugf("Successfully sent a sign-in request at %v", InURL)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Http Status: %v", resp.StatusCode)
	}
	return "You have checked in", nil
}
