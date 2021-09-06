package wallpaper

import (
	"fmt"

	wallpaperLib "github.com/reujab/wallpaper"
	"github.com/sirupsen/logrus"
)

func SetWallpaper() {
	wallpaperPath, err := GetWallpaper("3400", "1440", "")
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(wallpaperPath)
	err = wallpaperLib.SetFromFile(wallpaperPath)
	if err != nil {
		logrus.Errorln(err)
	}
}
