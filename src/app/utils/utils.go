package utils

import (
	"app/models"
	"app/settings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/grokify/html-strip-tags-go"
	"github.com/satori/go.uuid"
	"image"
	"io"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GetImageExtension(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return
}

func GetTmpDir(subDir string) string {
	return path.Join("/tmp", subDir)
}

func GetPostOnViewList(o orm.Ormer, p []models.Post) (m []models.PostOnView) {
	for _, v := range p {
		user, err := models.FindUserById(o, v.UserId)
		if err != nil {
			beego.Error(err)
		}
		l := models.PostOnView{
			Id:          v.Id,
			Title:       v.Title,
			Description: getPostDescription(v.Content),
			Image:       parseImageFromHtml(v.Content),
			User:        user,
			CreatedAt:   formatTime(v.CreatedAt),
		}
		m = append(m, l)
	}
	return
}

func getPostDescription(html string) (result string) {
	c := strip.StripTags(html)
	words := strings.Split(c, " ")
	counter := 0
	var d []string
	for _, v := range words {
		if counter == settings.PostDescriptionWordsCount {
			break
		}
		d = append(d, v)
		counter++
	}
	result = strings.Join(d, " ") + " ... "
	return
}

func parseImageFromHtml(html string) (result string) {
	images := findImages(html)
	if len(images) > 0 {
		result = "<img src='" + images[0] + "' alt=''>"
	}
	return
}

var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

// if your img's are properly formed with doublequotes then use this, it's more efficient.
// var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)
func findImages(htm string) []string {
	imgs := imgRE.FindAllStringSubmatch(htm, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

func formatTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	year, month, day := t.Date()
	return strconv.Itoa(day) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(year)
}
