package actions

import (
	"fmt"

	"github.com/chinatsu/mwt/client"
	"github.com/chinatsu/mwt/config"
)

func CheckOut(cfg config.Config) (string, error) {
	resp, err := client.Request(cfg, OutURL)
	if err != nil {
		return "", err
	}

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

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Http Status: %v", resp.StatusCode)
	}
	return "You have checked in", nil
}
