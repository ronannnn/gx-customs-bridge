package internal

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/ronannnn/infra/cfg"
	"github.com/ronannnn/infra/db"
	"github.com/ronannnn/infra/i18n"
	"github.com/ronannnn/infra/log"
	"github.com/ronannnn/infra/mq/rabbitmq"
	"github.com/ronannnn/infra/services/jwt/accesstoken"
	"github.com/ronannnn/infra/services/jwt/refreshtoken"
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
	Sys  cfg.Sys  `mapstructure:"sys"`
	User cfg.User `mapstructure:"user"`

	Log          log.Cfg          `mapstructure:"log"`
	AccessToken  accesstoken.Cfg  `mapstructure:"access-token"`
	RefreshToken refreshtoken.Cfg `mapstructure:"refresh-token"`
	Db           db.Cfg           `mapstructure:"db"`
	Rabbitmq     rabbitmq.Cfg     `mapstructure:"rabbitmq"`
	I18n         i18n.Cfg         `mapstructure:"i18n"`

	Customs CustomsCfg `mapstructure:"customs"`
}

func ProvideSysCfg(cfg *Cfg) *cfg.Sys {
	return &cfg.Sys
}

func ProvideUserCfg(cfg *Cfg) *cfg.User {
	return &cfg.User
}

func ProvideLogCfg(cfg *Cfg) *log.Cfg {
	return &cfg.Log
}

func ProvideDbCfg(cfg *Cfg) *db.Cfg {
	return &cfg.Db
}

func ProvideAccessTokenCfg(cfg *Cfg) *accesstoken.Cfg {
	return &cfg.AccessToken
}

func ProvideRefreshTokenCfg(cfg *Cfg) *refreshtoken.Cfg {
	return &cfg.RefreshToken
}

func ProvideRabbitmqCfg(cfg *Cfg) *rabbitmq.Cfg {
	return &cfg.Rabbitmq
}

func ProvideI18nCfg(cfg *Cfg) *i18n.Cfg {
	return &cfg.I18n
}

func ProvideCustomsCfg(cfg *Cfg) *CustomsCfg {
	return &cfg.Customs
}

func ProvideCfg() (*Cfg, error) {
	return NewCfg(provideConfigFilepath())
}

func provideConfigFilepath() string {
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
