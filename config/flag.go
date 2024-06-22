package config

import "flag"

var (
	Postgresql bool
	Dev        bool
	Redis      bool
)

func InitFlag() {
	flag.BoolVar(&Postgresql, "postgresql", false, "Whether to use postgresql as database")
	flag.BoolVar(&Dev, "dev", false, "Whether to run fontend in dev mode")
	flag.BoolVar(&Redis, "redis", false, "Whether to use redis as cache")
	flag.Parse()
}
