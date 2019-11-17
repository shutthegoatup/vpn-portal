package app

import (
	"flag"
)

var config string

func setupVars() {

	flag.StringVar(&config, "config", "configs/conf.yaml", "Path to the config file")
	flag.Parse()

}
