package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
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
}

type ErrorResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
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
	accessToken := &AccessToken{}
	resp, err := http.Get(fmt.Sprintf(AccessTokenUrlFormat, appId, appSecret))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	wechatResponse := &AccessTokenResponse{}
	err = json.Unmarshal(body, wechatResponse)
	if err != nil {
		// convert to normal response failed, try convert to errresp
		errResp := &ErrorResponse{}
		errToErrStruct := json.Unmarshal(body, errResp)
		if errToErrStruct != nil {
			return nil, errToErrStruct
		}
		return nil, errors.New(errResp.ErrorMessage)
	}
	accessToken.Token = wechatResponse.AccessToken
	accessToken.ExpiredAt = currentTimeStamp + wechatResponse.ExpiresIn
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

	postBody, err := json.Marshal(wechatUnionRequest)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(fmt.Sprintf(UnionIDUrlFormat, accessToken), "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		return "", nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	wechatUnionIDResponse := &UnionIDResponse{}
	err = json.Unmarshal(body, wechatUnionIDResponse)
	if err != nil {
		// convert to normal response failed, try convert to errresp
		errResp := &ErrorResponse{}
		errToErrStruct := json.Unmarshal(body, errResp)
		if errToErrStruct != nil {
			return "", errToErrStruct
		}
		return "", errors.New(errResp.ErrorMessage)
	}

	return wechatUnionIDResponse.UnionID, nil
}
