package wecom

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

type TextMessage struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentId int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int `json:"safe"`
	EnableIDTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

type MessageResponse struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
	MsgId        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

const (
	SendMessageUrlFormat = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%v"
)

func SendTextMessageByID(msg string, recipients []string, agentId int, accessToken string) error {
	receiver := strings.Join(recipients, "|")

	textMessage := TextMessage{
		ToUser:  receiver,
		MsgType: "text",
		AgentId: agentId,
		Text: struct {
			Content string `json:"content"`
		}{
			Content: msg,
		},
		Safe:                   1,
		EnableIDTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 0,
	}

	url := fmt.Sprintf(SendMessageUrlFormat, accessToken)
	client := resty.New()
	messageResponse := &MessageResponse{}
	resp, err := client.R().SetResult(messageResponse).SetBody(textMessage).Post(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(fmt.Sprintf("status_code_%d", resp.StatusCode()))
	}

	if messageResponse.ErrCode != 0 {
		return errors.New(messageResponse.ErrMsg)
	}

	return nil
}

func SendTextMessageByTag(msg string, recipients []string, agentId int, accessToken string) error {
	receiver := strings.Join(recipients, "|")

	textMessage := TextMessage{
		ToTag:   receiver,
		MsgType: "text",
		AgentId: agentId,
		Text: struct {
			Content string `json:"content"`
		}{
			Content: msg,
		},
		Safe:                   1,
		EnableIDTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 0,
	}

	url := fmt.Sprintf(SendMessageUrlFormat, accessToken)
	client := resty.New()
	messageResponse := &MessageResponse{}
	resp, err := client.R().SetResult(messageResponse).SetBody(textMessage).Post(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(fmt.Sprintf("status_code_%d", resp.StatusCode()))
	}

	if messageResponse.ErrCode != 0 {
		return errors.New(messageResponse.ErrMsg)
	}

	return nil
}
