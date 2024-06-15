package config

import "flag"

var (
	Postgresql bool
)

func InitFlag() {
	flag.BoolVar(&Postgresql, "postgresql", true, "Whether to use postgresql as database")
	flag.Parse()
}
