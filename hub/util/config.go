package logrusHelper

import (
	"io"
	"log"

	mate "github.com/heralight/logrus_mate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// UnmarshalConfiguration read configuration from viper.
// It returns a logrus_mate logger configuration instance.
func UnmarshalConfiguration(viper *viper.Viper) (conf mate.LoggerConfig) {
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	if err = conf.Validate(); err != nil {
		panic(err)
	}
	return
}

// SetConfig take a logrus logger instance and a conf mate.LoggerConfig.
// then apply conf specification to logrus instance.
// It returns an error if failed.
func SetConfig(logger *logrus.Logger, conf mate.LoggerConfig) (err error) {

	if conf.Out.Name == "" {
		conf.Out.Name = "stdout"
		conf.Out.Options = nil
	}

	var out io.Writer
	if out, err = mate.NewWriter(conf.Out.Name, conf.Out.Options); err != nil {
		return
	}

	logger.Out = out

	if conf.Formatter.Name == "" {
		conf.Formatter.Name = "text"
		conf.Formatter.Options = nil
	}

	var formatter logrus.Formatter
	if formatter, err = mate.NewFormatter(conf.Formatter.Name, conf.Formatter.Options); err != nil {
		return
	}

	logger.Formatter = formatter

	if conf.Hooks != nil {
		for _, hookConf := range conf.Hooks {
			var hook logrus.Hook
			if hook, err = mate.NewHook(hookConf.Name, hookConf.Options); err != nil {
				panic(err)
			}
			logger.Hooks.Add(hook)
		}
	}

	var lvl = logrus.DebugLevel
	if lvl, err = logrus.ParseLevel(conf.Level); err != nil {
		return
	} else {
		logger.Level = lvl
	}

	return
}
