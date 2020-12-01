package config

const (
	AppId     = "wx903b2abfe0f1a32f"
	AppSecret = "07eeb2a109fef7bc2009d924411288e4"
)

// Config
type Config struct {
}

type WechatLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
