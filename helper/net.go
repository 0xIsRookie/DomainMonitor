package helper

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Headers http.Header
	// Cookies    http.Cookie // 功能未实现
	Body       string
	Url        string
	IP         string
	StatusCode int
}

func (r *Response) sendRequest(url string, data string, timeout int, headers map[string]string, method string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[NetworkError] - [%s]: %v\n", url, r)
			return
		}
	}()
	var req *http.Request
	switch method {
	case http.MethodGet:
		req, _ = http.NewRequest(http.MethodGet, url, nil)
		break
	case http.MethodPost:
		req, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(data)))
		break
	default:
		log.Fatal("[Err] 请求参数异常:[", method, "]请求方法未定义")
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err := net.Dial(network, addr)
				if conn != nil {
					req.RemoteAddr = conn.RemoteAddr().String()
				} else {
					req.RemoteAddr = ""
				}
				return conn, err
			},
		},
	}

	// 格式化 headers 请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, _ := client.Do(req)
	// 函数结束时关闭相关链接
	defer resp.Body.Close()

	r.StatusCode = resp.StatusCode
	r.Headers = resp.Header
	r.IP = strings.Split(req.RemoteAddr, ":")[0]
	body, _ := ioutil.ReadAll(resp.Body)
	r.Body = string(body)
	r.Url = url
}

func (r *Response) Get(url string, timeout int, headers map[string]string) {
	r.sendRequest(url, "", timeout, headers, http.MethodGet)
}

func (r *Response) Post(url string, data string, timeout int, headers map[string]string) {
	r.sendRequest(url, "", timeout, headers, http.MethodPost)
}
