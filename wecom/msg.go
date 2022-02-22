package wecom

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

//type MessageRequest struct {
//	ToUser  string `json:"touser"`
//	ToParty string `json:"toparty"`
//	ToTag   string `json:"totag"`
//	MsgType string `json:"msgtype"`
//	AgentId int    `json:"agentid"`
//	Text    struct {
//		Content string `json:"content"`
//	} `json:"text"`
//	Safe                   int `json:"safe"`
//	EnableIDTrans          int `json:"enable_id_trans"`
//	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
//	DuplicateCheckInterval int `json:"duplicate_check_interval"`
//}

type MessageRequest struct {
	ToUser                 string        `json:"touser,omitempty"`
	ToParty                string        `json:"toparty,omitempty"`
	ToTag                  string        `json:"totag,omitempty"`
	MsgType                string        `json:"msgtype,omitempty"`
	AgentId                int           `json:"agentid,omitempty"`
	Text                   *Text         `json:"text,omitempty"`
	TemplateCard           *TemplateCard `json:"template_card,omitempty"`
	Image                  *Media        `json:"image,omitempty"`
	Voice                  *Media        `json:"voice,omitempty"`
	File                   *Media        `json:"file,omitempty"`
	TextCard               *TextCard     `json:"text_card,omitempty"`
	News                   *News         `json:"news,omitempty"`
	Markdown               *Text         `json:"markdown,omitempty"`
	EnableIDTrans          int           `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int           `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int           `json:"duplicate_check_interval,omitempty"`
}

type News struct {
	Articles []Articles `json:"articles,omitempty"`
}
type Articles struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	PicUrl      string `json:"picurl,omitempty"`
	Appid       string `json:"appid,omitempty"`
	PagePath    string `json:"pagepath,omitempty"`
}

type TextCard struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	BtnTxt      string `json:"btntxt,omitempty"`
}

type Media struct {
	MediaID string `json:"media_id,omitempty"`
}

type Text struct {
	Content string `json:"content,omitempty"`
}

type Source struct {
	IconURL   string `json:"icon_url,omitempty"`
	Desc      string `json:"desc,omitempty"`
	DescColor int    `json:"desc_color,omitempty"`
}
type ActionList struct {
	Text string `json:"text,omitempty"`
	Key  string `json:"key,omitempty"`
}
type ActionMenu struct {
	Desc       string       `json:"desc,omitempty"`
	ActionList []ActionList `json:"action_list,omitempty"`
}
type MainTitle struct {
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
}
type QuoteArea struct {
	Type      int    `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	Title     string `json:"title,omitempty"`
	QuoteText string `json:"quote_text,omitempty"`
}
type ImageTextArea struct {
	Type     int    `json:"type,omitempty"`
	URL      string `json:"url,omitempty"`
	Title    string `json:"title,omitempty"`
	Desc     string `json:"desc,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}
type CardImage struct {
	URL         string  `json:"url,omitempty"`
	AspectRatio float64 `json:"aspect_ratio,omitempty"`
}
type VerticalContentList struct {
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
}
type HorizontalContentList struct {
	KeyName string `json:"keyname,omitempty"`
	Value   string `json:"value,omitempty"`
	Type    int    `json:"type,omitempty"`
	URL     string `json:"url,omitempty"`
	MediaID string `json:"media_id,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
type JumpList struct {
	Type     int    `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	URL      string `json:"url,omitempty"`
	Appid    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}
type CardAction struct {
	Type     int    `json:"type,omitempty"`
	URL      string `json:"url,omitempty"`
	Appid    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}
type TemplateCard struct {
	CardType              string                  `json:"card_type,omitempty"`
	Source                *Source                 `json:"source,omitempty"`
	ActionMenu            *ActionMenu             `json:"action_menu,omitempty"`
	TaskID                string                  `json:"task_id,omitempty"`
	MainTitle             *MainTitle              `json:"main_title,omitempty"`
	QuoteArea             *QuoteArea              `json:"quote_area,omitempty"`
	ImageTextArea         *ImageTextArea          `json:"image_text_area,omitempty"`
	CardImage             *CardImage              `json:"card_image,omitempty"`
	VerticalContentList   []VerticalContentList   `json:"vertical_content_list,omitempty"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list,omitempty"`
	JumpList              []JumpList              `json:"jump_list,omitempty"`
	CardAction            *CardAction             `json:"card_action,omitempty"`
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
	sendMessageUrlFormat    = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%v"
	MessageTypeTemplateCard = "template_card"
	MessageTypeText         = "text"
	MessageTypeImage        = "image"
	MessageTypeVoice        = "voice"
	MessageTypeVideo        = "video"
	MessageTypeFile         = "file"
	MessageTypeTextCard     = "textcard"
	MessageTypeNews         = "news"
	MessageTypeMarkdown     = "markdown"
	CardTypeNewsNotice      = "news_notice"
)

func SendMessage(messageRequest MessageRequest, accessToken string) error {
	url := fmt.Sprintf(sendMessageUrlFormat, accessToken)
	client := resty.New()
	messageResponse := &MessageResponse{}
	resp, err := client.R().SetResult(messageResponse).SetBody(messageRequest).Post(url)
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
