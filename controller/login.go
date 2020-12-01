package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wechartTest/config"
	"wechartTest/utils"

	"github.com/gin-gonic/gin"
)

var WechatLogin = &config.WechatLoginResp{}

// login 微信登录
func Login(c *gin.Context) {
	code, ok := c.GetQuery("code")
	if !ok {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "code not send",
		})
		return
	}

	// https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + config.AppId + "&secret=" + config.AppSecret + "&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("http 去请求微接口j错误 错误", err)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ioutil.ReadAll 读取错误", err)
		return
	}

	err = json.Unmarshal(buf, WechatLogin)
	if err != nil {
		fmt.Println("json.Unmarshal 错误", err)
		return
	}

	token, err := utils.CreateToken(WechatLogin.OpenId)
	if err != nil {
		fmt.Println("utils.CustomerClaims 错误", err)
		return
	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  WechatLogin.ErrMsg,
		"data": token,
	})
}
