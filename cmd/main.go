package main

import (
	"fmt"
	"os"

	"github.com/chinatsu/mwt/actions"
	"github.com/chinatsu/mwt/config"
	"github.com/chinatsu/mwt/logger"
	flag "github.com/spf13/pflag"
)

var cfg = config.DefaultConfig()

func init() {
	flag.StringVar(&cfg.AspNetSessionId, "asp-net-session-id", os.Getenv("ASP_NET_SESSION_ID"), "The cookie named ASP.NET_SessionId")
	flag.StringVar(&cfg.MRHSession, "mrh-session", os.Getenv("MRH_SESSION"), "The cookie named MRHSession")
	flag.StringVar(&cfg.EmployeeID, "employee-id", os.Getenv("EMPLOYEE_ID"), "Employee ID (found in any POST request body made to MinWinTid)")
	flag.StringVar(&cfg.PositionID, "position-id", os.Getenv("POSITION_ID"), "Position ID (found in any POST request body made to MinWinTid")
	flag.StringVar(&cfg.LogLevel, "loglevel", "info", "logging level")
	flag.Parse()

	logger.Setup(cfg.LogLevel)

	if flag.NArg() > 0 {
		cfg.Action = flag.Arg(0)
	} else {
		println("Missing required positional argument [action]")
		os.Exit(1)
	}

	if err := cfg.Verify(); err != nil {
		println(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

func main() {
	ret, err := actions.PerformAction(cfg)
	if err != nil {
		println(fmt.Sprintf("Error: %v", err))
		os.Exit(1)
	}
	fmt.Printf("%+v\n", ret)
}
