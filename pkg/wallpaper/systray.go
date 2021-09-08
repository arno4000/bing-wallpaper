package wallpaper

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/getlantern/systray"
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var AutoUpdate bool
var AutoUpdateConfig bool

func Systray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	bounds := screenshot.GetDisplayBounds(0)
	_, wallpaper, err := GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), 0, "", false)
	icon, err := os.ReadFile("logo.png")
	if err != nil {
		logrus.Errorln(err)
	}
	systray.SetIcon(icon)
	mChooseWallpaper := systray.AddMenuItem("Choose wallpaper", "Choose a wallpaper of the last 7 days")
	mAutoMode := systray.AddMenuItemCheckbox("Always use newest wallpaper", "", AutoUpdateConfig)
	AutoUpdate = mAutoMode.Checked()

	mWallpaper0 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[0].Title, "")
	mWallpaper1 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[1].Title, "")
	mWallpaper2 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[2].Title, "")
	mWallpaper3 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[3].Title, "")
	mWallpaper4 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[4].Title, "")
	mWallpaper5 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[5].Title, "")
	mWallpaper6 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[6].Title, "")
	mWallpaper7 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[7].Title, "")
	for {
		select {
		case <-mWallpaper0.ClickedCh:
			setbackground(0, bounds)
		case <-mWallpaper1.ClickedCh:
			setbackground(1, bounds)
		case <-mWallpaper2.ClickedCh:
			setbackground(2, bounds)
		case <-mWallpaper3.ClickedCh:
			setbackground(3, bounds)
		case <-mWallpaper4.ClickedCh:
			setbackground(4, bounds)
		case <-mWallpaper5.ClickedCh:
			setbackground(5, bounds)
		case <-mWallpaper6.ClickedCh:
			setbackground(6, bounds)
		case <-mWallpaper7.ClickedCh:
			setbackground(7, bounds)
		case <-mAutoMode.ClickedCh:
			var config Config
			configName := os.Getenv("HOME") + "/.bing-wallpaper.yaml"
			config = Config{
				Daemon: AutoUpdate,
			}
			fmt.Println(config)
			fmt.Println(mAutoMode.Checked())
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
}

func onExit() {
	fmt.Println("KEKBye")
}

func setbackground(daysback int, bounds image.Rectangle) {
	wallpaperImage, _, err := GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), daysback, os.Getenv("HOME")+"/.cache", true)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(wallpaperImage)
	SetWallpaper(wallpaperImage)
}
