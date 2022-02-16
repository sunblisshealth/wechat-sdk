package wecom

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type AccessTokenResponse struct {
	ErrorCode   int    `json:"errorcode"`
	ErrorMsg    string `json:"errormsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type AccessToken struct {
	AccessToken string
	ExpiresAt   int64
}

const (
	AccessTokenUrlFormat = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%v&corpsecret=%v"
)

func GetAccessToken(corpId, corpSecret string) (*AccessToken, error) {
	client := resty.New()
	now := time.Now().Unix()
	url := fmt.Sprintf(AccessTokenUrlFormat, corpId, corpSecret)

	accessTokenResponse := &AccessTokenResponse{}

	resp, err := client.R().SetResult(accessTokenResponse).Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(fmt.Sprintf("status_code_%d", resp.StatusCode()))
	}
	accessToken := &AccessToken{
		AccessToken: accessTokenResponse.AccessToken,
		ExpiresAt:   now + accessTokenResponse.ExpiresIn,
	}

	return accessToken, nil
}
