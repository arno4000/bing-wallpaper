package wallpaper

import (
	"fmt"
	"image"
	"os"

	"github.com/getlantern/systray"
	"github.com/kbinani/screenshot"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var AutoUpdate bool

func Systray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	bounds := screenshot.GetDisplayBounds(0)
	_, wallpaper, err := GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), 0, "", false)
	if err != nil {
		logrus.Errorln(err)
	}
	systray.SetIcon(Logo)
	mChooseWallpaper := systray.AddMenuItem("Choose wallpaper", "Choose a wallpaper of the last 7 days")
	mQuit := systray.AddMenuItem("Quit", "")
	mWallpaper0 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[0].Title, "")
	mWallpaper1 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[1].Title, "")
	mWallpaper2 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[2].Title, "")
	mWallpaper3 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[3].Title, "")
	mWallpaper4 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[4].Title, "")
	mWallpaper5 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[5].Title, "")
	mWallpaper6 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[6].Title, "")
	mWallpaper7 := mChooseWallpaper.AddSubMenuItem(wallpaper.Images[7].Title, "")

	c := cron.New()
	c.AddFunc("*/10 * * * *", func() {
		_, wallpaper, err = GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), 0, "", false)
		if err != nil {
			logrus.Errorln(err)
		}
		mWallpaper0.SetTitle(wallpaper.Images[0].Title)
		mWallpaper1.SetTitle(wallpaper.Images[1].Title)
		mWallpaper2.SetTitle(wallpaper.Images[2].Title)
		mWallpaper3.SetTitle(wallpaper.Images[3].Title)
		mWallpaper4.SetTitle(wallpaper.Images[4].Title)
		mWallpaper5.SetTitle(wallpaper.Images[5].Title)
		mWallpaper6.SetTitle(wallpaper.Images[6].Title)
		mWallpaper7.SetTitle(wallpaper.Images[7].Title)
		if AutoUpdate {
			setbackground(0, bounds)
		}
	})
	c.Start()

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
		case <-mQuit.ClickedCh:
			systray.Quit()

		}
	}
}

func onExit() {
	os.Exit(0)
}

func setbackground(daysback int, bounds image.Rectangle) {
	wallpaperImage, _, err := GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), daysback, os.Getenv("HOME")+"/.cache", true)
	if err != nil {
		logrus.Errorln(err)
	}
	SetWallpaper(wallpaperImage)
}
