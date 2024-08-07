package internal

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ronannnn/infra/cfg"
)

const (
	// app
	ApplicationName    = "GX-CUSTOMS-BRIDGE"
	ApplicationVersion = "0.0.1"

	// config
	ConfigDir         = "configs"
	ConfigEnvKey      = "GX_CUSTOMS_BRIDGE_CONFIG"
	ConfigDefaultFile = "config.default.toml"
	ConfigTestFile    = "config.test.toml"
	ConfigReleaseFile = "config.release.toml"
)

type CustomsCfg struct {
	ImpPath         string `mapstructure:"imp-path"`           // 海关文件夹目录
	Inv101SysId     string `mapstructure:"inv101-sys-id"`      // inv101中的sysId
	Sas121DclErConc string `mapstructure:"sas121-dcl-er-conc"` // sas121中的申请人，即卡号对应的人
	IcCardNo        string `mapstructure:"ic-card-no"`         // 操作卡号
	OperCusRegCode  string `mapstructure:"oper-cus-reg-code"`  // 操作卡的海关十位
}

type Cfg struct {
	Sys     cfg.Sys    `mapstructure:"sys"`
	Log     cfg.Log    `mapstructure:"log"`
	Auth    cfg.Auth   `mapstructure:"auth"`
	Db      cfg.Db     `mapstructure:"db"`
	User    cfg.User   `mapstructure:"user"`
	Customs CustomsCfg `mapstructure:"customs"`
}

func ProvideLogCfg(cfg *Cfg) *cfg.Log {
	return &cfg.Log
}

func ProvideSysCfg(cfg *Cfg) *cfg.Sys {
	return &cfg.Sys
}

func ProvideDbCfg(cfg *Cfg) *cfg.Db {
	return &cfg.Db
}

func ProvideAuthCfg(cfg *Cfg) *cfg.Auth {
	return &cfg.Auth
}

func ProvideUserCfg(cfg *Cfg) *cfg.User {
	return &cfg.User
}

func ProvideCustomsCfg(cfg *Cfg) *CustomsCfg {
	return &cfg.Customs
}

func ProvideCfg() (*Cfg, error) {
	return NewCfg(provideConfigFilepath())
}

func provideConfigFilepath() string {
	env := os.Environ()
	for _, e := range env {
		pair := strings.Split(e, "=")
		key := pair[0]
		value := pair[1]
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	var configFilepath string
	if configEnv := os.Getenv(ConfigEnvKey); configEnv == "" { // env not found
		configFilepath = path.Join(ConfigDir, ConfigDefaultFile)
		fmt.Printf("config filepath: %s\n", configFilepath)
	} else { // internal.ConfigEnvKey 常量存储的环境变量不为空 将值赋值于config
		configFilepath = configEnv
		fmt.Printf("env: %s, config filepath: %s\n", ConfigEnvKey, configFilepath)
	}
	return configFilepath
}

func NewCfg(configFilePath string) (configs *Cfg, err error) {
	// reset time zone
	time.Local = time.FixedZone("utc", 0)
	configs = &Cfg{}
	err = cfg.ReadFromFile(configFilePath, configs)
	return
}
