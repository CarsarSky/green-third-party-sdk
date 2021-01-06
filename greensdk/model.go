package greensdk

type Profile struct {
	AccessKeyId     string
	AccessKeySecret string
}

// 文本类型任务
type TextTask struct {
	DataId  string `json:"dataId"`
	Content string `json:"content"`
}

type Request struct {
	Scenes []string    `json:"scenes"`
	Tasks  interface{} `json:"tasks"`
}
