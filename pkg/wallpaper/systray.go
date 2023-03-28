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
	mWallpapers := make([]*systray.MenuItem, 8)
	for i := 0; i < 8; i++ {
		mWallpapers[i] = mChooseWallpaper.AddSubMenuItem(wallpaper.Images[i].Title, "")
	}

	c := cron.New()
	c.AddFunc("0 */10 * * * *", func() {
		_, wallpaper, err = GetWallpaper(fmt.Sprint(bounds.Dx()), fmt.Sprint(bounds.Dy()), 0, "", false)
		if err != nil {
			logrus.Errorln(err)
		}
		for i := 0; i < 8; i++ {
			mWallpapers[i].SetTitle(wallpaper.Images[i].Title)
		}
		if AutoUpdate {
			setbackground(0, bounds)
		}
	})
	c.Start()

	for i := 0; i < 8; i++ {
		index := i
		go func() {
			for {
				select {
				case <-mWallpapers[index].ClickedCh:
					setbackground(index, bounds)
				}
			}
		}()
	}

	<-mQuit.ClickedCh
	systray.Quit()
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
