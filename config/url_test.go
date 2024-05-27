package config_test

import (
	"log"
	"net/url"
	"testing"
)

func TestGetPath(t *testing.T) {
	path := "movies/(o^^o)#感觉 #可爱老婆.mp4"
	log.Println(url.PathEscape(path))
}
