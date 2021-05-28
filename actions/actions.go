package actions

import (
	"fmt"

	"github.com/chinatsu/mwt/config"
)

const (
	StatusAction   = "status"
	CheckInAction  = "in"
	CheckOutAction = "out"
)

const (
	StatusURL = "https://minwintid.adeo.no/MinWintid/Registration/GetInitialData"
	OutURL    = "https://minwintid.adeo.no/MinWintid/Registration/RegisterOut"
	InURL     = "https://minwintid.adeo.no/MinWintid/Registration/RegisterIn"
)

func PerformAction(cfg config.Config) (string, error) {
	switch cfg.Action {
	case StatusAction:
		return GetStatus(cfg)

	case CheckOutAction:
		return CheckOut(cfg)
	case CheckInAction:
		return CheckIn(cfg)
	default:
		return "", fmt.Errorf("Unimplemented")
	}
}
