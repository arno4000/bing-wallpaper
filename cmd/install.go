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
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a systemd service file and install the binary on the system",

	Run: runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)

}

func runInstall(c *cobra.Command, args []string) {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		logrus.Fatalln("HAHA USE LINUX, DONT USE", runtime.GOOS, "KEKW")
	}
	path := os.Getenv("PATH")
	home := os.Getenv("HOME")
	var installLocation string
	binaryPath, err := os.Executable()
	exeName := filepath.Base(binaryPath)
	if err != nil {
		logrus.Errorln(err)
	}
	if strings.Contains(path, home+"/.local/bin") {
		installLocation = home + "/.local/bin"
	} else {
		logrus.Fatalln("Please add", home+".local/bin to your PATH environment variable and retry the installation")
	}
	systemDService := `[Unit]
Description=Bing Wallpaper Service
After=network.target
[Service]
Type=simple
ExecStart=` + installLocation + `/` + exeName + ` --daemon
Restart=on-failure
RestartSec=3
[Install]
WantedBy=default.target`

	if _, err := os.Stat(home); os.IsNotExist(err) {
		logrus.Warnln(installLocation, "does not exist, creating it. Please make shure that it is in the PATH variable")
		err = os.Mkdir(installLocation, 0777)
		if err != nil {
			logrus.Errorln(err)
		}
	}
	binarySrc, err := os.Open(binaryPath)
	if err != nil {
		logrus.Errorln(err)
	}
	defer binarySrc.Close()
	binaryDst, err := os.Create(installLocation + "/" + exeName)
	if err != nil {
		logrus.Errorln(err)
	}
	defer binaryDst.Close()
	_, err = io.Copy(binaryDst, binarySrc)
	if err != nil {
		logrus.Errorln(err)
	}
	err = os.Chmod(installLocation+"/"+exeName, 0755)
	if err != nil {
		logrus.Errorln(err)
	}
	err = os.Remove(binaryPath)
	if err != nil {
		logrus.Errorln(err)
	}
	err = ioutil.WriteFile(home+"/.config/systemd/user/bing-wallpaper.service", []byte(systemDService), 0644)
	if err != nil {
		logrus.Errorln(err)
	}
	cmd := exec.Command("systemctl --user daemon-reload")
	cmd.Run()
	cmd = exec.Command("systemctl start --user bing-wallpaper.service")
	cmd.Run()

}
