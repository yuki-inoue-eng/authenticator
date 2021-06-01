package configs

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rakyll/statik/fs"
	_ "github.com/yuki-inoue-eng/authenticator/configs/statik"
	"github.com/yuki-inoue-eng/authenticator/pkg/safe"
	"gopkg.in/yaml.v3"
)

type Configs interface {
	Load() error
	Get() *configs
}

type configs struct {
	Server   Server
	Database Database
	Jwt      Jwt
}

type Server struct {
	Port             string `yaml:"port"`
	ShutdownWaitTIme int    `yaml:"shutdown_wait_time"`
}

type Database struct {
	HostName     string `yaml:"host_name"`
	Port         string `yaml:"port"`
	UserName     string `yaml:"user_name"`
	UserPassword string `yaml:"user_password"`
	MaxOpenCon   int    `yaml:"max_open_connection"`
	MaxIdleCon   int    `yaml:"max_idle_connection"`
	MaxLifeTime  int    `yaml:"max_life_time"`
	Name         string `yaml:"name"`
}

type Jwt struct {
	Exp int64 `yaml:"exp"`
}

func New() (Configs, error) {
	cfg := &configs{}
	if err := cfg.Load(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *configs) Load() error {
	statikFs, err := fs.New()
	if err != nil {
		return fmt.Errorf("failed to generate fs: %v", err)
	}

	os.Getenv("env")
	fle, err := statikFs.Open("/configs.yml")
	if err != nil {
		return err
	}
	defer safe.Close(fle)
	byts, err := ioutil.ReadAll(fle)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(byts, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *configs) Get() *configs {
	return c
}
