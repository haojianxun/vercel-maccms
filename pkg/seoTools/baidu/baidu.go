package baidu

import (
	"fmt"
	"goapi/pkg/config"
	"io/ioutil"
	"net/http"
	"strings"
)

func SubmitURLBatch(siteURL string, urlList []string) (error, string) {
	api := fmt.Sprintf("http://data.zz.baidu.com/urls?site=%s&token=%s", siteURL, config.GetString("seo.baidu.apikey"))
	payload := strings.Join(urlList, "\n")

	client := &http.Client{}
	req, err := http.NewRequest("POST", api, strings.NewReader(payload))
	if err != nil {
		return err, ""
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	return nil, string(body)
}
