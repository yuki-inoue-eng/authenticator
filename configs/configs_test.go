package configs_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/yuki-inoue-eng/authenticator/configs"
)

func TestConfigs_Load(t *testing.T) {
	cfgs, err := configs.New()
	if err != nil {
		t.Error(err)
	}
	err = cfgs.Load()
	if err != nil {
		t.Error(err)
	}
	spew.Dump(cfgs)
}
