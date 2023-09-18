package configHandler

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Log struct {
	Level   string `mapstructure:"level"`
	Methods bool   `mapstructure:"methods"`
}

type Database struct {
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Db_name    string `mapstructure:"db_name"`
	Ssl_mode   string `mapstructure:"ssl_mode"`
	Table_name string `mapstructure:"table_name"`
}

type Config struct {
	Log      Log
	Database Database
}

func Init() *Config {
	config := Config{}
	config.Log.Level = "info"
	config.Log.Methods = false
	return &config
}

func LoadConfiguration() *Config {
	vpr := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))

	// Specify the name and type of the configuration file
	vpr.SetConfigName("config")
	vpr.SetConfigType("yaml")

	vpr.AddConfigPath("/etc/inventory-service")
	vpr.AddConfigPath("./config")
	vpr.AddConfigPath(".")

	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("cannot access executable directory, error : %v", err)
	}
	vpr.AddConfigPath(path.Dir(exe))

	if err := vpr.ReadInConfig(); err != nil {
		log.Fatalf("config file reading error : %v", err)
	}

	vpr.SetEnvPrefix("INVENTORY_SERVICE")
	vpr.AutomaticEnv()

	config := Init()
	if err := vpr.Unmarshal(config, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc()))); err != nil {
		log.Fatalf("failed to decode configuration to struct with error: %v", err)
	}

	var level log.Level
	if err := level.UnmarshalText([]byte(config.Log.Level)); err != nil {
		log.Fatal("log level not valid, error:", err)
	}
	log.SetLevel(level)
	log.SetReportCaller(config.Log.Methods)

	return config
}

func (config *Config) GetConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Db_name,
		config.Database.Ssl_mode,
	)
}
