package wechat

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type UnionIDRequest struct {
	UserList []User `json:"user_list"`
}

type User struct {
	OpenID string `json:"openid"`
	Lang   string `json:"lang"`
}

type UnionIDResponse struct {
	Subscribe      int    `json:"subscribe"`
	OpenID         string `json:"openid"`
	Language       string `json:"language"`
	SubscribeTime  int64  `json:"subscribe_time"`
	UnionID        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupID        int    `json:"groupid"`
	TagIDList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QRScene        int    `json:"qr_scene"`
	QRSceneStr     string `json:"qr_scene_str"`
	ErrorCode      int    `json:"errcode"`
	ErrorMessage   string `json:"errmsg"`
}

const (
	AccessTokenUrlFormat = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	UnionIDUrlFormat     = "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=%s"
)

type AccessToken struct {
	Token     string
	ExpiredAt int64
}

func GetAccessToken(appId, appSecret string) (*AccessToken, error) {
	currentTimeStamp := time.Now().Unix()
	accessTokenResponse := &AccessTokenResponse{}
	url := fmt.Sprintf(AccessTokenUrlFormat, appId, appSecret)
	client := resty.New()
	resp, err := client.R().SetResult(accessTokenResponse).Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("http_status_%v", resp.StatusCode()))
	}
	if accessTokenResponse.ErrorMessage != "" {
		return nil, errors.New(accessTokenResponse.ErrorMessage)
	}

	accessToken := &AccessToken{
		Token:     accessTokenResponse.AccessToken,
		ExpiredAt: currentTimeStamp + accessTokenResponse.ExpiresIn,
	}

	return accessToken, nil
}

func GetUnionID(openID string, accessToken string) (string, error) {
	wechatUnionRequest := &UnionIDRequest{
		UserList: []User{
			{
				OpenID: openID,
			},
		},
	}

	url := fmt.Sprintf(UnionIDUrlFormat, accessToken)
	client := resty.New()

	unionIdResponse := &UnionIDResponse{}
	resp, err := client.R().SetResult(unionIdResponse).SetHeader("content-type", "application/json").SetBody(wechatUnionRequest).Post(url)

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "nil", errors.New(fmt.Sprintf("http_status_%v", resp.StatusCode()))
	}
	if unionIdResponse.ErrorMessage != "" {
		return "", errors.New(unionIdResponse.ErrorMessage)
	}

	return unionIdResponse.UnionID, nil
}
