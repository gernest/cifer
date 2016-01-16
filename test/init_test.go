package test

import (
	"log"
	"os"
	"testing"

	"github.com/gernest/cifer/core"
)

var testConfig *core.Config

func TestMain(m *testing.M) {
	c, err := config()
	if err != nil {
		log.Fatal("some fish configuring test environment ", err)
	}
	testConfig = c
	status := m.Run()
	testConfig.CloseSession()
	os.Exit(status)
}
func config() (*core.Config, error) {
	return &core.Config{
		DatabaseName: "cifer_test",
		Conn:         "mongodb://localhost",
	}, nil
}
