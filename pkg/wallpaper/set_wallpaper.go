package wallpaper

import (
	wallpaperLib "github.com/arno4000/wallpaper"
)

func SetWallpaper(path string) {
	wallpaperLib.SetFromFile(path)
}
