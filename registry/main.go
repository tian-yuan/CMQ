package main

import (
	_ "github.com/heralight/logrus_mate/hooks/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"fmt"
	"github.com/tian-yuan/CMQ/registry/commands"
	"github.com/tian-yuan/iot-common/util"
	"runtime"
	"strings"
	"os"
)

func initLogger() {

	// ########## Init Viper
	var viper = viper.New()

	viper.SetConfigName("registry") // name of config file (without extension), here we use some logrus_mate sample
	viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath("./conf")               // optionally look for config in the working directory
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// ########### End Init Viper

	// Read configuration
	var c = util.UnmarshalConfiguration(viper) // Unmarshal configuration from Viper
	util.SetConfig(logrus.StandardLogger(), c) // for e.g. apply it to logrus default instance

	// ### End Read Configuration
}

func writePid() {
	pathName, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("get executable path failed : %s\n", err))
		return
	}
	index := strings.LastIndex(pathName, "/")
	path := string(pathName[0:index])
	executableName := string(pathName[index+1 : len(pathName)])
	pidFile := path + "/pid/" + executableName + ".pid"
	if err = util.WritePidFile(pidFile); err != nil {
		panic(fmt.Errorf("write pid file failed, pid file : %s, err : %s\n", pidFile, err))
	}
	logrus.Infof("write pid file %s success.", pidFile)
}

func main() {
	initLogger()
	runtime.GOMAXPROCS(runtime.NumCPU())

	writePid()
	commands.Execute()

	stopCh := util.SetupSignalHandler()
	<-stopCh

	logrus.Infof("hub stop")
	commands.Stop()
}

