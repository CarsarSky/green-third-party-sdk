package greensdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AliYunClient struct {
	Profile Profile
}

func (greenClient AliYunClient) GetResponse(path string, reqData Request) (string, error) {
	reqDataByte, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, host+path, strings.NewReader(string(reqDataByte)))

	if err != nil {
		fmt.Printf("build request err. req:%v,err:%v", reqData, err)
		return "", err
	}
	headers := buildRequestHeader(string(reqDataByte), path, greenClient.Profile.AccessKeyId, greenClient.Profile.AccessKeySecret)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("get resp err. req:%v,err:%v", reqData, err)
		return "", err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("close body err. err:%v", err)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read body err. req:%v,err:%v", reqData, err)
		return "", err
	}

	return string(body), nil
}
