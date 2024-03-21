package v1

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/config"
	"goapi/pkg/maccms"
	"goapi/pkg/mysql"
	"goapi/pkg/page"
	"net/http"
	"time"
)

type SeoController struct {
	BaseController
}

type CDATA struct {
	Text string `xml:",cdata"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName     struct{} `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Language    string   `xml:"language"`
	Docs        string   `xml:"docs"`
	Generator   string   `xml:"generator"`
	Image       struct {
		URL string `xml:"url"`
	} `xml:"image"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	PubDate     string `xml:"pubDate"`
	Description CDATA  `xml:"description"`
}

func (h *SeoController) Bing(c *gin.Context) {
	startTime := time.Date(2024, 03, 20, 0, 0, 0, 0, time.FixedZone("CST", 8*3600))
	timeDiff := time.Now().Unix() - startTime.Unix()
	day := timeDiff / 86400
	CurrentPage := day
	// 获取视屏数据
	var pageList page.PageList
	where := cmap.New().Items()
	where["vod_status"] = 1
	models.MacVodMgr(mysql.DB).Where(where).Count(&pageList.Total)
	pageList.CurrentPage = CurrentPage
	pageList.PageSize = 100
	page.InitPageList(&pageList)
	var listResult []models.MacVod
	models.MacVodMgr(mysql.DB).
		Where(where).
		Offset(int(pageList.Offset)).
		Limit(int(pageList.PageSize)).
		Order("vod_id asc").
		Find(&listResult)

	macData := config.Get("maccms")
	site := macData.(map[string]interface{})["site"]
	siteInfo := site.(config.StrMap)
	channel := Channel{
		Title:       siteInfo["site_name"].(string),
		Description: siteInfo["site_description"].(string),
		Link:        siteInfo["site_url"].(string),
		Language:    "zh-cn",
		Docs:        siteInfo["site_name"].(string),
		Generator:   "Rss Powered By " + siteInfo["site_url"].(string),
		Image: struct {
			URL string `xml:"url"`
		}{
			URL: siteInfo["site_logo"].(string),
		},
	}

	for _, data := range listResult {
		channel.Items = append(channel.Items, Item{
			Title:   data.VodName,
			Link:    fmt.Sprintf("https://%s/show-%s.html", siteInfo["site_url"], maccms.EncryptID(data.VodID)),
			Author:  data.VodAuthor,
			PubDate: data.VodPubdate,
			Description: CDATA{
				Text: data.VodContent,
			},
		})
	}

	rss := Rss{
		Version: "2.0",
		Channel: channel,
	}

	xmlData, err := xml.MarshalIndent(rss, "", "    ")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	finalXml := xml.Header + string(xmlData)
	c.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(finalXml))
}
