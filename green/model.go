package greensdk

type ClientInfo struct {
	SdkVersion  string `json:"sdkVersion"`
	CfgVersion  string `json:"cfgVersion"`
	UserType    string `json:"userType"`
	UserId      string `json:"userId"`
	UserNick    string `json:"userNick"`
	Avatar      string `json:"avatar"`
	Imei        string `json:"imei"`
	Imsi        string `json:"imsi"`
	Umid        string `json:"umid"`
	Ip          string `json:"ip"`
	Os          string `json:"os"`
	Channel     string `json:"channel"`
	HostAppName string `json:"hostAppName"`
	HostPackage string `json:"hostPackage"`
	HostVersion string `json:"hostVersion"`
}

// Request
type BizData struct {
	BizType  string   `json:"bizType,omitempty"`
	Scenes   []string `json:"scenes,omitempty"`
	Tasks    []*Task  `json:"tasks" valid:"required"`
	Callback string   `json:"callback,omitempty"`
}

type Task struct {
	DataId   string `json:"dataId" valid:"required,range(1|99999999999999999)"`
	Url      string `json:"url"`
	Interval int    `json:"interval"`
	Source   string `json:"source" valid:"required,in(qupost|quduopai)"`
}

type Request struct {
	Date string `json:"date"`
	BizData
}

// Response
type Response struct {
	DataSlice []*Data `bson:"data,omitempty,omitempty" json:"data,omitempty,omitempty"`
	Date      string  `bson:"date,omitempty" json:"date,omitempty"`
	Source    string  `bson:"source,omitempty" json:"source,omitempty"`
	Type      string  `bson:"type,omitempty" json:"type,omitempty"`
}

type Data struct {
	Code    int       `bson:"code" json:"code"`
	Msg     string    `bson:"msg" json:"msg"`
	TaskID  string    `bson:"taskId" json:"taskId"`
	DataID  string    `bson:"dataId" json:"dataId"`
	URL     string    `bson:"url" json:"url"`
	Results []*Result `bson:"results,omitempty,omitempty" json:"results,omitempty,omitempty"`
}

type Result struct {
	Label      string   `bson:"label" json:"label"`
	Rate       float64  `bson:"rate" json:"rate"`
	Scene      string   `bson:"scene" json:"scene"`
	Suggestion string   `bson:"suggestion" json:"suggestion"`
	Frames     []*Frame `bson:"frames,omitempty,omitempty" json:"frames,omitempty,omitempty"`
}

type Frame struct {
	Label  string  `bson:"label" json:"label"`
	Offset int     `bson:"offset" json:"offset"`
	Rate   float64 `bson:"rate" json:"rate"`
	URL    string  `bson:"url" json:"url"`
}
