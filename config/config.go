package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database
	NoClean  bool     // Don't clean comments and blank lines from output
	NoHeader bool     // Don't write the "generated by" header comment
	PGDump   string   // pg_dump executable name (defaults to "pg_dump")
	Schemas  []string // Schemas to dump (default is to dump all schemas)
	Writer   *os.File // Writer to write schema dump to (defaults to stdout)
	Redis    Redis
	SMTP     SMTP
}
type Database struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	DbName    string `yaml:"dbName"`
	DbUser    string `yaml:"dbuser"`
	SslEnable string `yaml:"sslenable"`
	Password  string `yaml:"password"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type SMTP struct {
	HostName    string `yaml:"hostName"`
	MailAddress string `yaml:"mailAddress"`
	Password    string `yaml:"password"`
	SmtpHost    string `yaml:"smtpHost"`
	SmtpPort    string `yaml:"smtpPort"`
}

var Cfg Config

func LoadConfig() error {
	viper.AddConfigPath("./config")
	// export ENV=production // To load the production env
	if os.Getenv("ENV") == "production" {
		viper.SetConfigName("production")
	} else {
		viper.SetConfigName("dev")
	}
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	viper.AutomaticEnv()
	// viper.BindEnv()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Cfg)
	return err
}
