package main

import (
	"fmt"
	"green-third-party-sdk/greensdk"

	uuid "github.com/satori/go.uuid"
)

const (
	accessKeyId     = "Your accessKeyId"
	accessKeySecret = "Your accessKeySecret"
)

// 阿里绿网文档地址：http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/pdf/lvwang-APIguide-cn-zh-2017-05-12.pdf
func main() {
	profile := greensdk.Profile{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret}
	path := "/green/text/scan"

	scenes := []string{"antispam"}
	task := greensdk.TextTask{DataId: uuid.NewV4().String(), Content: "test text scan"}
	tasks := []greensdk.TextTask{task}
	request := greensdk.Request{Scenes: scenes, Tasks: tasks}

	client := greensdk.AliYunClient{Profile: profile}

	fmt.Println(client.GetResponse(path, request))
}
