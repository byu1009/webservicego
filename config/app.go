package config

import "os"

type ConfSistem struct {
	AddFarmasi string
}

var Sistem = ConfSistem{
	AddFarmasi: os.Getenv("ADD_FARMASI"),
}