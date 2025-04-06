package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//go:embed *.yaml
var configs embed.FS

func InitConfig() {
	env := os.Getenv("env")
	user := os.Getenv("user")
	password := os.Getenv("password")
	ip := os.Getenv("ip")
	vp := viper.New()
	configStream, err := configs.ReadFile("application." + env + ".yaml")
	if err != nil {
		panic(err)
	}
	vp.SetConfigType("yaml")
	err = vp.ReadConfig(bytes.NewReader(configStream))
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("app", &App)
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("database", &DB)
	if err != nil {
		panic(err)
	}
	DB.Master.Dsn = fmt.Sprintf(DB.Master.Dsn, user, password, ip)
	// 临时写成一样的
	DB.Slave.Dsn = fmt.Sprintf(DB.Slave.Dsn, user, password, ip)
}
