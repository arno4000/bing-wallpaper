/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bing-wallpaper/pkg/wallpaper"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	wallpaperLib "github.com/reujab/wallpaper"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "bing-wallpaper",
	Short:            "bing-wallpaper is a tool written in go (golang) to get the daily wallper from bing and set it as wallpaper",
	TraverseChildren: true,
	Run:              runRoot,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bing-wallpaper.yaml)")
	rootCmd.PersistentFlags().IntP("daysback", "d", 0, "Number of days in the past to get the wallpaper from")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".bing-wallpaper" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".bing-wallpaper")
		configName := os.Getenv("HOME") + "/.bing-wallpaper.yaml"
		if _, err := os.Stat(configName); os.IsNotExist(err) {
			var config wallpaper.Config
			config = wallpaper.Config{
				Daemon: true,
			}
			b, err := yaml.Marshal(config)
			if err != nil {
				logrus.Errorln(err)
			}
			err = ioutil.WriteFile(configName, b, 0755)
			if err != nil {
				logrus.Errorln(err)
			}
		}

	}

	viper.AutomaticEnv() // d in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func runRoot(c *cobra.Command, args []string) {
	bounds := screenshot.GetDisplayBounds(0)
	daysBack, err := c.Flags().GetInt("daysback")
	if err != nil {
		logrus.Errorln(err)
	}
	wallpaper.AutoUpdateConfig = viper.Get("daemon").(bool)
	wallpaperPath, wallpaperStruct, err := wallpaper.GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), daysBack, "", true)
	if err != nil {
		logrus.Errorln(err)
	}
	currentWallpaper, err := wallpaperLib.Get()
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(currentWallpaper)
	if strings.Contains(currentWallpaper, wallpaperStruct.Images[1].Startdate) {
		wallpaper.SetWallpaper(wallpaperPath)
	}
}
