package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/adapter/record"
	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/config"
)

const apiUrl = "https://dnsapi.cn"

// const apiUrl = "http://127.0.0.1:9004"

type HttpClient struct {
	conf *config.Config
}

func NewHttpClient(conf *config.Config) *HttpClient {
	i := new(HttpClient)

	i.conf = conf
	return i
}

func (p HttpClient) Request(adapterInstance record.IAdapter) (buf []byte, err error) {

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return
	}

	// data
	//// public data
	postData := url.Values{}

	// check
	err = adapterInstance.Check()
	if err != nil {
		return
	}
	//// adapter data
	for k, v := range adapterInstance.GetPostData() {
		postData.Set(k, v)
	}

	postData.Set("login_token", fmt.Sprintf("%s,%s", p.conf.TokenId, p.conf.Token))
	postData.Set("format", "json")
	postData.Set("lang", p.conf.Lang)
	errorOnEmpty := "no"
	if p.conf.ErrorOnEmpty {
		errorOnEmpty = "yes"
	}
	postData.Set("error_on_empty", errorOnEmpty)

	// request

	// res

	u.Path = adapterInstance.Method()
	fmt.Println(u.String())
	newRequest, err := http.NewRequest("POST", u.String(), strings.NewReader(postData.Encode()))
	if err != nil {
		return
	}

	email := "example@example.cn"
	if len(p.conf.Email) > 0 {
		email = p.conf.Email
	}

	newRequest.Header.Add("UserAgent", fmt.Sprintf("LOPY DDNS Client/0.0.0 (%s)", email))
	// 这句不能要
	// newRequest.Header.Add("User-Agent", fmt.Sprintf("LOPY DDNS Client/0.0.0 (%s)", email))
	newRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
	}
	// client.do
	r, err := client.Do(newRequest)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	// err

	// this status code
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("http status error %d", r.StatusCode)
	}
	// return

	buf, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	// res

	res := record.Response{}
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return
	}

	code, err := res.Status.Code.Int64()
	if err != nil {
		return
	}

	if code != 1 {
		err = fmt.Errorf("api error: %v", res.Status.Message)
		return
	}

	err = adapterInstance.SetResponseJson(buf)
	return

}
