package main

import (
	"WebMonitor/helper"
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	timeout    int
	sleep      int
	apiAddress string
)

// ApiPost 请求数据
type ApiPost struct {
	Username string `json:"username"` // 用户名（登陆账户自动获取）
	Key      string `json:"key"`      // 用户密钥（登陆账户自动获取）
	Url      string `json:"url"`      // 需要检测的URL地址
}

// ApiResult 返回结果
type ApiResult struct {
	Code     string `json:"code"`     // 1001为域名正常 1002为已被拦截
	Msg      string `json:"msg"`      // 系统返回提示信息！
	Statu    string `json:"statu"`    // 异常为false，正常为true！
	Count    int    `json:"count"`    // 次数包，如需扣次数则返回！
	Reason   string `json:"reason"`   // 拦截原因，如有则返回！
	Describe string `json:"describe"` // 拦截描述，如有则返回！
	Url      string `json:"url"`      // 检测的地址！
}

// ApiMonitorInterface 请求接口
type ApiMonitorInterface interface {
	Check()
	FormatOutput()
}

type API struct {
	url     string
	Type    string     // 请求地址
	ApiPost ApiPost    // API 请求数据
	result  *ApiResult // 返回结果
}

func (r API) Check() {
	api := fmt.Sprintf("%s%s?username=%s&key=%s&url=%s",
		apiAddress,
		r.Type,
		r.ApiPost.Username,
		r.ApiPost.Key,
		r.ApiPost.Url)
	resp := helper.Response{}
	resp.Get(api, timeout, nil)
	err := json.Unmarshal([]byte(resp.Body), r.result)
	if err != nil {
		log.Println(resp.Body)
		log.Fatal("[Err] 解析JSON失败, ", err)
	}
}

func (r API) FormatOutput() {
	switch r.result.Code {
	case "1001":
		fmt.Printf("%6s%2s|", "正常", "")
		break
	case "1002":
		fmt.Printf("%6s%2s|", "拦截", "")
		break
	default:
		fmt.Printf("%5s%1s|", "数据异常", "")
		break
	}
	// 休息1秒
	time.Sleep(time.Duration(sleep) * time.Second)
}

func main() {
	fmt.Printf("|%90s%92s|\n", "R00kieT00ls 全自动浏览器异常排查工具", "Version 1.2")

	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Fatal("加载配置文件失败:", err)
	}

	timeout, err = cfg.Section("").Key("timeout").Int()
	if err != nil {
		log.Fatal("加载配置文件失败:", err)
	}
	sleep, err = cfg.Section("").Key("sleep").Int()
	if err != nil {
		log.Fatal("加载配置文件失败:", err)
	}

	apiAddress = cfg.Section("").Key("apiAddress").String()

	fi, err := os.Open("./urls.txt")
	if err != nil {
		log.Fatal("加载 urls.txt 错误: ", err)
	}
	defer fi.Close()

	fmt.Printf("|%20s%20s|%6s%2s|%6s%1s|%5s%1s|%6s%2s|%6s%4s|%6s%4s|%8s%2s|%8s%2s|%6s%2s|%6s%2s|%6s%2s|%6s%2s|%6s%0s|%7s%3s|\n",
		"Url", "",
		"长城", "",
		"移动墙", "",
		"移动污染", "",
		"微信", "",
		"QQ", "",
		"UC", "",
		"OPPO", "",
		"VIVO", "",
		"夸克", "",
		"搜狗", "",
		"华为", "",
		"小米", "",
		"净网云剑", "",
		"DNS ", "")

	br := bufio.NewReader(fi)
	for {
		url, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if url[0] == '#' {
			continue
		}
		// 跳过空行
		tmp := strings.TrimSpace(strings.Trim(string(url), "\n"))
		if len(tmp) == 0 {
			continue
		}
		apiPost := ApiPost{
			Username: cfg.Section("").Key("username").String(),
			Key:      cfg.Section("").Key("key").String(),
			Url:      tmp,
		}
		fmt.Printf("|%-40s|", url)
		qiang := API{Type: "qiang", ApiPost: apiPost, result: &ApiResult{}}
		qiang.Check()
		qiang.FormatOutput()

		ydjc := API{Type: "ydjc", ApiPost: apiPost, result: &ApiResult{}}
		ydjc.Check()
		ydjc.FormatOutput()

		yddnsjc := API{Type: "yddnsjc", ApiPost: apiPost, result: &ApiResult{}}
		yddnsjc.Check()
		yddnsjc.FormatOutput()

		weixin := API{Type: "wx", ApiPost: apiPost, result: &ApiResult{}}
		weixin.Check()
		weixin.FormatOutput()

		qq := API{Type: "qq", ApiPost: apiPost, result: &ApiResult{}}
		qq.Check()
		qq.FormatOutput()

		uc := API{Type: "uc", ApiPost: apiPost, result: &ApiResult{}}
		uc.Check()
		uc.FormatOutput()

		oppo := API{Type: "oppo", ApiPost: apiPost, result: &ApiResult{}}
		oppo.Check()
		oppo.FormatOutput()

		vivo := API{Type: "vivo", ApiPost: apiPost, result: &ApiResult{}}
		vivo.Check()
		vivo.FormatOutput()

		quark := API{Type: "quark", ApiPost: apiPost, result: &ApiResult{}}
		quark.Check()
		quark.FormatOutput()

		sogou := API{Type: "sogou", ApiPost: apiPost, result: &ApiResult{}}
		sogou.Check()
		sogou.FormatOutput()

		huawei := API{Type: "huawei", ApiPost: apiPost, result: &ApiResult{}}
		huawei.Check()
		huawei.FormatOutput()

		mi := API{Type: "mi", ApiPost: apiPost, result: &ApiResult{}}
		mi.Check()
		mi.FormatOutput()

		jwyj := API{Type: "jwyj", ApiPost: apiPost, result: &ApiResult{}}
		jwyj.Check()
		jwyj.FormatOutput()

		dns := API{Type: "dns", ApiPost: apiPost, result: &ApiResult{}}
		dns.Check()
		dns.FormatOutput()

		fmt.Println("")
	}
}
