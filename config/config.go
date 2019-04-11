package config

import (
	"bytes"
	"rilutham/stovia/lib/log"
	"rilutham/stovia/lib/utils"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
)

var (
	Path = []string{
		"$GOPATH/src/rilutham/stovia",
		"/var/app/current/stovia",
		"/var/app/current",
		".",
	}
	envReplacer = strings.NewReplacer(".", "_")
	encodingMap = map[string]string{
		":tab:": "\t",
	}
)

// DatabaseDSN :nodoc:
func DatabaseDSN() string {
	return viper.GetString("database_dsn")
}

// Version :nodoc:
func Version() string {
	v, err := utils.Asset("resources/VERSION")
	if err == nil {
		return string(bytes.TrimSpace(v))
	}

	return "Unknown"
}

func init() {
	var errors *multierror.Error
	log.For("config", "init").Info("Finding configuration on defined location")
	viper.SetDefault("config_name", "pnq")

	for _, in := range Path {
		viper.AddConfigPath(in)
	}

	viper.SetEnvPrefix("PRICESRV")
	viper.SetEnvKeyReplacer(envReplacer)
	viper.AutomaticEnv()

	err := viper.MergeInConfig()
	if err != nil {
		errors = multierror.Append(errors, err)
	}

	if errors != nil {
		log.For("config", "init").Infof("Error while reading config file %v", errors)
	}
}
