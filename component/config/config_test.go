package config

import (
	"testing"
)

func TestConfigBasic(t *testing.T) {
	Init("test.conf", "../..")
}
