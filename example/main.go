package main

import (
	"fmt"
	"github.com/sunblisshealth/wechat-sdk/wecom"
)

func main() {
	token, err := wecom.GetAccessToken("coprid", "corpsecret")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(token.AccessToken)
	}
	messageRequest := wecom.MessageRequest{
		ToUser:  "useid1|useid2",
		MsgType: wecom.MessageTypeTemplateCard,
		AgentId: 1000003,

		TemplateCard: &wecom.TemplateCard{
			CardType: wecom.CardTypeNewsNotice,
			MainTitle: &wecom.MainTitle{
				Title: "客户刚刚进行了扫码",
			},
			ImageTextArea: &wecom.ImageTextArea{
				Type:     0,
				URL:      "",
				Title:    "xxx",
				ImageURL: "https://avatars.githubusercontent.com/u/4082168?v=4",
			},
			HorizontalContentList: []wecom.HorizontalContentList{
				{
					KeyName: "产品",
					Value:   "小金旦",
				},
				{
					KeyName: "防伪码",
					Value:   "123123123123",
				},
				{
					KeyName: "查询时间",
					Value:   "2022年2月22日",
				},
				{
					KeyName: "查询次数",
					Value:   "10",
				},
				{
					KeyName: "领取红包",
					Value:   "未领取",
				},
			},
			CardAction: &wecom.CardAction{
				Type:     1,
				URL:      "https://www.baidu.com",
				Appid:    "",
				PagePath: "",
			},
		},

		EnableIDTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 0,
	}

	err = wecom.SendMessage(messageRequest, token.AccessToken)
	if err != nil {
		fmt.Println(err.Error())
	}
}
