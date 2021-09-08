package wallpaper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetWallpaper(resolutionWidth string, resolutionHeight string, daysBack int, path string, saveImage bool) (string, Wallpaper, error) {
	if daysBack > 7 {
		logrus.Fatalln("Only 7 Days back are supported by the API!")
	}
	if daysBack < 0 {
		logrus.Fatalln("Number must be Between 0 and 7")
	}
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
	imageURL := wallpaper.Images[daysBack].URL
	if saveImage {
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
		if path == "" {
			if runtime.GOOS == "windows" {
				path = os.Getenv("TEMP") + `\`
			} else if runtime.GOOS == "linux" {
				path = "/tmp/"
			} else {
				logrus.Fatalln(runtime.GOOS, "is currently not supported")
			}
		}
		if runtime.GOOS == "windows" {
			if !strings.HasSuffix(path, `\`) {
				path = path + `\`
			}
		} else {
			if !strings.HasSuffix(path, "/") {
				path = path + "/"
			}
		}
		err = ioutil.WriteFile(path+wallpaper.Images[daysBack].Startdate+".jpg", image, 0777)
		path = path + wallpaper.Images[daysBack].Startdate + ".jpg"
	}
	return path, wallpaper, err

}
