package initialize

import (
	"daigou-admin/config"
	"daigou-admin/global"
	"daigou-admin/utils/zaplog"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	confPath string //配置文件路径
)

func unmarshal(v *viper.Viper) {
	conf := &config.Server{}
	if err := v.Unmarshal(conf); err != nil {
		zaplog.Panicf("OnConfigChange err:%s", err)
	}
	global.G_Config = conf
}

func initConfig() {
	pflag.StringVar(&confPath, "confPath", "config.yaml", "set config path")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	v := viper.New()
	v.SetConfigFile(viper.GetString("confPath"))
	err := v.ReadInConfig()
	if err != nil {
		zaplog.Panicf("Fatal error config file: %s ", err)
	}
	unmarshal(v)
	global.GVA_VP = v
}
