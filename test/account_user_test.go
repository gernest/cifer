package test

import (
	"testing"

	"github.com/gernest/cifer/account"
	"github.com/gernest/cifer/core"
)

func TestAccount_Create(t *testing.T) {
	usr := &core.User{
		Email:    "test@cifer.io",
		Password: "pass",
	}
	result, err := account.Create(testConfig, usr)
	if err != nil {
		t.Fatal(err)
	}
	if result.ID == "" {
		t.Error("expected a id")
	}
}
