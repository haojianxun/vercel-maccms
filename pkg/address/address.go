package address

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Address struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

/**
 * . 提取请求IP的波段
 * @param ip
 * @return
 */

func IpDomainByIp(ip string, num int) string {
	str := ""
	if 0 == len(ip) {
		return str
	}
	nums := strings.Split(ip, ".")
	for index, string_ := range nums {
		if index < num {
			if index == num-1 {
				str = str + string_
			} else {
				str = str + string_ + "."
			}
		}
	}
	return str
}

// 缓存 ip

func IpSet(ip string, isLimit bool) {
	if isLimit {
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:ip:%v", ip), "true", 60*60*24)
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:noallow:%v", ip), time.Now().Format("2006-01-02 15:04:05"), 60*60*24)
	} else {
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:ip:%v", ip), "false", 60*60*24)
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:allow:%v", ip), time.Now().Format("2006-01-02 15:04:05"), 60*60*24)
	}
}

/*
  - {
    "status": "success",
    "country": "Indonesia",
    "countryCode": "ID",
    "region": "SU",
    "regionName": "North Sumatra",
    "city": "Medan",
    "zip": "",
    "lat": 3.5847,
    "lon": 98.6629,
    "timezone": "Asia/Jakarta",
    "isp": "SMART-TELECOM",
    "org": "",
    "as": "AS18004 PT WIRELESS INDONESIA ( WIN )",
    "query": "114.79.56.80"
    }
    *
*/
func GetMessageByIp(ip string) Address {
	var address Address
	url := fmt.Sprintf("http://ip-api.com/json/%v", ip)
	Client := new(http.Client)
	response, err := Client.Get(url)
	if err != nil {
		logger.Error(err)
		return address
	}
	if response.StatusCode != 200 {
		return address
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err)
		return address
	}
	err = json.Unmarshal(body, &address)
	if err != nil {
		logger.Error(err)
	}
	return address
}
