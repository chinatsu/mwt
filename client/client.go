package client

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chinatsu/mwt/config"
)

func Request(cfg config.Config, url string) (*http.Response, error) {
	httpClient := http.Client{}

	body := fmt.Sprintf("{\"activeIdentity\":{\"EmployeeId\":%v,\"PositionId\":\"%v\",\"IsPrimary\":true,\"Description\":\"%v\"},\"activeRole\":1}", cfg.EmployeeID, cfg.PositionID, cfg.PositionID)
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%v=%v;%v=%v", config.ASPNETSessionIDName, cfg.AspNetSessionId, config.MRHSessionName, cfg.MRHSession))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
