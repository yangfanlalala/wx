package wx

import "net/http"

// api Document https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html

const ApiContentCheckText = "https://api.weixin.qq.com/wxa/msg_sec_check"

type ContentLabel uint16
type ContentScene uint16

const (
	// 场景值
	ContentSceneData    ContentScene = 1
	ContentSceneComment ContentScene = 2
	ContentSceneBBS     ContentScene = 3
	ContentScenePost    ContentScene = 4
	// 建议处理结果
	ContentSuggestRicky  = "ricky"
	ContentSuggestPass   = "pass"
	ContentSuggestReview = "review"
	// 内容标签
	ContentLableOk        ContentLabel = 100
	ContentLableAd        ContentLabel = 10001
	ContentLabelPolitics  ContentLabel = 20001
	ContentLabelPorno     ContentLabel = 20002
	ContentLabelAbuse     ContentLabel = 20003
	ContentLabelCrime     ContentLabel = 20006
	ContentLabelCheat     ContentLabel = 20008
	ContentLabelVulgar    ContentLabel = 20012
	ContentLabelCopyright ContentLabel = 20013
	ContentLabelOther     ContentLabel = 21000
)

func (client *WeChatClient) ContentCheckText(data *ContentCheckTextRequest) (*ContentCheckTextResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiContentCheckText).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ContentCheckTextResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ContentCheckTextResponse, nil
}

func (client *WeChatClient) BuildContentCheckTextRequest() *ContentCheckTextRequest {
	return &ContentCheckTextRequest{
		Version: 2,
	}
}

func TransContentLabel(i ContentLabel) string {
	var result string
	switch i {
	case ContentLableOk:
		result = "正常"
	case ContentLableAd:
		result = "广告"
	case ContentLabelPolitics:
		result = "时政"
	case ContentLabelPorno:
		result = "色情"
	case ContentLabelAbuse:
		result = "辱骂"
	case ContentLabelCrime:
		result = "违法犯罪"
	case ContentLabelCheat:
		result = "欺诈"
	case ContentLabelVulgar:
		result = "低俗"
	case ContentLabelCopyright:
		result = "版权"
	case ContentLabelOther:
		result = "其他"
	default:
		result = "待确认"
	}
	return result
}

type ContentCheckTextRequest struct {
	AccessToken string       `position:"query" json:"-" name:"access_token"`
	Content     string       `position:"body" json:"-" name:"content"`
	Version     uint16       `position:"body" json:"" name:"version"`
	Scene       ContentScene `position:"body" json:"" name:"scene"`
	OpenID      string       `position:"body" json:"" name:"openid"`
	Title       string       `position:"body" json:"" name:"title"`
	Nickname    string       `position:"body" json:"" name:"nickname"`
	Signature   string       `position:"body" json:"" name:"signature"`
}
type ContentCheckResult struct {
	Suggest string `json:"suggest"`
	Label   string `json:"label"`
}
type ContentCheckDetail struct {
	Strategy  string `json:"strategy"`
	ErrorCode string `json:"errcode"`
	Suggest   string `json:"suggest"`
	Label     string `json:"label"`
	KeyWord   string `json:"keyword"`
	Prob      uint16 `json:"prob"`
}
type ContentCheckTextResponse struct {
	Detail  []*ContentCheckDetail `json:"detail"`
	TraceID string                `json:"trace_id"`
	Result  *ContentCheckResult   `json:"result"`
}
