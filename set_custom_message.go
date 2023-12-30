package wx

import "net/http"

// api document https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/kf-mgnt/kf-message/sendCustomMessage.html

const ApiSendCustomMessage = "https://api.weixin.qq.com/cgi-bin/message/custom/send"

type CustomMessageType string

const (
	CustomMessageTypeText            = "text"
	CustomMessageTypeImage           = "image"
	CustomMessageTypeLink            = "link"
	CustomMessageTypeMiniProgramPage = "miniprogrampage"
)

func (client *WeChatClient) SendCustomMessage(data *SendCustomMessageRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiSetBetaWeAppNickname).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) BuildSendCustomMessageRequest() *SendCustomMessageRequest {
	return &SendCustomMessageRequest{}
}

type SendCustomMessageRequest struct {
	AccessToken     string                        `position:"query" name:"access_token" json:"-"`
	ToUser          string                        `position:"body" name:"touser" json:"touser"`
	MessageType     CustomMessageType             `position:"body" name:"msgtype" json:"msgtype"`
	Text            *CustomTextMessage            `position:"body" name:"text" json:"text,omitempty"`
	Image           *CustomImageMessage           `position:"body" name:"image" json:"image,omitempty"`
	Link            *CustomLinkMessage            `position:"body" name:"link" json:"link,omitempty"`
	MiniProgramPage *CustomMiniProgramPageMessage `position:"body" name:"miniprogrampage" json:"miniprogrampage,omitempty"`
}

type CustomTextMessage struct {
	Content string `json:"content"`
}
type CustomImageMessage struct {
	MediaID string `json:"media_id"`
}
type CustomLinkMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}
type CustomMiniProgramPageMessage struct {
	Title        string `json:"title"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

type SendCustomMessageResponse struct {
}
