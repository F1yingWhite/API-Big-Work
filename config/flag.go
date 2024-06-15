package config

import "flag"

var (
	Postgresql bool
	Dev        bool
)

func InitFlag() {
	flag.BoolVar(&Postgresql, "postgresql", false, "Whether to use postgresql as database")
	flag.BoolVar(&Dev, "dev", false, "Whether to run fontend in dev mode")
	flag.Parse()
}
