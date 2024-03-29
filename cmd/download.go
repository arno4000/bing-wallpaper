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

	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Only download the wallpaper",
	Args:  cobra.ExactArgs(1),
	Run:   runDownload,
}

func init() {
	rootCmd.AddCommand(downloadCmd)

}
func runDownload(c *cobra.Command, args []string) {
	daysBack, err := c.Flags().GetInt("daysback")
	if err != nil {
		logrus.Errorln(err)
	}
	bounds := screenshot.GetDisplayBounds(0)
	_, _, err = wallpaper.GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), daysBack, args[0], true)
	if err != nil {
		logrus.Errorln(err)
	}
}
