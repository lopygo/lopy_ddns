package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/common"
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

func (p HttpClient) Request(adapterInstance common.IAdapter) ([]byte, error) {

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}

	// data
	//// public data
	postData := url.Values{}

	// check
	err = adapterInstance.Check()
	if err != nil {
		return nil, err
	}
	//// adapter data
	for k, v := range adapterInstance.GetData() {
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
		return nil, err
	}

	email := "example@example.cn"
	if len(p.conf.Email) > 0 {
		email = p.conf.Email
	}

	newRequest.Header.Add("UserAgent", fmt.Sprintf("LOPY DDNS Client/0.0.0 (%s)", email))
	// 这句不能要
	// newRequest.Header.Add("User-Agent", fmt.Sprintf("LOPY DDNS Client/0.0.0 (%s)", email))
	newRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	// client.do
	r, err := client.Do(newRequest)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	// err

	// this status code
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("http status error %d", r.StatusCode)
	}
	// return

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
