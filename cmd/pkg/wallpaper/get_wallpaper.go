package wallpaper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/sirupsen/logrus"
)

func GetWallpaper(resolutionWidth string, resolutionHeight string, date string) (string, error) {
	url := "https://bingwallpaper.microsoft.com/api/BWC/getHPImages?screenWidth=" + resolutionWidth + "&screenHeight=" + resolutionHeight + "&env=live"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Errorln(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorln(err)
	}
	defer res.Body.Close()
	var wallpaper Wallpaper
	err = json.Unmarshal(body, &wallpaper)
	if err != nil {
		logrus.Errorln(err)
	}
	imageURL := wallpaper.Images[0].URL
	req, err = http.NewRequest("GET", imageURL, nil)
	if err != nil {
		logrus.Errorln(err)
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorln(err)
	}
	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorln(err)
	}
	defer res.Body.Close()
	var imagePath string
	if runtime.GOOS == "windows" {
		imagePath = /*os.Getenv("TEMP") + "\\*/ "wallpaper.jpg"
	} else if runtime.GOOS == "linux" {
		imagePath = "/tmp/wallpaper.jpg"
	} else {
		logrus.Fatalln(runtime.GOOS, "is currently not supported")
	}
	ioutil.WriteFile(imagePath, image, 0777)
	return imagePath, err

}
