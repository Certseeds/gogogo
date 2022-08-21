package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func GetRequester(url string, token string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	// FIXME: method这个参数的类型检查做的太差了,还不如typescript的字符串枚举
	// 传入"get"既不报错,也不转化成GET,而是默认就这么往下走
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	request.Header.Set("User-Agent", runtime.Version())
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Host", "api.github.com")
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	serviceResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return serviceResp, nil
}
