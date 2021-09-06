package wallpaper

import (
	wallpaperLib "github.com/reujab/wallpaper"
)

func SetWallpaper(path string) {
	wallpaperLib.SetFromFile(path)
}
